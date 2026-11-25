// Package changeguard coordinates advisory scans; it does not intercept execution.
package changeguard
import("context";"errors";"time")
type ScanCycle struct{Signals uint64;Started,Finished time.Time;Error string}
// RunScans serializes scans, preserves signals arriving during a scan, and periodically
// reconciles missed notifications. Caller owns watcher lifetime and chosen root.
func RunScans(ctx context.Context,tr *Tracker,tick,reconcile time.Duration,scan func(context.Context)error,report func(ScanCycle))error{
 if tr==nil||tick<=0||reconcile<tick||scan==nil{return errors.New("invalid scan scheduler configuration")}
 timer:=time.NewTicker(tick);defer timer.Stop()
 var handled uint64;last:=time.Time{}
 for{
  select{case <-ctx.Done():return ctx.Err();case <-timer.C:}
  if ctx.Err()!=nil{return ctx.Err()}
  snapshot:=tr.Snapshot()
  if snapshot.Signals==handled&&!last.IsZero()&&time.Since(last)<reconcile{continue}
  cycle:=ScanCycle{Signals:snapshot.Signals,Started:time.Now().UTC()}
  tr.ScanRequest(cycle.Started)
  if err:=scan(ctx);err!=nil{cycle.Error=err.Error()}else{handled=snapshot.Signals}
  cycle.Finished=time.Now().UTC();last=cycle.Finished
  if report!=nil{report(cycle)}
 }
}
