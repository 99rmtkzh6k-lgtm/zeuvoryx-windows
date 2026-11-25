package engine

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

// Verdicts are advisory: suspicious files are never removed automatically.
type Finding struct {
	Path    string   `json:"path"`
	SHA256  string   `json:"sha256,omitempty"`
	Risk    int      `json:"risk"`
	Level   string   `json:"level"`
	Reasons []string `json:"reasons"`
 EvidenceKind string `json:"evidence_kind"`
 AdvisoryOnly bool `json:"advisory_only"`
	Size    int64    `json:"size"`
}
type Report struct {
	Started   time.Time `json:"started"`
	Finished  time.Time `json:"finished"`
	Root      string    `json:"root"`
	Files     int       `json:"files"`
	Bytes     int64     `json:"bytes"`
	Skipped   int       `json:"skipped"`
	Findings  []Finding `json:"findings"`
	Errors    []string  `json:"errors,omitempty"`
	Cancelled bool      `json:"cancelled"`
}
type Scanner struct {
	IOC               map[string]string
	MaxFileBytes      int64
	MaxArchiveEntry   int64
	MaxArchiveEntries int
	MaxArchiveTotal   int64
}

func New() *Scanner {
	return &Scanner{IOC: map[string]string{}, MaxFileBytes: 32 << 20, MaxArchiveEntry: 8 << 20, MaxArchiveEntries: 100, MaxArchiveTotal: 64 << 20}
}
func (s *Scanner) LoadIOC(path string) error {
	b, e := os.ReadFile(path)
	if e != nil {
		return e
	}
	for n, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 1 || len(parts[0]) != 64 {
			return fmt.Errorf("linha IOC %d inválida", n+1)
		}
		if _, e := hex.DecodeString(parts[0]); e != nil {
			return fmt.Errorf("linha IOC %d: %w", n+1, e)
		}
		label := "Assinatura SHA-256 conhecida"
		if len(parts) > 1 {
			label = strings.Join(parts[1:], " ")
		}
		s.IOC[strings.ToLower(parts[0])] = label
	}
	return nil
}

var doubleExt = regexp.MustCompile(`(?i)\.(pdf|docx?|xlsx?|jpe?g|png|txt|mp4)\.(exe|scr|com|bat|cmd|ps1|vbs|js)$`)
var riskyExtensions = map[string]bool{".ps1": true, ".vbs": true, ".vbe": true, ".js": true, ".jse": true, ".hta": true, ".wsf": true, ".scr": true, ".bat": true, ".cmd": true}
var rules = []struct {
	phrase string
	weight int
	label  string
}{
	{"-encodedcommand", 35, "PowerShell com comando codificado"},
	{"frombase64string(", 25, "Decodificação Base64 em script"},
	{"invoke-expression", 25, "Execução dinâmica de comando"},
	{"iex(", 20, "Execução dinâmica abreviada"},
	{"downloadstring(", 25, "Download e execução em script"},
	{"wscript.shell", 15, "Automação de shell"},
	{"-windowstyle hidden", 20, "Tentativa de ocultar janela"},
	{"add-mppreference -exclusionpath", 45, "Tentativa de criar exclusão de antivírus"},
	{"set-mppreference -disablerealtimemonitoring", 60, "Tentativa de desabilitar proteção"},
	{"vssadmin delete shadows", 50, "Tentativa de apagar cópias de sombra"},
	{"wbadmin delete catalog", 50, "Tentativa de apagar catálogo de backup"},
	{"bcdedit /set", 15, "Alteração de configuração de boot"},
	{"rundll32.exe javascript:", 45, "Execução suspeita via rundll32"},
}

func entropy(b []byte) float64 {
	if len(b) == 0 {
		return 0
	}
	var counts [256]int
	for _, v := range b {
		counts[v]++
	}
	var out float64
	for _, c := range counts {
		if c > 0 {
			p := float64(c) / float64(len(b))
			out -= p * math.Log2(p)
		}
	}
	return out
}
func tier(score int) string {
	if score >= 80 {
		return "critical"
	}
	if score >= 55 {
		return "high"
	}
	if score >= 30 {
		return "medium"
	}
	return "low"
}
func analyze(path string, data []byte, fullHash string, ioc map[string]string) *Finding {
	score := 0
 evidenceKind := "heuristic_advisory"
	var reasons []string
	name := filepath.Base(path)
	ext := strings.ToLower(filepath.Ext(name))
	if doubleExt.MatchString(name) {
		score += 50
		reasons = append(reasons, "Extensão dupla simulando documento")
	}
	if riskyExtensions[ext] {
		score += 10
		reasons = append(reasons, "Tipo de script executável merece cautela")
	}
	lower := strings.ToLower(string(data))
	for _, rule := range rules {
		if strings.Contains(lower, rule.phrase) {
			score += rule.weight
			reasons = append(reasons, rule.label)
		}
	}
	if len(data) >= 4096 && riskyExtensions[ext] && entropy(data) > 7.8 {
		score += 15
		reasons = append(reasons, "Entropia elevada em script")
	}
	if bytes.Contains(data, []byte("X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*")) {
		score = 100
 evidenceKind = "test_artifact"
		reasons = append(reasons, "Assinatura de teste EICAR detectada (não é malware real)")
	}
	if label, ok := ioc[strings.ToLower(fullHash)]; ok {
		score = 100
 if evidenceKind != "test_artifact" { evidenceKind = "known_hash_match" }
		reasons = append(reasons, "Correspondência SHA-256: "+label)
	}
	if score > 100 {
		score = 100
	}
	if pe := analyzePE(path, data); pe != nil {
		score += pe.Risk
		reasons = append(reasons, pe.Reasons...)
	}
	if score > 100 {
		score = 100
	}
	if score < 30 {
		return nil
	}
	return &Finding{Path: path, SHA256: fullHash, Risk: score, Level: tier(score), Reasons: reasons, Size: int64(len(data)), EvidenceKind:evidenceKind, AdvisoryOnly:true}
}
func (s *Scanner) examine(path string, size int64) ([]Finding, error) {
	f, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	// Hash the complete file in one sequential pass while retaining only the bounded sample.
 h := sha256.New()
 limit := s.MaxFileBytes
 if limit <= 0 {limit = 32 << 20}
 b,e:=io.ReadAll(io.TeeReader(io.LimitReader(f,limit),h))
 if e!=nil{return nil,e}
 if _,e=io.Copy(h,f);e!=nil{return nil,e}
 hash:=hex.EncodeToString(h.Sum(nil))
	var found []Finding
	if hit := analyze(path, b, hash, s.IOC); hit != nil {
		hit.Size = size
		found = append(found, *hit)
	}
	if strings.EqualFold(filepath.Ext(path), ".zip") {
		archive, e := zip.OpenReader(path)
		if e == nil {
			defer archive.Close()
			maxEntries := s.MaxArchiveEntries
			if maxEntries <= 0 {
				maxEntries = 100
			}
			entryMax := s.MaxArchiveEntry
			if entryMax <= 0 {
				entryMax = 8 << 20
			}
			var totalUncompressed uint64
			totalLimit := s.MaxArchiveTotal
			if totalLimit <= 0 {
				totalLimit = 64 << 20
			}
			for idx, z := range archive.File {
				if idx >= maxEntries {
					break
				}
				if z.FileInfo().IsDir() || z.UncompressedSize64 > uint64(entryMax) || z.UncompressedSize64 > uint64(totalLimit)-totalUncompressed {
					continue
				}
				totalUncompressed += z.UncompressedSize64
				r, e := z.Open()
				if e != nil {
					continue
				}
				data, e := io.ReadAll(io.LimitReader(r, entryMax+1))
				r.Close()
				if e != nil || int64(len(data)) > entryMax {
					continue
				}
				sum := sha256.Sum256(data)
				if hit := analyze(path+" :: "+z.Name, data, hex.EncodeToString(sum[:]), s.IOC); hit != nil {
					found = append(found, *hit)
				}
			}
		}
	}
	return found, nil
}
func (s *Scanner) Scan(ctx context.Context, root string, progress func(int, int64)) (Report, error) {
	abs, e := filepath.Abs(root)
	if e != nil {
		return Report{}, e
	}
	info, e := os.Stat(abs)
	if e != nil {
		return Report{}, e
	}
	report := Report{Started: time.Now(), Root: abs, Findings: []Finding{}}
	defer func() {}()
	if ctx.Err() != nil {
		report.Cancelled = true
		report.Finished = time.Now()
		return report, nil
	}
	if info.Mode().IsRegular() {
		hits, e := s.examine(abs, info.Size())
		report.Files = 1
		report.Bytes = info.Size()
		report.Findings = append(report.Findings, hits...)
		report.Finished = time.Now()
		return report, e
	}
	if !info.IsDir() {
		return report, errors.New("caminho não é arquivo ou pasta")
	}
	// Bounded worker pool; errors and accounting are synchronized.
	paths := make(chan string, 64)
	var wg sync.WaitGroup
	var mu sync.Mutex
	workers := runtime.NumCPU()
	if workers > 4 {
		workers = 4
	}
	if workers < 1 {
		workers = 1
	}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range paths {
				if ctx.Err() != nil {
					return
				}
				st, e := os.Stat(path)
				if e != nil {
					mu.Lock()
					report.Skipped++
					mu.Unlock()
					continue
				}
				hits, e := s.examine(path, st.Size())
				mu.Lock()
				if e != nil {
					report.Skipped++
					if len(report.Errors) < 30 {
						report.Errors = append(report.Errors, path+": "+e.Error())
					}
				} else {
					report.Files++
					report.Bytes += st.Size()
					report.Findings = append(report.Findings, hits...)
				}
				n, b := report.Files, report.Bytes
				mu.Unlock()
				if progress != nil {
					// Serialize callbacks for thread-safe UI consumers.
					mu.Lock()
					progress(n, b)
					mu.Unlock()
				}
			}
		}()
	}
	walkErr := filepath.WalkDir(abs, func(path string, d os.DirEntry, e error) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if e != nil {
			mu.Lock()
			report.Skipped++
			if len(report.Errors) < 30 {
				report.Errors = append(report.Errors, path+": "+e.Error())
			}
			mu.Unlock()
			return nil
		}
		if d.Type()&os.ModeSymlink != 0 {
			return nil
		}
		if !d.Type().IsRegular() {
			return nil
		}
		select {
		case paths <- path:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})
	close(paths)
	wg.Wait()
	report.Finished = time.Now()
	report.Cancelled = ctx.Err() != nil
	sort.Slice(report.Findings, func(i, j int) bool { return report.Findings[i].Risk > report.Findings[j].Risk })
	if walkErr != nil && !errors.Is(walkErr, context.Canceled) {
		return report, walkErr
	}
	return report, nil
}
