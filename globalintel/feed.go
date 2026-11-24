// Package globalintel provides advisory-only network intelligence from fixed official endpoints.
// It does not contact indicator URLs, upload files, alter firewall rules or authorize remediation.
package globalintel

import (
 "bufio"
 "bytes"
 "context"
 "encoding/json"
 "errors"
 "fmt"
 "io"
 "net/http"
 "net/netip"
 "strings"
 "time"
)

const maxBytes=8<<20
const maxEntries=20000
const drop4="https://www.spamhaus.org/drop/drop_v4.json"
const drop6="https://www.spamhaus.org/drop/drop_v6.json"

type NetworkList struct {
 Prefixes []netip.Prefix `json:"-"`
 Timestamp int64 `json:"timestamp"`
 Copyright string `json:"copyright"`
}
type Drop struct { V4,V6 NetworkList }
type Status struct {
 IPv4 int `json:"ipv4"`
 IPv6 int `json:"ipv6"`
 Updated string `json:"updated"`
 Source string `json:"source"`
 Copyright4 string `json:"copyright4"`
 Copyright6 string `json:"copyright6"`
}
func client()*http.Client{return &http.Client{Timeout:25*time.Second,CheckRedirect:func(*http.Request,[]*http.Request)error{return http.ErrUseLastResponse}}}

func fetch(ctx context.Context,c *http.Client,endpoint string)([]byte,error){
 req,e:=http.NewRequestWithContext(ctx,http.MethodGet,endpoint,nil);if e!=nil{return nil,e}
 req.Header.Set("Accept","application/json, text/json, application/x-ndjson")
 req.Header.Set("User-Agent","ZEUVORYX-LAB/3.0.5")
 r,e:=c.Do(req);if e!=nil{return nil,e};defer r.Body.Close()
 if r.StatusCode!=200{return nil,fmt.Errorf("Spamhaus HTTP %d",r.StatusCode)}
 b,e:=io.ReadAll(io.LimitReader(r.Body,maxBytes+1));if e!=nil{return nil,e};if len(b)>maxBytes{return nil,errors.New("feed exceeds limit")};return b,nil
}

func parseDrop(b []byte,ipv6 bool,now time.Time)(NetworkList,error){
 var out NetworkList
 if len(b)>maxBytes{return out,errors.New("feed exceeds limit")}
 seen:=map[netip.Prefix]bool{}
 sc:=bufio.NewScanner(bytes.NewReader(b));sc.Buffer(make([]byte,4096),16384)
 meta:=false
 for sc.Scan(){
  if len(bytes.TrimSpace(sc.Bytes()))==0{continue}
  if meta{return NetworkList{},errors.New("unexpected record after metadata")}
  var r struct{ CIDR string `json:"cidr"`; Timestamp int64 `json:"timestamp"`; Copyright string `json:"copyright"` }
  if e:=json.Unmarshal(sc.Bytes(),&r);e!=nil{return NetworkList{},e}
  if r.CIDR==""{
   if r.Timestamp<=0||len(r.Copyright)>512||!strings.Contains(strings.ToLower(r.Copyright),"spamhaus"){return NetworkList{},errors.New("invalid source metadata")}
   t:=time.Unix(r.Timestamp,0)
   if t.After(now.Add(10*time.Minute))||now.Sub(t)>26*time.Hour{return NetworkList{},errors.New("stale or future feed")}
   out.Timestamp=r.Timestamp;out.Copyright=r.Copyright;meta=true;continue
  }
  p,e:=netip.ParsePrefix(r.CIDR)
  if e!=nil||p.Addr().Is4In6()||p.Addr().Is6()!=ipv6||p!=p.Masked()||!p.Addr().IsGlobalUnicast()||p.Addr().IsPrivate()||p.Bits()==0{return NetworkList{},errors.New("invalid prefix")}
  if !seen[p]{seen[p]=true;out.Prefixes=append(out.Prefixes,p)}
  if len(seen)>maxEntries{return NetworkList{},errors.New("too many prefixes")}
 }
 if e:=sc.Err();e!=nil{return NetworkList{},e}
 if !meta||len(out.Prefixes)==0{return NetworkList{},errors.New("incomplete source")}
 return out,nil
}
func Refresh(ctx context.Context)(Drop,error){return refresh(ctx,client(),time.Now())}
func refresh(ctx context.Context,c *http.Client,now time.Time)(Drop,error){
 var d Drop
 a,e:=fetch(ctx,c,drop4);if e!=nil{return d,e};v4,e:=parseDrop(a,false,now);if e!=nil{return d,e}
 b,e:=fetch(ctx,c,drop6);if e!=nil{return d,e};v6,e:=parseDrop(b,true,now);if e!=nil{return d,e}
 return Drop{V4:v4,V6:v6},nil
}
func(d Drop) Status() Status{
 t:=d.V4.Timestamp;if d.V6.Timestamp<t{t=d.V6.Timestamp}
 s:=Status{IPv4:len(d.V4.Prefixes),IPv6:len(d.V6.Prefixes),Source:"The Spamhaus Project DROP",Copyright4:d.V4.Copyright,Copyright6:d.V6.Copyright}
 if t>0{s.Updated=time.Unix(t,0).UTC().Format(time.RFC3339)};return s
}
func(d Drop) Match(raw string,now time.Time)(bool,error){
 a,e:=netip.ParseAddr(strings.TrimSpace(raw));if e!=nil||a.Zone()!=""{return false,errors.New("enter one valid IP address")}
 a=a.Unmap();v:=d.V6;if a.Is4(){v=d.V4}
 if len(v.Prefixes)==0||v.Timestamp<=0||now.Sub(time.Unix(v.Timestamp,0))>26*time.Hour||time.Unix(v.Timestamp,0).After(now.Add(10*time.Minute)){return false,errors.New("feed unavailable or stale")}
 for _,p:=range v.Prefixes{if p.Contains(a){return true,nil}}
 return false,nil
}
