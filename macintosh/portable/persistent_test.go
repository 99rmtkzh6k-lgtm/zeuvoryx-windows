package portable

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestSnapshotRoundtrip(t *testing.T) {
	p := filepath.Join(t.TempDir(), "intel", "snapshot.json")
	v := newIntelSnapshot()
	v.Hashes[strings.Repeat("a", 64)] = "test"
	v.URLs["https://example.invalid/a"] = true
	v.Updated = time.Now()
	if e := saveSnapshot(p, v); e != nil {
		t.Fatal(e)
	}
	r, e := loadSnapshot(p)
	if e != nil || len(r.Hashes) != 1 || len(r.URLs) != 1 {
		t.Fatal(r, e)
	}
	v.Hashes[strings.Repeat("b", 64)] = "second"
	if e = saveSnapshot(p, v); e != nil {
		t.Fatal(e)
	}
	r, e = loadSnapshot(p)
	if e != nil || len(r.Hashes) != 2 {
		t.Fatal(r, e)
	}
}
func TestInvalidSnapshotDoesNotOverwrite(t *testing.T) {
	p := filepath.Join(t.TempDir(), "snapshot.json")
	v := newIntelSnapshot()
	if e := saveSnapshot(p, v); e != nil {
		t.Fatal(e)
	}
	v.Hashes["bad"] = "test"
	if e := saveSnapshot(p, v); e == nil {
		t.Fatal("invalid accepted")
	}
	r, e := loadSnapshot(p)
	if e != nil || len(r.Hashes) != 0 {
		t.Fatal(r, e)
	}
	os.WriteFile(p, []byte("broken"), 0600)
	if _, e = loadSnapshot(p); e == nil {
		t.Fatal("corruption hidden")
	}
}
func TestReputationExact(t *testing.T) {
	m, e := parseReputationList("# example\nhttps://EXAMPLE.invalid/a\n")
	if e != nil {
		t.Fatal(e)
	}
	r, e := assessWithReputation("https://example.invalid/a#x", m)
	if e != nil || len(r.Signals) != 1 {
		t.Fatal(r, e)
	}
	r, e = assessWithReputation("https://example.invalid/b", m)
	if e != nil || len(r.Signals) != 0 {
		t.Fatal(r, e)
	}
}
func TestReputationInvalid(t *testing.T) {
	for _, s := range []string{"file:///tmp/a", "https://x.invalid@other.invalid/", "https://x.invalid/a b", ""} {
		if _, e := parseReputationList(s); e == nil {
			t.Fatal(s)
		}
	}
}
