package portable

import (
	"bufio"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const maxIntelBytes = 8 << 20
const maxIndicators = 100000
const bazaarURL = "https://mb-api.abuse.ch/api/v1/"
const kevURL = "https://www.cisa.gov/sites/default/files/feeds/known_exploited_vulnerabilities.json"

func intelClient() *http.Client {
	return &http.Client{Timeout: 45 * time.Second, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
}
func intelBody(c *http.Client, req *http.Request) ([]byte, error) {
	resp, e := c.Do(req)
	if e != nil {
		return nil, fmt.Errorf("Consulta externa indisponível: %w", e)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fonte externa respondeu HTTP %d", resp.StatusCode)
	}
	b, e := io.ReadAll(io.LimitReader(resp.Body, maxIntelBytes+1))
	if e != nil {
		return nil, e
	}
	if len(b) > maxIntelBytes {
		return nil, fmt.Errorf("Resposta acima do limite de segurança")
	}
	return b, nil
}
func validHash(h string) bool {
	if len(h) != 64 {
		return false
	}
	_, e := hex.DecodeString(h)
	return e == nil
}

// Only strict SHA-256 indicators. No paths, samples, scripts or commands are imported.
func parseHashList(raw string) (map[string]string, error) {
	if len(raw) > maxIntelBytes {
		return nil, fmt.Errorf("Lista acima do limite de 8 MiB")
	}
	out := map[string]string{}
	sc := bufio.NewScanner(strings.NewReader(raw))
	sc.Buffer(make([]byte, 4096), 4096)
	for n := 1; sc.Scan(); n++ {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		h := strings.ToLower(parts[0])
		if !validHash(h) {
			return nil, fmt.Errorf("SHA-256 inválido na linha %d", n)
		}
		out[h] = "Indicador importado pelo usuário (fonte não verificada pelo aplicativo)"
		if len(out) > maxIndicators {
			return nil, fmt.Errorf("Limite de indicadores excedido")
		}
	}
	if e := sc.Err(); e != nil {
		return nil, e
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("Nenhum indicador válido")
	}
	return out, nil
}
func parseBazaar(b []byte) (map[string]string, error) {
	var r struct {
		Status string `json:"query_status"`
		Data   []struct {
			Hash string `json:"sha256_hash"`
		} `json:"data"`
	}
	if e := json.Unmarshal(b, &r); e != nil {
		return nil, fmt.Errorf("Resposta MalwareBazaar inválida")
	}
	if r.Status != "ok" || len(r.Data) == 0 || len(r.Data) > 100 {
		return nil, fmt.Errorf("Consulta MalwareBazaar sem resultado válido; confira a chave e disponibilidade")
	}
	out := map[string]string{}
	for _, v := range r.Data {
		h := strings.ToLower(v.Hash)
		if !validHash(h) {
			return nil, fmt.Errorf("Indicador inválido na fonte MalwareBazaar")
		}
		out[h] = "MalwareBazaar: indicador SHA-256 publicado"
	}
	return out, nil
}
func fetchBazaar(ctx context.Context, c *http.Client, key string) (map[string]string, error) {
	key = strings.TrimSpace(key)
	if len(key) < 8 || len(key) > 256 || strings.ContainsAny(key, "\r\n\t ") {
		return nil, fmt.Errorf("Informe uma Auth-Key válida do MalwareBazaar")
	}
	form := url.Values{"query": {"get_recent"}, "selector": {"100"}}
	req, e := http.NewRequestWithContext(ctx, http.MethodPost, bazaarURL, strings.NewReader(form.Encode()))
	if e != nil {
		return nil, e
	}
	req.Header.Set("Auth-Key", key)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	b, e := intelBody(c, req)
	if e != nil {
		return nil, e
	}
	return parseBazaar(b)
}

// Return a new map or fail without modifying the existing scanner database.
func mergeIntel(old, added map[string]string) (map[string]string, error) {
	out := make(map[string]string, len(old)+len(added))
	for k, v := range old {
		out[k] = v
	}
	for k, v := range added {
		if !validHash(k) {
			return nil, fmt.Errorf("Indicador inválido")
		}
		if prior, ok := out[k]; !ok || strings.HasPrefix(prior, "Indicador importado") {
			out[k] = v
		}
		if len(out) > maxIndicators {
			return nil, fmt.Errorf("Limite de 100 mil indicadores atingido")
		}
	}
	return out, nil
}

var cvePattern = regexp.MustCompile(`^CVE-[0-9]{4}-[0-9]{4,19}$`)

type kevEntry struct {
	ID         string `json:"cveID"`
	Vendor     string `json:"vendorProject"`
	Product    string `json:"product"`
	Name       string `json:"vulnerabilityName"`
	Action     string `json:"requiredAction"`
	Ransomware string `json:"knownRansomwareCampaignUse"`
}
type kevResult struct {
	Version string    `json:"version"`
	Count   int       `json:"count"`
	Found   bool      `json:"found"`
	Entry   *kevEntry `json:"entry,omitempty"`
}

func parseKEV(b []byte, id string) (kevResult, error) {
	id = strings.ToUpper(strings.TrimSpace(id))
	if !cvePattern.MatchString(id) {
		return kevResult{}, fmt.Errorf("Informe um identificador CVE válido")
	}
	var r struct {
		Version string     `json:"catalogVersion"`
		Count   int        `json:"count"`
		Items   []kevEntry `json:"vulnerabilities"`
	}
	if e := json.Unmarshal(b, &r); e != nil {
		return kevResult{}, fmt.Errorf("Catálogo KEV inválido")
	}
	if r.Version == "" || r.Count <= 0 || r.Count != len(r.Items) {
		return kevResult{}, fmt.Errorf("Catálogo KEV incompleto")
	}
	res := kevResult{Version: r.Version, Count: r.Count}
	for _, v := range r.Items {
		if v.ID == id {
			res.Found = true
			res.Entry = &v
			break
		}
	}
	return res, nil
}
func fetchKEV(ctx context.Context, c *http.Client, id string) (kevResult, error) {
	id = strings.ToUpper(strings.TrimSpace(id))
	if !cvePattern.MatchString(id) {
		return kevResult{}, fmt.Errorf("Informe um identificador CVE válido")
	}
	req, e := http.NewRequestWithContext(ctx, http.MethodGet, kevURL, nil)
	if e != nil {
		return kevResult{}, e
	}
	req.Header.Set("Accept", "application/json")
	b, e := intelBody(c, req)
	if e != nil {
		return kevResult{}, e
	}
	return parseKEV(b, id)
}
