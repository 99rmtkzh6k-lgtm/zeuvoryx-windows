package signatureav

import (
	"context"
	"os/exec"
	"time"
)

// Activity describes observed process counters. It does not estimate scan coverage.
type Activity struct {
	ObservedAt       time.Time `json:"observed_at"`
	LastActivityAt   time.Time `json:"last_activity_at"`
	CPU100ns         uint64    `json:"cpu_100ns"`
	ReadBytes        uint64    `json:"read_bytes"`
	WriteBytes       uint64    `json:"write_bytes"`
	OtherBytes       uint64    `json:"other_bytes"`
	WorkingSetBytes  *uint64   `json:"working_set_bytes,omitempty"`
	PeakWorkingBytes *uint64   `json:"peak_working_set_bytes,omitempty"`
	IdleSeconds      int64     `json:"idle_seconds"`
	IdleAdvisory     bool      `json:"idle_advisory"`
}

const activityInterval = 5 * time.Second
const idleAdvisoryAfter = 2 * time.Minute

type activityObserverKey struct{}

// WithActivityObserver observes sampled ClamScan activity without exposing paths.
func WithActivityObserver(ctx context.Context, observer func(Activity)) context.Context {
	return context.WithValue(ctx, activityObserverKey{}, observer)
}

func updateActivity(last *Activity, current Activity) Activity {
	current.LastActivityAt = current.ObservedAt
	if last != nil && current.CPU100ns == last.CPU100ns && current.ReadBytes == last.ReadBytes &&
		current.WriteBytes == last.WriteBytes && current.OtherBytes == last.OtherBytes {
		current.LastActivityAt = last.LastActivityAt
	}
	idle := current.ObservedAt.Sub(current.LastActivityAt)
	current.IdleSeconds = int64(idle / time.Second)
	current.IdleAdvisory = idle >= idleAdvisoryAfter
	return current
}

func runWithActivity(ctx context.Context, cmd *exec.Cmd) (*Activity, error) {
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	sample, release := openProcessSampler(cmd.Process.Pid)
	if sample == nil {
		return nil, cmd.Wait()
	}
	defer release()

	observer, _ := ctx.Value(activityObserverKey{}).(func(Activity))
	done, completed := make(chan struct{}), make(chan struct{})
	var last *Activity
	capture := func() {
		current, ok := sample()
		if !ok {
			return
		}
		next := updateActivity(last, current)
		last = &next
		if observer != nil {
			observer(next)
		}
	}
	capture()
	go func() {
		defer close(completed)
		tick := time.NewTicker(activityInterval)
		defer tick.Stop()
		for {
			select {
			case <-done:
				capture()
				return
			case <-tick.C:
				capture()
			}
		}
	}()
	err := cmd.Wait()
	close(done)
	<-completed
	return last, err
}
