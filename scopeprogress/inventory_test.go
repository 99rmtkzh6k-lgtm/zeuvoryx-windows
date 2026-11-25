package scopeprogress
import("context";"os";"path/filepath";"testing")
func TestInventory314(t *testing.T){
 d:=t.TempDir();a:=filepath.Join(d,"a");b:=filepath.Join(d,"b");z:=filepath.Join(d,"zero")
 for p,v:=range map[string]string{a:"abc",b:"12345",z:""}{if e:=os.WriteFile(p,[]byte(v),0600);e!=nil{t.Fatal(e)}}
 r:=Enumerate(context.Background(),[]string{d},func(p string)bool{return p==b},nil)
 if r.Files!=2||r.Bytes!=3||r.Excluded!=1||r.State!="ready"{t.Fatalf("%+v",r)}
 r=Enumerate(context.Background(),[]string{a,filepath.Join(d,"missing")},nil,nil);if r.State!="partial"||r.Errors!=1||r.Files!=1{t.Fatalf("%+v",r)}
 ctx,cancel:=context.WithCancel(context.Background());cancel();if r=Enumerate(ctx,[]string{d},nil,nil);r.State!="cancelled"||r.Files!=0{t.Fatal(r)}
}
