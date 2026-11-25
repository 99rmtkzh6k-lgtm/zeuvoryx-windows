package main
import("os";"path/filepath";"testing";"time")
func Test310SelectedTempCleanup(t *testing.T){
 root:=t.TempDir();now:=time.Now();past:=now.Add(-9*24*time.Hour)
 for _,n:=range []string{"old.tmp","keep.tmp","changed.tmp","fresh.tmp","personal.txt"}{p:=filepath.Join(root,n);if e:=os.WriteFile(p,[]byte("abc"),0600);e!=nil{t.Fatal(e)};if n!="fresh.tmp"{if e:=os.Chtimes(p,past,past);e!=nil{t.Fatal(e)}}}
 p,e:=opt310Preview(root,now);if e!=nil||len(p.Items)!=3{t.Fatalf("preview %v %v",p,e)}
 ids:=map[string]int{};for _,v:=range p.Items{ids[v.Name]=v.ID}
 if _,e=opt310Apply(p,[]int{999},now);e==nil{t.Fatal("bad selection accepted")}
 if _,e=opt310Apply(p,[]int{ids["old.tmp"]},now.Add(6*time.Minute));e==nil{t.Fatal("expired plan accepted")}
 if e=os.WriteFile(filepath.Join(root,"changed.tmp"),[]byte("new content"),0600);e!=nil{t.Fatal(e)}
 r,e:=opt310Apply(p,[]int{ids["old.tmp"],ids["changed.tmp"]},now);if e!=nil||r.Removed!=1||r.Skipped!=1||r.Bytes!=3{t.Fatalf("apply %+v %v",r,e)}
 if _,e=os.Stat(filepath.Join(root,"old.tmp"));!os.IsNotExist(e){t.Fatal("old file not removed",e)}
 for _,n:=range []string{"keep.tmp","changed.tmp","fresh.tmp","personal.txt"}{if _,e=os.Stat(filepath.Join(root,n));e!=nil{t.Fatal(n,e)}}
}
