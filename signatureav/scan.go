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
 "github.com/zeuvoryx/zeuvoryx/auditmetrics"
)

type Finding struct {
 Path string `json:"path"`
 Signature string `json:"signature"`
 Kind string `json:"kind"`
}
type Report struct {
 Engine string `json:"engine"`
 Version string `json:"engine_database_version,omitempty"`
 DatabaseStatus string `json:"database_status,omitempty"`
 Root string `json:"root"`
 Roots []string `json:"roots,omitempty"`
 UnavailableRoots []string `json:"unavailable_roots,omitempty"`
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
 Evidence auditmetrics.SignatureEvidence `json:"evidence_315"`
}
type bounded struct { b []byte; overflow bool; observer lineObserver314 }
func (w *bounded) Write(p []byte)(int,error) {
 w.observer.feed(p)
 const max = 8<<20
 n:=max-len(w.b)
 if n<len(p) { w.overflow=true; if n>0 {w.b=append(w.b,p[:n]...)} } else {w.b=append(w.b,p...)}
 return len(p),nil
}
func command(ctx context.Context, exe string,args ...string)*exec.Cmd {
 c:=exec.CommandContext(ctx,exe,args...)
 c.WaitDelay=5*time.Second
 c.Dir=filepath.Dir(exe)
 configure(c)
 return c
}
// BundledDatabase is the same directory written by the integrated FreshClam.
func BundledDatabase(exe string) string {return filepath.Join(filepath.Dir(exe),"database")}
func officialBundleOnly(exe,db string)bool{return db!=""&&strings.EqualFold(filepath.Clean(db),filepath.Clean(BundledDatabase(exe)))}
func databaseArgs(db string)([]string,error){
 if db==""{return nil,nil}
 if !filepath.IsAbs(db){return nil,errors.New("absolute ClamAV DB required")}
 return []string{"--database="+db},nil
}
func Version(ctx context.Context, exe string)(string,error){return VersionAt(ctx,exe,"")}
func VersionAt(ctx context.Context, exe,db string)(string,error){
 if !filepath.IsAbs(exe){return "",errors.New("absolute ClamAV executable required")}
 args,e:=databaseArgs(db);if e!=nil{return "",e}
 c:=command(ctx,exe,append(args,"--version")...)
 var b bounded
 c.Stdout=&b;c.Stderr=&b
 e=c.Run()
 v:=strings.TrimSpace(string(b.b))
 if e!=nil{return "",e}
 if b.overflow||!strings.HasPrefix(v,"ClamAV "){return "",errors.New("unexpected ClamAV version response")}
 return v,nil
}
// Version freshness is a conservative check, not proof of database authenticity.
// DatabaseFresh validates a reported daily database version and date; not an authenticity proof.
func DatabaseFresh(v string, now time.Time) bool {return recentDatabase(v,now)}
// DatabaseStatus distinguishes a known older database from absent or implausible metadata.
// A valid-looking header is not proof that every signature file can be loaded.
func DatabaseStatus(v string,now time.Time)string{
 parts:=strings.SplitN(v,"/",3);if len(parts)!=3||!strings.HasPrefix(parts[0],"ClamAV "){return "unknown"}
 if n,e:=strconv.Atoi(parts[1]);e!=nil||n<=0{return "unknown"}
 date,e:=time.Parse("Mon Jan 2 15:04:05 2006",strings.TrimSpace(parts[2]));if e!=nil{return "unknown"}
 age:=now.UTC().Sub(date)
 if age < -24*time.Hour{return "unknown"}
 if age>7*24*time.Hour{return "stale"}
 return "fresh"
}
func recentDatabase(v string,now time.Time)bool{return DatabaseStatus(v,now)=="fresh"}
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
// Desktop callers supply their integrated database explicitly; generic callers may use ClamAV defaults.
func Scan(ctx context.Context, exe, root, db string) (Report,error) {
 return ScanRoots(ctx,exe,[]string{root},db)
}
// ScanRoots loads the database once for all eligible roots. Counts belong to the batch, not each root.
func ScanRoots(ctx context.Context, exe string, roots []string, db string) (r Report, err error) {
 r=Report{Engine:"ClamAV / independent signatures",Started:time.Now().UTC(),ExitCode:-1,Scanned:-1,Infected:-1,KnownSignatures:-1,Errors:-1,Findings:[]Finding{},Partial:true,Status:"inconclusive_or_partial",Coverage:"ClamAV has size, archive, format and access limits; no memory or resident protection. No detections is not a guarantee. No files are sent or removed."}
 defer func(){r.Finished=time.Now().UTC();if err!=nil{r.Problem=err.Error();r.Partial=true;r.Status="inconclusive_or_partial"};kinds:=make([]string,0,len(r.Findings));for _,f:=range r.Findings{kinds=append(kinds,f.Kind)};r.Evidence=auditmetrics.Signature(r.Scanned,r.Infected,r.KnownSignatures,r.Errors,kinds,r.Partial);if r.Evidence.Partial{r.Partial=true;r.Status="inconclusive_or_partial";if err==nil{err=errors.New("ClamAV evidence inconsistent or partial");r.Problem=err.Error()}}}()
 if !filepath.IsAbs(exe){return r,errors.New("absolute ClamAV executable required")}
 if len(roots)==0{return r,errors.New("no roots supplied")}
 seen:=make(map[string]bool)
 for _,input:=range roots{
  root,e:=filepath.Abs(input)
  if e!=nil||strings.HasPrefix(root,`\\`){r.UnavailableRoots=append(r.UnavailableRoots,input);continue}
  if seen[root]{continue};seen[root]=true
  st,e:=os.Lstat(root)
  if e!=nil||(!st.Mode().IsRegular()&&!st.IsDir()){r.UnavailableRoots=append(r.UnavailableRoots,input);continue}
  r.Roots=append(r.Roots,root)
 }
 r.Root=strings.Join(r.Roots,"; ")
 if len(r.Roots)==0{return r,errors.New("no accessible regular roots")}
 vc,cancel:=context.WithTimeout(ctx,15*time.Second);r.Version,err=VersionAt(vc,exe,db);cancel()
 if err!=nil{return r,fmt.Errorf("ClamAV version: %w",err)}
 r.DatabaseStatus=DatabaseStatus(r.Version,time.Now())
 args:=[]string{"--stdout","--infected","--recursive","--alert-exceeds-max=yes","--follow-dir-symlinks=0","--follow-file-symlinks=0"}
 dbArgs,e:=databaseArgs(db);if e!=nil{return r,e};args=append(args,dbArgs...)
 if officialBundleOnly(exe,db){args=append(args,"--official-db-only=yes")}
 args=append(args,"--");args=append(args,r.Roots...)
 cmd:=command(ctx,exe,args...)
 var out bounded
 if fn,ok:=ctx.Value(observerKey314{}).(func(Finding));ok{out.observer.fn=fn}
 // Both streams share one synchronized os/exec writer.
 cmd.Stdout=&out;cmd.Stderr=&out
 e=cmd.Run();code:=0
 if e!=nil{var ee *exec.ExitError;if errors.As(e,&ee){code=ee.ExitCode()}else{code=-1}}
 parse(&r,string(out.b),code,out.overflow)
 if ctx.Err()!=nil{return r,ctx.Err()}
 if len(r.UnavailableRoots)>0{return r,errors.New("one or more requested roots inaccessible; batch coverage partial")}
 if e!=nil&&code!=1{return r,fmt.Errorf("ClamAV scan: %w",e)}
 if r.DatabaseStatus!="fresh"{return r,fmt.Errorf("ClamAV: database freshness %s; scan results retained, coverage reduced",r.DatabaseStatus)}
 if r.Partial{return r,errors.New("ClamAV report partial or inconsistent; see engine output")}
 return r,nil
}
var _ io.Writer=(*bounded)(nil)
