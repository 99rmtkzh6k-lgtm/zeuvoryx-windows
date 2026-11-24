"""Integrate local audit and report accounting after 3.0.4 generators."""
from pathlib import Path
import re
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
    fields=[("ScanID","string","scan_id"),("ScanMode","string","scan_mode"),("ErrorCount","int","error_count"),("OmittedErrorDetails","int","omitted_error_details"),("ScanError","string","scan_error,omitempty"),("AuditError","string","audit_error,omitempty")]
    s=one(s,'type Report struct {','type Report struct {\n'+''.join('    '+a+' '+b+' '+chr(96)+'json:"'+c+'"'+chr(96)+'\n' for a,b,c in fields))
    return rx(s,r'if len\(report\.Errors\) < 30 \{','report.ErrorCount++\nif len(report.Errors) < 10000 {',2)
def summary(s):
    s=rx(s,r'const ScanEngineVersion = "[^"]+"','const ScanEngineVersion = "3.0.5-lab"')
    return one(s,'r.Partial = r.Cancelled || r.Skipped > 0 || len(r.Errors) > 0','''if r.ErrorCount < len(r.Errors) { r.ErrorCount = len(r.Errors) }
    r.OmittedErrorDetails = r.ErrorCount - len(r.Errors)
    r.Partial = r.Cancelled || r.Skipped > 0 || r.ErrorCount > 0 || r.ScanError != ""''')
def main(s):
    s=one(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/audit"')
    s=rx(s,r'if len\(r\.Errors\) < 30 \{','r.ErrorCount++\nif len(r.Errors) < 10000 {')
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
