package updater

import (
 "context"
 "encoding/json"
 "fmt"
 "io"
 "net/http"
 "strconv"
 "strings"
 "time"
)

const Current = "3.0.10"
const ReleasesURL = "https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/releases"
const endpoint = "https://api.github.com/repos/99rmtkzh6k-lgtm/zeuvoryx-windows/releases"

type Status struct {
 Current string `json:"current"`
 Latest string `json:"latest"`
 Available bool `json:"available"`
 Found bool `json:"found"`
 Preview bool `json:"preview"`
 DownloadURL string `json:"download_url,omitempty"`
}
type release struct {
 Tag string `json:"tag_name"`
 Draft bool `json:"draft"`
 Pre bool `json:"prerelease"`
 Assets []struct { Name string `json:"name"` } `json:"assets"`
}
func version(s string) ([3]uint64, bool) {
 var v [3]uint64
 s = strings.TrimPrefix(s, "v")
 p := strings.Split(s, ".")
 if len(p)!=3 { return v,false }
 for i,x:=range p {
  if x=="" || (len(x)>1 && x[0]=='0') {return v,false}
  for _,c:=range x {if c<'0'||c>'9' {return v,false}}
  n,e:=strconv.ParseUint(x,10,32);if e!=nil{return v,false};v[i]=n
 }
 return v,true
}
func newer(a,b [3]uint64) bool {
 for i:=range a {if a[i]!=b[i]{return a[i]>b[i]}}
 return false
}
func selectRelease(list []release, current, platform, arch string, includePreview bool) Status {
 result:=Status{Current:current}
 installed,ok:=version(current);if !ok{return result}
 var best [3]uint64
 for _,r:=range list {
  if r.Draft || !r.Pre || !includePreview || !strings.HasPrefix(r.Tag,"v") || !strings.HasSuffix(r.Tag,"-lab") {continue}
  tag := strings.TrimSuffix(r.Tag,"-lab")
  v,ok:=version(tag);if !ok {continue}
  name:=""
  switch {
  case platform=="windows" && arch=="amd64":name="ZEUVORYX-LAB-Setup-"+strings.TrimPrefix(tag,"v")+".exe"
  case platform=="darwin" && (arch=="amd64"||arch=="arm64"):name="ZEUVORYX-LAB-macOS-"+arch+"-"+strings.TrimPrefix(tag,"v")+".pkg"
  default:continue
  }
  for _,a:=range r.Assets {
   if a.Name==name && (!result.Found||newer(v,best)){
    result.Found=true;result.Preview=r.Pre;best=v;result.Latest=strings.TrimPrefix(tag,"v");result.DownloadURL=ReleasesURL+"/download/"+r.Tag+"/"+name
   }
  }
 }
 result.Available=result.Found&&newer(best,installed)
 return result
}
func Check(ctx context.Context,platform,arch string,includePreview bool)(Status,error){
 client:=&http.Client{Timeout:8*time.Second,CheckRedirect:func(*http.Request,[]*http.Request)error{return http.ErrUseLastResponse}}
 return check(ctx,client,endpoint,platform,arch,includePreview)
}
func check(ctx context.Context,client *http.Client,url,platform,arch string,includePreview bool)(Status,error){
 result:=Status{Current:Current}
 req,e:=http.NewRequestWithContext(ctx,http.MethodGet,url,nil);if e!=nil{return result,e}
 req.Header.Set("Accept","application/vnd.github+json")
 req.Header.Set("User-Agent","ZEUVORYX-update-check")
 resp,e:=client.Do(req);if e!=nil{return result,e};defer resp.Body.Close()
 if resp.StatusCode!=http.StatusOK{return result,fmt.Errorf("GitHub: HTTP %d",resp.StatusCode)}
 const max=2<<20
 b,e:=io.ReadAll(io.LimitReader(resp.Body,max+1));if e!=nil{return result,e};if len(b)>max{return result,fmt.Errorf("resposta excede limite")}
 var list []release
 if e=json.Unmarshal(b,&list);e!=nil{return result,e}
 return selectRelease(list,Current,platform,arch,includePreview),nil
}
