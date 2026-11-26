//go:build windows

package main
import("archive/zip";"context";"crypto/sha256";"encoding/hex";"os";"path/filepath";"testing";"syscall";"time";"encoding/json";"github.com/zeuvoryx/zeuvoryx/familygate";"github.com/zeuvoryx/zeuvoryx/audit";"github.com/zeuvoryx/zeuvoryx/engine")
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

func TestP0WindowsLockedFileContinuesAndPersists(t *testing.T){
 root:=t.TempDir()
 locked:=filepath.Join(root,"locked.tmp")
 if err:=os.WriteFile(locked,[]byte("benign locked fixture"),0600);err!=nil{t.Fatal(err)}
 if err:=os.WriteFile(filepath.Join(root,"readable.txt"),[]byte("benign readable fixture"),0600);err!=nil{t.Fatal(err)}
 name,err:=syscall.UTF16PtrFromString(locked);if err!=nil{t.Fatal(err)}
 h,err:=syscall.CreateFile(name,syscall.GENERIC_READ,0,nil,syscall.OPEN_EXISTING,syscall.FILE_ATTRIBUTE_NORMAL,0);if err!=nil{t.Fatal(err)}
 defer syscall.CloseHandle(h)
 ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second);defer cancel()
 sc:=engine.New()
 for i:=0;i<2;i++{
  r,e:=sc.Scan(ctx,root,nil);r.Summarize()
  if e!=nil||r.Cancelled||!r.Partial||r.Files!=1||r.Skipped!=1||r.ErrorCount!=1||len(r.SkippedDetails)!=1||r.SkippedDetails[0].Path!=locked||r.SkippedDetails[0].CodigoWindows!=32{t.Fatalf("scan %d: files=%d skipped=%d errors=%d details=%+v err=%v",i,r.Files,r.Skipped,r.ErrorCount,r.SkippedDetails,e)}
  store,e:=audit.Open(t.TempDir());if e!=nil{t.Fatal(e)}
  id,e:=audit.NewID();if e!=nil{t.Fatal(e)}
  if e=store.Finish(id,"custom",r,audit.Event{Files:r.Files,Skipped:r.Skipped,ErrorCount:r.ErrorCount,Partial:r.Partial});e!=nil{t.Fatal(e)}
  raw,e:=store.ReadReport(id);if e!=nil{t.Fatal(e)}
  var out engine.Report;if e=json.Unmarshal(raw,&out);e!=nil{t.Fatal(e)}
  if len(out.SkippedDetails)!=1||out.SkippedDetails[0].CodigoWindows!=32{t.Fatal(string(raw))}
  data,_:=json.Marshal(out.SkippedDetails);t.Logf("Windows actual exclusive handle persisted skipped_details=%s",data)
 }
 q:=filepath.Join(t.TempDir(),"quarantine")
 st:=&State{scanner:sc,quarantine:&engine.Quarantine{Dir:q}}
 r:=st.scanRoots(ctx,[]string{root},"",nil);r.Summarize()
 if r.Cancelled||r.Files!=1||r.Skipped!=1||r.ErrorCount!=1||len(r.SkippedDetails)!=1||r.SkippedDetails[0].CodigoWindows!=32{t.Fatalf("root scan coverage: %+v",r)}
}

func TestP0GameRecordsPersistAcrossReads(t *testing.T){
 t.Setenv("APPDATA",t.TempDir())
 previous:=family314;family314=familygate.Open(filepath.Join(t.TempDir(),"family.json"));defer func(){family314=previous}()
 for _,id:=range []string{"zvx-tetris","zvx-poker","zvx-mines","zvx-memory"}{
  ok,e:=saveGameRecord315(id,42);if e!=nil||!ok{t.Fatal(id,e)}
  ok,e=saveGameRecord315(id,12);if e!=nil||!ok{t.Fatal(id,e)}
  v,e:=getGameRecord315(id);if e!=nil||v!=42{t.Fatal(id,v,e)}
 }
 path,e:=records315Path();if e!=nil{t.Fatal(e)}
 b,e:=os.ReadFile(path);if e!=nil{t.Fatal(e)};var records map[string]int
 if e=json.Unmarshal(b,&records);e!=nil||len(records)!=4{t.Fatal(string(b),e)}
 if ok,e:=saveGameRecord315("../invalid",99);e==nil||ok{t.Fatal("invalid game accepted")}
 t.Logf("persisted game records: %s",b)
}
