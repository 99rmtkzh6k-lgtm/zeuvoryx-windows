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

func TestWikiLinkRegression(t *testing.T){
 cases:=[]struct{raw string;want int;invalid bool}{
 {"https://wikileaks.org/ ",0,false},
 {"https://www.wikileaks.org/",0,false},
 {"https://wikileaks.org./",0,false},
 {"http://wikileaks.org/",1,false},
 {"https://wikileaks.org@other.invalid/",1,false},
 {"https://wikileaks.org.other.invalid/",1,false},
 {"https://wikileaks-support.invalid/",1,false},
 {"https://wikileakz.invalid/",1,false},
 {"https://wikileak.invalid/",1,false},
 {"https://wiki-leaks.invalid/",1,false},
 {"https://wikilekas.invalid/",1,false},
 {"https://example.invalid/news/wikileaks",0,false},
 {"https://example.invalid/?q=wikileaks",0,false},
 {"https://wikileaks.org/?redirect=https%3A%2F%2Fother.invalid",1,false},
 {"https://wikileaks.org/?%72edirect=https%3A%2F%2Fother.invalid",1,false},
 {"https://wikileaks.org/?%2572edirect=https%3A%2F%2Fother.invalid",1,false},
 {"https://wikileaks.org/?redirect%3Dhttps%253A%252F%252Fother.invalid",1,false},
 {"https://wikileaks.org/?redirect_uri=https%3A%2F%2Fother.invalid",1,false},
 {"https://wikileaks.org/?next=%2Fnews",1,false},
 {"https://example.invalid/?myurl=value",0,false},
 {"file:///wikileaks.org",0,true},
 }
 for _,c:=range cases{
  r,e:=assessLink(c.raw);if c.invalid{if e==nil{t.Fatalf("expected invalid: %s",c.raw)};continue}
  if e!=nil||len(r.Signals)!=c.want{t.Fatalf("%s: got %+v %v want %d signals",c.raw,r,e,c.want)}
 }
 t.Logf("%d structural URL cases; no network access or active blocking",len(cases))
}
