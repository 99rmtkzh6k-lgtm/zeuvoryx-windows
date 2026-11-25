package engine
import("testing";"context";"os";"path/filepath";"github.com/zeuvoryx/zeuvoryx/adaptive")
func Test308ActualAIOnCleanScan(t *testing.T){
 dir:=t.TempDir()
 for _,name:=range []string{"one.txt","two.txt","three.txt"}{if e:=os.WriteFile(filepath.Join(dir,name),[]byte("ordinary document "+name),0600);e!=nil{t.Fatal(e)}}
 a:=adaptive.Status();r,e:=New().Scan(context.Background(),dir,nil)
 if e!=nil||r.Files!=3||len(r.Findings)!=0{t.Fatalf("scan: %+v %v",r,e)}
 b:=adaptive.Status()
 if b.Observations-a.Observations!=3||b.Successful-a.Successful!=3||b.Errors!=a.Errors||b.LastScore==nil{t.Fatalf("model did not evaluate clean files: %+v -> %+v",a,b)}
}
