package adaptive
import("testing";"sync")
func TestActivityEvidence(t *testing.T){
 a:=Status();b:=SelfCheck()
 if !b.SelfTestOK||b.SelfTests!=a.SelfTests+1||b.Observations!=a.Observations{t.Fatal("self-check accounting")}
 var wg sync.WaitGroup;for i:=0;i<8;i++{wg.Add(1);go func(){defer wg.Done();if _,e:=Observe(Vector{1,1,0,0,0});e!=nil{t.Error(e)}}()};wg.Wait()
 _,_ = Observe(Vector{2,0,0,0,0});c:=Status()
 if c.Observations!=a.Observations+9||c.Successful!=a.Successful+8||c.Errors!=a.Errors+1||c.Running!=0||c.LastScore==nil||c.Last==""{t.Fatalf("bad accounting %+v",c)}
}
