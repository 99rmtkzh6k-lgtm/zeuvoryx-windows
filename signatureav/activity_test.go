package signatureav

import (
	"testing"
	"time"
)

func TestActivityAdvisoryUsesObservedCounters(t *testing.T) {
	start := time.Date(2026, 11, 26, 12, 0, 0, 0, time.UTC)
	first := updateActivity(nil, Activity{ObservedAt: start, CPU100ns: 100})
	still := updateActivity(&first, Activity{ObservedAt: start.Add(idleAdvisoryAfter), CPU100ns: 100})
	if !still.IdleAdvisory || still.IdleSeconds != 120 || !still.LastActivityAt.Equal(start) {
		t.Fatalf("inactivity not identified: %+v", still)
	}
	for _, next := range []Activity{
		{CPU100ns: 101}, {CPU100ns: 100, ReadBytes: 1},
		{CPU100ns: 100, WriteBytes: 1}, {CPU100ns: 100, OtherBytes: 1},
	} {
		next.ObservedAt = still.ObservedAt.Add(activityInterval)
		result := updateActivity(&still, next)
		if result.IdleAdvisory || result.IdleSeconds != 0 || !result.LastActivityAt.Equal(next.ObservedAt) {
			t.Fatalf("observed activity did not clear warning: %+v", result)
		}
	}
}
