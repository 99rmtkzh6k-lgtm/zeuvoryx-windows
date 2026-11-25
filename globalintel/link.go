package globalintel
import("errors";"net/netip";"net/url";"strings";"time")
type LinkVerdict struct{Host,State string;AdvisoryOnly bool}
// InspectLink never resolves or visits a host; only explicit IP addresses can be
// checked by DROP. Names require a separate domain reputation source.
func(d Drop) InspectLink(raw string,now time.Time)(LinkVerdict,error){
 result:=LinkVerdict{State:"unverified",AdvisoryOnly:true}
 if len(raw)>4096{return result,errors.New("URL too long")}
 u,err:=url.Parse(strings.TrimSpace(raw));if err!=nil||u==nil||u.Opaque!=""||u.User!=nil||u.Hostname()==""||(u.Scheme!="https"&&u.Scheme!="http"){return result,errors.New("unsupported or invalid URL")}
 result.Host=u.Hostname()
 if _,err=netip.ParseAddr(result.Host);err!=nil{result.State="domain_reputation_unavailable";return result,nil}
 hit,err:=d.Match(result.Host,now);if err!=nil{return result,err}
 if hit{result.State="listed_ip"}else{result.State="not_listed_not_proof_of_safety"};return result,nil
}
