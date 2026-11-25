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
 r,e:=Scan(ctx,exe,clean,os.Getenv("CLAMAV_REAL_DB"));if e!=nil||r.Status!="no_detection_reported"{t.Fatalf("clean: %+v %v",r,e)}
 // Official EICAR harmless test fixture, not a live virus.
 fixture:="X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"
 path:=filepath.Join(root,"test.txt");if e=os.WriteFile(path,[]byte(fixture),0600);e!=nil{t.Fatal(e)}
 r,e=Scan(ctx,exe,root,os.Getenv("CLAMAV_REAL_DB"));if e!=nil||r.Status!="detections_reported"||r.Scanned<2||len(r.Findings)!=1||r.Findings[0].Kind!="test_eicar"||!strings.Contains(strings.ToLower(r.Findings[0].Signature),"eicar"){t.Fatalf("actual engine: %+v %v",r,e)}
 if b,e:=os.ReadFile(path);e!=nil||string(b)!=fixture{t.Fatal("scanner altered fixture")}
 t.Logf("ClamAV actual engine: %s, database signatures: %d, detection: %s",r.Version,r.KnownSignatures,r.Findings[0].Signature)
}

func TestDatabaseAge(t *testing.T){
 now:=time.Date(2026,11,25,12,0,0,0,time.UTC)
 if !recentDatabase("ClamAV 1.5.4/28194/Wed Nov 25 06:26:59 2026",now){t.Fatal("fresh database rejected")}
 for _,v:=range []string{"ClamAV 1.5.4","ClamAV 1.5.4/0/Wed Nov 25 06:26:59 2026","ClamAV 1.5.4/1/Wed Nov 11 06:26:59 2026","ClamAV 1.5.4/1/Mon Dec  7 06:26:59 2026"}{if recentDatabase(v,now){t.Fatal("invalid date accepted",v)}}
}

func TestRealClamAVBatch(t *testing.T){
 exe:=os.Getenv("CLAMAV_REAL_EXECUTABLE");if exe==""{t.Skip("real ClamAV executable not configured")}
 ctx,cancel:=context.WithTimeout(context.Background(),3*time.Minute);defer cancel()
 a,b:=filepath.Join(t.TempDir(),"a.txt"),filepath.Join(t.TempDir(),"b.txt")
 for _,p:=range []string{a,b}{if e:=os.WriteFile(p,[]byte("ordinary harmless content"),0600);e!=nil{t.Fatal(e)}}
 r,e:=ScanRoots(ctx,exe,[]string{a,b,a},os.Getenv("CLAMAV_REAL_DB"))
 if e!=nil||r.Partial||r.Scanned!=2||len(r.Roots)!=2{t.Fatalf("batch %+v %v",r,e)}
 r,e=ScanRoots(ctx,exe,[]string{a,filepath.Join(t.TempDir(),"missing.txt")},os.Getenv("CLAMAV_REAL_DB"))
 if e==nil||!r.Partial||len(r.UnavailableRoots)!=1||r.Scanned!=1{t.Fatalf("missing root %+v %v",r,e)}
}

func TestExplicitDatabaseArguments(t *testing.T){
 db:=filepath.Join(t.TempDir(),"database with spaces")
 args,e:=databaseArgs(db);if e!=nil||len(args)!=1||args[0]!="--database="+db{t.Fatalf("args=%v error=%v",args,e)}
 if _,e=databaseArgs("relative-db");e==nil{t.Fatal("relative path accepted")}
 if _,e=VersionAt(context.Background(),filepath.Join(t.TempDir(),"missing-exe"),"relative-db");e==nil||!strings.Contains(e.Error(),"absolute ClamAV DB"){t.Fatalf("wrong validation: %v",e)}
 if BundledDatabase(filepath.Join(filepath.Dir(db),"clamscan.exe"))!=filepath.Join(filepath.Dir(db),"database"){t.Fatal("wrong bundled path")}
}

func TestOfflineDatabaseStatus(t *testing.T){
 now:=time.Date(2026,11,25,12,0,0,0,time.UTC)
 for _,c:=range []struct{v,want string}{
 {"ClamAV 1.5.4/28194/Wed Nov 25 06:26:59 2026","fresh"},
 {"ClamAV 1.5.4/1/Wed Nov 11 06:26:59 2026","stale"},
 {"ClamAV 1.5.4","unknown"},
 {"ClamAV 1.5.4/1/Mon Dec  7 06:26:59 2026","unknown"},
 }{if got:=DatabaseStatus(c.v,now);got!=c.want{t.Fatalf("%q: %s != %s",c.v,got,c.want)}}
}
