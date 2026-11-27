// Signed LAB updater. The production public key is pinned at build time.
// A missing key disables automatic execution. No private key belongs in this package.
package updater
import("context";"crypto/ed25519";"encoding/hex";"errors";"fmt";"io";"net/http";"net/url";"os";"path/filepath";"strings";"time";"github.com/zeuvoryx/zeuvoryx/releaseguard")
var OwnerPublicKeyHex string
const ManifestName="zeuvoryx-lab-manifest.json"
const SignatureName="zeuvoryx-lab-manifest.sig"
func ownerKey()(ed25519.PublicKey,error){b,e:=hex.DecodeString(OwnerPublicKeyHex);if e!=nil||len(b)!=ed25519.PublicKeySize{return nil,errors.New("chave oficial ainda não configurada")};return ed25519.PublicKey(b),nil}
func SignedEnabled()bool{_,e:=ownerKey();return e==nil}
func Sequence(v string)(uint64,error){
 a,ok:=version(v);if !ok||a[0]>999999||a[1]>999999||a[2]>999999{return 0,errors.New("invalid version sequence")}
 n:=a[0]*1000000000000+a[1]*1000000+a[2];if n==0{return 0,errors.New("zero version")};return n,nil
}
func assetBase(v string)string{return ReleasesURL+"/download/v"+v+"-lab/"}
func InstallerName(v string)string{return "ZEUVORYX-LAB-Setup-"+v+".exe"}
func verifyLab(raw,sig []byte,key ed25519.PublicKey,current,target string,now time.Time)(releaseguard.Manifest,error){
 seq,e:=Sequence(current);if e!=nil{return releaseguard.Manifest{},e}
 want,e:=Sequence(target);if e!=nil||want<=seq{return releaseguard.Manifest{},errors.New("update must be newer")}
 m,e:=releaseguard.Verify(raw,sig,key,seq,"lab",now);if e!=nil{return m,e}
 if m.Sequence!=want||m.Version!=target||m.URL!=assetBase(target)+InstallerName(target){return releaseguard.Manifest{},errors.New("release identity mismatch")}
 return m,nil
}
func allowedURL(u *url.URL)bool{
 if u==nil||u.Scheme!="https"||u.User!=nil||u.Port()!=""||u.Fragment!=""{return false}
 switch u.Hostname(){case "github.com","release-assets.githubusercontent.com","objects.githubusercontent.com":return true};return false
}
func downloadClient()*http.Client{return &http.Client{Timeout:5*time.Minute,CheckRedirect:func(req *http.Request,via []*http.Request)error{if len(via)>=5||!allowedURL(req.URL){return errors.New("unapproved update redirect")};return nil}}}
type fetcher func(context.Context,string,int64,io.Writer)error
func fetchWith(c *http.Client)fetcher{return func(ctx context.Context,addr string,limit int64,dst io.Writer)error{
 u,e:=url.Parse(addr);if e!=nil||!allowedURL(u){return errors.New("unapproved update URL")}
 req,e:=http.NewRequestWithContext(ctx,http.MethodGet,addr,nil);if e!=nil{return e};req.Header.Set("User-Agent","ZEUVORYX-signed-update")
 r,e:=c.Do(req);if e!=nil{return e};defer r.Body.Close()
 if r.StatusCode!=http.StatusOK{return fmt.Errorf("update HTTP %d",r.StatusCode)}
 if r.ContentLength>limit{return errors.New("oversized response")}
 n,e:=io.Copy(dst,io.LimitReader(r.Body,limit+1));if e!=nil{return e};if n>limit{return errors.New("oversized response")};return nil
}}
type staged struct{Path string; Manifest releaseguard.Manifest}
func stage(ctx context.Context,fetch fetcher,key ed25519.PublicKey,current,target,dir string,now time.Time)(staged,error){
 var result staged
 if len(key)!=ed25519.PublicKeySize{return result,errors.New("missing trusted key")}
 cur,e:=Sequence(current);if e!=nil{return result,e};next,e:=Sequence(target);if e!=nil||next<=cur{return result,errors.New("version not newer")}
 var raw,sig strings.Builder
 if e=fetch(ctx,assetBase(target)+ManifestName,releaseguard.MaxManifest,&raw);e!=nil{return result,e}
 if e=fetch(ctx,assetBase(target)+SignatureName,ed25519.SignatureSize,&sig);e!=nil{return result,e}
 m,e:=verifyLab([]byte(raw.String()),[]byte(sig.String()),key,current,target,now);if e!=nil{return result,e}
 tmp,e:=os.MkdirTemp(dir,"zeuvoryx-verified-");if e!=nil{return result,e}
 ok:=false;defer func(){if !ok{os.RemoveAll(tmp)}}()
 path:=filepath.Join(tmp,InstallerName(target))
 f,e:=os.OpenFile(path,os.O_CREATE|os.O_EXCL|os.O_WRONLY,0600);if e!=nil{return result,e}
 e=fetch(ctx,m.URL,m.Size,f);if e==nil{e=f.Sync()};ce:=f.Close();if e==nil{e=ce};if e!=nil{return result,e}
 in,e:=os.Open(path);if e!=nil{return result,e};e=releaseguard.VerifyPackage(in,m);ce=in.Close();if e==nil{e=ce};if e!=nil{return result,e}
 ok=true;return staged{path,m},nil
}
// stageGate prevents two in-process requests from staging the same installer.
// It cannot lock Safari, other app instances, or external GitHub downloads.
var stageGate DownloadGate

func StageAuthorized(ctx context.Context,target string)(staged,error){
 key,e:=ownerKey();if e!=nil{return staged{},e}
 release,ok:=stageGate.Begin(InstallerName(target))
 if !ok{return staged{},errors.New(DuplicateDownloadMessage)}
 defer release()
 return stage(ctx,fetchWith(downloadClient()),key,Current,target,os.TempDir(),time.Now())
}

// CompletedVersion checks the installer-provided version against this executable.
// This is a UI receipt, not a cryptographic attestation of the installation.
func CompletedVersion(target string) string { if target==Current {return Current};return "" }
