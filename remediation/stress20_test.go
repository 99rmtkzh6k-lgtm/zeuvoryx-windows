package remediation
import("os";"testing";"time")
func TestZStressPolicy5M(t *testing.T){
 if os.Getenv("ZEUVORYX_QA20")!="1"{t.Skip("opt-in workload")}
 start:=time.Now();actions:=[]Action{Observe,Quarantine,StopProcess,CleanPersistence,Repair,Action("unknown")}
 for i:=0;i<5000000;i++{
  b:=i%128; a:=actions[(i/128)%len(actions)]
  e:=Evidence{Verified:b&1!=0,HeuristicOnly:b&2!=0,CriticalOrUnknown:b&4!=0,CurrentIdentityVerified:b&8!=0,UserConsent:b&16!=0,RecoveryReady:b&32!=0,TestArtifact:b&64!=0}
  safe:=b&63==57
  if Allowed(a,e)!=(a==Observe||(a==Quarantine&&safe)){t.Fatalf("unsafe policy case %d",i)}
  if EligibleForProductionMetric(e)!=(safe&&b&64==0){t.Fatalf("metric gate %d",i)}
 }
 t.Logf("QA20 remediation: 5000000 iterations over repeated finite policy combinations; elapsed=%s",time.Since(start))
}
