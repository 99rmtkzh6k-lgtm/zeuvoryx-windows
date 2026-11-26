package engine

import (
	"crypto/rand"
	"crypto/sha256"
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
	SHA256   string    `json:"sha256,omitempty"`
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
	abs, e := filepath.Abs(path)
	if e != nil {
		return false
	}
	dir, e := filepath.Abs(q.Dir)
	if e != nil {
		return false
	}
	rel, e := filepath.Rel(dir, abs)
	return e == nil && (rel == "." || (!strings.HasPrefix(rel, ".."+string(os.PathSeparator)) && rel != ".."))
}

func copyClosed(from, to, expectedHash string, expectedSize int64) (string, error) {
	src, e := os.Open(from)
	if e != nil {
		return "", e
	}
	defer src.Close()
	dst, e := os.OpenFile(to, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if e != nil {
		return "", e
	}

	hash := sha256.New()
	written, copyErr := io.Copy(io.MultiWriter(dst, hash), src)
	digest := hex.EncodeToString(hash.Sum(nil))
	if copyErr == nil && (written != expectedSize || (expectedHash != "" && digest != expectedHash)) {
		copyErr = errors.New("integridade ou tamanho do arquivo em quarentena não confere")
	}
	var syncErr error
	if copyErr == nil {
		syncErr = dst.Sync()
	}
	if e := errors.Join(copyErr, syncErr, dst.Close()); e != nil {
		return "", errors.Join(e, os.Remove(to))
	}
	return digest, nil
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

	if e = os.MkdirAll(q.Dir, 0700); e != nil {
		return QuarantineItem{}, e
	}

	var random [16]byte
	if _, e = rand.Read(random[:]); e != nil {
		return QuarantineItem{}, e
	}

	id := hex.EncodeToString(random[:])

	item := QuarantineItem{
		ID:       id,
		Original: abs,
		When:     time.Now(),
		Size:     st.Size(),
	}

	bin := filepath.Join(q.Dir, id+".bin")
	metaPath := filepath.Join(q.Dir, id+".json")

	item.SHA256, e = copyClosed(abs, bin, "", st.Size())
	if e != nil {
		return QuarantineItem{}, e
	}
	latest, e := os.Lstat(abs)
	if e != nil || !os.SameFile(st, latest) || st.Size() != latest.Size() || !st.ModTime().Equal(latest.ModTime()) {
		return QuarantineItem{}, errors.Join(errors.New("arquivo de origem mudou durante o isolamento"), e, os.Remove(bin))
	}
	meta, e := json.MarshalIndent(item, "", "  ")
	if e != nil {
		return QuarantineItem{}, errors.Join(e, os.Remove(bin))
	}

	if e = os.WriteFile(metaPath, meta, 0600); e != nil {
		return QuarantineItem{}, errors.Join(e, os.Remove(bin))
	}

	if e = os.Remove(abs); e != nil {
		return QuarantineItem{}, errors.Join(
			e,
			os.Remove(bin),
			os.Remove(metaPath),
		)
	}

	return item, nil
}

func validID(id string) bool {
	if len(id) != 32 {
		return false
	}
	_, e := hex.DecodeString(id)
	return e == nil
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

		if json.Unmarshal(b, &item) == nil &&
			validID(item.ID) &&
			v.Name() == item.ID+".json" {
			result = append(result, item)
		}
	}

	return result, nil
}

func (q *Quarantine) Restore(id string) error {
	if !validID(id) {
		return errors.New("identificador inválido")
	}

	metaPath := filepath.Join(q.Dir, id+".json")
	bin := filepath.Join(q.Dir, id+".bin")

	b, e := os.ReadFile(metaPath)
	if e != nil {
		return e
	}

	var item QuarantineItem

	if e = json.Unmarshal(b, &item); e != nil {
		return e
	}

	if item.ID != id ||
		item.Original == "" ||
		!filepath.IsAbs(item.Original) ||
		q.IsQuarantined(item.Original) {
		return errors.New("metadados inconsistentes")
	}

	if _, e = os.Lstat(item.Original); e == nil {
		return errors.New("já existe arquivo no destino; restauração cancelada")
	}

	if !os.IsNotExist(e) {
		return e
	}

	if e = os.MkdirAll(filepath.Dir(item.Original), 0700); e != nil {
		return e
	}

	if _, e = copyClosed(bin, item.Original, item.SHA256, item.Size); e != nil {
		return e
	}

	if e = os.Remove(bin); e != nil {
		return fmt.Errorf("arquivo restaurado, falha ao limpar cópia: %w", e)
	}

	if e = os.Remove(metaPath); e != nil {
		return fmt.Errorf("arquivo restaurado, falha ao limpar metadados: %w", e)
	}

	return nil
}
