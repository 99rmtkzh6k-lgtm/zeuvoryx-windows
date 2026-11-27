//go:build windows && build48gate

package main

import (
 "bytes"
 "context"
 "os"
 "path/filepath"
 "testing"
 "time"
 "github.com/zeuvoryx/zeuvoryx/engine"
)

func TestBuild48QuickControlled(t *testing.T) {
 home:=t.TempDir()
 temp:=t.TempDir()
 t.Setenv("USERPROFILE",home)
 t.Setenv("TMP",temp)
 t.Setenv("TEMP",temp)
 t.Setenv("APPDATA",t.TempDir())
 t.Setenv("ProgramData",t.TempDir())
 for _,p:=range []string{filepath.Join(home,"Downloads"),filepath.Join(home,"Desktop"),temp}{
  if err:=os.MkdirAll(p,0700);err!=nil{t.Fatal(err)}
  if err:=os.WriteFile(filepath.Join(p,"fixture.txt"),[]byte("benign Build48 fixture"),0600);err!=nil{t.Fatal(err)}
 }
 st:=&State{scanner:engine.New(),quarantine:&engine.Quarantine{Dir:t.TempDir()}}
 ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second);defer cancel()
 r,err:=st.quickScan(ctx,nil)
 if err!=nil||r.Cancelled||r.Files!=3{t.Fatalf("controlled quick scan: files=%d cancelled=%v err=%v",r.Files,r.Cancelled,err)}
}

func TestBuild48CancelActual(t *testing.T) {
 root:=t.TempDir()
 payload:=bytes.Repeat([]byte("Build48 benign cancellation fixture\n"),32768)
 for i:=0;i<128;i++{p:=filepath.Join(root,time.Unix(int64(i),0).UTC().Format("150405")+".bin");if err:=os.WriteFile(p,payload,0600);err!=nil{t.Fatal(err)}}
 st:=&State{scanner:engine.New(),quarantine:&engine.Quarantine{Dir:t.TempDir()}}
 if !st.startScan(root,false){t.Fatal("native scan did not start")}
 time.Sleep(25*time.Millisecond)
 st.mu.Lock()
 if !st.running||st.cancel==nil{st.mu.Unlock();t.Fatal("scan completed before cancellation fixture was exercised")}
 started:=time.Now();st.cancel();st.mu.Unlock()
 deadline:=started.Add(2*time.Second)
 for time.Now().Before(deadline){
  st.mu.Lock();running,cancelled:=st.running,st.report.Cancelled;st.mu.Unlock()
  if !running{if !cancelled{t.Fatal("native cancellation not confirmed by report")};t.Logf("Native cancellation confirmed in %v",time.Since(started));return}
  time.Sleep(10*time.Millisecond)
 }
 t.Fatal("native scan failed the two-second cancellation gate")
}

func TestBuild48QuarantineRoundTrip(t *testing.T) {
 path:=filepath.Join(t.TempDir(),"benign-fixture.txt");data:=[]byte("safe Build48 quarantine integrity fixture")
 if err:=os.WriteFile(path,data,0600);err!=nil{t.Fatal(err)}
 q:=&engine.Quarantine{Dir:t.TempDir()}
 item,err:=q.Add(path);if err!=nil{t.Fatal(err)}
 if _,err=os.Stat(path);!os.IsNotExist(err){t.Fatal("original remains after isolation")}
 if err=q.Restore(item.ID);err!=nil{t.Fatal(err)}
 got,err:=os.ReadFile(path);if err!=nil||!bytes.Equal(got,data){t.Fatal("restored data differs",err)}
}
