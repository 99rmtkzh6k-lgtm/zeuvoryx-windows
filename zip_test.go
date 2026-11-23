package engine

import (
	"archive/zip"
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestZipScan(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "archive.zip")
	f, e := os.Create(p)
	if e != nil {
		t.Fatal(e)
	}
	w := zip.NewWriter(f)
	z, e := w.Create("danger.pdf.ps1")
	if e != nil {
		t.Fatal(e)
	}
	z.Write([]byte("Invoke-Expression foo -EncodedCommand xyz"))
	w.Close()
	f.Close()
	r, e := New().Scan(context.Background(), p, nil)
	if e != nil || len(r.Findings) == 0 {
		t.Fatalf("ZIP falhou: %+v %v", r, e)
	}
}
