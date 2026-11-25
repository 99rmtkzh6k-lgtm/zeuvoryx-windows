package scanplan
import "time"
// Await bounds collection of optional results. Producers use buffered channels.
func Await[T any](ch <-chan T, grace time.Duration)(T,bool){timer:=time.NewTimer(grace);defer timer.Stop();select{case v,ok:=<-ch:return v,ok;case <-timer.C:var zero T;return zero,false}}
