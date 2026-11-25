//go:build windows

package main
import("bytes";"context";"fmt";"os";"os/exec";"path/filepath";"sync";"syscall";"time";"github.com/zeuvoryx/zeuvoryx/fieldlearn")
var fieldSession struct{sync.Mutex;start time.Time;busy bool}
func fieldStart()(string,error){
 fieldSession.Lock();defer fieldSession.Unlock()
 if fieldSession.busy||!fieldSession.start.IsZero(){return "",fmt.Errorf("Observação já iniciada")}
 fieldSession.start=time.Now().UTC();return fieldSession.start.Format(time.RFC3339),nil
}
func fieldActive()string{fieldSession.Lock();defer fieldSession.Unlock();if fieldSession.start.IsZero(){return ""};return fieldSession.start.Format(time.RFC3339)}
func fieldCancel()error{fieldSession.Lock();defer fieldSession.Unlock();if fieldSession.busy{return fmt.Errorf("Exportação em andamento")};fieldSession.start=time.Time{};return nil}
type fieldLimit struct{b bytes.Buffer;overflow bool}
func(w *fieldLimit)Write(p []byte)(int,error){n:=fieldlearn.MaxBytes-w.b.Len();if n<len(p){w.overflow=true;if n>0{_,_=w.b.Write(p[:n])}}else{_,_=w.b.Write(p)};return len(p),nil}
func fieldCollect(start,end time.Time)(fieldlearn.EventLog,error){
 root:=os.Getenv("SystemRoot");if !filepath.IsAbs(root){return fieldlearn.EventLog{},fmt.Errorf("SystemRoot indisponível")}
 if end.Sub(start)>24*time.Hour{return fieldlearn.EventLog{},fmt.Errorf("Sessão superior a 24 horas; inicie uma nova")}
 ctx,cancel:=context.WithTimeout(context.Background(),25*time.Second);defer cancel()
 q:=fmt.Sprintf("*[System[(EventID=1000 or EventID=1001 or EventID=1002 or EventID=1116 or EventID=1117 or EventID=1118 or EventID=1119) and TimeCreated[@SystemTime >= '%s' and @SystemTime <= '%s']]]",start.Format(time.RFC3339Nano),end.Format(time.RFC3339Nano))
 cmd:=exec.CommandContext(ctx,filepath.Join(root,"System32","wevtutil.exe"),"qe","Microsoft-Windows-Windows Defender/Operational","/q:"+q,"/f:xml","/e:Events","/c:501","/rd:false")
 cmd.SysProcAttr=&syscall.SysProcAttr{HideWindow:true};var out fieldLimit;cmd.Stdout=&out
 if e:=cmd.Run();e!=nil||out.overflow{return fieldlearn.EventLog{State:"unavailable"},fmt.Errorf("Registro Defender indisponível, acesso negado ou limite atingido")}
 return fieldlearn.Parse(bytes.NewReader(out.b.Bytes()),start,end)
}
func fieldExport(s *State)(string,error){
 fieldSession.Lock();if fieldSession.start.IsZero()||fieldSession.busy{fieldSession.Unlock();return "",fmt.Errorf("Inicie a observação primeiro")}
 start:=fieldSession.start;fieldSession.busy=true;fieldSession.Unlock()
 defer func(){fieldSession.Lock();fieldSession.busy=false;fieldSession.Unlock()}()
 end:=time.Now().UTC();d,e:=fieldCollect(start,end);problem:="";if e!=nil{problem=e.Error()}
 motors:=[]fieldlearn.Motor{}
 if r,ok:=s.snapshotForExport();ok{
  motors=append(motors,fieldlearn.Motor{Name:"ZEUVORYX",State:"scan_report",Started:r.Started,Finished:r.Finished,Files:fieldlearn.Count(int64(r.Files)),Bytes:fieldlearn.Count(r.Bytes),Alerts:fieldlearn.Count(int64(len(r.Findings))),Partial:r.Partial||r.InProgress||r.Cancelled})
  for _,q:=range r.SignatureReports{motors=append(motors,fieldlearn.Motor{Name:"ClamAV",State:q.Status,Version:q.Version,Started:q.Started,Finished:q.Finished,Files:fieldlearn.Count(q.Scanned),Alerts:fieldlearn.Count(q.Infected),Errors:fieldlearn.Count(q.Errors),Signatures:fieldlearn.Count(int64(q.Evidence.SignatureRecords)),Tests:fieldlearn.Count(int64(q.Evidence.TestRecords)),Limits:fieldlearn.Count(int64(q.Evidence.HeuristicOrLimitRecords)),Partial:q.Partial})}
 }
 report:=fieldlearn.Build(start,end,d,motors,problem)
 path,e:=jsonReport315("zeuvoryx-aprendizado",report);if e!=nil{return "",e}
 fieldSession.Lock();fieldSession.start=time.Time{};fieldSession.Unlock()
 return path,nil
}
