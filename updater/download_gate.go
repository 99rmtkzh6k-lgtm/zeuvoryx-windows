package updater

import "sync"

// DownloadGate serializes requests for the same artifact inside one process.
// It does not control downloads initiated in Safari or other processes.
type DownloadGate struct {
 mu sync.Mutex
 active map[string]struct{}
}

// Begin reserves an artifact until the returned release function is called.
// An empty identifier is rejected so unrelated downloads cannot share a lock.
func (g *DownloadGate) Begin(id string) (release func(), ok bool) {
 if id == "" { return nil, false }
 g.mu.Lock()
 if g.active == nil { g.active = make(map[string]struct{}) }
 if _, exists := g.active[id]; exists {
  g.mu.Unlock()
  return nil, false
 }
 g.active[id] = struct{}{}
 g.mu.Unlock()
 var once sync.Once
 return func() { once.Do(func() {
  g.mu.Lock()
  delete(g.active,id)
  g.mu.Unlock()
 }) }, true
}

const DuplicateDownloadMessage = "Já está baixando"
