package engine
import "context"
type liveReportKey struct{}
func WithLiveReport(ctx context.Context,fn func(Report))context.Context{return context.WithValue(ctx,liveReportKey{},fn)}
func publishLiveReport(ctx context.Context,r Report){
 if fn,ok:=ctx.Value(liveReportKey{}).(func(Report));ok&&fn!=nil{fn(r)}
}
