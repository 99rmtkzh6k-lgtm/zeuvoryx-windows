package portable

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"unicode"
)

type LinkAssessment struct {
	Host       string   `json:"host"`
	Signals    []string `json:"signals"`
	Assessment string   `json:"assessment"`
}

// Local structural triage only. Never resolves, fetches, or opens the URL.
func assessLink(raw string) (LinkAssessment, error) {
	r := LinkAssessment{Signals: []string{}, Assessment: "Inconclusivo: nenhum indicador estrutural identificado"}
	raw = strings.TrimSpace(raw)
	if len(raw) == 0 || len(raw) > 8192 {
		return r, fmt.Errorf("Informe uma URL HTTP ou HTTPS de até 8192 bytes")
	}
	if strings.Contains(raw, "\\") || strings.IndexFunc(raw, unicode.IsControl) >= 0 || strings.IndexFunc(raw, unicode.IsSpace) >= 0 {
		return r, fmt.Errorf("URL ambígua ou inválida")
	}
	u, e := url.Parse(raw)
	if e != nil || u.Hostname() == "" || u.Opaque != "" || (u.Scheme != "http" && u.Scheme != "https") {
		return r, fmt.Errorf("Informe um endereço HTTP ou HTTPS válido")
	}
	host := strings.TrimSuffix(strings.ToLower(u.Hostname()), ".")
	r.Host = host
	add := func(s string) { r.Signals = append(r.Signals, s) }
	if u.Scheme == "http" {
		add("HTTP sem criptografia: não envie senhas ou dados bancários")
	}
	if u.User != nil {
		add("Identificação antes de @ pode disfarçar o domínio real")
	}
	if net.ParseIP(host) != nil {
		add("Endereço IP em vez de nome de domínio")
	}
	if strings.Contains(host, "xn--") || strings.IndexFunc(host, func(r rune) bool { return r > 127 }) >= 0 {
		add("Domínio internacionalizado: verifique possíveis caracteres parecidos")
	}
	if len(strings.Split(host, ".")) >= 5 {
		add("Muitos subdomínios: confira o domínio registrado")
	}
	if p := u.Port(); p != "" && p != "80" && p != "443" {
		add("Porta não convencional")
	}
	if strings.Contains(strings.ToLower(u.RawQuery), "redirect=") || strings.Contains(strings.ToLower(u.RawQuery), "url=") {
		add("Parâmetro de possível redirecionamento")
	}
	if len(r.Signals) > 0 {
		r.Assessment = "Atenção: indicadores estruturais; não comprovam fraude"
	}
	return r, nil
}
