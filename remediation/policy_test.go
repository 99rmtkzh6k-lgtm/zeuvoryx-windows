package remediation
import "testing"
func TestFailClosed(t *testing.T){
 good:=Evidence{Verified:true,CurrentIdentityVerified:true,UserConsent:true,RecoveryReady:true}
 if !Allowed(Quarantine,good){t.Fatal("verified reversible action rejected")}
 for _,a:=range []Action{StopProcess,CleanPersistence,Repair,"delete","unknown"}{if Allowed(a,good){t.Fatal("unimplemented action allowed",a)}}
 for mask:=0;mask<64;mask++{e:=Evidence{Verified:mask&1!=0,HeuristicOnly:mask&2!=0,CriticalOrUnknown:mask&4!=0,CurrentIdentityVerified:mask&8!=0,UserConsent:mask&16!=0,RecoveryReady:mask&32!=0};want:=mask==57;if Allowed(Quarantine,e)!=want{t.Fatal(mask)}}
}

func TestTestArtifactsNeverCount(t *testing.T){
 e:=Evidence{Verified:true,CurrentIdentityVerified:true,UserConsent:true,RecoveryReady:true}
 if !EligibleForProductionMetric(e){t.Fatal("valid prerequisite rejected")}
 e.TestArtifact=true
 if EligibleForProductionMetric(e){t.Fatal("test artifact counted")}
 e.TestArtifact=false;e.HeuristicOnly=true
 if EligibleForProductionMetric(e){t.Fatal("heuristic counted")}
}
