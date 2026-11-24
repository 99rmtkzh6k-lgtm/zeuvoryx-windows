package portable

import (
	"bufio"
	"fmt"
	"net/url"
	"strings"
	"unicode"
)

// Exact URLs only. Never opens links or sends them outside the device.
func validReputationURL(s string) bool {
	if len(s) == 0 || len(s) > 8192 || strings.Contains(s, "\\") || strings.IndexFunc(s, unicode.IsSpace) >= 0 || strings.IndexFunc(s, unicode.IsControl) >= 0 {
		return false
	}
	u, e := url.Parse(s)
	return e == nil && (u.Scheme == "http" || u.Scheme == "https") && u.Hostname() != "" && u.User == nil && u.Opaque == "" && u.Fragment == ""
}
func canonicalReputationURL(s string) string {
	s = strings.TrimSpace(s)
	if strings.Contains(s, "\\") || strings.IndexFunc(s, unicode.IsSpace) >= 0 || strings.IndexFunc(s, unicode.IsControl) >= 0 {
		return ""
	}
	u, e := url.Parse(s)
	if e != nil {
		return ""
	}
	u.Fragment = ""
	u.RawFragment = ""
	u.Host = strings.ToLower(u.Host)
	if u.Path == "" {
		u.Path = "/"
	}
	v := u.String()
	if !validReputationURL(v) {
		return ""
	}
	return v
}
func parseReputationList(raw string) (map[string]bool, error) {
	if len(raw) > maxIntelBytes {
		return nil, fmt.Errorf("Lista acima do limite de 8 MiB")
	}
	out := map[string]bool{}
	s := bufio.NewScanner(strings.NewReader(raw))
	s.Buffer(make([]byte, 4096), 16384)
	for s.Scan() {
		v := strings.TrimSpace(s.Text())
		if v == "" || strings.HasPrefix(v, "#") {
			continue
		}
		k := canonicalReputationURL(v)
		if k == "" {
			return nil, fmt.Errorf("Lista contém URL inválida")
		}
		out[k] = true
		if len(out) > maxReputationURLs {
			return nil, fmt.Errorf("Limite de URLs atingido")
		}
	}
	if e := s.Err(); e != nil {
		return nil, e
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("Lista vazia")
	}
	return out, nil
}
func assessWithReputation(raw string, urls map[string]bool) (LinkAssessment, error) {
	r, e := assessLink(raw)
	if e != nil {
		return r, e
	}
	if urls[canonicalReputationURL(raw)] {
		r.Signals = append(r.Signals, "Correspondência exata na lista local importada; confirme a procedência e não abra o link")
		r.Assessment = "Atenção: URL presente na base de reputação importada"
	}
	return r, nil
}
