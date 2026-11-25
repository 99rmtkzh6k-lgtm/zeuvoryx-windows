package globalintel
import("testing";"time")
func TestInspectLinkLab(t *testing.T){
 now:=time.Unix(1800000000,0);v,e:=parseDrop([]byte(fixture("192.0.2.0/24",now.Unix())),false,now);if e!=nil{t.Fatal(e)}
 d:=Drop{V4:v}
 for _,tc:=range []struct{u,want string}{{"https://192.0.2.1/a","listed_ip"},{"https://198.51.100.1/","not_listed_not_proof_of_safety"},{"https://example.org/","domain_reputation_unavailable"}}{
  got,e:=d.InspectLink(tc.u,now);if e!=nil||got.State!=tc.want||!got.AdvisoryOnly{t.Fatal(got,e)}
 }
 for _,u:=range []string{"file:///test","https://name@example.org/","https://"}{if _,e:=d.InspectLink(u,now);e==nil{t.Fatal("accepted invalid URL",u)}}
 if _,e:=d.InspectLink("https://192.0.2.1",now.Add(48*time.Hour));e==nil{t.Fatal("stale feed accepted")}
}
