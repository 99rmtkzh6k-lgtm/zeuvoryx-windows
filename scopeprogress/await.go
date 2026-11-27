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

// AwaitResultContext stops waiting for auxiliary results on scan cancellation.
// It does not mark an unresponsive worker as complete.
func AwaitResultContext[T any](ctx context.Context, ch <-chan T, budget time.Duration) (T, bool) {
    timer := time.NewTimer(budget)
    defer timer.Stop()
    select {
    case value, ok := <-ch:
        return value, ok
    case <-ctx.Done():
        select {
        case value, ok := <-ch:
            return value, ok
        default:
            var zero T
            return zero, false
        }
    case <-timer.C:
        var zero T
        return zero, false
    }
}
