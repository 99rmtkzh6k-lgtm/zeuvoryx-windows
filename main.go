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
	// Visual assets are static and contain no user data.
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
	// Requests from other websites cannot operate the local scanner.
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
		reply(w, map[string]string{"downloads": filepath.Join(home, "Downloads")})
	case r.URL.Path == "/api/status" && r.Method == "GET":
		s.mu.Lock()
		data := map[string]any{"running": s.running, "files": s.files, "bytes": s.bytes, "report": s.report, "error": s.err}
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
		s.mu.Lock()
		if s.running {
			s.mu.Unlock()
			fail(w, 409, "varredura em andamento")
			return
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
			rep, e := s.scanner.Scan(ctx, p.Path, func(n int, b int64) { s.mu.Lock(); s.files = n; s.bytes = b; s.mu.Unlock() })
			s.mu.Lock()
			s.report = rep
			s.running = false
			if e != nil {
				s.err = e.Error()
			}
			s.mu.Unlock()
			cancel()
		}()
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
		// Safety: can quarantine only a file actually detected by the completed scan.
		s.mu.Lock()
		allowed := !s.running
		matched := false
		for _, f := range s.report.Findings {
			if f.Path == p.Path {
				matched = true
				break
			}
		}
		s.mu.Unlock()
		if !allowed || !matched {
			fail(w, 403, "arquivo não consta na última varredura concluída")
			return
		}
		item, e := s.quarantine.Add(p.Path)
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
		if e := s.quarantine.Restore(p.ID); e != nil {
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
func main() {
	scan := flag.String("scan", "", "escaneia um arquivo ou diretório via linha de comando")
	ioc := flag.String("ioc", "", "caminho opcional para assinaturas SHA-256")
	flag.Parse()
	s := engine.New()
	if *ioc != "" {
		if e := s.LoadIOC(*ioc); e != nil {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
		}
	} else {
		if exe, e := os.Executable(); e == nil {
			s.LoadIOC(filepath.Join(filepath.Dir(exe), "signatures.sha256"))
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
	state := &State{scanner: s, quarantine: q, token: hex.EncodeToString(key[:])}
	listener, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		log.Fatal(e)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", state.serve)
	server := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	url := fmt.Sprintf("http://127.0.0.1:%d/?token=%s", listener.Addr().(*net.TCPAddr).Port, state.token)
	fmt.Println("ZEUVORYX: " + url)
	go func() {
		time.Sleep(300 * time.Millisecond)
		switch runtime.GOOS {
		case "windows":
			exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", url).Start()
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
