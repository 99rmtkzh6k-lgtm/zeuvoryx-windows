package audit
import("os";"path/filepath";"testing";"encoding/json")
func TestAuditHistory(t *testing.T){
 d:=t.TempDir();s,e:=Open(d);if e!=nil{t.Fatal(e)};id,e:=NewID();if e!=nil{t.Fatal(e)}
 if _,e=s.Add(Event{Kind:"scan_started",ScanID:id,Mode:"quick"});e!=nil{t.Fatal(e)}
 if e=s.Finish(id,"quick",map[string]int{"files":42},Event{Files:42,Bytes:120,Partial:true});e!=nil{t.Fatal(e)}
 if _,e=s.Rating(0);e==nil{t.Fatal("bad rating")};if _,e=s.Rating(5);e!=nil{t.Fatal(e)};if _,e=s.Rating(3);e!=nil{t.Fatal(e)}
 s,e=Open(d);if e!=nil{t.Fatal(e)};h,e:=s.History();if e!=nil||h.Finished!=1||h.Files!=42||h.Rating!=3||len(h.Events)!=4{t.Fatalf("%+v %v",h,e)}
 r,e:=s.ReadReport(id);if e!=nil||!json.Valid(r){t.Fatal(e)};if e=s.SaveReport(id,map[string]int{"files":1});e==nil{t.Fatal("overwrite")}
 if _,e=s.ReadReport("../secret");e==nil{t.Fatal("traversal")}
 if e=os.WriteFile(filepath.Join(d,"event-"+h.Events[0].ID+".json"),[]byte("{}"),0600);e!=nil{t.Fatal(e)}
 h,e=s.History();if e!=nil||h.Invalid!=1{t.Fatalf("%+v %v",h,e)}
}
func TestWrongKey(t *testing.T){a,e:=Open(t.TempDir());if e!=nil{t.Fatal(e)};b,e:=Open(t.TempDir());if e!=nil{t.Fatal(e)};v,e:=a.seal([]byte("{\"x\": 1}"));if e!=nil{t.Fatal(e)};if _,e=b.verify(v);e==nil{t.Fatal("wrong key accepted")};if _,e=a.verify(v);e!=nil{t.Fatal(e)}}
