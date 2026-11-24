package portable

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const diagMaxSamples = 1800

// Explicit allowlist: never serialize paths, raw reports, URLs, credentials or error text.
type diagSample struct {
	At                time.Time `json:"at_utc"`
	ScanRunning       bool      `json:"scan_running"`
	Files             int       `json:"files"`
	Bytes             int64     `json:"bytes"`
	ScanFinished      bool      `json:"scan_finished"`
	ScanCancelled     bool      `json:"scan_cancelled"`
	Skipped           int       `json:"skipped"`
	Findings          int       `json:"findings"`
	ScanError         bool      `json:"scan_error"`
	MonitorEnabled    bool      `json:"monitor_enabled"`
	Indicators        int       `json:"sha256_indicators"`
	GoHeapBytes       uint64    `json:"go_heap_bytes"`
	DefenderRunning   bool      `json:"defender_running"`
	DefenderOperation string    `json:"defender_operation,omitempty"`
	DefenderError     bool      `json:"defender_error"`
	ActiveBefore      *int      `json:"defender_active_before,omitempty"`
	ActiveAfter       *int      `json:"defender_active_after,omitempty"`
}

type diagReport struct {
	Schema    string       `json:"schema"`
	Version   string       `json:"version"`
	Recording bool         `json:"recording"`
	Started   time.Time    `json:"started_utc"`
	Stopped   *time.Time   `json:"stopped_utc,omitempty"`
	Dropped   int          `json:"old_samples_discarded"`
	Samples   []diagSample `json:"samples"`
}

type diagRecorder struct {
	mu      sync.Mutex
	active  bool
	started time.Time
	stopped *time.Time
	dropped int
	next    int
	samples []diagSample
}

func (d *diagRecorder) begin(now time.Time, sample func() diagSample) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.active {
		return false
	}
	d.active = true
	d.started = now.UTC()
	d.stopped = nil
	d.dropped = 0
	d.next = 0
	d.samples = make([]diagSample, 0, diagMaxSamples)
	d.appendLocked(sample())
	return true
}
func (d *diagRecorder) appendLocked(s diagSample) {
	if len(d.samples) < diagMaxSamples {
		d.samples = append(d.samples, s)
		return
	}
	d.samples[d.next] = s
	d.next = (d.next + 1) % diagMaxSamples
	d.dropped++
}
func (d *diagRecorder) capture(sample func() diagSample) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.active {
		d.appendLocked(sample())
	}
}
func (d *diagRecorder) stop(now time.Time, sample func() diagSample) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.active {
		return false
	}
	d.appendLocked(sample())
	d.active = false
	t := now.UTC()
	d.stopped = &t
	return true
}
func (d *diagRecorder) report() diagReport {
	d.mu.Lock()
	defer d.mu.Unlock()
	out := make([]diagSample, 0, len(d.samples))
	out = append(out, d.samples[d.next:]...)
	out = append(out, d.samples[:d.next]...)
	return diagReport{Schema: "zeuvoryx-diagnostics-v1", Version: "3.0-diagnostico", Recording: d.active, Started: d.started, Stopped: d.stopped, Dropped: d.dropped, Samples: out}
}
func diagOp(op string) string {
	switch op {
	case "quick", "full", "update", "history", "updates", "remediate", "network-block", "folder-block":
		return op
	}
	return ""
}
func diagCount(v any) *int {
	var n int
	switch x := v.(type) {
	case int:
		n = x
	case float64:
		if x < 0 || x > 1000000000 || x != float64(int(x)) {
			return nil
		}
		n = int(x)
	default:
		return nil
	}
	if n < 0 || n > 1000000000 {
		return nil
	}
	return &n
}
func writeDiagReport(dir string, r diagReport) (string, error) {
	if len(r.Samples) == 0 {
		return "", fmt.Errorf("Nenhuma amostra: inicie uma sessão de diagnóstico")
	}
	if e := os.MkdirAll(dir, 0700); e != nil {
		return "", e
	}
	f, e := os.CreateTemp(dir, "ZEUVORYX-diagnostico-*.json")
	if e != nil {
		return "", e
	}
	name := f.Name()
	ok := false
	defer func() {
		if !ok {
			f.Close()
			os.Remove(name)
		}
	}()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if e = enc.Encode(r); e != nil {
		return "", e
	}
	if e = f.Close(); e != nil {
		return "", e
	}
	ok = true
	return name, nil
}
func exportDiag(r diagReport) (string, error) {
	base, e := os.UserConfigDir()
	if e != nil {
		return "", e
	}
	return writeDiagReport(filepath.Join(base, "ZEUVORYX Macintosh LAB", "diagnosticos"), r)
}
