package labcheck
import "testing"
func TestIntegratedFoundationDiagnostics(t *testing.T){
 r,e:=Run()
 if e!=nil||!r.Update||!r.Statistics||!r.Policy{t.Fatalf("diagnostic failed: %+v %v",r,e)}
}
