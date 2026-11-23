package engine

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type QuarantineItem struct {
	ID       string    `json:"id"`
	Original string    `json:"original"`
	When     time.Time `json:"when"`
	Size     int64     `json:"size"`
}
type Quarantine struct{ Dir string }

func NewQuarantine() (*Quarantine, error) {
	base, e := os.UserConfigDir()
	if e != nil {
		return nil, e
	}
	dir := filepath.Join(base, "ZEUVORYX", "quarantine")
	if e = os.MkdirAll(dir, 0700); e != nil {
		return nil, e
	}
	return &Quarantine{Dir: dir}, nil
}
func (q *Quarantine) IsQuarantined(path string) bool {
	abs, _ := filepath.Abs(path)
	dir, _ := filepath.Abs(q.Dir)
	rel, e := filepath.Rel(dir, abs)
	return e == nil && (rel == "." || (!strings.HasPrefix(rel, ".."+string(os.PathSeparator)) && rel != ".."))
}
func (q *Quarantine) Add(path string) (QuarantineItem, error) {
	abs, e := filepath.Abs(path)
	if e != nil {
		return QuarantineItem{}, e
	}
	if q.IsQuarantined(abs) {
		return QuarantineItem{}, errors.New("arquivo já está em quarentena")
	}
	st, e := os.Lstat(abs)
	if e != nil {
		return QuarantineItem{}, e
	}
	if !st.Mode().IsRegular() {
		return QuarantineItem{}, errors.New("somente arquivos comuns")
	}
	var random [16]byte
	if _, e = rand.Read(random[:]); e != nil {
		return QuarantineItem{}, e
	}
	id := hex.EncodeToString(random[:])
	item := QuarantineItem{ID: id, Original: abs, When: time.Now(), Size: st.Size()}
	source, e := os.Open(abs)
	if e != nil {
		return QuarantineItem{}, e
	}
	defer source.Close()
	destPath := filepath.Join(q.Dir, id+".bin")
	dest, e := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if e != nil {
		return QuarantineItem{}, e
	}
	_, e = io.Copy(dest, source)
	closeErr := dest.Close()
	if e != nil || closeErr != nil {
		os.Remove(destPath)
		return QuarantineItem{}, fmt.Errorf("falha ao copiar: %v %v", e, closeErr)
	}
	meta, e := json.MarshalIndent(item, "", "  ")
	if e != nil {
		os.Remove(destPath)
		return QuarantineItem{}, e
	}
	if e = os.WriteFile(filepath.Join(q.Dir, id+".json"), meta, 0600); e != nil {
		os.Remove(destPath)
		return QuarantineItem{}, e
	}
	// Original is deleted only after the secure copy and metadata have been persisted.
	if e = os.Remove(abs); e != nil {
		os.Remove(destPath)
		os.Remove(filepath.Join(q.Dir, id+".json"))
		return QuarantineItem{}, e
	}
	return item, nil
}
func (q *Quarantine) List() ([]QuarantineItem, error) {
	entries, e := os.ReadDir(q.Dir)
	if e != nil {
		return nil, e
	}
	result := []QuarantineItem{}
	for _, v := range entries {
		if !strings.HasSuffix(v.Name(), ".json") {
			continue
		}
		b, e := os.ReadFile(filepath.Join(q.Dir, v.Name()))
		if e != nil {
			continue
		}
		var item QuarantineItem
		if json.Unmarshal(b, &item) == nil && validID(item.ID) {
			result = append(result, item)
		}
	}
	return result, nil
}
func validID(id string) bool {
	if len(id) != 32 {
		return false
	}
	_, e := hex.DecodeString(id)
	return e == nil
}
func (q *Quarantine) Restore(id string) error {
	if !validID(id) {
		return errors.New("identificador inválido")
	}
	b, e := os.ReadFile(filepath.Join(q.Dir, id+".json"))
	if e != nil {
		return e
	}
	var item QuarantineItem
	if e = json.Unmarshal(b, &item); e != nil {
		return e
	}
	if item.ID != id || item.Original == "" {
		return errors.New("metadados inconsistentes")
	}
	if _, e = os.Lstat(item.Original); e == nil {
		return errors.New("já existe arquivo no destino; restauração cancelada")
	}
	if !os.IsNotExist(e) {
		return e
	}
	src, e := os.Open(filepath.Join(q.Dir, id+".bin"))
	if e != nil {
		return e
	}
	defer src.Close()
	if e = os.MkdirAll(filepath.Dir(item.Original), 0700); e != nil {
		return e
	}
	dst, e := os.OpenFile(item.Original, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if e != nil {
		return e
	}
	_, e = io.Copy(dst, src)
	ce := dst.Close()
	if e != nil || ce != nil {
		os.Remove(item.Original)
		return fmt.Errorf("restauração incompleta: %v %v", e, ce)
	}
	os.Remove(filepath.Join(q.Dir, id+".json"))
	os.Remove(filepath.Join(q.Dir, id+".bin"))
	return nil
}
