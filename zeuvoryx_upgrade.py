#!/usr/bin/env python3
"""ZEUVORYX 3.0.1 experimental upgrade, applied AFTER existing workflow code generation.
No network calls. Abort on unexpected source. Use a separate Git branch.
"""
from pathlib import Path
import argparse, shutil, sys, re

EMBEDDED_GO = {'scan_upgrade.go': 'package engine\n'
                    '\n'
                    'import (\n'
                    '\t"context"\n'
                    '\t"crypto/sha256"\n'
                    '\t"encoding/hex"\n'
                    '\t"errors"\n'
                    '\t"io"\n'
                    '\t"os"\n'
                    '\t"path/filepath"\n'
                    '\t"strings"\n'
                    '\t"sync"\n'
                    ')\n'
                    '\n'
                    'const ScanEngineVersion = "3.0.1-test"\n'
                    'const maxSampleBytes int64 = 32 << 20\n'
                    '\n'
                    'var hashBuffers = sync.Pool{New: func() any { return make([]byte, 128<<10) }}\n'
                    '\n'
                    'type upgradeReader struct {\n'
                    '\tctx context.Context\n'
                    '\tr   io.Reader\n'
                    '}\n'
                    '\n'
                    'func (r upgradeReader) Read(b []byte) (int, error) {\n'
                    '\tif e := r.ctx.Err(); e != nil {\n'
                    '\t\treturn 0, e\n'
                    '\t}\n'
                    '\treturn r.r.Read(b)\n'
                    '}\n'
                    '\n'
                    '// Allocate only the required prefix; hash the whole file in one pass, even with a '
                    'short prefix.\n'
                    'func readScanContent(ctx context.Context, r io.Reader, size, limit int64) ([]byte, '
                    'string, error) {\n'
                    '\tif size < 0 {\n'
                    '\t\treturn nil, "", errors.New("tamanho inválido")\n'
                    '\t}\n'
                    '\tif err := ctx.Err(); err != nil {\n'
                    '\t\treturn nil, "", err\n'
                    '\t}\n'
                    '\tif limit <= 0 || limit > maxSampleBytes {\n'
                    '\t\tlimit = maxSampleBytes\n'
                    '\t}\n'
                    '\tn := size\n'
                    '\tif n > limit {\n'
                    '\t\tn = limit\n'
                    '\t}\n'
                    '\tsample := make([]byte, int(n))\n'
                    '\trd := upgradeReader{ctx, r}\n'
                    '\th := sha256.New()\n'
                    '\tif _, e := io.ReadFull(rd, sample); e != nil {\n'
                    '\t\treturn nil, "", e\n'
                    '\t}\n'
                    '\t_, _ = h.Write(sample)\n'
                    '\tbuf := hashBuffers.Get().([]byte)\n'
                    '\tdefer hashBuffers.Put(buf)\n'
                    '\ttail, e := io.CopyBuffer(h, io.LimitReader(rd, size-n), buf)\n'
                    '\tif e != nil {\n'
                    '\t\treturn nil, "", e\n'
                    '\t}\n'
                    '\tif tail != size-n {\n'
                    '\t\treturn nil, "", errors.New("arquivo mudou de tamanho durante a leitura")\n'
                    '\t}\n'
                    '\tvar extra [1]byte\n'
                    '\tif count, err := io.ReadFull(rd, extra[:]); count != 0 || err != io.EOF {\n'
                    '\t\tif err != nil && err != io.EOF {\n'
                    '\t\t\treturn nil, "", err\n'
                    '\t\t}\n'
                    '\t\treturn nil, "", errors.New("arquivo cresceu durante a leitura")\n'
                    '\t}\n'
                    '\tif e = ctx.Err(); e != nil {\n'
                    '\t\treturn nil, "", e\n'
                    '\t}\n'
                    '\treturn sample, hex.EncodeToString(h.Sum(nil)), nil\n'
                    '}\n'
                    '\n'
                    'func stableFile(path string, before, after os.FileInfo) bool {\n'
                    '\tif before == nil || after == nil || !after.Mode().IsRegular() || !os.SameFile(before, '
                    'after) || before.Size() != after.Size() || !before.ModTime().Equal(after.ModTime()) {\n'
                    '\t\treturn false\n'
                    '\t}\n'
                    '\tcurrent, e := os.Lstat(path)\n'
                    '\treturn e == nil && current.Mode().IsRegular() && os.SameFile(after, current) && '
                    'current.Size() == after.Size() && current.ModTime().Equal(after.ModTime())\n'
                    '}\n'
                    '\n'
                    'func evidenceFor(reasons []string) (string, bool) {\n'
                    '\tfor _, r := range reasons {\n'
                    '\t\tif strings.HasPrefix(r, "Correspondência SHA-256: ") {\n'
                    '\t\t\treturn "ioc_match", true\n'
                    '\t\t}\n'
                    '\t}\n'
                    '\tfor _, r := range reasons {\n'
                    '\t\tif r == "Assinatura de teste EICAR detectada (não é malware real)" {\n'
                    '\t\t\treturn "test_eicar", true\n'
                    '\t\t}\n'
                    '\t}\n'
                    '\treturn "heuristic", false\n'
                    '}\n'
                    '\n'
                    'func isUnderPath(root, path string) bool {\n'
                    '\trel, e := filepath.Rel(root, path)\n'
                    '\treturn e == nil && rel != ".." && !strings.HasPrefix(rel, '
                    '".."+string(os.PathSeparator))\n'
                    '}\n'
                    '\n'
                    '// Resolve both paths and refuse ambiguous paths. No filename or directory is a malware '
                    'allowlist.\n'
                    'func checkProtected(path string, roots []string) error {\n'
                    '\tabsolute, e := filepath.Abs(path)\n'
                    '\tif e != nil {\n'
                    '\t\treturn e\n'
                    '\t}\n'
                    '\tresolved, e := filepath.EvalSymlinks(absolute)\n'
                    '\tif e != nil {\n'
                    '\t\treturn e\n'
                    '\t}\n'
                    '\tfor _, root := range roots {\n'
                    '\t\tif root == "" {\n'
                    '\t\t\tcontinue\n'
                    '\t\t}\n'
                    '\t\troot, e = filepath.Abs(root)\n'
                    '\t\tif e != nil {\n'
                    '\t\t\treturn e\n'
                    '\t\t}\n'
                    '\t\tif isUnderPath(root, absolute) || isUnderPath(root, resolved) {\n'
                    '\t\t\treturn errors.New("isolamento manual bloqueado: diretório protegido do sistema")\n'
                    '\t\t}\n'
                    '\t\treal, err := filepath.EvalSymlinks(root)\n'
                    '\t\tif err == nil && isUnderPath(real, resolved) {\n'
                    '\t\t\treturn errors.New("isolamento manual bloqueado: diretório protegido do sistema")\n'
                    '\t\t}\n'
                    '\t}\n'
                    '\treturn nil\n'
                    '}\n'
                    '\n'
                    'func quarantineHash(path string, expected string) error {\n'
                    '\tif len(expected) != 64 {\n'
                    '\t\treturn errors.New("SHA-256 obrigatório")\n'
                    '\t}\n'
                    '\tif _, e := hex.DecodeString(expected); e != nil {\n'
                    '\t\treturn e\n'
                    '\t}\n'
                    '\tf, e := os.Open(path)\n'
                    '\tif e != nil {\n'
                    '\t\treturn e\n'
                    '\t}\n'
                    '\tdefer f.Close()\n'
                    '\tbefore, e := f.Stat()\n'
                    '\tif e != nil {\n'
                    '\t\treturn e\n'
                    '\t}\n'
                    '\tif !before.Mode().IsRegular() {\n'
                    '\t\treturn errors.New("arquivo não regular")\n'
                    '\t}\n'
                    '\th := sha256.New()\n'
                    '\tn, e := io.Copy(h, f)\n'
                    '\tif e != nil {\n'
                    '\t\treturn e\n'
                    '\t}\n'
                    '\tafter, e := f.Stat()\n'
                    '\tif e != nil {\n'
                    '\t\treturn e\n'
                    '\t}\n'
                    '\tif n != before.Size() || !stableFile(path, before, after) || '
                    '!strings.EqualFold(hex.EncodeToString(h.Sum(nil)), expected) {\n'
                    '\t\treturn errors.New("arquivo alterado ou SHA-256 divergente; refaça a análise")\n'
                    '\t}\n'
                    '\treturn nil\n'
                    '}\n'
                    '\n'
                    '// Derive the expected authorization from a fresh scan finding, not a UI-provided '
                    'score.\n'
                    'func (q *Quarantine) AddVerified(hit Finding) (QuarantineItem, error) {\n'
                    '\t_, ok := evidenceFor(hit.Reasons)\n'
                    '\tif !ok || !hit.QuarantineEligible || strings.Contains(hit.Path, " :: ") {\n'
                    '\t\treturn QuarantineItem{}, errors.New("indício heurístico: revisão necessária; '
                    'quarentena indisponível neste teste")\n'
                    '\t}\n'
                    '\treturn q.addWithHash(hit.Path, hit.SHA256)\n'
                    '}\n'
                    '\n'
                    'func (r *Report) Summarize() {\n'
                    '\tr.EngineVersion = ScanEngineVersion\n'
                    '\tr.CoverageNote = "SHA-256 integral dos arquivos lidos; heurística limitada ao prefixo '
                    'de até 32 MiB; ZIP limitado por entradas e volume; conteúdo não acessível, formatos não '
                    'suportados e memória não são inspecionados por este motor."\n'
                    '\tif !r.Started.IsZero() && !r.Finished.IsZero() {\n'
                    '\t\tr.DurationMS = r.Finished.Sub(r.Started).Milliseconds()\n'
                    '\t}\n'
                    '\tr.Partial = r.Cancelled || r.Skipped > 0 || len(r.Errors) > 0\n'
                    '\tseen := map[string]struct{}{}\n'
                    '\tr.HeuristicAlerts = 0\n'
                    '\tr.IOCMatches = 0\n'
                    '\tr.EICARTests = 0\n'
                    '\tfor _, f := range r.Findings {\n'
                    '\t\tif f.SHA256 != "" {\n'
                    '\t\t\tseen[f.SHA256] = struct{}{}\n'
                    '\t\t}\n'
                    '\t\tswitch f.Evidence {\n'
                    '\t\tcase "ioc_match":\n'
                    '\t\t\tr.IOCMatches++\n'
                    '\t\tcase "test_eicar":\n'
                    '\t\t\tr.EICARTests++\n'
                    '\t\tdefault:\n'
                    '\t\t\tr.HeuristicAlerts++\n'
                    '\t\t}\n'
                    '\t}\n'
                    '\tr.UniqueFindingHashes = len(seen)\n'
                    '}\n',
 'scan_upgrade_integration_test.go': 'package engine\n'
                                     '\n'
                                     'import (\n'
                                     '\t"context"\n'
                                     '\t"crypto/sha256"\n'
                                     '\t"encoding/hex"\n'
                                     '\t"os"\n'
                                     '\t"path/filepath"\n'
                                     '\t"testing"\n'
                                     ')\n'
                                     '\n'
                                     'func TestUpgradeIntegratedIOCAndRestore(t *testing.T) {\n'
                                     '\tdir := t.TempDir()\n'
                                     '\tp := filepath.Join(dir, "sample.bin")\n'
                                     '\tdata := []byte("harmless test marker")\n'
                                     '\tif e := os.WriteFile(p, data, 0600); e != nil {\n'
                                     '\t\tt.Fatal(e)\n'
                                     '\t}\n'
                                     '\tsum := sha256.Sum256(data)\n'
                                     '\ts := New()\n'
                                     '\ts.MaxFileBytes = 2\n'
                                     '\ts.IOC[hex.EncodeToString(sum[:])] = "synthetic local indicator"\n'
                                     '\tr, e := s.Scan(context.Background(), p, nil)\n'
                                     '\tif e != nil || r.Files != 1 || len(r.Findings) != 1 || r.IOCMatches '
                                     '!= 1 || !r.Findings[0].QuarantineEligible {\n'
                                     '\t\tt.Fatalf("%+v %v", r, e)\n'
                                     '\t}\n'
                                     '\tq := &Quarantine{Dir: filepath.Join(dir, "isolation")}\n'
                                     '\thit := r.Findings[0]\n'
                                     '\tif e := os.WriteFile(p, []byte("changed content"), 0600); e != nil '
                                     '{\n'
                                     '\t\tt.Fatal(e)\n'
                                     '\t}\n'
                                     '\tif _, e := q.AddVerified(hit); e == nil {\n'
                                     '\t\tt.Fatal("changed file allowed")\n'
                                     '\t}\n'
                                     '\tif e := os.WriteFile(p, data, 0600); e != nil {\n'
                                     '\t\tt.Fatal(e)\n'
                                     '\t}\n'
                                     '\titem, e := q.AddVerified(hit)\n'
                                     '\tif e != nil {\n'
                                     '\t\tt.Fatal(e)\n'
                                     '\t}\n'
                                     '\tif item.SHA256 != hit.SHA256 {\n'
                                     '\t\tt.Fatal("metadata hash missing")\n'
                                     '\t}\n'
                                     '\tif _, e := os.Stat(p); !os.IsNotExist(e) {\n'
                                     '\t\tt.Fatal("original not isolated")\n'
                                     '\t}\n'
                                     '\tif e := q.Restore(item.ID); e != nil {\n'
                                     '\t\tt.Fatal(e)\n'
                                     '\t}\n'
                                     '\tb, e := os.ReadFile(p)\n'
                                     '\tif e != nil || string(b) != string(data) {\n'
                                     '\t\tt.Fatal("restoration failed", e)\n'
                                     '\t}\n'
                                     '}\n'
                                     'func TestUpgradeIntegratedAdvisory(t *testing.T) {\n'
                                     '\tf := analyze("example.ps1", []byte("Invoke-Expression example"), "", '
                                     'nil)\n'
                                     '\tif f == nil || f.Evidence != "heuristic" || f.QuarantineEligible {\n'
                                     '\t\tt.Fatalf("%+v", f)\n'
                                     '\t}\n'
                                     '}\n'
                                     'func TestUpgradeProgressMonotonic(t *testing.T) {\n'
                                     '\tdir := t.TempDir()\n'
                                     '\tfor i := 0; i < 100; i++ {\n'
                                     '\t\tp := filepath.Join(dir, hex.EncodeToString([]byte{byte(i)}))\n'
                                     '\t\tif e := os.WriteFile(p, []byte("x"), 0600); e != nil {\n'
                                     '\t\t\tt.Fatal(e)\n'
                                     '\t\t}\n'
                                     '\t}\n'
                                     '\tlast := 0\n'
                                     '\tbad := false\n'
                                     '\tr, e := New().Scan(context.Background(), dir, func(n int, b int64) '
                                     '{\n'
                                     '\t\tif n < last || int64(n) != b {\n'
                                     '\t\t\tbad = true\n'
                                     '\t\t}\n'
                                     '\t\tlast = n\n'
                                     '\t})\n'
                                     '\tif e != nil || bad || r.Files != 100 || last != 100 {\n'
                                     '\t\tt.Fatalf("%+v %v %d", r, e, last)\n'
                                     '\t}\n'
                                     '}\n',
 'scan_upgrade_other.go': '//go:build !windows\n'
                          '\n'
                          'package engine\n'
                          '\n'
                          'func protectedRoots() ([]string, error)                     { return nil, nil }\n'
                          'func maybeSignature(path string, data []byte, hit *Finding) {}\n',
 'scan_upgrade_test.go': 'package engine\n'
                         '\n'
                         'import (\n'
                         '\t"bytes"\n'
                         '\t"context"\n'
                         '\t"crypto/sha256"\n'
                         '\t"encoding/hex"\n'
                         '\t"errors"\n'
                         '\t"os"\n'
                         '\t"path/filepath"\n'
                         '\t"testing"\n'
                         ')\n'
                         '\n'
                         'func TestUpgradeFullHash(t *testing.T) {\n'
                         '\tfor _, size := range []int{0, 1, 1023, 1024, 1025, 300000} {\n'
                         '\t\tb := bytes.Repeat([]byte{42}, size)\n'
                         '\t\tp, h, e := readScanContent(context.Background(), bytes.NewReader(b), '
                         'int64(size), 1024)\n'
                         '\t\tsum := sha256.Sum256(b)\n'
                         '\t\twant := size\n'
                         '\t\tif want > 1024 {\n'
                         '\t\t\twant = 1024\n'
                         '\t\t}\n'
                         '\t\tif e != nil || len(p) != want || h != hex.EncodeToString(sum[:]) {\n'
                         '\t\t\tt.Fatalf("size %d: %d %v", size, len(p), e)\n'
                         '\t\t}\n'
                         '\t}\n'
                         '}\n'
                         'func TestUpgradeReadChanged(t *testing.T) {\n'
                         '\tfor _, size := range []int64{1, 4} {\n'
                         '\t\tif _, _, e := readScanContent(context.Background(), '
                         'bytes.NewReader([]byte("abc")), size, 2); e == nil {\n'
                         '\t\t\tt.Fatal("accepted inconsistent size")\n'
                         '\t\t}\n'
                         '\t}\n'
                         '}\n'
                         'func TestUpgradeCancel(t *testing.T) {\n'
                         '\tctx, c := context.WithCancel(context.Background())\n'
                         '\tc()\n'
                         '\t_, _, e := readScanContent(ctx, bytes.NewReader(nil), 0, 1)\n'
                         '\tif !errors.Is(e, context.Canceled) {\n'
                         '\t\tt.Fatal(e)\n'
                         '\t}\n'
                         '}\n'
                         'func TestUpgradeEvidence(t *testing.T) {\n'
                         '\tfor _, x := range []struct {\n'
                         '\t\tr    []string\n'
                         '\t\tkind string\n'
                         '\t\tok   bool\n'
                         '\t}{{[]string{"PE: combinação de APIs de injeção remota (indicador, não prova)"}, '
                         '"heuristic", false}, {[]string{"Assinatura de teste EICAR detectada (não é malware '
                         'real)"}, "test_eicar", true}, {[]string{"Correspondência SHA-256: teste"}, '
                         '"ioc_match", true}} {\n'
                         '\t\ta, b := evidenceFor(x.r)\n'
                         '\t\tif a != x.kind || b != x.ok {\n'
                         '\t\t\tt.Fatal(a, b)\n'
                         '\t\t}\n'
                         '\t}\n'
                         '}\n'
                         'func TestUpgradeProtected(t *testing.T) {\n'
                         '\td := t.TempDir()\n'
                         '\troot := filepath.Join(d, "Windows")\n'
                         '\tif e := os.Mkdir(root, 0700); e != nil {\n'
                         '\t\tt.Fatal(e)\n'
                         '\t}\n'
                         '\tp := filepath.Join(root, "kernel32.dll")\n'
                         '\toutside := filepath.Join(d, "clean.txt")\n'
                         '\tfor _, x := range []string{p, outside} {\n'
                         '\t\tif e := os.WriteFile(x, []byte("x"), 0600); e != nil {\n'
                         '\t\t\tt.Fatal(e)\n'
                         '\t\t}\n'
                         '\t}\n'
                         '\tif checkProtected(p, []string{root}) == nil {\n'
                         '\t\tt.Fatal("protected path accepted")\n'
                         '\t}\n'
                         '\tif e := checkProtected(outside, []string{root}); e != nil {\n'
                         '\t\tt.Fatal(e)\n'
                         '\t}\n'
                         '\talias := filepath.Join(d, "alias")\n'
                         '\tif os.Symlink(root, alias) == nil && checkProtected(filepath.Join(alias, '
                         '"kernel32.dll"), []string{root}) == nil {\n'
                         '\t\tt.Fatal("symlink bypass")\n'
                         '\t}\n'
                         '}\n'
                         'func TestUpgradeHashMismatch(t *testing.T) {\n'
                         '\tp := filepath.Join(t.TempDir(), "sample.txt")\n'
                         '\tif e := os.WriteFile(p, []byte("abc"), 0600); e != nil {\n'
                         '\t\tt.Fatal(e)\n'
                         '\t}\n'
                         '\th := sha256.Sum256([]byte("abc"))\n'
                         '\tif e := quarantineHash(p, hex.EncodeToString(h[:])); e != nil {\n'
                         '\t\tt.Fatal(e)\n'
                         '\t}\n'
                         '\th[0] ^= 1\n'
                         '\tif quarantineHash(p, hex.EncodeToString(h[:])) == nil {\n'
                         '\t\tt.Fatal("wrong hash accepted")\n'
                         '\t}\n'
                         '}\n'
                         'func TestUpgradeSummary(t *testing.T) {\n'
                         '\tr := Report{Skipped: 1, Findings: []Finding{{SHA256: "a", Evidence: '
                         '"heuristic"}, {SHA256: "a", Evidence: "heuristic"}, {SHA256: "b", Evidence: '
                         '"test_eicar"}}}\n'
                         '\tr.Summarize()\n'
                         '\tif !r.Partial || r.UniqueFindingHashes != 2 || r.HeuristicAlerts != 2 || '
                         'r.EICARTests != 1 || r.EngineVersion == "" {\n'
                         '\t\tt.Fatalf("%+v", r)\n'
                         '\t}\n'
                         '}\n'
                         'func TestUpgradeQuarantineHeuristic(t *testing.T) {\n'
                         '\tq := &Quarantine{Dir: t.TempDir()}\n'
                         '\tp := filepath.Join(t.TempDir(), "safe.txt")\n'
                         '\tif e := os.WriteFile(p, []byte("safe"), 0600); e != nil {\n'
                         '\t\tt.Fatal(e)\n'
                         '\t}\n'
                         '\tif _, e := q.AddVerified(Finding{Path: p, Risk: 100, QuarantineEligible: true, '
                         'Reasons: []string{"heuristic"}}); e == nil {\n'
                         '\t\tt.Fatal("heuristic allowed")\n'
                         '\t}\n'
                         '\tif _, e := os.Stat(p); e != nil {\n'
                         '\t\tt.Fatal(e)\n'
                         '\t}\n'
                         '}\n'
                         'func BenchmarkUpgradeHash(b *testing.B) {\n'
                         '\tdata := bytes.Repeat([]byte{23}, 1<<20)\n'
                         '\tb.SetBytes(int64(len(data)))\n'
                         '\tb.ReportAllocs()\n'
                         '\tfor i := 0; i < b.N; i++ {\n'
                         '\t\t_, _, e := readScanContent(context.Background(), bytes.NewReader(data), '
                         'int64(len(data)), 64<<10)\n'
                         '\t\tif e != nil {\n'
                         '\t\t\tb.Fatal(e)\n'
                         '\t\t}\n'
                         '\t}\n'
                         '}\n',
 'scan_upgrade_windows.go': '//go:build windows\n'
                            '\n'
                            'package engine\n'
                            '\n'
                            'import (\n'
                            '\t"errors"\n'
                            '\t"os"\n'
                            '\t"path/filepath"\n'
                            '\t"runtime"\n'
                            '\t"strings"\n'
                            '\t"sync"\n'
                            '\t"syscall"\n'
                            '\t"unsafe"\n'
                            ')\n'
                            '\n'
                            'var systemKernel = syscall.NewLazyDLL("kernel32.dll")\n'
                            'var trustOnce sync.Once\n'
                            'var trustProc *syscall.LazyProc\n'
                            '\n'
                            'func wintrustProc() *syscall.LazyProc {\n'
                            '\ttrustOnce.Do(func() {\n'
                            '\t\tvar buf [32768]uint16\n'
                            '\t\tgetDir := systemKernel.NewProc("GetSystemDirectoryW")\n'
                            '\t\tif getDir.Find() != nil {\n'
                            '\t\t\treturn\n'
                            '\t\t}\n'
                            '\t\tn, _, _ := getDir.Call(uintptr(unsafe.Pointer(&buf[0])), '
                            'uintptr(len(buf)))\n'
                            '\t\tif n == 0 || n >= uintptr(len(buf)) {\n'
                            '\t\t\treturn\n'
                            '\t\t}\n'
                            '\t\tdir := syscall.UTF16ToString(buf[:n])\n'
                            '\t\tif !filepath.IsAbs(dir) {\n'
                            '\t\t\treturn\n'
                            '\t\t}\n'
                            '\t\tproc := syscall.NewLazyDLL(filepath.Join(dir, '
                            '"wintrust.dll")).NewProc("WinVerifyTrust")\n'
                            '\t\tif proc.Find() == nil {\n'
                            '\t\t\ttrustProc = proc\n'
                            '\t\t}\n'
                            '\t})\n'
                            '\treturn trustProc\n'
                            '}\n'
                            '\n'
                            'func protectedRoots() ([]string, error) {\n'
                            '\tvar buf [32768]uint16\n'
                            '\tproc := systemKernel.NewProc("GetSystemWindowsDirectoryW")\n'
                            '\tif e := proc.Find(); e != nil {\n'
                            '\t\treturn nil, e\n'
                            '\t}\n'
                            '\tn, _, _ := proc.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))\n'
                            '\tif n == 0 || n >= uintptr(len(buf)) {\n'
                            '\t\treturn nil, errors.New("não foi possível localizar diretório protegido do '
                            'Windows")\n'
                            '\t}\n'
                            '\twin := syscall.UTF16ToString(buf[:n])\n'
                            '\tif !filepath.IsAbs(win) {\n'
                            '\t\treturn nil, errors.New("diretório do Windows inválido")\n'
                            '\t}\n'
                            '\troots := []string{win, filepath.Join(filepath.VolumeName(win)+"\\\\", '
                            '"Program Files"), filepath.Join(filepath.VolumeName(win)+"\\\\", "Program Files '
                            '(x86)"), filepath.Join(filepath.VolumeName(win)+"\\\\", "ProgramData", '
                            '"Microsoft", "Windows Defender")}\n'
                            '\tfor _, k := range []string{"ProgramFiles", "ProgramFiles(x86)", '
                            '"ProgramW6432"} {\n'
                            '\t\tif p := os.Getenv(k); filepath.IsAbs(p) {\n'
                            '\t\t\troots = append(roots, p)\n'
                            '\t\t}\n'
                            '\t}\n'
                            '\tif p := os.Getenv("ProgramData"); filepath.IsAbs(p) {\n'
                            '\t\troots = append(roots, filepath.Join(p, "Microsoft", "Windows Defender"), '
                            'filepath.Join(p, "Microsoft", "Windows Defender Advanced Threat Protection"))\n'
                            '\t}\n'
                            '\treturn roots, nil\n'
                            '}\n'
                            '\n'
                            'type trustGUID struct {\n'
                            '\tA    uint32\n'
                            '\tB, C uint16\n'
                            '\tD    [8]byte\n'
                            '}\n'
                            'type trustFile struct {\n'
                            '\tSize    uint32\n'
                            '\tPath    *uint16\n'
                            '\tHandle  syscall.Handle\n'
                            '\tSubject *trustGUID\n'
                            '}\n'
                            'type trustData struct {\n'
                            '\tSize                   uint32\n'
                            '\tPolicy, SIP            uintptr\n'
                            '\tUI, Revocation, Choice uint32\n'
                            '\tFile                   *trustFile\n'
                            '\tStateAction            uint32\n'
                            '\tState                  syscall.Handle\n'
                            '\tURL                    *uint16\n'
                            '\tFlags, UIContext       uint32\n'
                            '\tSignatureSettings      uintptr\n'
                            '}\n'
                            '\n'
                            '// File-choice Authenticode verification, offline/cache only. Catalog-only and '
                            'other\n'
                            '// unverifiable cases remain unknown; a valid signature never suppresses an '
                            'IOC.\n'
                            'func signatureStatus(path string) string {\n'
                            '\tproc := wintrustProc()\n'
                            '\tif proc == nil {\n'
                            '\t\treturn "unverified"\n'
                            '\t}\n'
                            '\tp, e := syscall.UTF16PtrFromString(path)\n'
                            '\tif e != nil {\n'
                            '\t\treturn "unverified"\n'
                            '\t}\n'
                            '\tguid := trustGUID{0x00aac56b, 0xcd44, 0x11d0, [8]byte{0x8c, 0xc2, 0x00, 0xc0, '
                            '0x4f, 0xc2, 0x95, 0xee}}\n'
                            '\tf := trustFile{Path: p}\n'
                            '\tf.Size = uint32(unsafe.Sizeof(f))\n'
                            '\t// WTD_UI_NONE, WTD_REVOKE_NONE, WTD_CHOICE_FILE, WTD_STATEACTION_VERIFY,\n'
                            '\t// WTD_CACHE_ONLY_URL_RETRIEVAL | WTD_REVOCATION_CHECK_NONE.\n'
                            '\td := trustData{UI: 2, Choice: 1, File: &f, StateAction: 1, Flags: 0x1000 | '
                            '0x10}\n'
                            '\td.Size = uint32(unsafe.Sizeof(d))\n'
                            '\tresult, _, _ := proc.Call(^uintptr(0), uintptr(unsafe.Pointer(&guid)), '
                            'uintptr(unsafe.Pointer(&d)))\n'
                            '\td.StateAction = 2\n'
                            '\t_, _, _ = proc.Call(^uintptr(0), uintptr(unsafe.Pointer(&guid)), '
                            'uintptr(unsafe.Pointer(&d)))\n'
                            '\truntime.KeepAlive(p)\n'
                            '\truntime.KeepAlive(f)\n'
                            '\truntime.KeepAlive(d)\n'
                            '\truntime.KeepAlive(guid)\n'
                            '\tif uint32(result) == 0 {\n'
                            '\t\treturn "authenticode_valid_offline"\n'
                            '\t}\n'
                            '\treturn "unverified"\n'
                            '}\n'
                            '\n'
                            'func maybeSignature(path string, data []byte, hit *Finding) {\n'
                            '\tif hit == nil || strings.Contains(path, " :: ") {\n'
                            '\t\treturn\n'
                            '\t}\n'
                            '\text := strings.ToLower(filepath.Ext(path))\n'
                            '\tif len(data) >= 2 && data[0] == \'M\' && data[1] == \'Z\' || ext == ".ps1" || '
                            'ext == ".psm1" || ext == ".psd1" {\n'
                            '\t\thit.Trust = signatureStatus(path)\n'
                            '\t}\n'
                            '}\n'}

MARK = '      - name: Validate source code'
STEP = '''      - name: Apply ZEUVORYX 3.0.1 scanner hardening
        shell: pwsh
        run: |
          python zeuvoryx_upgrade.py --apply
          if ($LASTEXITCODE -ne 0) { throw 'ZEUVORYX upgrade failed' }
          gofmt -w main.go desktop_windows.go engine/scanner.go engine/pe.go engine/quarantine.go engine/scan_upgrade.go engine/scan_upgrade_windows.go engine/scan_upgrade_other.go engine/scan_upgrade_test.go engine/scan_upgrade_integration_test.go
          if ($LASTEXITCODE -ne 0) { throw 'Upgrade formatting failed' }
'''

def one(s,a,b):
    count=s.count(a)
    if count!=1: raise RuntimeError(f'Expected one source anchor, got {count}: {a[:140]!r}')
    return s.replace(a,b,1)

def scanner(s):
    s=one(s,'type Finding struct {','''type Finding struct {
    Evidence string `json:"evidence"`
    Trust string `json:"trust,omitempty"`
    QuarantineEligible bool `json:"quarantine_eligible"`''')
    s=one(s,'type Report struct {','''type Report struct {
    EngineVersion string `json:"engine_version"`
    DurationMS int64 `json:"duration_ms"`
    Partial bool `json:"partial"`
    CoverageNote string `json:"coverage_note"`
    UniqueFindingHashes int `json:"unique_finding_hashes"`
    HeuristicAlerts int `json:"heuristic_alerts"`
    IOCMatches int `json:"ioc_matches"`
    EICARTests int `json:"eicar_tests"`''')
    s=one(s,'return &Finding{Path: path, SHA256: fullHash, Risk: score, Level: tier(score), Reasons: reasons, Size: int64(len(data))}', '''kind,eligible:=evidenceFor(reasons)
    return &Finding{Path: path, SHA256: fullHash, Risk: score, Level: tier(score), Reasons: reasons, Size: int64(len(data)), Evidence:kind,QuarantineEligible:eligible&&!strings.Contains(path," :: ")}''')
    # Existing workflow creates this one-pass implementation; replace it without touching its generation logic.
    a=s.index('\t// Capture the heuristic prefix and compute the whole-file hash in a single pass.')
    b=s.index('\tvar found []Finding',a)
    s=s[:a]+'''\tbefore,e:=f.Stat()
    if e!=nil{return nil,e};if !before.Mode().IsRegular()||before.Size()!=size{return nil,errors.New("arquivo mudou antes da leitura")}
    b,hash,e:=readScanContent(ctx,f,size,s.MaxFileBytes);if e!=nil{return nil,e}
    after,e:=f.Stat();if e!=nil{return nil,e};if !stableFile(path,before,after){return nil,errors.New("arquivo mudou durante a leitura")}
'''+s[b:]
    s=one(s,'hit.Size = size','hit.Size = size\n\t\tmaybeSignature(path,b,hit)')
    # Same open descriptor for ZIP; use total budget of actually decompressed bytes.
    s=one(s,'archive, e := zip.OpenReader(path)','archive, e := zip.NewReader(f,size)')
    s=one(s,'\t\t\tdefer archive.Close()\n','')
    s=one(s,'\t\t\t\ttotalUncompressed += z.UncompressedSize64\n','')
    s=one(s,'data, e := io.ReadAll(io.LimitReader(guardedReader{ctx, r}, entryMax+1))','''remaining:=uint64(totalLimit)-totalUncompressed
                allowed:=entryMax;if uint64(allowed)>remaining{allowed=int64(remaining)}
                data,e:=io.ReadAll(io.LimitReader(guardedReader{ctx,r},allowed+1))''')
    s=one(s,'if e != nil || int64(len(data)) > entryMax {','''if int64(len(data))>allowed{break}
                totalUncompressed+=uint64(len(data))
                if e != nil || uint64(len(data))!=z.UncompressedSize64 {''')
    s=one(s,'\treturn found, nil\n}', '''\tif ctx.Err()!=nil{return nil,ctx.Err()}
    after,e=f.Stat();if e!=nil{return nil,e};if !stableFile(path,before,after){return nil,errors.New("arquivo mudou durante a análise")}
    return found, nil
}''')
    s=one(s,'(Report, error) {\n\tabs, e := filepath.Abs(root)','(report Report, err error) {\n\tabs, e := filepath.Abs(root)')
    s=one(s,'report := Report{Started:', 'report = Report{Started:')
    s=one(s,'defer func() {}()','defer func(){report.Summarize()}()')
    s=one(s,'report.Files = 1','''if e!=nil{report.Skipped=1;report.Errors=append(report.Errors,abs+": "+e.Error());report.Finished=time.Now();return report,e}
        report.Files = 1''')
    s=one(s,'\t\t\t\tn, b := report.Files, report.Bytes\n\t\t\t\tmu.Unlock()\n\t\t\t\tif progress != nil {\n\t\t\t\t\t// Serialize callbacks for thread-safe UI consumers.\n\t\t\t\t\tmu.Lock()\n\t\t\t\t\tprogress(n, b)\n\t\t\t\t\tmu.Unlock()\n\t\t\t\t}', '''\t\t\t\tif progress != nil {progress(report.Files,report.Bytes)}
                mu.Unlock()''')
    # Lstat avoids following a symlink explicitly supplied as root or changed during traversal.
    s=one(s,'info, e := os.Stat(abs)','info, e := os.Lstat(abs)')
    s=one(s,'st, e := os.Stat(path)','st, e := os.Lstat(path)')
    s=one(s,'if e != nil {\n\t\t\t\t\tmu.Lock()\n\t\t\t\t\treport.Skipped++','if e != nil || !st.Mode().IsRegular() {\n\t\t\t\t\tmu.Lock()\n\t\t\t\t\treport.Skipped++')
    return s

def pe(s):
    # Keep combination visible (medium); never treat APIs alone as high-confidence malware.
    s=one(s,'score += 45','score += 30')
    s=one(s,'score += 10','score += 5')
    return s

def quarantine(s):
    s=one(s,'type QuarantineItem struct {','type QuarantineItem struct {\n SHA256 string `json:"sha256,omitempty"`')
    s=one(s,'func (q *Quarantine) Add(path string) (QuarantineItem, error) {','''func (q *Quarantine) Add(path string) (QuarantineItem,error){return q.addWithHash(path,"")}
func (q *Quarantine) addWithHash(path,expected string) (QuarantineItem, error) {''')
    s=one(s,'\tif q.IsQuarantined(abs) {','''    roots,e:=protectedRoots();if e!=nil{return QuarantineItem{},e};if e=checkProtected(abs,roots);e!=nil{return QuarantineItem{},e}
    if expected!=""{if e=quarantineHash(abs,expected);e!=nil{return QuarantineItem{},e}}
\tif q.IsQuarantined(abs) {''')
    s=one(s,'Size:     st.Size(),','Size:     st.Size(),\n SHA256:expected,')
    s=one(s,'\tif e = os.WriteFile(metaPath, meta, 0600); e != nil {','''    if expected!=""{if e=quarantineHash(bin,expected);e!=nil{return QuarantineItem{},errors.Join(e,os.Remove(bin))}}
\tif e = os.WriteFile(metaPath, meta, 0600); e != nil {''')
    s=one(s,'\tif e = os.Remove(abs); e != nil {','''    current,e:=os.Lstat(abs)
    if e!=nil||!stableFile(abs,st,current){return QuarantineItem{},errors.Join(errors.New("arquivo original mudou"),os.Remove(bin),os.Remove(metaPath))}
    if expected!=""{if e=quarantineHash(abs,expected);e!=nil{return QuarantineItem{},errors.Join(e,os.Remove(bin),os.Remove(metaPath))}}
\tif e = os.Remove(abs); e != nil {''')
    s=one(s,'\tif e = copyClosed(bin, item.Original); e != nil {','''    if item.SHA256!=""{if e=quarantineHash(bin,item.SHA256);e!=nil{return e}}
\tif e = copyClosed(bin, item.Original); e != nil {''')
    return s

def main(s):
    s=one(s,'matched := false','matched := -1')
    s=one(s,'for _, f := range s.report.Findings {\n\t\t\tif f.Path == p.Path {\n\t\t\t\tmatched = true','for i, f := range s.report.Findings {\n\t\t\tif f.Path == p.Path && f.QuarantineEligible {\n\t\t\t\tmatched = i')
    s=one(s,'if !allowed || !matched {','if !allowed || matched < 0 {')
    s=one(s,'s.quarantine.Add(p.Path)','s.quarantine.AddVerified(s.report.Findings[matched])')
    s=one(s,'s.report = rep','rep.Summarize()\n\t\ts.report = rep')
    return s

def desktop(s):
    # No need for a new import; use index into the current report, under the existing mutex.
    a=s.index('if !bind("isolate"');b=s.index('if !bind("restoreItem"',a)
    x=s[a:b]
    x=one(x,'var hash string','var hash string\n hitIndex:=-1')
    prefix,tail=x.split('if hash == "" {',1)
    prefix=one(prefix,'for _, f := range s.report.Findings {','for i, f := range s.report.Findings {')
    x=prefix+'if hash == "" {'+tail
    x=one(x,'if f.Path == path {','if f.Path == path && f.QuarantineEligible {\n hitIndex=i')
    x=one(x,'if hash == "" {','if hash == "" || hitIndex<0 {')
    x=one(x,'s.quarantine.Add(path)','s.quarantine.AddVerified(s.report.Findings[hitIndex])')
    x=one(x,'s.report.Findings = findings','s.report.Findings = findings\n s.report.Summarize()') if 's.report.Findings = findings' in x else one(x,'s.report.Findings=findings','s.report.Findings=findings\n s.report.Summarize()')
    return s[:a]+x+s[b:]

def html(s):
    s=one(s,"if(!f.path.includes(' :: ')){let b=node('button','Isolar arquivo');", "row.append(node('p',f.evidence==='ioc_match'?'Correspondência com indicador local SHA-256':f.evidence==='test_eicar'?'Arquivo de teste EICAR (inofensivo)':'Indício heurístico: não comprova infecção'));if(f.trust)row.append(node('p',f.trust==='authenticode_valid_offline'?'Authenticode válido na verificação local (não garante segurança)':'Assinatura não verificada; pode estar em catálogo do Windows'));if(!f.quarantine_eligible)row.append(node('p','Revisão manual: isolamento bloqueado neste teste.'));if(f.quarantine_eligible&&!f.path.includes(' :: ')){let b=node('button','Isolar arquivo');")
    return s

def apply(root):
    sources={'engine/scanner.go':scanner,'engine/pe.go':pe,'engine/quarantine.go':quarantine,'main.go':main,'desktop_windows.go':desktop,'desktop.html':html}
    modified={}
    for name,fn in sources.items():
        p=root/name;modified[p]=fn(p.read_text(encoding='utf-8-sig'))
    for name,source in EMBEDDED_GO.items():
        dest=root/'engine'/name
        if dest.exists(): raise RuntimeError(f'Existing generated file: {dest}')
        modified[dest]=source
    if len(modified)!=11: raise RuntimeError('Incomplete upgrade bundle')
    for p,s in modified.items(): p.write_text(s,encoding='utf-8',newline='\n')
    print('Upgrade applied: ', ', '.join(str(p.relative_to(root)) for p in modified))

def install(root):
    wf=root/'.github/workflows/windows-build.yml';s=wf.read_text(encoding='utf-8-sig')
    if STEP in s:print('Workflow step already installed');return
    new=one(s,MARK,STEP+MARK)
    backup=wf.with_suffix('.yml.pre-upgrade.bak')
    if backup.exists(): raise RuntimeError('Backup already exists; inspect it first')
    shutil.copy2(wf,backup);wf.write_text(new,encoding='utf-8',newline='\n');print('Updated workflow; backup:',backup)


# Competitive UX and performance revision. Builds on the original 3.0.1 hardening.
_old_scanner, _old_main, _old_html = scanner, main, html
_old_pe = pe

def scanner(s):
    s=_old_scanner(s)
    s=one(s,'type Report struct {','''type Report struct {
    FilesPerSecond float64 `json:"files_per_second"`
    MBPerSecond float64 `json:"mb_per_second"`
    DuplicateAlertPaths int `json:"duplicate_alert_paths"`''')
    return s

def main(s):
    s=_old_main(s)
    # The generated desktop previously built a full Report and summary for every file.
    s=one(s,'one, err := s.scanner.Scan(ctx, path, nil)','hits, err := s.scanner.ExamineRegular(ctx, path, info.Size())')
    s=one(s,'if one.Cancelled {','if ctx.Err()!=nil {')
    s=one(s,'r.Files += one.Files','r.Files++')
    s=one(s,'r.Bytes += one.Bytes','r.Bytes += info.Size()')
    s=one(s,'r.Skipped += one.Skipped','')
    s=one(s,'r.Findings = append(r.Findings, one.Findings...)','r.Findings = append(r.Findings, hits...)')
    return s

def pe(s):
    s=_old_pe(s)
    # Exact API symbol lookup instead of repeated substring scans over all imports.
    a=s.index('\t\tseen := map[string]bool{}')
    b=s.index('\n\t\t// These APIs',a)
    s=s[:a]+'''        seen := indexImportSymbols(imports)
        has := func(parts ...string) bool {
            for _, p := range parts { if seen[p] { return true } }
            return false
        }'''+s[b:]
    return s

EMBEDDED_GO['scan_upgrade.go']=EMBEDDED_GO['scan_upgrade.go'].replace('3.0.1-test','3.0.1-competitive-beta')
EMBEDDED_GO['scan_upgrade.go']=one(EMBEDDED_GO['scan_upgrade.go'], '\tbuf := hashBuffers.Get().([]byte)\n\tdefer hashBuffers.Put(buf)\n\ttail, e := io.CopyBuffer(h, io.LimitReader(rd, size-n), buf)', '''    var tail int64
    var e error
    if size > n {
        buf:=hashBuffers.Get().([]byte)
        tail,e=io.CopyBuffer(h,io.LimitReader(rd,size-n),buf)
        hashBuffers.Put(buf)
    }''')
EMBEDDED_GO['scan_upgrade.go'] += '''
// Fast path for the already bounded desktop worker pool. The caller accounts for errors.
func (s *Scanner) ExamineRegular(ctx context.Context,path string,size int64)([]Finding,error){
    if err:=ctx.Err();err!=nil{return nil,err}
    return s.examine(ctx,path,size)
}

// ImportedSymbols returns "symbol:dll". Match exact case-insensitive symbol names.
func indexImportSymbols(imports []string)map[string]bool{
    seen:=make(map[string]bool,len(imports))
    for _,imp:=range imports{symbol,_,_:=strings.Cut(imp,":");seen[strings.ToLower(symbol)]=true}
    return seen
}
'''
EMBEDDED_GO['scan_upgrade.go']=one(EMBEDDED_GO['scan_upgrade.go'],'\tr.UniqueFindingHashes = len(seen)','''    r.UniqueFindingHashes = len(seen)
    r.DuplicateAlertPaths=0
    hashed:=0
    for _,f:=range r.Findings{if f.SHA256!=""{hashed++}}
    r.DuplicateAlertPaths=hashed-len(seen)
    r.FilesPerSecond=0;r.MBPerSecond=0
    if dt:=r.Finished.Sub(r.Started).Seconds();!r.Started.IsZero()&&!r.Finished.IsZero()&&dt>0{
        r.FilesPerSecond=float64(r.Files)/dt;r.MBPerSecond=float64(r.Bytes)/1e6/dt
    }''')
EMBEDDED_GO['scan_upgrade_test.go']=one(EMBEDDED_GO['scan_upgrade_test.go'],'\t"testing"','\t"testing"\n\t"time"')
EMBEDDED_GO['scan_upgrade_test.go'] += '''
func TestCompetitiveSummary(t *testing.T){
    start:=time.Now();r:=Report{Started:start,Finished:start.Add(2*time.Second),Files:10,Bytes:4000000,Findings:[]Finding{{SHA256:"a"},{SHA256:"a"},{SHA256:"b"}}}
    r.Summarize();if r.FilesPerSecond!=5||r.MBPerSecond!=2||r.DuplicateAlertPaths!=1||r.UniqueFindingHashes!=2{t.Fatal(r)}
}
func TestCompetitiveExactImportNames(t *testing.T){
    m:=indexImportSymbols([]string{"WriteProcessMemory:KERNEL32.dll","NotCreateRemoteThread:example.dll","IsDebuggerPresent:KERNEL32.dll"})
    if !m["writeprocessmemory"]||!m["isdebuggerpresent"]||m["createremotethread"]{t.Fatal(m)}
}
'''
EMBEDDED_GO['scan_upgrade_integration_test.go'] += '''
func TestCompetitiveFastPath(t *testing.T){
    dir:=t.TempDir();p:=filepath.Join(dir,"example.txt");data:=[]byte("harmless content")
    if err:=os.WriteFile(p,data,0600);err!=nil{t.Fatal(err)}
    sc:=New();sum:=sha256.Sum256(data);sc.IOC[hex.EncodeToString(sum[:])]="test"
    hits,err:=sc.ExamineRegular(context.Background(),p,int64(len(data)))
    if err!=nil||len(hits)!=1||hits[0].SHA256!=hex.EncodeToString(sum[:]){t.Fatal(hits,err)}
    ctx,cancel:=context.WithCancel(context.Background());cancel()
    if _,err=sc.ExamineRegular(ctx,p,int64(len(data)));err==nil{t.Fatal("cancel not observed")}
}
'''

COMPETITIVE_CSS=r'''
/* ZEUVORYX — original dark graphite, ice-blue and amber; no heavy image rendering. */
:root{--bg:#101419;--panel:#192127;--line:#34414b;--muted:#b5c1cb;--cyan:#6ce3cf;--blue:#8fb9ff}
body{background:var(--bg);color:#eef4f8;font-family:'Segoe UI',sans-serif}
aside{background:#131a20;border-color:var(--line)}
.brand{letter-spacing:1.2px}.tag{color:#a9b9c5}.zeus-art,.mythic{display:none}
.hero{background:#1b2a30;border:1px solid #40545b;box-shadow:none}
.card,.entry{background:var(--panel);border-color:var(--line)}
button{background:#28343e;border-color:#53636f;color:#f2f6f8}button.primary{background:#6ce3cf;color:#102723}
.nav button.active{background:#243d40;border-color:#426566}.nav button{color:#c5d0d7}
.risk{color:#ffd38b}.competitive-status{border-left:3px solid #8fb9ff;margin-bottom:10px}
.competitive-status[data-kind="review"]{border-left-color:#ffd38b}
.competitive-status[data-kind="limited"]{border-left-color:#ffd38b}
.competitive-status h2{margin-bottom:6px}.competitive-status p{line-height:1.55}
.competitive-filters{display:flex;flex-wrap:wrap;align-items:center;gap:7px;margin:12px 0}
.competitive-filters select{font:inherit;color:#eef4f8;background:#28343e;border:1px solid #53636f;border-radius:8px;padding:8px}
#zAlertMore{margin-top:10px}#zMetrics{font-variant-numeric:tabular-nums}
@media(prefers-reduced-motion:reduce){*{animation:none!important;scroll-behavior:auto!important}}
'''
COMPETITIVE_JS=r'''
// Render only the visible page, cap DOM creation, and retain complete export in Go.
let zState=null,zLimit=50;
const zBaseRender=renderFindings;
function zResults(){
 if(!zState||el('alerts').classList.contains('hidden'))return;
 const r=zState.report||{},all=r.findings||[],mode=el('zEvidence').value;
 const filtered=mode==='all'?all:all.filter(f=>f.evidence===mode);
 const view=filtered.slice(0,zLimit);
 zBaseRender({...zState,report:{...r,findings:view}});
 if(!filtered.length&&all.length)el('findings').replaceChildren(node('p','Nenhum alerta nesta categoria.','muted'));
 el('zAlertCount').textContent=fmt(view.length)+' de '+fmt(filtered.length)+' resultados nesta categoria · '+fmt(all.length)+' no relatório completo';
 el('zAlertMore').hidden=view.length>=filtered.length;
}
renderFindings=function(s){zState=s;zLimit=50;zResults();};
el('zEvidence').onchange=()=>{zLimit=50;zResults()};
el('zAlertMore').onclick=()=>{zLimit+=50;zResults()};
document.querySelector('[data-page="alerts"]').addEventListener('click',zResults);
function zSummary(s){
 const r=s.report||{},n=(r.findings||[]).length;
 const state=el('zSummary');let title,detail,kind='normal';
 if(s.running){title='Análise em andamento';detail='Aguardando o relatório final. A proteção residente depende do Windows.';}
 else if(s.error||r.cancelled){title='Análise interrompida';detail='Não é possível concluir sobre todo o escopo.';kind='limited';}
 else if(!s.done){title='Proteção residente: consulte o Windows';detail='ZEUVORYX ainda não executou uma análise nesta sessão.';}
 else if(n){title='Resultados precisam de revisão';detail=fmt(r.ioc_matches)+' correspondências SHA-256 · '+fmt(r.heuristic_alerts)+' indícios heurísticos · '+fmt(r.eicar_tests)+' testes EICAR.';kind='review';}
 else{title='Nenhum alerta no conteúdo examinado';detail='Isso não comprova ausência de malware.';}
 if(s.done&&(r.partial||r.skipped)){detail+=' Cobertura parcial: '+fmt(r.skipped)+' itens ignorados.';kind='limited';}
 state.dataset.kind=kind;el('zTitle').textContent=title;el('zDetail').textContent=detail;
 const has=s.done&&Number.isFinite(r.files_per_second);
 el('zMetrics').textContent=has?fmt(Math.round(r.duration_ms/1000))+' s · '+r.files_per_second.toFixed(1)+' arquivos/s · '+r.mb_per_second.toFixed(2)+' MB/s · '+fmt(r.unique_finding_hashes)+' hashes distintos':'Métricas disponíveis ao fim da análise.';
}
'''

def html(s):
    s=_old_html(s)
    # Preserve packaged art on disk, remove only decorative base64 blobs from the active dashboard.
    s=re.sub(r'<img class="(?:zeus-art|mythic)" src="data:image/jpeg;base64,[A-Za-z0-9+/=]+" alt="[^"]*">','',s)
    s=re.sub(r',url\("data:image/jpeg;base64,[A-Za-z0-9+/=]+"\) center/cover','',s)
    s=one(s,'</style></head>',COMPETITIVE_CSS+'</style></head>')
    s=one(s,'OLYMPUS 3.0</div>','OLYMPUS 3.0.1 BETA</div>')
    s=one(s,'<section id="overview"><div class="hero"','<section id="overview"><div id="zSummary" class="card competitive-status" role="status"><h2 id="zTitle">Estado da análise</h2><p id="zDetail" class="sub">Aguardando consulta.</p><p id="zMetrics" class="sub"></p></div><div class="hero"')
    s=one(s,'<div id="findings" class="list"></div>','<div class="competitive-filters"><label for="zEvidence">Filtrar alertas</label><select id="zEvidence"><option value="all">Todos</option><option value="ioc_match">Indicador SHA-256</option><option value="heuristic">Heurística</option><option value="test_eicar">Teste EICAR</option></select><span id="zAlertCount" class="muted" aria-live="polite"></span></div><div id="findings" class="list"></div><button id="zAlertMore" hidden>Mostrar mais 50</button>')
    s=one(s,'let running=s.running;','let running=s.running;zSummary(s);')
    # Define handlers before the original initial refresh; no second polling loop.
    s=one(s,'refresh();setInterval(refresh,2000);',COMPETITIVE_JS+'\nrefresh();setInterval(refresh,2000);')
    # Do not display internal level enum strings as a user-facing risk verdict.
    s=one(s,'`${f.level} · ${f.risk}/100`', '`${({low:"Baixo",medium:"Moderado",high:"Alto",critical:"Crítico"})[f.level]||"Indeterminado"} · ${f.risk}/100`')
    return s

if __name__=='__main__':
    parser=argparse.ArgumentParser();g=parser.add_mutually_exclusive_group(required=True);g.add_argument('--apply',action='store_true');g.add_argument('--install',action='store_true');parser.add_argument('--root',default='.');a=parser.parse_args()
    try:(apply if a.apply else install)(Path(a.root).resolve())
    except Exception as e:print('ABORTED:',e,file=sys.stderr);sys.exit(1)
