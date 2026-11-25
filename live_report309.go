package main
import (
 "time"
 "github.com/zeuvoryx/zeuvoryx/engine"
)
const liveFindingLimit=500
const liveErrorLimit=20
// Source owns its worker mutex. Copy only new, immutable finding and error entries.
func(s *State) publishLive(src engine.Report) {
 s.mu.Lock();defer s.mu.Unlock()
 if !s.running{return}
 old:=s.report
 if src.Files<old.Files||src.Bytes<old.Bytes||len(src.Findings)<len(old.Findings)||len(src.Errors)<len(old.Errors){return}
 src.ScanID=old.ScanID;src.ScanMode=old.ScanMode
 if src.Inventory314.State==""{src.Inventory314=old.Inventory314}
 if src.SignatureState==""{src.SignatureState=old.SignatureState;src.SignatureProblem=old.SignatureProblem;src.SignatureRoots=old.SignatureRoots;src.SignatureReports=old.SignatureReports}
 if src.WindowsLocal314==nil{src.WindowsLocal314=old.WindowsLocal314;src.WindowsProblem314=old.WindowsProblem314}
 if src.Started.IsZero(){src.Started=old.Started}
 src.Finished=time.Time{}
 src.InProgress=true;src.Partial=true;src.SnapshotAt=time.Now()
 src.EngineVersion=engine.ScanEngineVersion;src.BuildCommit=old.BuildCommit
 src.Findings=append(old.Findings,src.Findings[len(old.Findings):]...)
 src.Errors=append(old.Errors,src.Errors[len(old.Errors):]...)
 src.ErrorCategories=cloneErrorCounts(src.ErrorCategories)
 if src.ErrorCount<len(src.Errors){src.ErrorCount=len(src.Errors)}
 src.OmittedErrorDetails=src.ErrorCount-len(src.Errors)
 s.report=src;s.files=src.Files;s.bytes=src.Bytes
}
func cloneErrorCounts(m map[string]int)map[string]int{
 if m==nil{return nil};n:=make(map[string]int,len(m));for k,v:=range m{n[k]=v};return n
}
// Caller holds State.mu. Copy the mutable slices before returning to WebView/JSON.
func(s *State) reportSnapshotLocked(view bool)(engine.Report,int){
 r:=s.report;total:=len(r.Findings)
 if view&&s.running&&len(r.Findings)>liveFindingLimit{r.Findings=r.Findings[len(r.Findings)-liveFindingLimit:]}
 if view&&len(r.Errors)>liveErrorLimit{r.Errors=r.Errors[len(r.Errors)-liveErrorLimit:]}
 r.Findings=append([]engine.Finding(nil),r.Findings...)
 r.Errors=append([]string(nil),r.Errors...)
 r.ErrorCategories=cloneErrorCounts(r.ErrorCategories)
 if s.running{r.InProgress=true;r.Partial=true;r.Finished=time.Time{};r.SnapshotAt=time.Now()}
 return r,total
}
func(s *State) snapshotForExport()(engine.Report,bool){
 s.mu.Lock();r,_:=s.reportSnapshotLocked(false);ok:=!r.Started.IsZero();s.mu.Unlock()
 if ok&&r.InProgress{r.Finished=r.SnapshotAt;r.Summarize();r.Finished=time.Time{};r.Partial=true}
 return r,ok
}
