package scopeprogress

import (
	"context"
	"testing"
	"time"
)

func TestAwaitShutdown(t *testing.T) {
	t.Run("normal result", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		results := make(chan int, 1)
		results <- 42
		if got, ok := AwaitShutdown(ctx, results, 0); !ok || got != 42 {
			t.Fatalf("result = %d, %t", got, ok)
		}
	})
	t.Run("cooperative cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		results := make(chan int, 1)
		go func() {
			<-ctx.Done()
			results <- 42
		}()
		cancel()
		if got, ok := AwaitShutdown(ctx, results, time.Second); !ok || got != 42 {
			t.Fatalf("result = %d, %t", got, ok)
		}
	})
	t.Run("unresponsive stage", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		if _, ok := AwaitShutdown(ctx, make(chan int), time.Millisecond); ok {
			t.Fatal("missing result was reported as available")
		}
	})
}
