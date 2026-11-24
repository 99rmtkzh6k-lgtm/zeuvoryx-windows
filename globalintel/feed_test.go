package globalintel
import("context";"fmt";"io";"net/http";"strings";"testing";"time")
type transport func(*http.Request)(*http.Response,error)
func(t transport)RoundTrip(r *http.Request)(*http.Response,error){return t(r)}
func fixture(cidr string,ts int64)string{return fmt.Sprintf("{\"cidr\":%q}\n{\"timestamp\":%d,\"copyright\":\"The Spamhaus Project - test fixture\"}\n",cidr,ts)}
func TestDropValidation(t *testing.T){
 now:=time.Unix(1800000000,0)
 b:=fixture("192.0.2.0/24",now.Unix())
 for _,tc:=range []struct{raw string;v6,ok bool}{{b,false,true},{b,true,false},{fixture("0.0.0.0/0",now.Unix()),false,false},{fixture("10.0.0.0/8",now.Unix()),false,false},{fixture("192.0.2.0/24",now.Add(-48*time.Hour).Unix()),false,false},{fixture("192.0.2.0/24",now.Add(time.Hour).Unix()),false,false},{b+"{}",false,false},{"{\"cidr\":\"192.0.2.0/24\"}",false,false}}{
  _,e:=parseDrop([]byte(tc.raw),tc.v6,now);if (e==nil)!=tc.ok{t.Fatalf("validation: %v",e)}
 }
 v,e:=parseDrop([]byte(b),false,now);if e!=nil{t.Fatal(e)};d:=Drop{V4:v}
 ok,e:=d.Match("192.0.2.1",now);if e!=nil||!ok{t.Fatal(ok,e)}
 ok,e=d.Match("198.51.100.1",now);if e!=nil||ok{t.Fatal(ok,e)}
 if _,e=d.Match("192.0.2.1",now.Add(48*time.Hour));e==nil{t.Fatal("stale match accepted")}
}
func TestDropFetch(t *testing.T){
 now:=time.Unix(1800000000,0);n:=0
 c:=&http.Client{Transport:transport(func(r *http.Request)(*http.Response,error){
  n++;if r.Method!="GET"||r.URL.Scheme!="https"||r.Header.Get("Auth-Key")!=""{t.Fatal("unexpected request")}
  v:="192.0.2.0/24";if r.URL.String()==drop6{v="2001:db8::/32"}else if r.URL.String()!=drop4{t.Fatal("unexpected endpoint")}
  return &http.Response{StatusCode:200,Body:io.NopCloser(strings.NewReader(fixture(v,now.Unix())))},nil
 })}
 d,e:=refresh(context.Background(),c,now);if e!=nil||n!=2||d.Status().IPv4!=1||d.Status().IPv6!=1{t.Fatal(d,e,n)}
 c.Transport=transport(func(*http.Request)(*http.Response,error){return &http.Response{StatusCode:403,Body:io.NopCloser(strings.NewReader("denied"))},nil})
 d,e=refresh(context.Background(),c,now);if e==nil||d.Status().IPv4!=0{t.Fatal("partial failure accepted")}
}
