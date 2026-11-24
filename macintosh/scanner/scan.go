// Read-only Macintosh laboratory scanner. No elevation, network or remediation.
package scanner
import("context";"crypto/sha256";"encoding/hex";"errors";"io";"os";"path/filepath";"strings";"syscall";"time")

const maxFile int64=256<<20
const maxTotal int64=2<<30
const maxFiles=10000
type Finding struct{Path string `json:"path"`;Kind string `json:"kind"`; SHA256 string `json:"sha256"`}
type Report struct{Files int `json:"files"`;Bytes int64 `json:"bytes"`;Skipped int `json:"skipped"`;Partial bool `json:"partial"`;Findings []Finding `json:"findings"`}
var marker=[]byte("X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*")
type contextReader struct{ctx context.Context;r io.Reader}
func(r contextReader)Read(p []byte)(int,error){if e:=r.ctx.Err();e!=nil{return 0,e};return r.r.Read(p)}
func inspect(ctx context.Context,path string,st os.FileInfo)(*Finding,error){
 f,e:=os.OpenFile(path,os.O_RDONLY|syscall.O_NOFOLLOW|syscall.O_NONBLOCK,0);if e!=nil{return nil,e};defer f.Close()
 before,e:=f.Stat();if e!=nil||!before.Mode().IsRegular()||!os.SameFile(st,before)||before.Size()!=st.Size(){return nil,errors.New("file changed")}
 h:=sha256.New()
 n,e:=io.Copy(h,io.LimitReader(contextReader{ctx,f},st.Size()+1));if e!=nil||n!=st.Size(){return nil,errors.New("incomplete hash")}
 after,e:=f.Stat();if e!=nil||after.Size()!=before.Size()||!after.ModTime().Equal(before.ModTime()){return nil,errors.New("file changed")}
 hash:=hex.EncodeToString(h.Sum(nil))
 test:=sha256.Sum256(marker)
 if hash==hex.EncodeToString(test[:]){return &Finding{path,"test_eicar",hash},nil}
 name:=strings.ToLower(filepath.Base(path))
 for _,ext:=range []string{".pdf.app",".pdf.command",".jpg.command",".txt.command",".docx.app"}{if strings.HasSuffix(name,ext){return &Finding{path,"advisory_double_extension",hash},nil}}
 return nil,nil
}
func Scan(ctx context.Context,root string)(Report,error){
 r:=Report{Findings:[]Finding{}}
 abs,e:=filepath.Abs(root);if e!=nil{return r,e}
 abs,e=filepath.EvalSymlinks(abs);if e!=nil{return r,e}
 // Explicit user selection only. Never automatically traverse protected system paths.
 ctx,cancel:=context.WithTimeout(ctx,2*time.Minute);defer cancel()
 stop:=errors.New("scan budget reached")
 e=filepath.WalkDir(abs,func(path string,d os.DirEntry,err error)error{
  if ctx.Err()!=nil{return ctx.Err()}
  if err!=nil{r.Skipped++;return nil}
  if d.Type()&os.ModeSymlink!=0{r.Skipped++;return nil}
  if d.IsDir(){return nil}
  if !d.Type().IsRegular(){r.Skipped++;return nil}
  st,e:=d.Info();if e!=nil{r.Skipped++;return nil}
  if r.Files>=maxFiles||r.Bytes>=maxTotal{return stop}
  if st.Size()<0||st.Size()>maxFile||st.Size()>maxTotal-r.Bytes{r.Skipped++;return nil}
  hit,e:=inspect(ctx,path,st);if e!=nil{r.Skipped++;return nil}
  r.Files++;r.Bytes+=st.Size()
  if hit!=nil{if len(r.Findings)>=1000{return stop};r.Findings=append(r.Findings,*hit)}
  return nil
 })
 r.Partial=r.Skipped>0||e!=nil
 if e!=nil&&!errors.Is(e,stop)&&!errors.Is(e,context.Canceled)&&!errors.Is(e,context.DeadlineExceeded){return r,e}
 return r,nil
}
