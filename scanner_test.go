package engine

import (
	"archive/zip"
	"context"
	"crypto/sha256"
	"debug/pe"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEICAR(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.txt")
	eicar := "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"
	if e := os.WriteFile(file, []byte(eicar), 0600); e != nil {
		t.Fatal(e)
	}
	r, e := New().Scan(context.Background(), dir, nil)
	if e != nil {
		t.Fatal(e)
	}
	if len(r.Findings) != 1 || r.Findings[0].Risk != 100 {
		t.Fatalf("EICAR não detectado: %+v", r)
	}
}
func TestClean(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "hello.txt"), []byte("olá mundo, documento simples"), 0600)
	r, e := New().Scan(context.Background(), dir, nil)
	if e != nil || len(r.Findings) != 0 {
		t.Fatalf("falso positivo: %+v %v", r, e)
	}
}
func TestScriptHeuristic(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "invoice.pdf.ps1")
	os.WriteFile(p, []byte("powershell -EncodedCommand abc; Invoke-Expression $x"), 0600)
	r, e := New().Scan(context.Background(), dir, nil)
	if e != nil || len(r.Findings) != 1 || r.Findings[0].Risk < 80 {
		t.Fatalf("heurística falhou: %+v %v", r, e)
	}
}
func TestIOC(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "known.bin")
	os.WriteFile(p, []byte("exemplo"), 0600)
	sum := sha256.Sum256([]byte("exemplo"))
	ioc := filepath.Join(t.TempDir(), "ioc.txt")
	os.WriteFile(ioc, []byte(hex.EncodeToString(sum[:])+" Teste IOC\n"), 0600)
	s := New()
	if e := s.LoadIOC(ioc); e != nil {
		t.Fatal(e)
	}
	r, e := s.Scan(context.Background(), dir, nil)
	if e != nil || len(r.Findings) != 1 || r.Findings[0].Risk != 100 {
		t.Fatalf("IOC falhou: %+v %v", r, e)
	}
}
func TestQuarantineRestore(t *testing.T) {
	dir := t.TempDir()
	q := &Quarantine{Dir: filepath.Join(dir, "q")}
	os.MkdirAll(q.Dir, 0700)
	p := filepath.Join(dir, "sample.txt")
	os.WriteFile(p, []byte("hello"), 0600)
	item, e := q.Add(p)
	if e != nil {
		t.Fatal(e)
	}
	if _, e = os.Stat(p); !os.IsNotExist(e) {
		t.Fatal("original não removido")
	}
	items, e := q.List()
	if e != nil || len(items) != 1 {
		t.Fatal(items, e)
	}
	if e = q.Restore(item.ID); e != nil {
		t.Fatal(e)
	}
	b, e := os.ReadFile(p)
	if e != nil || strings.TrimSpace(string(b)) != "hello" {
		t.Fatal("restauração inválida")
	}
}

func TestFalsePositivePowerShellDocumentation(t *testing.T) {
	// Documentation mentioning one risky token is not sufficient by itself.
	dir := t.TempDir()
	p := filepath.Join(dir, "guide.txt")
	os.WriteFile(p, []byte("Use Invoke-Expression only in trusted scripts; this is a tutorial."), 0600)
	r, err := New().Scan(context.Background(), dir, nil)
	if err != nil || len(r.Findings) != 0 {
		t.Fatalf("unexpected result: %+v %v", r, err)
	}
}

func TestCancellationBeforeFileScan(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "a.bin")
	os.WriteFile(p, []byte("harmless"), 0600)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	r, err := New().Scan(ctx, p, nil)
	if err != nil || !r.Cancelled || r.Files != 0 {
		t.Fatalf("cancellation: %+v %v", r, err)
	}
}

func TestPEFileWithoutHeader(t *testing.T) {
	if analyzePE("document.pdf", []byte("not a portable executable")) != nil {
		t.Fatal("false PE")
	}
}

func TestArchiveEntryLimits(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "sample.zip")
	f, err := os.Create(p)
	if err != nil {
		t.Fatal(err)
	}
	w := zip.NewWriter(f)
	z, err := w.Create("invoice.pdf.ps1")
	if err != nil {
		t.Fatal(err)
	}
	z.Write([]byte("powershell -EncodedCommand A; Invoke-Expression B"))
	w.Close()
	f.Close()
	s := New()
	r, err := s.Scan(context.Background(), p, nil)
	if err != nil || len(r.Findings) != 1 {
		t.Fatalf("zip detection: %+v %v", r, err)
	}
	s.MaxArchiveTotal = 1
	r, err = s.Scan(context.Background(), p, nil)
	if err != nil || len(r.Findings) != 0 {
		t.Fatalf("zip limit: %+v %v", r, err)
	}
}

func TestConcurrentScans(t *testing.T) {
	dir := t.TempDir()
	for i := 0; i < 75; i++ {
		os.WriteFile(filepath.Join(dir, fmt.Sprintf("clean-%03d.txt", i)), []byte("safe text"), 0600)
	}
	for i := 0; i < 5; i++ {
		os.WriteFile(filepath.Join(dir, fmt.Sprintf("invoice-%03d.pdf.ps1", i)), []byte("-EncodedCommand x; Invoke-Expression y"), 0600)
	}
	s := New()
	for n := 0; n < 10; n++ {
		r, err := s.Scan(context.Background(), dir, nil)
		if err != nil || r.Files != 80 || len(r.Findings) != 5 {
			t.Fatalf("pass %d: %+v %v", n, r, err)
		}
	}
}

func Test307BoundedImports(t *testing.T){
 for _,width:=range []int{4,8}{
  for _,scenario:=range []string{"valid","cross-section","fallback","truncated","bad-thunk","ordinal","bad-dir","empty","no-terminator","overflow"}{
   t.Run(scenario+string(rune('0'+width)),func(t *testing.T){
    b:=make([]byte,512); f:=&pe.File{Sections:[]*pe.Section{{SectionHeader:pe.SectionHeader{VirtualAddress:0x1000,Size:128,Offset:32}},{SectionHeader:pe.SectionHeader{VirtualAddress:0x2000,Size:256,Offset:256}}}}
    d:=pe.DataDirectory{VirtualAddress:0x1000,Size:40}; thunk:=uint32(0x1040); off:=96; sym:=uint32(0x1060); symOff:=128
    if scenario=="cross-section"{thunk=0x2000;off=256;sym=0x2040;symOff=320}
    binary.LittleEndian.PutUint32(b[32:],thunk);binary.LittleEndian.PutUint32(b[44:],0x1030);binary.LittleEndian.PutUint32(b[48:],thunk);copy(b[80:],"KERNEL32.dll\x00");copy(b[symOff+2:],"WriteProcessMemory\x00")
    if width==4{binary.LittleEndian.PutUint32(b[off:],sym)}else{binary.LittleEndian.PutUint64(b[off:],uint64(sym))}
    wantErr:=false; want:=1
    switch scenario{case "fallback":binary.LittleEndian.PutUint32(b[32:],0);case "truncated":b=b[:100];wantErr=true;case "bad-thunk":binary.LittleEndian.PutUint32(b[32:],0x9000);wantErr=true;case "ordinal":if width==4{binary.LittleEndian.PutUint32(b[off:],0x80000001)}else{binary.LittleEndian.PutUint64(b[off:],0x8000000000000001)};want=0;case "bad-dir":d.Size=19;wantErr=true;case "empty":d=pe.DataDirectory{};want=0;case "no-terminator":d.Size=20;wantErr=true;case "overflow":d.VirtualAddress=0xfffffff0;wantErr=true}
    if width==4{h:=&pe.OptionalHeader32{NumberOfRvaAndSizes:16};h.DataDirectory[1]=d;f.OptionalHeader=h}else{h:=&pe.OptionalHeader64{NumberOfRvaAndSizes:16};h.DataDirectory[1]=d;f.OptionalHeader=h}
    got,err:=boundedImportedSymbols(f,b);if (err!=nil)!=wantErr{t.Fatalf("got %v %v want error=%v",got,err,wantErr)};if !wantErr&&(len(got)!=want||(want==1&&got[0]!="WriteProcessMemory:KERNEL32.dll")){t.Fatalf("got %v",got)}
   })
  }
 }
}


func TestQuarantineRejectsTamperedCopy(t *testing.T) {
 dir := t.TempDir()
 q := &Quarantine{Dir: filepath.Join(dir, "quarantine")}
 path := filepath.Join(dir, "fixture.txt")
 data := []byte("benign")
 if err := os.WriteFile(path, data, 0600); err != nil { t.Fatal(err) }
 item, err := q.Add(path)
 if err != nil { t.Fatal(err) }
 hash := sha256.Sum256(data)
 if item.SHA256 != hex.EncodeToString(hash[:]) { t.Fatalf("incorrect digest: %q", item.SHA256) }
 stored := filepath.Join(q.Dir, item.ID+".bin")
 if err := os.WriteFile(stored, []byte("alterd"), 0600); err != nil { t.Fatal(err) }
 if err := q.Restore(item.ID); err == nil { t.Fatal("restored modified file") }
 if _, err := os.Stat(path); !os.IsNotExist(err) { t.Fatalf("restored path remains: %v", err) }
 if _, err := os.Stat(stored); err != nil { t.Fatal(err) }
 if err := os.WriteFile(stored, data, 0600); err != nil { t.Fatal(err) }
 if err := q.Restore(item.ID); err != nil { t.Fatal(err) }
}
