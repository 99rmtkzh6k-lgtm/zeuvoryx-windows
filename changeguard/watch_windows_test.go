package changeguard
import("context";"os";"path/filepath";"testing";"time")
func TestNativeDirectoryNotification(t *testing.T){
 root:=t.TempDir();tr:=New([]string{root});ctx,cancel:=context.WithCancel(context.Background());defer cancel();done:=make(chan error,1);go func(){done<-Watch(ctx,root,tr)}()
 deadline:=time.Now().Add(5*time.Second)
 for len(tr.Snapshot().ActiveRoots)==0&&time.Now().Before(deadline){select{case e:=<-done:t.Fatalf("watch failed: %v",e);case <-time.After(20*time.Millisecond):}}
 if len(tr.Snapshot().ActiveRoots)!=1{t.Fatal("watch not ready")}
 if e:=os.WriteFile(filepath.Join(root,"benign.txt"),[]byte("harmless"),0600);e!=nil{t.Fatal(e)}
 deadline=time.Now().Add(5*time.Second);for tr.Snapshot().Signals==0&&time.Now().Before(deadline){time.Sleep(20*time.Millisecond)}
 if tr.Snapshot().Signals==0{t.Fatal("no notification for real file creation")}
 cancel();select{case e:=<-done:if e!=nil{t.Fatal(e)};case <-time.After(3*time.Second):t.Fatal("watch cancellation failed")}
}
