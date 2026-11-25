package main
import("testing";"time";"context";"os";"path/filepath";"github.com/zeuvoryx/zeuvoryx/engine";"github.com/zeuvoryx/zeuvoryx/signatureav")
func Test309LiveReportSnapshots(t *testing.T){
 s:=&State{running:true,report:engine.Report{ScanID:"test",ScanMode:"full",Started:time.Now()}}
 raw:=engine.Report{Root:"test",Started:time.Now(),Files:1,Bytes:10,Findings:[]engine.Finding{{Path:"first",Risk:100}},Errors:[]string{"denied"},ErrorCount:1,Skipped:1,ErrorCategories:map[string]int{"access_denied":1}}
 s.publishLive(raw)
 a,n:=s.reportSnapshotLocked(true)
 if !a.InProgress||!a.Partial||!a.Finished.IsZero()||n!=1||a.ErrorCount!=1||a.ScanID!="test"{t.Fatalf("invalid partial %+v",a)}
 raw.ErrorCategories["access_denied"]=2;raw.Findings=append(raw.Findings,engine.Finding{Path:"second"})
 raw.Files=2;raw.Bytes=20;s.publishLive(raw)
 b,ok:=s.snapshotForExport();if !ok||!b.InProgress||!b.Partial||!b.Finished.IsZero()||len(b.Findings)!=2||b.Files!=2{t.Fatalf("export %+v",b)}
 if len(a.Findings)!=1||a.ErrorCategories["access_denied"]!=1{t.Fatal("snapshot mutated")}
 for i:=0;i<600;i++{raw.Findings=append(raw.Findings,engine.Finding{Path:"test"})};s.publishLive(raw)
 c,total:=s.reportSnapshotLocked(true);d,_:=s.snapshotForExport()
 if len(c.Findings)!=500||total!=602||len(d.Findings)!=602{t.Fatal("display cap lost full report")}
 s.running=false;s.publishLive(engine.Report{Files:1000});if s.report.Files!=2{t.Fatal("late publication")}
}
func Test309RootLiveBeforeCompletion(t *testing.T){
 dir:=t.TempDir();q:=filepath.Join(t.TempDir(),"q");if e:=os.Mkdir(q,0700);e!=nil{t.Fatal(e)}
 data:=[]byte("X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*")
 if e:=os.WriteFile(filepath.Join(dir,"test.txt"),data,0600);e!=nil{t.Fatal(e)}
 s:=&State{running:true,scanner:engine.New(),quarantine:&engine.Quarantine{Dir:q},report:engine.Report{Started:time.Now()}}
 observed:=false;r:=s.scanRoots(context.Background(),[]string{dir},"",func(n int,b int64){if n>0{v,_:=s.snapshotForExport();if v.InProgress&&len(v.Findings)==1{observed=true}}})
 if !observed||len(r.Findings)!=1||r.Files!=1{t.Fatalf("no live alert before end: %v %+v",observed,r)}
}
func Test309CustomLiveBeforeCompletion(t *testing.T){
 dir:=t.TempDir();if e:=os.WriteFile(filepath.Join(dir,"plain.txt"),[]byte("ordinary file"),0600);e!=nil{t.Fatal(e)}
 s:=&State{running:true,report:engine.Report{Started:time.Now()}}
 r,e:=engine.New().Scan(engine.WithLiveReport(context.Background(),s.publishLive),dir,nil)
 v,_:=s.snapshotForExport()
 if e!=nil||r.Files!=1||v.Files!=1||!v.InProgress{t.Fatalf("custom live %+v %+v %v",r,v,e)}
}

func Test314SignatureReportSurvivesScannerSnapshot(t *testing.T){
 s:=&State{running:true,report:engine.Report{ScanID:"test314",Started:time.Now()}}
 roots:=[]string{"root-a","root-b"}
 rep:=engine.Report{SignatureState:"running",SignatureRoots:roots,SignatureReports:[]signatureav.Report{{Root:"root-a",Partial:true,Scanned:-1,Findings:[]signatureav.Finding{{Kind:"heuristic_or_limit",Signature:"Heuristics.Limits.Exceeded.MaxScanTime"}}}}}
 s.publishSignatureProgress314(&rep)
 s.publishLive(engine.Report{Files:4,Bytes:10})
 snap,ok:=s.snapshotForExport()
 if !ok||snap.SignatureState!="running"||len(snap.SignatureReports)!=1||snap.SignatureReports[0].Scanned!=-1||!snap.Partial{t.Fatalf("lost signature progress: %+v",snap)}
 rep.SignatureReports=append(rep.SignatureReports,signatureav.Report{Root:"root-b",Scanned:2})
 rep.SignatureState="partial";s.publishSignatureProgress314(&rep)
 after,_:=s.snapshotForExport()
 if len(after.SignatureReports)!=2||after.SignatureState!="partial"||len(snap.SignatureReports)!=1{t.Fatal("snapshot changed or missing second root")}
 s.running=false;rep.SignatureState="finished";s.publishSignatureProgress314(&rep)
 if s.report.SignatureState!="partial"{t.Fatal("late report changed finished session")}
}
