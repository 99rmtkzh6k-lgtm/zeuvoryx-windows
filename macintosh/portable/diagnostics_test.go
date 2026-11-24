package portable

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestDiagnosticsOptIn(t *testing.T) {
	var d diagRecorder
	now := time.Now()
	n := 0
	sample := func() diagSample { n++; return diagSample{At: now, Files: n} }
	d.capture(sample)
	if n != 0 {
		t.Fatal("captured without permission")
	}
	if !d.begin(now, sample) || d.begin(now, sample) {
		t.Fatal("start")
	}
	d.capture(sample)
	if !d.stop(now, sample) || d.stop(now, sample) {
		t.Fatal("stop")
	}
	d.capture(sample)
	r := d.report()
	if n != 3 || r.Recording || len(r.Samples) != 3 || r.Stopped == nil {
		t.Fatal(r)
	}
	if !d.begin(now, sample) || len(d.report().Samples) != 1 {
		t.Fatal("reset")
	}
}
func TestDiagnosticsRing(t *testing.T) {
	var d diagRecorder
	n := 0
	sample := func() diagSample { n++; return diagSample{Files: n} }
	d.begin(time.Now(), sample)
	for i := 0; i < diagMaxSamples+4; i++ {
		d.capture(sample)
	}
	r := d.report()
	if len(r.Samples) != diagMaxSamples || r.Dropped != 5 || r.Samples[0].Files != 6 || r.Samples[len(r.Samples)-1].Files != n {
		t.Fatal("wrong bounded chronology")
	}
}
func TestDiagnosticsAllowlist(t *testing.T) {
	if diagOp("C:/Users/secret") != "" || diagOp("quick") != "quick" || diagCount("secret") != nil || diagCount(1.5) != nil || diagCount(-1) != nil || *diagCount(0.0) != 0 {
		t.Fatal("invalid values")
	}
	r := diagReport{Samples: []diagSample{{DefenderOperation: diagOp("SECRET"), ActiveAfter: diagCount(0)}}}
	b, e := json.Marshal(r)
	if e != nil || strings.Contains(string(b), "SECRET") || !strings.Contains(string(b), `"defender_active_after":0`) {
		t.Fatal(string(b), e)
	}
}
func TestDiagnosticsExport(t *testing.T) {
	var d diagRecorder
	dir := t.TempDir()
	if _, e := writeDiagReport(dir, d.report()); e == nil {
		t.Fatal("empty export")
	}
	d.begin(time.Now(), func() diagSample { return diagSample{Files: 42} })
	p, e := writeDiagReport(dir, d.report())
	if e != nil {
		t.Fatal(e)
	}
	b, e := os.ReadFile(p)
	var r diagReport
	if e != nil || json.Unmarshal(b, &r) != nil || r.Schema != "zeuvoryx-diagnostics-v1" || r.Samples[0].Files != 42 {
		t.Fatal("invalid export")
	}
}
func TestDiagnosticsConcurrent(t *testing.T) {
	var d diagRecorder
	sample := func() diagSample { return diagSample{Files: 1} }
	d.begin(time.Now(), sample)
	var w sync.WaitGroup
	for i := 0; i < 8; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			for j := 0; j < 50; j++ {
				d.capture(sample)
				_ = d.report()
			}
		}()
	}
	w.Wait()
	if len(d.report().Samples) != 401 {
		t.Fatal("lost sample")
	}
}
