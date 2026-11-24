package portable

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestHashListStrict(t *testing.T) {
	h := strings.Repeat("A", 64)
	r, e := parseHashList("# header\n" + h + " label\n" + h + "\n")
	if e != nil || len(r) != 1 || r[strings.ToLower(h)] == "" {
		t.Fatal(r, e)
	}
	for _, x := range []string{"", h + "\ninvalid", strings.Repeat("z", 64), strings.Repeat("a", 4097)} {
		if _, e := parseHashList(x); e == nil {
			t.Fatal("accepted invalid list")
		}
	}
}
func TestBazaarStrict(t *testing.T) {
	h := strings.Repeat("b", 64)
	b := fmt.Sprintf(`{"query_status":"ok","data":[{"sha256_hash":"%s"}]}`, h)
	r, e := parseBazaar([]byte(b))
	if e != nil || len(r) != 1 {
		t.Fatal(r, e)
	}
	for _, x := range []string{`{"query_status":"no_results"}`, `{"query_status":"ok","data":[{"sha256_hash":"bad"}]}`, `not json`} {
		if _, e := parseBazaar([]byte(x)); e == nil {
			t.Fatal(x)
		}
	}
}
func TestIntelMerge(t *testing.T) {
	a, b := strings.Repeat("a", 64), strings.Repeat("b", 64)
	old := map[string]string{a: "known"}
	out, e := mergeIntel(old, map[string]string{a: "other", b: "new"})
	if e != nil || len(old) != 1 || len(out) != 2 || out[a] != "known" {
		t.Fatal(out, e)
	}
	if _, e := mergeIntel(old, map[string]string{"bad": "bad"}); e == nil {
		t.Fatal("accepted invalid")
	}
}
func TestKEVLookup(t *testing.T) {
	b := []byte(`{"catalogVersion":"2026.01.01","count":1,"vulnerabilities":[{"cveID":"CVE-2025-1234","product":"Example"}]}`)
	r, e := parseKEV(b, "cve-2025-1234")
	if e != nil || !r.Found || r.Entry.Product != "Example" {
		t.Fatal(r, e)
	}
	r, e = parseKEV(b, "CVE-2025-12345")
	if e != nil || r.Found {
		t.Fatal(r, e)
	}
	if _, e := parseKEV(b, "bad"); e == nil {
		t.Fatal("invalid CVE")
	}
	if _, e := parseKEV([]byte(`{"catalogVersion":"x","count":1}`), "CVE-2025-1234"); e == nil {
		t.Fatal("incomplete catalog")
	}
}

type fakeIntelTransport func(*http.Request) (*http.Response, error)

func (f fakeIntelTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestIntelHTTPContract(t *testing.T) {
	h := strings.Repeat("c", 64)
	c := &http.Client{Transport: fakeIntelTransport(func(r *http.Request) (*http.Response, error) {
		if r.URL.String() != bazaarURL || r.Method != "POST" || r.Header.Get("Auth-Key") != "example-key" {
			t.Fatal("incorrect request")
		}
		if e := r.ParseForm(); e != nil || r.Form.Get("query") != "get_recent" || r.Form.Get("selector") != "100" {
			t.Fatal("incorrect form")
		}
		return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(fmt.Sprintf(`{"query_status":"ok","data":[{"sha256_hash":"%s"}]}`, h)))}, nil
	})}
	r, e := fetchBazaar(context.Background(), c, "example-key")
	if e != nil || len(r) != 1 {
		t.Fatal(r, e)
	}
}
func TestIntelHTTPRejectsErrors(t *testing.T) {
	c := &http.Client{Transport: fakeIntelTransport(func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 403, Body: io.NopCloser(strings.NewReader("denied"))}, nil
	})}
	if _, e := fetchBazaar(context.Background(), c, "example-key"); e == nil {
		t.Fatal("accepted HTTP error")
	}
}
func TestIntelRedirectPolicy(t *testing.T) {
	c := intelClient()
	if c.CheckRedirect(nil, nil) != http.ErrUseLastResponse {
		t.Fatal("redirect allowed")
	}
}
