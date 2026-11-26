package scopeprogress

import (
	"context"
	"time"
)

// AwaitResult bounds only auxiliary diagnostic waits. The caller must prevent
// late publication to a completed session. It never fabricates a result.
func AwaitResult[T any](ch <-chan T, budget time.Duration) (T, bool) {
	timer := time.NewTimer(budget)
	defer timer.Stop()
	select {
	case value, ok := <-ch:
		return value, ok
	case <-timer.C:
		var zero T
		return zero, false
	}
}

// AwaitShutdown lets an active stage finish normally. After cancellation, it
// gives the stage time to return its final evidence before giving up.
func AwaitShutdown[T any](ctx context.Context, ch <-chan T, grace time.Duration) (T, bool) {
	select {
	case value, ok := <-ch:
		return value, ok
	case <-ctx.Done():
		return AwaitResult(ch, grace)
	}
}

// AwaitContext shares a single deadline among multiple shutdown consumers.
// Prefer a result already buffered even when the common deadline expired.
func AwaitContext[T any](ctx context.Context,ch <-chan T)(T,bool){
 select{case v,ok:=<-ch:return v,ok;default:}
 select{case v,ok:=<-ch:return v,ok;case <-ctx.Done():var zero T;return zero,false}
}
