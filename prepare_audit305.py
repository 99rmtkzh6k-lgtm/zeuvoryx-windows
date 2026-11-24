"""Integrate local audit and report accounting after 3.0.4 generators."""
from pathlib import Path
import re
import os
def one(s,a,b):
    if s.count(a)!=1: raise RuntimeError("Expected one anchor: "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
def rx(s,a,b,expected=1):
    s,n=re.subn(a,lambda _:b,s)
    if n!=expected: raise RuntimeError("Unexpected anchor count "+str(n)+": "+a)
    return s
def edit(name,fn):
    p=Path(name);p.write_text(fn(p.read_text(encoding="utf-8-sig")),encoding="utf-8")
def scanner(s):
    fields=[("ErrorCategories","map[string]int","error_categories,omitempty"),("BuildCommit","string","build_commit,omitempty"),("ScanID","string","scan_id"),("ScanMode","string","scan_mode"),("ErrorCount","int","error_count"),("OmittedErrorDetails","int","omitted_error_details"),("ScanError","string","scan_error,omitempty"),("AuditError","string","audit_error,omitempty")]
    s=one(s,'type Report struct {','type Report struct {\n'+''.join('    '+a+' '+b+' '+chr(96)+'json:"'+c+'"'+chr(96)+'\n' for a,b,c in fields))
    s=rx(s,r'func\s+\(s\s+\*Scanner\)\s+examine\(ctx\s+context\.Context,\s*path\s+string,\s*size\s+int64\)\s*\(\[\]Finding,\s*error\)\s*\{','func (s *Scanner) examineUnchecked(ctx context.Context, path string, size int64) ([]Finding, error) {')
    return rx(s,r'if len\(report\.Errors\) < 30 \{','report.ErrorCount++\nif len(report.Errors) < 10000 {',2)
def summary(s):
    s=rx(s,r'const ScanEngineVersion = "[^"]+"','const ScanEngineVersion = "3.0.5-lab"')
    commit=os.environ.get("GITHUB_SHA","")
    if commit and not re.fullmatch(r"[0-9a-f]{40}",commit):
        raise RuntimeError("Invalid build commit")
    s=one(s,'r.EngineVersion = ScanEngineVersion','r.EngineVersion = ScanEngineVersion\n    r.BuildCommit = "'+commit+'"')
    return one(s,'r.Partial = r.Cancelled || r.Skipped > 0 || len(r.Errors) > 0','''if r.ErrorCount < len(r.Errors) { r.ErrorCount = len(r.Errors) }
    r.OmittedErrorDetails = r.ErrorCount - len(r.Errors)
    r.Partial = r.Cancelled || r.Skipped > 0 || r.ErrorCount > 0 || r.ScanError != ""''')
def main(s):
    s=one(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/audit"')
    s=rx(s,r'if len\(r\.Errors\) < 30 \{','r.ErrorCount++\nengine.CountScanError(&r,err)\nif len(r.Errors) < 10000 {')
    s=one(s,'s.report = engine.Report{}','''s.report = engine.Report{}
    scanID,idErr:=audit.NewID()
    mode:="custom"; if automatic { mode="full"; if path==":quick:" {mode="quick"} }
    started:=time.Now()
    store,storeErr:=audit.Default()
    if idErr!=nil { storeErr=idErr }
    if storeErr==nil { _,storeErr=store.Add(audit.Event{Kind:"scan_started",ScanID:scanID,Mode:mode}) }
    ''')
    return rx(s,r'rep\.Summarize\(\)\s*\n\s*s\.report = rep','''rep.ScanID=scanID; rep.ScanMode=mode
    if rep.Started.IsZero(){rep.Started=started};if rep.Finished.IsZero(){rep.Finished=time.Now()}
    if e!=nil{rep.ScanError=e.Error()}
    rep.Summarize()
    if storeErr==nil { storeErr=store.Finish(scanID,mode,rep,audit.Event{Files:rep.Files,Bytes:rep.Bytes,Skipped:rep.Skipped,Alerts:len(rep.Findings),ErrorCount:rep.ErrorCount,Cancelled:rep.Cancelled,Partial:rep.Partial,Error:rep.ScanError}) }
    if storeErr!=nil{rep.AuditError=storeErr.Error()}
    s.report = rep''')
def desktop(s):
    s=one(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/audit"')
    return one(s,'if !bind("getScanState", func() map[string]any {','''if !bind("auditHistory",func()(audit.History,error){a,e:=audit.Default();if e!=nil{return audit.History{},e};return a.History()}){return}
    if !bind("auditRating",func(stars int)(audit.Event,error){a,e:=audit.Default();if e!=nil{return audit.Event{},e};return a.Rating(stars)}){return}
    if !bind("getScanState", func() map[string]any {''').replace('3.0.4 LAB','3.0.5 LAB')
def html(s):
    s=s.replace('3.0.4','3.0.5')
    s=one(s,'<section id="overview">','''<section id="overview"><div class="card" style="margin-bottom:10px"><h2 id="auditTitle"></h2><p id="auditNote" class="sub"></p><div class="actions"><button id="auditLoad"></button><button id="auditShowRating"></button></div><div id="auditRateBox" hidden><label id="auditRateLabel" for="auditStars"></label><select id="auditStars"><option value="1">★</option><option value="2">★★</option><option value="3">★★★</option><option value="4">★★★★</option><option value="5">★★★★★</option></select> <button id="auditSave"></button></div><p id="zLocalAuditState" role="status"></p><div id="auditList"></div></div>''')
    return one(s,'</script></body></html>','''
const auditCopy={
'pt-BR':['Histórico e avaliação local','Dados locais. Sem estatísticas globais.','Consultar histórico','Avaliar aplicativo','Avaliação voluntária (1–5)','Salvar avaliação','Varreduras finalizadas','Arquivos contabilizados','Avaliação local','Registros inválidos','Erro de auditoria','Últimos eventos (até 100)','Avaliação registrada','Sem avaliação','Modo','Erros','Detalhes omitidos'],
'en-US':['Local history and rating','Local data. No global statistics.','Load history','Rate application','Optional rating (1–5)','Save rating','Finished scans','Files counted','Local rating','Invalid records','Audit error','Recent events (up to 100)','Rating saved','Not rated','Mode','Errors','Omitted details'],
'es-ES':['Historial y evaluación local','Datos locales. Sin estadísticas globales.','Consultar historial','Evaluar aplicación','Evaluación voluntaria (1–5)','Guardar evaluación','Análisis finalizados','Archivos contabilizados','Evaluación local','Registros inválidos','Error de auditoría','Últimos eventos (hasta 100)','Evaluación guardada','Sin evaluación','Modo','Errores','Detalles omitidos']};
function auditText(){return auditCopy[document.documentElement.lang]||auditCopy['pt-BR']}
function auditPaint(){let t=auditText();['auditTitle','auditNote','auditLoad','auditShowRating','auditRateLabel','auditSave'].forEach((id,i)=>el(id).textContent=t[i])}
async function auditRefresh(){let t=auditText(),out=el('auditList');try{let h=await window.auditHistory();out.replaceChildren();out.append(node('p',t[6]+': '+h.Finished+' | '+t[7]+': '+fmt(h.Files)+' | '+t[8]+': '+(h.Rating?h.Rating+'/5':t[13])+' | '+t[9]+': '+h.Invalid));out.append(node('p',t[11]));for(let v of h.Events||[])out.append(node('p',new Date(v.At).toLocaleString()+' · '+v.Kind+' · '+(v.Mode||'')+' · '+(v.Stars?v.Stars+'/5':fmt(v.Files))))}catch(e){el('zLocalAuditState').textContent=t[10]+': '+String(e)}}
el('auditLoad').onclick=auditRefresh;
el('auditShowRating').onclick=()=>{el('auditRateBox').hidden=!el('auditRateBox').hidden};
el('auditSave').onclick=async()=>{try{await window.auditRating(Number(el('auditStars').value));el('zLocalAuditState').textContent=auditText()[12];await auditRefresh()}catch(e){el('zLocalAuditState').textContent=auditText()[10]+': '+String(e)}};
new MutationObserver(auditPaint).observe(document.documentElement,{attributes:true,attributeFilter:['lang']});auditPaint();
async function auditReportStatus(){try{let s=await window.getScanState(),r=s.report||{};if(r.scan_id&&!s.running){let t=auditText();el('auditNote').textContent=t[1]+' | '+t[14]+': '+r.scan_mode+' | '+t[15]+': '+r.error_count+' | '+t[16]+': '+r.omitted_error_details+(r.audit_error?' | '+t[10]+': '+r.audit_error:'')}}catch(e){}}
setInterval(auditReportStatus,5000);
</script></body></html>''')
edit("engine/scanner.go",scanner)
edit("engine/scan_upgrade.go",summary)
edit("main.go",main)
edit("desktop_windows.go",desktop)
edit("desktop.html",html)

# Native Windows error classification and bounded retry. No ACL bypass or privilege changes.
retry=Path("engine/scan_retry.go")
retry.write_text(r'''package engine

import (
 "context"
 "errors"
 "io/fs"
 "runtime"
 "syscall"
 "time"
)

// Contain recoverable Go panics in the current file only. Never treat it as clean.
var errAnalyzerPanic = errors.New("falha interna recuperavel durante a analise de arquivo")

func safeScanRead(read func()([]Finding,error))(hits []Finding,err error){
 defer func(){if recover()!=nil{hits=nil;err=errAnalyzerPanic}}()
 return read()
}
func (s *Scanner) examine(ctx context.Context,path string,size int64)([]Finding,error){
 return safeScanRead(func()([]Finding,error){return s.examineUnchecked(ctx,path,size)})
}

func ScanErrorCategory(err error) string {
 if err==nil {return "none"}
 if errors.Is(err,errAnalyzerPanic){return "analyzer_panic"}
 if errors.Is(err,context.Canceled)||errors.Is(err,context.DeadlineExceeded){return "cancelled"}
 if errors.Is(err,fs.ErrPermission){return "access_denied"}
 if errors.Is(err,fs.ErrNotExist){return "not_found"}
 if runtime.GOOS=="windows" {
  var code syscall.Errno
  if errors.As(err,&code) {
   switch code {
   case 5:return "access_denied"
   case 2,3:return "not_found"
   case 32:return "sharing_violation"
   case 33:return "lock_violation"
   }
  }
 }
 return "other"
}

// Caller owns the report lock. Counts final errors; retries are not counted as new files.
func CountScanError(r *Report,err error) {
 if err==nil{return}
 if r.ErrorCategories==nil {r.ErrorCategories=make(map[string]int)}
 r.ErrorCategories[ScanErrorCategory(err)]++
}

func retryScanRead(ctx context.Context,read func()([]Finding,error))([]Finding,error) {
 if e:=ctx.Err();e!=nil{return nil,e}
 hits,err:=read()
 // A live Windows desktop can hold temporary files longer than 25 ms.
 // At most two delayed retries, only for native sharing/lock conflicts.
 for _,delay:=range []time.Duration{50*time.Millisecond,150*time.Millisecond}{
  category:=ScanErrorCategory(err)
  if category!="sharing_violation"&&category!="lock_violation"{return hits,err}
  timer:=time.NewTimer(delay)
  select{case <-ctx.Done():timer.Stop();return nil,ctx.Err();case <-timer.C:}
  if e:=ctx.Err();e!=nil{return nil,e}
  hits,err=read()
 }
 return hits,err
}
''',encoding="utf-8")
edit("engine/scan_upgrade.go",lambda s:rx(s,r"return\s+s\.examine\(ctx,\s*path,\s*size\)","return retryScanRead(ctx,func()([]Finding,error){return s.examine(ctx,path,size)})"))
Path("engine/scan_retry_test.go").write_text(r'''package engine

import (
 "context"
 "errors"
 "io/fs"
 "runtime"
 "syscall"
 "testing"
)

func TestScanRecoverablePanic(t *testing.T){
 n:=0
 hits,err:=retryScanRead(context.Background(),func()([]Finding,error){
  n++
  return safeScanRead(func()([]Finding,error){panic("synthetic test only")})
 })
 if n!=1||hits!=nil||!errors.Is(err,errAnalyzerPanic){t.Fatal(n,hits,err)}
 var r Report
 CountScanError(&r,err)
 if r.ErrorCategories["analyzer_panic"]!=1{t.Fatal(r.ErrorCategories)}
 hits,err=safeScanRead(func()([]Finding,error){return []Finding{{Risk:50}},nil})
 if err!=nil||len(hits)!=1||hits[0].Risk!=50{t.Fatal(hits,err)}
}

func TestScanErrorCategories(t *testing.T) {
 for _,tc:=range []struct{err error; want string}{
 {fs.ErrPermission,"access_denied"},{fs.ErrNotExist,"not_found"},
 {context.Canceled,"cancelled"},{errors.New("unknown"),"other"},
 }{if got:=ScanErrorCategory(tc.err);got!=tc.want{t.Fatal(got,tc.want)}}
 var r Report
 CountScanError(&r,fs.ErrPermission);CountScanError(&r,fs.ErrPermission);CountScanError(&r,fs.ErrNotExist)
 if r.ErrorCategories["access_denied"]!=2||r.ErrorCategories["not_found"]!=1{t.Fatal(r.ErrorCategories)}
}
func TestScanRetryBounded(t *testing.T){
 n:=0
 _,err:=retryScanRead(context.Background(),func()([]Finding,error){n++;return nil,fs.ErrPermission})
 if n!=1||!errors.Is(err,fs.ErrPermission){t.Fatal(n,err)}
 n=0;_,err=retryScanRead(context.Background(),func()([]Finding,error){n++;return nil,fs.ErrNotExist})
 if n!=1||!errors.Is(err,fs.ErrNotExist){t.Fatal(n,err)}
 ctx,cancel:=context.WithCancel(context.Background());cancel()
 n=0;_,err=retryScanRead(ctx,func()([]Finding,error){n++;return nil,nil})
 if n!=0||!errors.Is(err,context.Canceled){t.Fatal(n,err)}
 if runtime.GOOS!="windows"{return}
 for _,code:=range []syscall.Errno{32,33}{
  n=0;_,err=retryScanRead(context.Background(),func()([]Finding,error){n++;if n==1{return nil,code};return nil,nil})
  if n!=2||err!=nil{t.Fatal(n,err)}
  n=0;_,err=retryScanRead(context.Background(),func()([]Finding,error){n++;return nil,code})
  if n!=3||err==nil{t.Fatal(n,err)}
  n=0;_,err=retryScanRead(context.Background(),func()([]Finding,error){n++;if n<3{return nil,code};return nil,nil})
  if n!=3||err!=nil{t.Fatal(n,err)}
  n=0;_,err=retryScanRead(context.Background(),func()([]Finding,error){n++;if n==1{return nil,code};return nil,fs.ErrNotExist})
  if n!=2||!errors.Is(err,fs.ErrNotExist){t.Fatal(n,err)}
  ctx,cancel:=context.WithCancel(context.Background())
  n=0;_,err=retryScanRead(ctx,func()([]Finding,error){n++;cancel();return nil,code})
  cancel()
  if n!=1||!errors.Is(err,context.Canceled){t.Fatal(n,err)}
 }
}
''',encoding="utf-8")


# ZVX-EVO-017 Windows candidate: lifecycle diagnostics and resource pressure.
# Runs on the actual generated Windows source, not just the legacy root scanner.
def stability_main(s):
    s=one(s,'func main() {','func main() {\n    defer startStabilityDiagnostics()()')
    s=rx(s,r'workers\s*:=\s*runtime\.NumCPU\(\)','workers := stabilityWorkers(runtime.NumCPU())')
    s=rx(s,r'runDesktop\(state,\s*\*auto\)','defer watchStability(state)()\n        runDesktop(state, *auto)')
    s=one(s,'scanID,idErr:=audit.NewID()','scanID,idErr:=audit.NewID()') if 'scanID,idErr:=audit.NewID()' in s else s
    s=rx(s,r'started\s*:=\s*time\.Now\(\)\s*\n\s*store,storeErr\s*:=\s*audit\.Default\(\)','started:=time.Now()\n    log.Printf("stability scan_start mode=%s",mode)\n    store,storeErr:=audit.Default()')
    s=one(s,'s.report = rep','log.Printf("stability scan_end files=%d bytes=%d cancelled=%t partial=%t errors=%d",rep.Files,rep.Bytes,rep.Cancelled,rep.Partial,rep.ErrorCount)\n    s.report = rep')
    return s
def stability_desktop(s):
    s=one(s,'import (','import (\n "log"')
    return one(s,'w.Run()','log.Print("stability webview_loop_enter")\n    w.Run()\n    log.Print("stability webview_loop_return")')
edit("main.go",stability_main)
edit("desktop_windows.go",stability_desktop)
