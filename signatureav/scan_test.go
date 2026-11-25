package signatureav
import ("context";"os";"path/filepath";"strings";"testing";"time")
func TestReport(t *testing.T){
 cases:=[]struct{raw string;code int;partial bool;kind string}{
 {"----------- SCAN SUMMARY -----------\nKnown viruses: 123\nScanned files: 1\nInfected files: 0\n",0,false,""},
 {"C:\\a.exe: Win.Test FOUND\n----------- SCAN SUMMARY -----------\nKnown viruses: 123\nScanned files: 1\nInfected files: 1\n",1,false,"signature_detection"},
 {"C:\\a.zip: Heuristics.Limits.Exceeded FOUND\n----------- SCAN SUMMARY -----------\nKnown viruses: 123\nScanned files: 1\nInfected files: 1\n",1,true,"heuristic_or_limit"},
 {"ERROR: DB missing",2,true,""},
 {"----------- SCAN SUMMARY -----------\nKnown viruses: 123\nScanned files: 0\nInfected files: 0\n",0,true,""},
 {"----------- SCAN SUMMARY -----------\nKnown viruses: 123\nScanned files: 2\nInfected files: 1\n",1,true,""},
 }
 for _,c:=range cases{r:=Report{};parse(&r,c.raw,c.code,false);if r.Partial!=c.partial{t.Fatalf("%+v",r)};if c.kind!=""&&(len(r.Findings)!=1||r.Findings[0].Kind!=c.kind){t.Fatalf("%+v",r)}}
 r:=Report{};parse(&r,cases[0].raw,0,true);if !r.Partial{t.Fatal("truncated output accepted")}
}
func TestRealClamAV(t *testing.T){
 exe:=os.Getenv("CLAMAV_REAL_EXECUTABLE");if exe==""{t.Skip("real ClamAV executable not configured")}
 ctx,cancel:=context.WithTimeout(context.Background(),3*time.Minute);defer cancel()
 root:=t.TempDir()
 clean:=filepath.Join(root,"clean.txt");if e:=os.WriteFile(clean,[]byte("ZEUVORYX harmless content"),0600);e!=nil{t.Fatal(e)}
 r,e:=Scan(ctx,exe,clean,"");if e!=nil||r.Status!="no_detection_reported"{t.Fatalf("clean: %+v %v",r,e)}
 // Official EICAR harmless test fixture, not a live virus.
 fixture:="X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"
 path:=filepath.Join(root,"test.txt");if e=os.WriteFile(path,[]byte(fixture),0600);e!=nil{t.Fatal(e)}
 r,e=Scan(ctx,exe,root,"");if e!=nil||r.Status!="detections_reported"||r.Scanned<2||len(r.Findings)!=1||r.Findings[0].Kind!="test_eicar"||!strings.Contains(strings.ToLower(r.Findings[0].Signature),"eicar"){t.Fatalf("actual engine: %+v %v",r,e)}
 if b,e:=os.ReadFile(path);e!=nil||string(b)!=fixture{t.Fatal("scanner altered fixture")}
 t.Logf("ClamAV actual engine: %s, database signatures: %d, detection: %s",r.Version,r.KnownSignatures,r.Findings[0].Signature)
}
