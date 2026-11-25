package globalintel
import("testing";"path/filepath";"time";"context";"fmt")
func TestDailyCache314(t *testing.T){
 p:=filepath.Join(t.TempDir(),"drop.json");d:=OpenDaily(p);now:=time.Now();d.state.Attempt=now
 d.state.V4=[]byte(fmt.Sprintf("{\"cidr\":\"8.8.8.0/24\"}\n{\"timestamp\":%d,\"copyright\":\"Spamhaus test fixture\"}\n",now.Unix()))
 d.state.V6=[]byte(fmt.Sprintf("{\"cidr\":\"2001:4860::/32\"}\n{\"timestamp\":%d,\"copyright\":\"Spamhaus test fixture\"}\n",now.Unix()))
 if e:=d.persist();e!=nil{t.Fatal(e)}
 r:=OpenDaily(p);if r.Status().IPv4!=1||r.Status().IPv6!=1||r.Status().Stale{t.Fatal(r.Status())}
 yes,e:=r.Match("8.8.8.8",now);if e!=nil||!yes{t.Fatal(yes,e)}
 if _,e=r.Refresh(context.Background());e==nil{t.Fatal("refresh interval not enforced")}
}
