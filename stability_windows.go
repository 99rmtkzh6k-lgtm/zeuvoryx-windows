//go:build windows

package main

import (
 "fmt"
 "io"
 "log"
 "os"
 "path/filepath"
 "runtime"
 "runtime/debug"
 "strings"
 "sync"
 "time"
)

// Local-only bounded log. Runtime crash stacks are separately opt-in.
type stabilityWriter struct {
 mu sync.Mutex
 out io.Writer
 remaining int
}
func (w *stabilityWriter) Write(p []byte)(int,error) {
 w.mu.Lock(); defer w.mu.Unlock()
 n:=len(p)
 if len(p)>w.remaining {p=p[:w.remaining]}
 if len(p)>0 {k,e:=w.out.Write(p);w.remaining-=k;if e!=nil{return k,e}}
 return n,nil
}
func stabilityWorkers(n int) int {
 if n<=2{return 1}
 if n>4{return 4}
 return n
}
func startStabilityDiagnostics() func() {
 old:=log.Writer()
 // Never raise an existing operator-configured limit. This is a soft Go limit,
 // not a process RAM cap; it does not cover WebView2.
 if runtime.NumCPU()<=2 {
  const budget=256<<20
  if debug.SetMemoryLimit(-1)>budget {debug.SetMemoryLimit(budget)}
 }
 base,e:=os.UserConfigDir()
 if e!=nil{return func(){}}
 dir:=filepath.Join(base,"ZEUVORYX","stability-lab")
 if e=os.MkdirAll(dir,0700);e!=nil{return func(){}}
 // Only retire our own expired, regular diagnostic files; do not touch reports.
 if entries,err:=os.ReadDir(dir);err==nil {
  for _,entry:=range entries {
   name:=entry.Name()
   if !strings.HasPrefix(name,"session-")||!(strings.HasSuffix(name,".log")||strings.HasSuffix(name,".crash")){continue}
   if info,err:=entry.Info();err==nil&&info.Mode().IsRegular()&&time.Since(info.ModTime())>7*24*time.Hour { _=os.Remove(filepath.Join(dir,name)) }
  }
 }
 prefix:=fmt.Sprintf("session-%d-%d",time.Now().UnixNano(),os.Getpid())
 f,e:=os.OpenFile(filepath.Join(dir,prefix+".log"),os.O_CREATE|os.O_EXCL|os.O_WRONLY,0600)
 if e!=nil{return func(){}}
 log.SetOutput(io.MultiWriter(&stabilityWriter{out:f,remaining:256<<10},old))
 // Native runtime stack may contain local paths; opt-in, never upload.
 crashEnabled:=false
 if os.Getenv("ZEUVORYX_CRASH_DETAIL")=="1" {
  if cf,err:=os.OpenFile(filepath.Join(dir,prefix+".crash"),os.O_CREATE|os.O_EXCL|os.O_WRONLY,0600);err==nil {
   crashEnabled=debug.SetCrashOutput(cf,debug.CrashOptions{})==nil
   _=cf.Close()
  }
 }
 log.Printf("stability session_start version=3.0.8-lab cpu_count=%d go=%s workers=%d soft_go_limit=%d",runtime.NumCPU(),runtime.Version(),stabilityWorkers(runtime.NumCPU()),debug.SetMemoryLimit(-1))
 return func(){
  log.Print("stability main_return")
  if crashEnabled { _=debug.SetCrashOutput(nil,debug.CrashOptions{}) }
  log.SetOutput(old);_ = f.Sync();_ = f.Close()
 }
}
// Writes only numeric progress, never filenames or contents. Samples only active scans.
func watchStability(s *State) func() {
 stop:=make(chan struct{});done:=make(chan struct{})
 go func(){
  defer close(done)
  ticker:=time.NewTicker(5*time.Second);defer ticker.Stop()
  for {
   select {
   case <-stop:return
   case <-ticker.C:
    if !s.mu.TryLock(){log.Print("stability state_busy");continue}
    active,n,b:=s.running,s.files,s.bytes
    s.mu.Unlock()
    if active {
     var m runtime.MemStats;runtime.ReadMemStats(&m)
     log.Printf("stability scan_progress files=%d bytes=%d heap=%d sys=%d goroutines=%d",n,b,m.HeapAlloc,m.Sys,runtime.NumGoroutine())
    }
   }
  }
 }()
 return func(){close(stop);<-done}
}
