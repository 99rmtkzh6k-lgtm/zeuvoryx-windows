package gamecatalog
import("context";"errors";"net/netip";"net/url";"testing";"time")
func TestReviewGate314(t *testing.T){
 p:=Page{Body:[]byte("<html>benign test</html>"),URL:"https://example.com",TLS:true,DROP:"no_match"}
 good:=ScanEvidence{ClamVersion:"ClamAV test",Signatures:10,ClamFiles:1,LocalFiles:1,LocalBytes:int64(len(p.Body))}
 f:=func(s ScanEvidence,e error)Review{return Analyze(context.Background(),p,func(context.Context,[]byte)(ScanEvidence,error){return s,e})}
 r:=f(good,nil);if !r.CanOpen(time.Now())||r.CanOpen(r.ExpiresAt)||r.CanOpen(r.CheckedAt.Add(-time.Second)){t.Fatal(r)}
 for _,s:=range []ScanEvidence{{}, {ClamHits:1}, {LocalAlerts:1}, {ClamPartial:true}, {LocalPartial:true}, {ClamHits:-1}}{if f(s,nil).CanOpen(time.Now()){t.Fatal("false pass",s)}}
 if f(good,errors.New("motor indisponível")).CanOpen(time.Now()){t.Fatal("error accepted")}
 p.TLS=false;if f(good,nil).CanOpen(time.Now()){t.Fatal("unverified TLS")}
 p.TLS=true;p.DROP="match";if f(good,nil).CanOpen(time.Now()){t.Fatal("DROP accepted")}
}
func TestOrigin314(t *testing.T){
 for _,a:=range []string{"127.0.0.1","10.1.2.3","192.168.0.1","169.254.1.1","100.64.0.1","::1","::ffff:127.0.0.1","2001:db8::1","64:ff9b::a00:1"}{if publicIP(netip.MustParseAddr(a)){t.Fatal(a)}}
 if !publicIP(netip.MustParseAddr("8.8.8.8")){t.Fatal("public blocked")}
 for _,raw:=range []string{"http://example.com","https://other.example","https://example.com:8443","https://a@example.com"}{u,_:=url.Parse(raw);if redirectOK(u,"example.com"){t.Fatal(raw)}}
 u,_:=url.Parse("https://example.com/game");if !redirectOK(u,"example.com"){t.Fatal(u)}
}
