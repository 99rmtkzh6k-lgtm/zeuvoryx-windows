package portable

import "testing"

func TestLinkTriage(t *testing.T) {
	for _, tc := range []struct {
		url, host string
		signals   int
		invalid   bool
	}{
		{"https://example.org/", "example.org", 0, false},
		{"http://example.org/", "example.org", 1, false},
		{"https://example.org@127.0.0.1:8080/", "127.0.0.1", 3, false},
		{"https://xn--example-9d0b.com/", "xn--example-9d0b.com", 1, false},
		{"https://a.b.c.d.example.org/", "a.b.c.d.example.org", 1, false},
		{"file:///C:/Windows/", "", 0, true},
		{"https://example.org\\@other.test/", "", 0, true},
		{"https://example.org/\nx", "", 0, true},
		{"https://example.org/?url=https%3A%2F%2Fother.test", "example.org", 1, false},
	} {
		r, e := assessLink(tc.url)
		if tc.invalid {
			if e == nil {
				t.Fatalf("expected invalid %q", tc.url)
			}
			continue
		}
		if e != nil || r.Host != tc.host || len(r.Signals) != tc.signals {
			t.Fatalf("%q: %+v %v", tc.url, r, e)
		}
	}
}
