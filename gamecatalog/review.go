package gamecatalog
import(
 "context";"crypto/sha256";"crypto/tls";"encoding/hex";"fmt";"io";"mime";"net";"net/http";"net/netip";"net/url";"strings";"time"
)
const MaxPage=2<<20
const Lifetime=5*time.Minute
type Review struct {
 Status string `json:"status"`
 CheckedAt time.Time `json:"checked_at"`
 ExpiresAt time.Time `json:"expires_at"`
 URL string `json:"url"`
 SHA256 string `json:"sha256,omitempty"`
 Bytes int `json:"bytes"`
 TLS bool `json:"tls"`
 DROP string `json:"drop"`
 ClamVersion string `json:"clam_version,omitempty"`
 Signatures int64 `json:"signatures"`
 LocalAlerts int `json:"local_alerts"`
 Problem string `json:"problem,omitempty"`
}
type Page struct{Body []byte; URL string; TLS bool; DROP string}
type Scanner func(context.Context,[]byte)(ScanEvidence,error)
type ScanEvidence struct{ClamVersion string; Signatures int64; ClamFiles int64; ClamHits int64; ClamPartial bool; LocalFiles int; LocalBytes int64; LocalAlerts int; LocalPartial bool}
func (r Review) CanOpen(now time.Time)bool{
 return r.Status=="no_detection_sample"&&r.TLS&&r.Bytes>0&&r.Signatures>0&&r.ClamVersion!=""&&r.LocalAlerts==0&&r.SHA256!=""&&!r.CheckedAt.IsZero()&&!now.Before(r.CheckedAt)&&now.Before(r.ExpiresAt)
}
func Analyze(ctx context.Context,p Page,scan Scanner) (r Review) {
 r=Review{Status:"inconclusive",URL:p.URL,TLS:p.TLS,DROP:p.DROP,Bytes:len(p.Body)}
 defer func(){r.CheckedAt=time.Now().UTC();if r.Status=="no_detection_sample"{r.ExpiresAt=r.CheckedAt.Add(Lifetime)}}()
 if !p.TLS||len(p.Body)==0||len(p.Body)>MaxPage||p.DROP=="match"{r.Problem="Página fora dos critérios de revisão";return}
 sum:=sha256.Sum256(p.Body);r.SHA256=hex.EncodeToString(sum[:])
 s,e:=scan(ctx,p.Body);r.ClamVersion=s.ClamVersion;r.Signatures=s.Signatures;r.LocalAlerts=s.LocalAlerts
 if s.ClamHits>0||s.LocalAlerts>0{r.Status="blocked";r.Problem="Detecção ou indício de risco; abertura não autorizada";return}
 if e!=nil{r.Problem=e.Error();return}
 if ctx.Err()!=nil||s.ClamPartial||s.LocalPartial||s.ClamFiles!=1||s.LocalFiles!=1||s.LocalBytes!=int64(len(p.Body))||s.Signatures<=0||s.ClamVersion==""||s.ClamHits!=0{r.Problem="Motores não confirmaram cobertura do conteúdo";return}
 r.Status="no_detection_sample";return
}
var denied=[]netip.Prefix{
 netip.MustParsePrefix("0.0.0.0/8"),netip.MustParsePrefix("100.64.0.0/10"),netip.MustParsePrefix("192.0.0.0/24"),netip.MustParsePrefix("192.0.2.0/24"),netip.MustParsePrefix("198.18.0.0/15"),netip.MustParsePrefix("198.51.100.0/24"),netip.MustParsePrefix("203.0.113.0/24"),netip.MustParsePrefix("240.0.0.0/4"),netip.MustParsePrefix("2001:db8::/32"),netip.MustParsePrefix("2001::/23"),netip.MustParsePrefix("2002::/16"),netip.MustParsePrefix("64:ff9b::/96"),netip.MustParsePrefix("64:ff9b:1::/48"),
}
func publicIP(a netip.Addr)bool{a=a.Unmap();if !a.IsGlobalUnicast()||a.IsPrivate()||a.IsLoopback()||a.IsLinkLocalUnicast()||a.Zone()!=""{return false};for _,p:=range denied{if p.Contains(a){return false}};return true}
func redirectOK(u *url.URL,host string)bool{return u.Scheme=="https"&&strings.EqualFold(u.Hostname(),host)&&u.Port()==""&&u.User==nil}
type Reputation func(string)(bool,error)
// Fetch reads only the initial HTML document, not scripts, frames, ads, or browser downloads.
// No cookies, proxy, user credentials or arbitrary URLs. DNS addresses are checked and pinned for each connection.
func Fetch(ctx context.Context,id string,reputation Reputation)(Page,error){
 var p Page
 raw,e:=Resolve(id);if e!=nil{return p,e};u,e:=url.Parse(raw);if e!=nil{return p,e};host:=u.Hostname()
 p.URL=raw;p.DROP="unavailable"
 if reputation!=nil{p.DROP="no_match"}
 tr:=&http.Transport{TLSClientConfig:&tls.Config{MinVersion:tls.VersionTLS12},TLSHandshakeTimeout:8*time.Second,ResponseHeaderTimeout:10*time.Second,DisableKeepAlives:true,MaxResponseHeaderBytes:128<<10,DialContext:func(ctx context.Context,network,address string)(net.Conn,error){
 h,port,e:=net.SplitHostPort(address);if e!=nil||!strings.EqualFold(h,host)||port!="443"{return nil,fmt.Errorf("Destino fora da origem aprovada")}
 ips,e:=net.DefaultResolver.LookupNetIP(ctx,"ip",h);if e!=nil||len(ips)==0||len(ips)>32{return nil,fmt.Errorf("DNS não confirmado")}
 for _,ip:=range ips{
 if !publicIP(ip){return nil,fmt.Errorf("DNS aponta para rede não pública")}
 if reputation!=nil{hit,err:=reputation(ip.Unmap().String());if hit{p.DROP="match";return nil,fmt.Errorf("IP consta na base DROP")};if err!=nil{p.DROP="unavailable"}}
 }
 d:=net.Dialer{Timeout:6*time.Second}
 for _,ip:=range ips{c,err:=d.DialContext(ctx,"tcp",net.JoinHostPort(ip.String(),port));if err==nil{return c,nil}}
 return nil,fmt.Errorf("Conexão com a origem não confirmada")
 }}
 defer tr.CloseIdleConnections()
 c:=&http.Client{Transport:tr,Timeout:25*time.Second,CheckRedirect:func(req *http.Request,via []*http.Request)error{if len(via)>=4||!redirectOK(req.URL,host){return fmt.Errorf("Redirecionamento fora da origem aprovada")};return nil}}
 req,e:=http.NewRequestWithContext(ctx,http.MethodGet,raw,nil);if e!=nil{return p,e}
 req.Header.Set("Accept","text/html, application/xhtml+xml");req.Header.Set("User-Agent","ZEUVORYX-LAB/3.0.14 catalog-review")
 res,e:=c.Do(req);if e!=nil{return p,e};defer res.Body.Close()
 p.TLS=res.TLS!=nil&&len(res.TLS.VerifiedChains)>0;p.URL=res.Request.URL.String()
 if !p.TLS||res.StatusCode!=200{return p,fmt.Errorf("Origem não confirmou HTML HTTPS (HTTP %d)",res.StatusCode)}
 mt,_,e:=mime.ParseMediaType(res.Header.Get("Content-Type"));if e!=nil||(mt!="text/html"&&mt!="application/xhtml+xml"){return p,fmt.Errorf("Tipo de conteúdo não aprovado")}
 if res.ContentLength>MaxPage{return p,fmt.Errorf("Página excede limite de análise")}
 b,e:=io.ReadAll(io.LimitReader(res.Body,MaxPage+1));if e!=nil{return p,e};if len(b)==0||len(b)>MaxPage{return p,fmt.Errorf("Tamanho de página fora do limite")}
 p.Body=b;return p,nil
}
