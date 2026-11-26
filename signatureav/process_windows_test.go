package signatureav

import (
	"os"
	"testing"
)

func TestWindowsSamplerQueriesRealProcess(t *testing.T) {
	sample, closeHandle := openProcessSampler(os.Getpid())
	if sample == nil {
		t.Fatal("cannot open test process for native observation")
	}
	defer closeHandle()
	a, ok := sample()
	if !ok || a.ObservedAt.IsZero() || a.WorkingSetBytes == nil || *a.WorkingSetBytes == 0 ||
		a.PeakWorkingBytes == nil || *a.PeakWorkingBytes < *a.WorkingSetBytes {
		t.Fatalf("native process counters missing or inconsistent: %+v", a)
	}
}
