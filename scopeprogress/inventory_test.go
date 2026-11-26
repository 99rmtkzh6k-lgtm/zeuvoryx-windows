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

func TestInventoryOverflow315(t *testing.T){
 const max=int64(^uint64(0)>>1)
 r:=Inventory{Files:1,Bytes:max-1}
 if e:=addFile(&r,1);e!=nil||r.Bytes!=max||r.Files!=2{t.Fatal(r,e)}
 if e:=addFile(&r,1);e==nil||r.Bytes!=max||r.Files!=2{t.Fatal("overflow changed counters",r,e)}
 r=Inventory{Files:max};if e:=addFile(&r,0);e==nil||r.Files!=max{t.Fatal("file counter wrapped",r,e)}
}

func TestP0EmptyDirectoryFinishes(t *testing.T) {
 ctx,cancel:=context.WithCancel(context.Background());defer cancel()
 r:=Enumerate(ctx,[]string{t.TempDir()},nil,nil)
 if r.State!="ready"||r.Files!=0||r.Bytes!=0||r.Errors!=0{t.Fatalf("empty scope: %+v",r)}
 if r.ReadyAt.IsZero(){t.Fatal("completion timestamp missing")}
}

func TestP0InventoryCompleteness322(t *testing.T) {
 cases := []struct{ state string; errors int64; want bool }{
 {"ready",0,false},{"ready",1,true},{"partial",0,true},{"partial",1,true},
 {"cancelled",0,true},{"unavailable",0,true},{"counting",0,true},{"",0,true},
 }
 for _, c := range cases { if got := (Inventory{State:c.state,Errors:c.errors}).Incomplete(); got != c.want { t.Errorf("%+v: got %v",c,got) } }
 r:=Enumerate(context.Background(),[]string{filepath.Join(t.TempDir(),"missing")},nil,nil)
 if !r.Incomplete(){t.Fatalf("missing root classified complete: %+v",r)}
}
