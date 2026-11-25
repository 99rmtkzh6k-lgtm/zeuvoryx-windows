// Package signatureav calls an independent ClamAV installation, without changing scanned files.
package signatureav

import (
 "bufio"
 "context"
 "errors"
 "fmt"
 "io"
 "os"
 "os/exec"
 "path/filepath"
 "strconv"
 "strings"
 "time"
)

type Finding struct {
 Path string `json:"path"`
 Signature string `json:"signature"`
 Kind string `json:"kind"`
}
type Report struct {
 Engine string `json:"engine"`
 Version string `json:"engine_database_version,omitempty"`
 Root string `json:"root"`
 Started time.Time `json:"started"`
 Finished time.Time `json:"finished"`
 ExitCode int `json:"exit_code"`
 Scanned int64 `json:"scanned_files"`
 Infected int64 `json:"engine_reported_infected_files"`
 KnownSignatures int64 `json:"known_signatures"`
 Errors int64 `json:"engine_reported_errors"`
 Findings []Finding `json:"findings"`
 Partial bool `json:"partial"`
 Status string `json:"status"`
 Raw string `json:"engine_output"`
 Problem string `json:"problem,omitempty"`
 Coverage string `json:"coverage_note"`
}
type bounded struct { b []byte; overflow bool }
func (w *bounded) Write(p []byte)(int,error) {
 const max = 8<<20
 n:=max-len(w.b)
 if n<len(p) { w.overflow=true; if n>0 {w.b=append(w.b,p[:n]...)} } else {w.b=append(w.b,p...)}
 return len(p),nil
}
func command(ctx context.Context, exe string,args ...string)*exec.Cmd {
 c:=exec.CommandContext(ctx,exe,args...)
 c.Dir=filepath.Dir(exe)
 configure(c)
 return c
}
func Version(ctx context.Context, exe string)(string,error){
 if !filepath.IsAbs(exe){return "",errors.New("absolute ClamAV executable required")}
 c:=command(ctx,exe,"--version")
 var b bounded
 c.Stdout=&b;c.Stderr=&b
 e:=c.Run()
 v:=strings.TrimSpace(string(b.b))
 if e!=nil{return "",e}
 if b.overflow||!strings.HasPrefix(v,"ClamAV "){return "",errors.New("unexpected ClamAV version response")}
 return v,nil
}
// Version freshness is a conservative check, not proof of database authenticity.
func recentDatabase(v string, now time.Time) bool {
 parts:=strings.SplitN(v,"/",3);if len(parts)!=3{return false}
 if n,e:=strconv.Atoi(parts[1]);e!=nil||n<=0{return false}
 date,e:=time.Parse("Mon Jan 2 15:04:05 2006",strings.TrimSpace(parts[2]))
 if e!=nil{return false}
 age:=now.UTC().Sub(date)
 return age>=-24*time.Hour && age<=7*24*time.Hour
}
func parse(r *Report, raw string, code int, overflow bool) {
 r.Raw=raw; r.ExitCode=code; r.Scanned=-1; r.Infected=-1; r.KnownSignatures=-1; r.Errors=-1
 r.Partial=overflow
 s:=bufio.NewScanner(strings.NewReader(raw));s.Buffer(make([]byte,4096),1<<20)
 summary:=false
 for s.Scan(){
  line:=strings.TrimSpace(s.Text())
  if strings.Contains(line,"SCAN SUMMARY"){summary=true;continue}
  if strings.HasSuffix(line," FOUND"){
   v:=strings.TrimSuffix(line," FOUND"); i:=strings.LastIndex(v,": ")
   if i<=0||i+2>=len(v){r.Partial=true;continue}
   label:=v[i+2:]; lower:=strings.ToLower(label); kind:="signature_detection"
   if strings.Contains(lower,"eicar"){kind="test_eicar"}else if strings.Contains(lower,"heuristics.")||strings.Contains(lower,"encrypted"){kind="heuristic_or_limit";r.Partial=true}
   r.Findings=append(r.Findings,Finding{Path:v[:i],Signature:label,Kind:kind})
  }
  if strings.HasSuffix(line," ERROR")||strings.HasPrefix(line,"ERROR:")||strings.HasPrefix(line,"WARNING:"){r.Partial=true}
  if !summary{continue}
  k,v,ok:=strings.Cut(line,":");if !ok{continue}
  n,e:=strconv.ParseInt(strings.TrimSpace(v),10,64);if e!=nil||n<0{continue}
  switch k{case "Scanned files":r.Scanned=n;case "Infected files":r.Infected=n;case "Known viruses":r.KnownSignatures=n;case "Total errors":r.Errors=n}
 }
 if s.Err()!=nil{r.Partial=true}
 if !summary||r.Scanned<=0||r.Infected<0||r.KnownSignatures<=0||r.Errors>0||r.Infected!=int64(len(r.Findings)){r.Partial=true}
 if code!=0&&code!=1{r.Partial=true}
 if (code==0&&r.Infected!=0)||(code==1&&r.Infected<=0){r.Partial=true}
 if r.Partial{r.Status="inconclusive_or_partial"}else if code==1{r.Status="detections_reported"}else{r.Status="no_detection_reported"}
}
// Scan uses a single ClamAV process for a local file or tree, never a shell, never --remove or --move.
// db is optional and only for explicitly selected laboratory databases; normal desktop use passes "".
func Scan(ctx context.Context, exe, root, db string) (r Report, err error) {
 r=Report{Engine:"ClamAV / independent signatures",Started:time.Now().UTC(),ExitCode:-1,Scanned:-1,Infected:-1,KnownSignatures:-1,Errors:-1,Findings:[]Finding{},Partial:true,Status:"inconclusive_or_partial",Coverage:"ClamAV has size, archive, format and access limits; no memory or resident protection. No detections is not a guarantee. No files are sent or removed."}
 defer func(){r.Finished=time.Now().UTC();if err!=nil{r.Problem=err.Error();r.Partial=true;r.Status="inconclusive_or_partial"}}()
 if !filepath.IsAbs(exe){return r,errors.New("absolute ClamAV executable required")}
 root,err=filepath.Abs(root);if err!=nil{return r,err};r.Root=root
 if strings.HasPrefix(root,`\\`){return r,errors.New("network paths not supported")}
 st,err:=os.Lstat(root);if err!=nil{return r,err}
 if !st.Mode().IsRegular()&&!st.IsDir(){return r,errors.New("regular file or directory required")}
 vc,cancel:=context.WithTimeout(ctx,15*time.Second);r.Version,err=Version(vc,exe);cancel()
 if err!=nil{return r,fmt.Errorf("ClamAV version: %w",err)}
 args:=[]string{"--stdout","--infected","--recursive","--alert-exceeds-max=yes","--follow-dir-symlinks=0","--follow-file-symlinks=0"}
 if db!=""{if !filepath.IsAbs(db){return r,errors.New("absolute laboratory DB required")};args=append(args,"--database="+db)}
 args=append(args,"--",root)
 cmd:=command(ctx,exe,args...)
 var out bounded
 // Both streams share one synchronized os/exec writer.
 cmd.Stdout=&out;cmd.Stderr=&out
 e:=cmd.Run();code:=0
 if e!=nil{var ee *exec.ExitError;if errors.As(e,&ee){code=ee.ExitCode()}else{code=-1}}
 parse(&r,string(out.b),code,out.overflow)
 if !recentDatabase(r.Version,time.Now()){r.Partial=true;r.Status="inconclusive_or_partial";return r,errors.New("ClamAV daily database date unknown or older than 7 days") }
 if ctx.Err()!=nil{return r,ctx.Err()}
 if e!=nil&&code!=1{return r,fmt.Errorf("ClamAV scan: %w",e)}
 if r.Partial{return r,errors.New("ClamAV report partial or inconsistent; see engine output")}
 return r,nil
}
var _ io.Writer=(*bounded)(nil)
