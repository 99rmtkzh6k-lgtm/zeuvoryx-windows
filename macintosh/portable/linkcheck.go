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
	if hasRedirectParameter(u.RawQuery) {
		add("Parâmetro de possível redirecionamento")
	}
	if imitatesWikiLeaks(host) {
		add("Domínio semelhante a wikileaks.org: confirme o endereço oficial; não comprova fraude")
	}
	if len(r.Signals) > 0 {
		r.Assessment = "Atenção: indicadores estruturais; não comprovam fraude"
	}
	return r, nil
}

// Bounded decoding of query keys only. Never visits the destination.
func hasRedirectParameter(raw string) bool {
 for round:=0;round<3;round++ {
  for _,part:=range strings.FieldsFunc(raw,func(r rune)bool{return r=='&'||r==';'}) {
   key,_,ok:=strings.Cut(part,"=");if !ok{continue}
   key,e:=url.QueryUnescape(key);if e!=nil{continue}
   switch strings.ToLower(key) {
   case "redirect","url","redirect_uri","redirect_url","return_url","returnurl","next","continue","target","destination","dest":return true
   }
  }
  decoded,e:=url.QueryUnescape(raw);if e!=nil||decoded==raw{break};raw=decoded
 }
 return false
}
// Advisory, deliberately limited to the identity requested for this regression.
// Only DNS labels are examined, not ordinary text or page paths.
func imitatesWikiLeaks(host string) bool {
 const official="wikileaks.org"
 if host==official||strings.HasSuffix(host,"."+official){return false}
 for _,label:=range strings.Split(host,".") {
  s:=strings.ReplaceAll(label,"-","")
  if strings.Contains(s,"wikileaks")||nearWikiLabel(s){return true}
 }
 return false
}
func nearWikiLabel(s string) bool {
 const brand="wikileaks"
 if len(s)<len(brand)-1||len(s)>len(brand)+1{return false}
 // Bounded edit distance (one insertion, deletion or substitution), including a
 // single adjacent transposition. ASCII only; other scripts have a separate signal.
 for _,r:=range s{if !(r>='a'&&r<='z'||r>='0'&&r<='9'){return false}}
 if len(s)==len(brand){
  diffs:=[]int{};for i:=range brand{if s[i]!=brand[i]{diffs=append(diffs,i)}}
  return len(diffs)<=1||(len(diffs)==2&&diffs[1]==diffs[0]+1&&s[diffs[0]]==brand[diffs[1]]&&s[diffs[1]]==brand[diffs[0]])
 }
 a,b:=s,brand;if len(a)>len(b){a,b=b,a}
 for i:=0;i<len(b);i++{if a==b[:i]+b[i+1:]{return true}}
 return false
}

// AssessLink performs offline structural advisory analysis only.
func AssessLink(raw string) (LinkAssessment, error) { return assessLink(raw) }
