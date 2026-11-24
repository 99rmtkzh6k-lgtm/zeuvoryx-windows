package scanner
import("context";"os";"path/filepath";"testing")
func TestReadOnlyScan(t *testing.T){
 d:=t.TempDir();p:=filepath.Join(d,"marker.txt")
 if e:=os.WriteFile(p,marker,0600);e!=nil{t.Fatal(e)}
 if e:=os.WriteFile(filepath.Join(d,"note.txt"),[]byte("harmless"),0600);e!=nil{t.Fatal(e)}
 if e:=os.Symlink(p,filepath.Join(d,"link"));e!=nil{t.Fatal(e)}
 r,e:=Scan(context.Background(),d)
 if e!=nil||r.Files!=2||r.Skipped!=1||!r.Partial||len(r.Findings)!=1||r.Findings[0].Kind!="test_eicar"{t.Fatal(r,e)}
 if b,e:=os.ReadFile(p);e!=nil||string(b)!=string(marker){t.Fatal("file modified")}
}
func TestAdvisoryAndCancel(t *testing.T){
 d:=t.TempDir();p:=filepath.Join(d,"invoice.pdf.command")
 if e:=os.WriteFile(p,[]byte("inert"),0600);e!=nil{t.Fatal(e)}
 r,e:=Scan(context.Background(),p);if e!=nil||len(r.Findings)!=1||r.Findings[0].Kind!="advisory_double_extension"{t.Fatal(r,e)}
 ctx,cancel:=context.WithCancel(context.Background());cancel()
 r,e=Scan(ctx,d);if e!=nil||!r.Partial||r.Files!=0{t.Fatal(r,e)}
}
