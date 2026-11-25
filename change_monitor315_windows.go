package main

import (
 "context"
 "os"
 "path/filepath"
 "sync"
 "time"
 "github.com/zeuvoryx/zeuvoryx/changeguard"
)
var changeMonitor315 *changeguard.Tracker
func changeMonitor315Snapshot()changeguard.Status{
 if changeMonitor315==nil{return changeguard.New(nil).Snapshot()}
 return changeMonitor315.Snapshot()
}
func initChangeMonitor315(){
 var roots []string
 if home,e:=os.UserHomeDir();e==nil&&home!=""{roots=[]string{filepath.Join(home,"Downloads"),filepath.Join(home,"Desktop")}}
 changeMonitor315=changeguard.New(roots)
}
// Signals accelerate an existing quick scan, without replacing a report with unresolved findings.
func runChangeMonitor315(ctx context.Context,enabled func()bool,request func()bool){
 tr:=changeMonitor315
 if tr==nil{return}
 tick:=time.NewTicker(time.Second);defer tick.Stop()
 var cancel context.CancelFunc
 var workers sync.WaitGroup
 stop:=func(){if cancel!=nil{cancel();workers.Wait();cancel=nil};tr.Enable(false)}
 defer stop()
 var last time.Time
 var handled uint64
 for{select{case <-ctx.Done():return;case now:=<-tick.C:
  if !enabled(){if cancel!=nil{stop()};continue}
  if cancel==nil{
   tr.Enable(true);last=now;handled=tr.Snapshot().Signals
   cancel=startChangeWatchers315(ctx,tr,&workers)
  }
  s:=tr.Snapshot()
  if now.Sub(last)>=5*time.Minute||(s.Signals>handled&&now.Sub(last)>=time.Minute){
   if request(){last=now;handled=s.Signals;tr.ScanRequest(now)}
  }
 }}
}

// The caller owns cancellation and joins the workers on disable or exit.
func startChangeWatchers315(ctx context.Context,tr *changeguard.Tracker,workers *sync.WaitGroup) context.CancelFunc {
 sub,cancel:=context.WithCancel(ctx)
 for _,root:=range tr.Snapshot().Roots{
  workers.Add(1)
  go func(root string){defer workers.Done();if e:=changeguard.Watch(sub,root,tr);e!=nil&&sub.Err()==nil{tr.Problem(root,e)}}(root)
 }
 return cancel
}
