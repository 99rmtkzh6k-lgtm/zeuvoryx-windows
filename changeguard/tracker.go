package changeguard

import ("sync";"time")

type Status struct {
 Enabled bool `json:"enabled"`
 Roots []string `json:"roots"`
 ActiveRoots []string `json:"active_roots"`
 Signals uint64 `json:"change_signals"`
 LastSignal time.Time `json:"last_signal"`
 LastScanRequest time.Time `json:"last_scan_request"`
 Problems map[string]string `json:"problems"`
 Note string `json:"note"`
}
type Tracker struct {mu sync.Mutex;s Status}
func New(roots []string)*Tracker{return &Tracker{s:Status{Roots:append([]string(nil),roots...),Problems:map[string]string{},Note:"Directory notifications are signals, not file counts or malware detections. No pre-execution blocking."}}}
func(t *Tracker) Enable(v bool){t.mu.Lock();defer t.mu.Unlock();t.s.Enabled=v}
func(t *Tracker) Active(root string,v bool){t.mu.Lock();defer t.mu.Unlock();for i,r:=range t.s.ActiveRoots{if r==root{if v{delete(t.s.Problems,root)};if !v{t.s.ActiveRoots=append(t.s.ActiveRoots[:i],t.s.ActiveRoots[i+1:]...)};return}};if v{t.s.ActiveRoots=append(t.s.ActiveRoots,root);delete(t.s.Problems,root)}}
func(t *Tracker) Problem(root string,e error){t.mu.Lock();defer t.mu.Unlock();if e!=nil{t.s.Problems[root]=e.Error()}}
func(t *Tracker) Signal(at time.Time){t.mu.Lock();defer t.mu.Unlock();if t.s.Signals<^uint64(0){t.s.Signals++};if at.After(t.s.LastSignal){t.s.LastSignal=at.UTC()}}
func(t *Tracker) ScanRequest(at time.Time){t.mu.Lock();defer t.mu.Unlock();if at.After(t.s.LastScanRequest){t.s.LastScanRequest=at.UTC()}}
func(t *Tracker) Snapshot() Status{t.mu.Lock();defer t.mu.Unlock();s:=t.s;s.Roots=append([]string(nil),s.Roots...);s.ActiveRoots=append([]string(nil),s.ActiveRoots...);s.Problems=map[string]string{};for k,v:=range t.s.Problems{s.Problems[k]=v};return s}
