//go:build windows

package main
import("archive/zip";"context";"crypto/sha256";"encoding/hex";"os";"path/filepath";"testing";"github.com/zeuvoryx/zeuvoryx/engine")
func Test306ControlledRoots(t *testing.T){
 root:=t.TempDir();q:=filepath.Join(root,"quarantine");if e:=os.Mkdir(q,0700);e!=nil{t.Fatal(e)}
 payload:=make([]byte,2<<20);for i:=range payload{payload[i]=byte(i%251)}
 sum:=sha256.Sum256(payload)
 for _,name:=range []string{"large.bin","other.bin","quarantine/excluded.bin"}{if e:=os.WriteFile(filepath.Join(root,name),payload,0600);e!=nil{t.Fatal(e)}}
 f,e:=os.Create(filepath.Join(root,"fixture.zip"));if e!=nil{t.Fatal(e)};z:=zip.NewWriter(f);entry,e:=z.Create("ordinary.txt");if e!=nil{t.Fatal(e)};if _,e=entry.Write([]byte("harmless fixture"));e!=nil{t.Fatal(e)};if e=z.Close();e!=nil{t.Fatal(e)};if e=f.Close();e!=nil{t.Fatal(e)}
 sc:=engine.New();sc.IOC[hex.EncodeToString(sum[:])]="synthetic fixture";st:=&State{scanner:sc,quarantine:&engine.Quarantine{Dir:q}}
 r:=st.scanRoots(context.Background(),[]string{root},"",nil)
 if r.Cancelled||r.Files!=3||len(r.Findings)!=2||r.Skipped!=0{t.Fatalf("unexpected root coverage: files=%d hits=%d skipped=%d cancelled=%v",r.Files,len(r.Findings),r.Skipped,r.Cancelled)}
 for _,hit:=range r.Findings{if hit.SHA256!=hex.EncodeToString(sum[:]){t.Fatal("hash changed")}}
}
