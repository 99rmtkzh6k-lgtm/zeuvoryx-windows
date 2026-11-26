package remediation

import("crypto/sha256";"encoding/hex";"errors";"os";"path/filepath";"testing")
func TestVerifiedOutcome(t *testing.T){
 d:=t.TempDir();original:=filepath.Join(d,"original.txt");payload:=filepath.Join(d,"copy.bin")
 data:=[]byte("harmless remediation regression fixture");h:=sha256.Sum256(data);digest:=hex.EncodeToString(h[:])
 if e:=os.WriteFile(payload,data,0600);e!=nil{t.Fatal(e)}
 o:=Outcome{Path:original,SHA256:digest,Detector:"test fixture",TestArtifact:true}
 if e:=os.WriteFile(original,data,0600);e!=nil{t.Fatal(e)}
 x:=OutcomeForQuarantine(o,payload,nil);if x.Status=="quarantine_verified"||x.OriginalAbsent{t.Fatal(x)}
 if e:=os.Remove(original);e!=nil{t.Fatal(e)}
 x=OutcomeForQuarantine(o,payload,nil);if x.Status!="quarantine_verified"||!x.OriginalAbsent||!x.CopyVerified||x.Executor!="ZEUVORYX"{t.Fatal(x)}
 log:=filepath.Join(d,"events");x,e:=Record(log,x);if e!=nil{t.Fatal(e)};history,e:=ReadHistory(log);if e!=nil||len(history.Events)!=1||history.Events[0].ID!=x.ID{t.Fatal(history,e)}
 if e=os.WriteFile(payload,[]byte("changed"),0600);e!=nil{t.Fatal(e)}
 x=OutcomeForQuarantine(o,payload,nil);if x.Status=="quarantine_verified"||x.CopyVerified{t.Fatal(x)}
 x=OutcomeForQuarantine(o,payload,errors.New("access denied"));if x.Executor!=""||x.Status=="quarantine_verified"{t.Fatal(x)}
}
func TestMissingPathDoesNotProveAction(t *testing.T){d:=t.TempDir();o:=Outcome{Path:filepath.Join(d,"absent"),SHA256:EICARSHA256};x:=OutcomeForQuarantine(o,filepath.Join(d,"missing"),nil);if x.Status=="quarantine_verified"||x.Executor!=""{t.Fatal(x)}}
