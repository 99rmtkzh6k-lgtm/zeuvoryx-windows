package changeguard
import("errors";"testing";"time")
func TestTracker(t *testing.T){t1:=New([]string{"Downloads"});t1.Enable(true);t1.Active("Downloads",true);t1.Signal(time.Now());t1.ScanRequest(time.Now());s:=t1.Snapshot();if !s.Enabled||s.Signals!=1||len(s.ActiveRoots)!=1||s.LastScanRequest.IsZero(){t.Fatal(s)};s.ActiveRoots[0]="changed";s.Roots[0]="changed";t1.Active("Downloads",false);t1.Problem("Downloads",errors.New("access denied"));s=t1.Snapshot();if len(s.ActiveRoots)!=0||s.Roots[0]!="Downloads"||s.Problems["Downloads"]!="access denied"{t.Fatal(s)};s.Problems["Downloads"]="changed";if t1.Snapshot().Problems["Downloads"]!="access denied"{t.Fatal("mutable snapshot")}}

func TestMonotonicEvidence315(t *testing.T){
 tr:=New(nil);now:=time.Now();tr.Signal(now);tr.Signal(now.Add(-time.Minute));tr.ScanRequest(now);tr.ScanRequest(now.Add(-time.Minute))
 s:=tr.Snapshot();if s.Signals!=2||!s.LastSignal.Equal(now)||!s.LastScanRequest.Equal(now){t.Fatal(s)}
}

func TestMonitorRecoveryAndSaturation315(t *testing.T){
 tr:=New([]string{"Downloads"});tr.Active("Downloads",true);tr.Problem("Downloads",errors.New("temporary"));tr.Active("Downloads",true)
 if _,ok:=tr.Snapshot().Problems["Downloads"];ok{t.Fatal("stale error after recovery")}
 tr.s.Signals=^uint64(0);tr.Signal(time.Now());if tr.Snapshot().Signals!=^uint64(0){t.Fatal("counter wrapped")}
}
