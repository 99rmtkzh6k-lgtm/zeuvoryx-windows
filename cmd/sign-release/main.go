// Offline owner signing tool. Run only on the owner's trusted computer.
// Never commit the private key. The public key is safe to pin in the application.
package main
import("crypto/ed25519";"crypto/rand";"crypto/sha256";"encoding/hex";"encoding/json";"flag";"fmt";"io";"os";"path/filepath";"time";"github.com/zeuvoryx/zeuvoryx/releaseguard";"github.com/zeuvoryx/zeuvoryx/updater")
func fatal(e error){if e!=nil{fmt.Fprintln(os.Stderr,e);os.Exit(1)}}
func main(){
 gen:=flag.Bool("generate",false,"generate private key outside repository")
 key:=flag.String("key","","private key file outside repository")
 ver:=flag.String("version","","target semantic version")
 pkg:=flag.String("installer","","exact installer file")
 out:=flag.String("out","","existing output directory")
 flag.Parse();if *key==""{fatal(fmt.Errorf("provide -key"))}
 if *gen{pub,priv,e:=ed25519.GenerateKey(rand.Reader);fatal(e);f,e:=os.OpenFile(*key,os.O_WRONLY|os.O_CREATE|os.O_EXCL,0600);fatal(e);_,e=f.Write(priv);fatal(e);fatal(f.Close());fmt.Println(hex.EncodeToString(pub));return}
 if *out==""||*pkg==""{fatal(fmt.Errorf("provide -out and -installer"))}
 priv,e:=os.ReadFile(*key);fatal(e);if len(priv)!=ed25519.PrivateKeySize{fatal(fmt.Errorf("wrong private key size"))}
 seq,e:=updater.Sequence(*ver);fatal(e);if filepath.Base(*pkg)!=updater.InstallerName(*ver){fatal(fmt.Errorf("wrong installer filename"))}
 f,e:=os.Open(*pkg);fatal(e);h:=sha256.New();n,e:=io.Copy(h,io.LimitReader(f,releaseguard.MaxPackage+1));fatal(e);fatal(f.Close());if n<=0||n>releaseguard.MaxPackage{fatal(fmt.Errorf("wrong installer size"))}
 m:=releaseguard.Manifest{Sequence:seq,Version:*ver,Channel:"lab",Platform:"windows-amd64",URL:updater.ReleasesURL+"/download/v"+*ver+"-lab/"+updater.InstallerName(*ver),SHA256:hex.EncodeToString(h.Sum(nil)),Size:n,Expires:time.Now().Add(7*24*time.Hour).Unix()}
 raw,e:=json.Marshal(m);fatal(e);sig:=ed25519.Sign(ed25519.PrivateKey(priv),append([]byte(releaseguard.Domain),raw...))
 _,e=releaseguard.Verify(raw,sig,ed25519.PrivateKey(priv).Public().(ed25519.PublicKey),seq-1,"lab",time.Now());fatal(e)
 fatal(os.WriteFile(filepath.Join(*out,updater.ManifestName),raw,0600));fatal(os.WriteFile(filepath.Join(*out,updater.SignatureName),sig,0600))
 fmt.Println("Manifest and signature ready; verify and publish only with owner approval.")
}
