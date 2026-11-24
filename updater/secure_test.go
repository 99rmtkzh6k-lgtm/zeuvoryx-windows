package updater
import("bytes";"context";"crypto/ed25519";"crypto/rand";"crypto/sha256";"encoding/hex";"encoding/json";"errors";"io";"net/url";"os";"testing";"time";"github.com/zeuvoryx/zeuvoryx/releaseguard")
func TestSignedStage(t *testing.T){
 pub,priv,e:=ed25519.GenerateKey(rand.Reader);if e!=nil{t.Fatal(e)}
 now:=time.Unix(1800000000,0);data:=[]byte("inert test package");sum:=sha256.Sum256(data);seq,_:=Sequence("3.0.6")
 m:=releaseguard.Manifest{Sequence:seq,Version:"3.0.6",Channel:"lab",Platform:"windows-amd64",URL:assetBase("3.0.6")+InstallerName("3.0.6"),SHA256:hex.EncodeToString(sum[:]),Size:int64(len(data)),Expires:now.Add(time.Hour).Unix()}
 raw,_:=json.Marshal(m);sig:=ed25519.Sign(priv,append([]byte(releaseguard.Domain),raw...))
 for _,tc:=range []struct{name string;body,sig,installer []byte;failURL string;current,target string;want bool}{
 {"valid",raw,sig,data,"","3.0.5","3.0.6",true},
 {"signature",raw,make([]byte,64),data,"","3.0.5","3.0.6",false},
 {"tamper",raw,sig,[]byte("inert test packagE"),"","3.0.5","3.0.6",false},
 {"truncated",raw,sig,data[:3],"","3.0.5","3.0.6",false},
 {"network",raw,sig,data,m.URL,"3.0.5","3.0.6",false},
 {"manifest network",raw,sig,data,assetBase("3.0.6")+ManifestName,"3.0.5","3.0.6",false},
 {"rollback",raw,sig,data,"","3.0.6","3.0.6",false},
 {"other release",raw,sig,data,"","3.0.5","3.0.7",false},
 }{t.Run(tc.name,func(t *testing.T){
 dir:=t.TempDir();f:=func(_ context.Context,u string,max int64,w io.Writer)error{
 if u==tc.failURL{return errors.New("simulated network failure")}
 b:=tc.installer;if u==assetBase(tc.target)+ManifestName{b=tc.body};if u==assetBase(tc.target)+SignatureName{b=tc.sig}
 if int64(len(b))>max{return errors.New("too large")};_,e:=w.Write(b);return e}
 got,e:=stage(context.Background(),f,pub,tc.current,tc.target,dir,now)
 if (e==nil)!=tc.want{t.Fatalf("%+v %v",got,e)}
 entries,_:=os.ReadDir(dir);if !tc.want&&len(entries)!=0{t.Fatal("failed update left a package")}
 if tc.want{b,e:=os.ReadFile(got.Path);if e!=nil||!bytes.Equal(b,data){t.Fatal("wrong staged file")}}
 })}
}
func TestUpdateTrustAndURL(t *testing.T){
 for _,x:=range []string{"http://github.com/test","https://github.com.evil.test/test","https://github.com:9999/test","https://x@github.com/test","https://github.com/test#x"}{u,_:=url.Parse(x);if allowedURL(u){t.Fatal(x)}}
 for _,x:=range []string{"https://github.com/test","https://release-assets.githubusercontent.com/test","https://objects.githubusercontent.com/test"}{u,_:=url.Parse(x);if !allowedURL(u){t.Fatal(x)}}
 for _,x:=range []string{"0.0.0","3.01.2","3.0.1000000","3.0.5-lab"}{if _,e:=Sequence(x);e==nil{t.Fatal(x)}}
 old:=OwnerPublicKeyHex;defer func(){OwnerPublicKeyHex=old}();OwnerPublicKeyHex="";if SignedEnabled(){t.Fatal("empty key")}
 OwnerPublicKeyHex=hex.EncodeToString(make([]byte,32));if !SignedEnabled(){t.Fatal("key shape")}
}

func TestCompletedVersion(t *testing.T){for _,x:=range []string{"","3.0.4","3.0.6","3.0.5-lab","v3.0.5"}{if CompletedVersion(x)!=""{t.Fatal(x)}};if CompletedVersion(Current)!=Current{t.Fatal("current version not acknowledged")}}
