package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/zeuvoryx/zeuvoryx/engine"
	"github.com/zeuvoryx/zeuvoryx/ui"
)

type State struct {
	mu         sync.Mutex
	running    bool
	files      int
	bytes      int64
	cancel     context.CancelFunc
	report     engine.Report
	err        string
	scanner    *engine.Scanner
	quarantine *engine.Quarantine
	token      string
}

func reply(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}

func fail(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	reply(w, map[string]string{"error": msg})
}

func (s *State) serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && (r.URL.Path == "/assets/cover.jpg" || r.URL.Path == "/assets/dashboard.jpg") {
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Cache-Control", "public, max-age=86400")
		if r.URL.Path == "/assets/cover.jpg" {
			_, _ = w.Write(ui.Cover)
		} else {
			_, _ = w.Write(ui.Dashboard)
		}
		return
	}

	if o := r.Header.Get("Origin"); o != "" && o != "http://"+r.Host {
		fail(w, 403, "origem bloqueada")
		return
	}

	if r.URL.Query().Get("token") != s.token && r.Header.Get("X-Zeuvoryx-Token") != s.token {
		fail(w, 403, "sessão inválida")
		return
	}

	switch {
	case r.URL.Path == "/" && r.Method == "GET":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		w.Write(ui.HTML)

	case r.URL.Path == "/api/default" && r.Method == "GET":
		home, err := os.UserHomeDir()
		if err != nil {
			fail(w, 500, err.Error())
			return
		}
		reply(w, map[string]string{
			"downloads": filepath.Join(home, "Downloads"),
		})

	case r.URL.Path == "/api/status" && r.Method == "GET":
		s.mu.Lock()
		data := map[string]any{
			"running": s.running,
			"files":   s.files,
			"bytes":   s.bytes,
			"report":  s.report,
			"error":   s.err,
		}
		s.mu.Unlock()
		reply(w, data)

	case r.URL.Path == "/api/start" && r.Method == "POST":
		var p struct {
			Path string `json:"path"`
		}

		if e := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4096)).Decode(&p); e != nil {
			fail(w, 400, "caminho inválido")
			return
		}

		if strings.TrimSpace(p.Path) == "" {
			fail(w, 400, "informe um caminho")
			return
		}

		if !s.startScan(p.Path, false) {
			fail(w, 409, "varredura em andamento")
			return
		}

		reply(w, map[string]bool{"started": true})

	case r.URL.Path == "/api/full" && r.Method == "POST":
		if !s.startScan("", true) {
			fail(w, 409, "varredura em andamento")
			return
		}

		reply(w, map[string]bool{"started": true})

	case r.URL.Path == "/api/stop" && r.Method == "POST":
		s.mu.Lock()
		if s.cancel != nil {
			s.cancel()
		}
		s.mu.Unlock()

		reply(w, map[string]bool{"stopping": true})

	case r.URL.Path == "/api/quarantine" && r.Method == "GET":
		items, e := s.quarantine.List()
		if e != nil {
			fail(w, 500, e.Error())
			return
		}

		reply(w, items)

	case r.URL.Path == "/api/quarantine" && r.Method == "POST":
		var p struct {
			Path string `json:"path"`
		}

		if e := json.NewDecoder(http.MaxBytesReader(w, r.Body, 8192)).Decode(&p); e != nil {
			fail(w, 400, "JSON inválido")
			return
		}

		if strings.Contains(p.Path, " :: ") {
			fail(w, 400, "não é possível isolar entradas internas de ZIP; extraia primeiro")
			return
		}

		s.mu.Lock()

		allowed := !s.running
		matched := false

		for _, f := range s.report.Findings {
			if f.Path == p.Path {
				matched = true
				break
			}
		}

		if !allowed || !matched {
			s.mu.Unlock()
			fail(w, 403, "arquivo não consta na última varredura concluída")
			return
		}

		item, e := s.quarantine.Add(p.Path)
		s.mu.Unlock()

		if e != nil {
			fail(w, 400, e.Error())
			return
		}

		reply(w, item)

	case r.URL.Path == "/api/restore" && r.Method == "POST":
		var p struct {
			ID string `json:"id"`
		}

		if e := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1024)).Decode(&p); e != nil {
			fail(w, 400, "JSON inválido")
			return
		}

		s.mu.Lock()

		if s.running {
			s.mu.Unlock()
			fail(w, 409, "aguarde ou cancele a varredura antes de restaurar")
			return
		}

		e := s.quarantine.Restore(p.ID)
		s.mu.Unlock()

		if e != nil {
			fail(w, 400, e.Error())
			return
		}

		reply(w, map[string]bool{"restored": true})

	case r.URL.Path == "/api/export" && r.Method == "GET":
		s.mu.Lock()
		rep := s.report
		s.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Disposition", "attachment; filename=zeuvoryx-report.json")
		json.NewEncoder(w).Encode(rep)

	default:
		fail(w, 404, "endpoint indisponível")
	}
}

func parseFixedRoots(out string) []string {
	var roots []string
	seen := map[string]bool{}

	for _, line := range strings.Fields(out) {
		line = strings.ToUpper(line)

		if len(line) == 2 &&
			line[0] >= 'A' &&
			line[0] <= 'Z' &&
			line[1] == ':' &&
			!seen[line] {

			roots = append(roots, line+`\`)
			seen[line] = true
		}
	}

	sort.Strings(roots)
	return roots
}

func fixedRoots(ctx context.Context) ([]string, error) {
	if runtime.GOOS != "windows" {
		return []string{string(os.PathSeparator)}, nil
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	out, err := exec.CommandContext(
		ctx,
		"powershell.exe",
		"-NoProfile",
		"-NonInteractive",
		"-Command",
		"(Get-CimInstance -ClassName Win32_LogicalDisk -Filter 'DriveType=3' -ErrorAction Stop).DeviceID",
	).Output()

	if err != nil {
		return nil, fmt.Errorf("falha ao listar discos locais: %w", err)
	}

	roots := parseFixedRoots(string(out))

	if len(roots) == 0 {
		return nil, fmt.Errorf("nenhum disco local encontrado")
	}

	return roots, nil
}

func samePath(a, b string) bool {
	if b == "" {
		return false
	}

	a, b = filepath.Clean(a), filepath.Clean(b)

	if runtime.GOOS == "windows" {
		return strings.EqualFold(a, b)
	}

	return a == b
}

func (s *State) scanRoots(
	ctx context.Context,
	roots []string,
	self string,
	progress func(int, int64),
) engine.Report {

	r := engine.Report{
		Root:     strings.Join(roots, "; "),
		Started:  time.Now(),
		Findings: []engine.Finding{},
	}

	addError := func(path string, err error) {
		r.Skipped++

		if len(r.Errors) < 30 {
			r.Errors = append(r.Errors, path+": "+err.Error())
		}
	}

	for _, root := range roots {
		if ctx.Err() != nil {
			break
		}

		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if ctx.Err() != nil {
				return ctx.Err()
			}

			if err != nil {
				addError(path, err)
				return nil
			}

			if s.quarantine.IsQuarantined(path) {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			if d.IsDir() {
				return nil
			}

			if !d.Type().IsRegular() || samePath(path, self) {
				r.Skipped++
				return nil
			}

			info, e := os.Lstat(path)

			if e != nil {
				addError(path, e)
				return nil
			}

			if !info.Mode().IsRegular() {
				r.Skipped++
				return nil
			}

			one, e := s.scanner.Scan(ctx, path, nil)

			if e != nil {
				addError(path, e)
				return nil
			}

			if one.Cancelled {
				return context.Canceled
			}

			r.Files += one.Files
			r.Bytes += one.Bytes
			r.Skipped += one.Skipped
			r.Findings = append(r.Findings, one.Findings...)

			if progress != nil {
				progress(r.Files, r.Bytes)
			}

			return nil
		})

		if err != nil && ctx.Err() == nil {
			addError(root, err)
		}
	}

	r.Cancelled = ctx.Err() != nil
	r.Finished = time.Now()

	sort.SliceStable(r.Findings, func(i, j int) bool {
		return r.Findings[i].Risk > r.Findings[j].Risk
	})

	return r
}

func (s *State) fullScan(
	ctx context.Context,
	progress func(int, int64),
) (engine.Report, error) {

	roots, e := fixedRoots(ctx)

	if e != nil {
		now := time.Now()

		return engine.Report{
			Root:      "Discos locais",
			Started:   now,
			Finished:  now,
			Cancelled: ctx.Err() != nil,
		}, e
	}

	self, _ := os.Executable()

	return s.scanRoots(ctx, roots, self, progress), nil
}

func (s *State) startScan(path string, automatic bool) bool {
	s.mu.Lock()

	if s.running {
		s.mu.Unlock()
		return false
	}

	ctx, cancel := context.WithCancel(context.Background())

	s.cancel = cancel
	s.running = true
	s.files = 0
	s.bytes = 0
	s.err = ""
	s.report = engine.Report{}

	s.mu.Unlock()

	go func() {
		defer cancel()

		progress := func(n int, b int64) {
			s.mu.Lock()
			s.files = n
			s.bytes = b
			s.mu.Unlock()
		}

		var rep engine.Report
		var e error

		if automatic {
			rep, e = s.fullScan(ctx, progress)
		} else {
			rep, e = s.scanner.Scan(ctx, path, progress)
		}

		s.mu.Lock()
		defer s.mu.Unlock()

		s.report = rep
		s.files = rep.Files
		s.bytes = rep.Bytes
		s.running = false
		s.cancel = nil

		if e != nil {
			s.err = e.Error()
		}
	}()

	return true
}

func main() {
	scan := flag.String(
		"scan",
		"",
		"escaneia um arquivo ou diretório via linha de comando",
	)

	ioc := flag.String(
		"ioc",
		"",
		"caminho opcional para assinaturas SHA-256",
	)

	auto := flag.Bool(
		"auto",
		true,
		"varredura automática dos discos locais ao abrir",
	)

	flag.Parse()

	s := engine.New()

	if *ioc != "" {
		if e := s.LoadIOC(*ioc); e != nil {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
		}
	} else {
		if exe, e := os.Executable(); e == nil {
			if e = s.LoadIOC(filepath.Join(filepath.Dir(exe), "signatures.sha256")); e != nil {
				log.Printf("Aviso: assinaturas não carregadas: %v", e)
			}
		}
	}

	if *scan != "" {
		r, e := s.Scan(context.Background(), *scan, nil)

		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
		}

		json.NewEncoder(os.Stdout).Encode(r)
		return
	}

	q, e := engine.NewQuarantine()

	if e != nil {
		log.Fatal(e)
	}

	var key [24]byte

	if _, e := rand.Read(key[:]); e != nil {
		log.Fatal(e)
	}

	state := &State{
		scanner:    s,
		quarantine: q,
		token:      hex.EncodeToString(key[:]),
	}

	listener, e := net.Listen("tcp", "127.0.0.1:0")

	if e != nil {
		log.Fatal(e)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", state.serve)

	server := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	url := fmt.Sprintf(
		"http://127.0.0.1:%d/?token=%s",
		listener.Addr().(*net.TCPAddr).Port,
		state.token,
	)

	fmt.Println("ZEUVORYX: " + url)

	if *auto {
		state.startScan("", true)
	}

	go func() {
		time.Sleep(300 * time.Millisecond)

		switch runtime.GOOS {
		case "windows":
			exec.Command(
				"rundll32.exe",
				"url.dll,FileProtocolHandler",
				url,
			).Start()

		case "darwin":
			exec.Command("open", url).Start()

		default:
			exec.Command("xdg-open", url).Start()
		}
	}()

	if e := server.Serve(listener); e != nil && e != http.ErrServerClosed {
		log.Fatal(e)
	}
}
