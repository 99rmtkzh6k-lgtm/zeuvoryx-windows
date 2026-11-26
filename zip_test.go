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

func TestDeepZIPEntriesPreserveEICARCoverage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested-folders.zip")
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	archive := zip.NewWriter(file)
	const control = "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"
	for _, name := range []string{
		"backup/2026/11/26/a/control.txt", "backup/2026/11/26/b/control.txt",
		"backup/2026/11/26/c/control.txt", "backup/2026/11/26/d/control.txt",
	} {
		entry, err := archive.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := entry.Write([]byte(control)); err != nil {
			t.Fatal(err)
		}
	}
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	report, err := New().Scan(context.Background(), path, nil)
	if err != nil || len(report.Findings) != 4 {
		t.Fatalf("want four control occurrences, got %+v (%v)", report.Findings, err)
	}
	hash := report.Findings[0].SHA256
	for _, finding := range report.Findings {
		if finding.Risk != 100 || hash == "" || finding.SHA256 != hash {
			t.Fatalf("inconsistent evidence for duplicated test content: %+v", finding)
		}
	}
}
