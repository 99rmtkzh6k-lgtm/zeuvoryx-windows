package updater

import (
 "context"
 "net/http"
 "net/http/httptest"
 "testing"
)
func TestSelectRelease(t *testing.T){
 var r release
 r.Tag="v3.1.0-lab"; r.Pre=true; r.Assets=append(r.Assets,struct{Name string `json:"name"`}{"ZEUVORYX-LAB-Setup-3.1.0.exe"})
 for _,tc:=range []struct{current,os,arch string; want bool}{
 {"3.0.2","windows","amd64",true},{"3.1.0","windows","amd64",false},{"4.0.0","windows","amd64",false},{"3.0.2","darwin","arm64",false},{"3.0.2","windows","arm64",false},
 }{if got:=selectRelease([]release{r},tc.current,tc.os,tc.arch,true).Available;got!=tc.want{t.Fatal(tc,got)}}
 r.Pre=true;if selectRelease([]release{r},Current,"windows","amd64",false).Found{t.Fatal("prerelease")}
 if !selectRelease([]release{r},Current,"windows","amd64",true).Available{t.Fatal("test channel")}
 r.Pre=false;if selectRelease([]release{r},Current,"windows","amd64",true).Found{t.Fatal("stable channel")}
 r.Pre=true;r.Draft=true;if selectRelease([]release{r},Current,"windows","amd64",true).Found{t.Fatal("draft")}
}
func TestVersions(t *testing.T){
 for _,s:=range []string{"3.0.2-beta","03.0.2","3.0","3.0.+2","3.0.4294967296","3.0.-2"}{if _,ok:=version(s);ok{t.Fatal(s)}}
 a,_:=version("3.10.0");b,_:=version("3.9.9");if !newer(a,b){t.Fatal("numeric ordering")}
}
func TestHTTP(t *testing.T){
 for _,tc:=range []struct{code int;body string;fail bool}{
 {200,"[]",false},{200,"[",true},{403,"{}",true},{500,"{}",true},
 }{
 srv:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.WriteHeader(tc.code);w.Write([]byte(tc.body))}))
 _,e:=check(context.Background(),srv.Client(),srv.URL,"windows","amd64",false);srv.Close();if (e!=nil)!=tc.fail{t.Fatal(tc,e)}
 }
}

func TestLabIsolation(t *testing.T){
 var r release; r.Pre=true;r.Tag="v3.0.6-lab";r.Assets=append(r.Assets,struct{Name string `json:"name"`}{"ZEUVORYX-Setup-3.0.6.exe"})
 if selectRelease([]release{r},Current,"windows","amd64",true).Found{t.Fatal("wrong installer identity")}
 r.Assets[0].Name="ZEUVORYX-LAB-Setup-3.0.6.exe"
 if !selectRelease([]release{r},Current,"windows","amd64",true).Available{t.Fatal("lab update missing")}
 r.Tag="v3.0.6";if selectRelease([]release{r},Current,"windows","amd64",true).Found{t.Fatal("wrong tag")}
}

func TestDownloadURL(t *testing.T) {
 var r release; r.Pre=true;r.Tag="v3.0.6-lab";r.Assets=append(r.Assets,struct{Name string `json:"name"`}{"ZEUVORYX-LAB-Setup-3.0.6.exe"})
 v:=selectRelease([]release{r},"3.0.5","windows","amd64",true)
 if !v.Available||v.DownloadURL!=ReleasesURL+"/download/v3.0.6-lab/ZEUVORYX-LAB-Setup-3.0.6.exe"{t.Fatal(v)}
 r.Tag="v3.0.6-lab/../evil";if selectRelease([]release{r},"3.0.5","windows","amd64",true).DownloadURL!=""{t.Fatal("invalid tag")}
}
