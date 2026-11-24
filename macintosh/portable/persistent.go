package portable

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type intelSnapshot struct {
	Version int               `json:"version"`
	Updated time.Time         `json:"updated"`
	Hashes  map[string]string `json:"hashes"`
	URLs    map[string]bool   `json:"urls"`
}

const maxSnapshotBytes = 48 << 20
const maxReputationURLs = 100000

func newIntelSnapshot() intelSnapshot {
	return intelSnapshot{Version: 1, Hashes: map[string]string{}, URLs: map[string]bool{}}
}
func intelSnapshotPath() (string, error) {
	d, e := os.UserConfigDir()
	if e != nil {
		return "", e
	}
	return filepath.Join(d, "ZEUVORYX Macintosh LAB", "intelligence", "snapshot.json"), nil
}
func validateSnapshot(v intelSnapshot) error {
	if v.Version != 1 || v.Hashes == nil || v.URLs == nil || len(v.Hashes) > maxIndicators || len(v.URLs) > maxReputationURLs {
		return fmt.Errorf("Base inválida ou acima do limite")
	}
	for k, label := range v.Hashes {
		if !validHash(k) || k != strings.ToLower(k) || len(label) > 512 {
			return fmt.Errorf("Indicador inválido na base")
		}
	}
	for k, b := range v.URLs {
		if !b || !validReputationURL(k) {
			return fmt.Errorf("URL inválida na base")
		}
	}
	return nil
}
func loadSnapshot(path string) (intelSnapshot, error) {
	v := newIntelSnapshot()
	f, e := os.Open(path)
	if os.IsNotExist(e) {
		return v, nil
	}
	if e != nil {
		return v, e
	}
	defer f.Close()
	b, e := io.ReadAll(io.LimitReader(f, maxSnapshotBytes+1))
	if e != nil {
		return v, e
	}
	if len(b) > maxSnapshotBytes {
		return v, fmt.Errorf("Base acima do limite")
	}
	if e = json.Unmarshal(b, &v); e != nil {
		return newIntelSnapshot(), e
	}
	if e = validateSnapshot(v); e != nil {
		return newIntelSnapshot(), e
	}
	return v, nil
}
func saveSnapshot(path string, v intelSnapshot) error {
	if e := validateSnapshot(v); e != nil {
		return e
	}
	b, e := json.Marshal(v)
	if e != nil {
		return e
	}
	if len(b) > maxSnapshotBytes {
		return fmt.Errorf("Base acima do limite")
	}
	dir := filepath.Dir(path)
	if e = os.MkdirAll(dir, 0700); e != nil {
		return e
	}
	f, e := os.CreateTemp(dir, ".snapshot-*")
	if e != nil {
		return e
	}
	defer os.Remove(f.Name())
	_, e = f.Write(b)
	if e == nil {
		e = f.Sync()
	}
	ce := f.Close()
	if e == nil {
		e = ce
	}
	if e != nil {
		return e
	}
	return os.Rename(f.Name(), path)
}
