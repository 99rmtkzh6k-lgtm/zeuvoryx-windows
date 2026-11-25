"""Live reports: incremental worker snapshots, bounded desktop, full local export."""
from pathlib import Path
import re
def once(s,a,b):
    if s.count(a)!=1: raise ValueError("live anchor: "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
def rx(s,a,b,n=1):
    s,c=re.subn(a,lambda _:b,s)
    if c!=n: raise ValueError("live regex: "+a[:100]+" count="+str(c))
    return s
def edit(p,fn):
    p=Path(p);p.write_text(fn(p.read_text(encoding="utf-8-sig")),encoding="utf-8")
def main(s):
    a=s.index('func (s *State) scanRoots(');b=s.index('func (s *State) quickScan(',a)
    x=s[a:b]
    x=once(x,'var mu sync.Mutex','s.publishLive(r)\n    var mu sync.Mutex')
    x=rx(x,r'mu\.Lock\(\)\s*\n?\s*defer mu\.Unlock\(\)', 'mu.Lock()\n    defer mu.Unlock()\n    defer func(){s.publishLive(r)}()')
    # Publish skip events as they occur, including symlinks and nonregular files.
    x,c=re.subn(r'r\.Skipped\+\+\s*\n?\s*mu\.Unlock\(\)', 'r.Skipped++\n    s.publishLive(r)\n    mu.Unlock()',x)
    if c<1: raise ValueError("skip anchors missing")
    x=rx(x,r'if progress != nil && time\.Since\(lastProgress\) >= 200\*time\.Millisecond \{','if len(hits)>0 || time.Since(lastProgress)>=200*time.Millisecond { s.publishLive(r) }\n    if progress != nil && time.Since(lastProgress) >= 200*time.Millisecond {')
    x=rx(x,r'if progress != nil \{\s*progress\(r\.Files,\s*r\.Bytes\)\s*\}', 's.publishLive(r)\n    if progress != nil {progress(r.Files,r.Bytes)}')
    s=s[:a]+x+s[b:]
    # Store scan identity before releasing the state lock.
    s=rx(s,r'(started\s*:=\s*time\.Now\(\))','started:=time.Now()\n    s.report=engine.Report{ScanID:scanID,ScanMode:mode,Started:started,InProgress:true,Partial:true,EngineVersion:engine.ScanEngineVersion,Root:path}\n    s.report.Summarize();s.report.Partial=true')
    s=rx(s,r's\.scanner\.Scan\(ctx,\s*path,\s*progress\)', 's.scanner.Scan(engine.WithLiveReport(ctx,s.publishLive),path,progress)')
    return s
def scanner(s):
    s=once(s,'type Report struct {','type Report struct {\n InProgress bool '+chr(96)+'json:"in_progress"'+chr(96)+'\n SnapshotAt time.Time '+chr(96)+'json:"snapshot_at"'+chr(96))
    a=s.index('func (s *Scanner) Scan(')
    x=s[a:]
    x=rx(x,r'defer func\(\) \{\s*report\.Summarize\(\)\s*\}\(\)', 'publishLiveReport(ctx,report)\n    defer func(){report.Summarize();publishLiveReport(ctx,report)}()')
    # In the generated Scan method each unlock occurs after a synchronized update.
    x,c=re.subn(r'(?<!defer )mu\.Unlock\(\)', 'publishLiveReport(ctx,report)\n    mu.Unlock()',x)
    if c<3: raise ValueError("scanner worker publication missing")
    return s[:a]+x
def desktop(s):
    s=once(s,'"report": s.report, "done":','"report": report, "finding_total": total, "done":')
    s=rx(s,r'(if !bind\("getScanState", func\(\) map\[string\]any \{\s*s\.mu\.Lock\(\)\s*defer s\.mu\.Unlock\(\))', 'if !bind("getScanState", func() map[string]any {\n    s.mu.Lock(); defer s.mu.Unlock()\n    report,total:=s.reportSnapshotLocked(true)')
    a=s.index('if !bind("exportReport",');b=s.index('home, e := os.UserHomeDir()',a)
    x=s[a:b];x=rx(x,r's\.mu\.Lock\(\)[\s\S]*s\.mu\.Unlock\(\)', 'r,ok:=s.snapshotForExport()\n    if !ok{return "",fmt.Errorf("Nenhuma varredura iniciada")}')
    s=s[:a]+x+s[b:]
    s=once(s,'if !bind("exportReport",','''if !bind("auditExportHistory",func(id string)(string,error){
    a,e:=audit.Default();if e!=nil{return "",e};data,e:=a.ReadReport(id);if e!=nil{return "",e}
    home,e:=os.UserHomeDir();if e!=nil{return "",e}
    f,e:=os.CreateTemp(home,"zeuvoryx-history-*.json");if e!=nil{return "",e};name:=f.Name()
    _,e=f.Write(data);ce:=f.Close();if e==nil{e=ce};if e!=nil{os.Remove(name);return "",e};return name,nil
    }){return}
    if !bind("exportReport",''')
    return s
def html(s):
    s=once(s,"el('hits').textContent=s.done?fmt((r.findings||[]).length):'—';","el('hits').textContent=(s.running||s.done)?fmt(s.finding_total??(r.findings||[]).length):'—';")
    s=once(s,"el('export').disabled=running||!s.done;","el('export').disabled=!running&&!s.done;")
    s=once(s,"String((r.findings||[]).length)","String(s.finding_total??(r.findings||[]).length)")
    s=s.replace("Aguardando o relatório final. A proteção residente depende do Windows.","Resultados parciais atualizados durante a análise. A proteção residente depende do Windows.")
    s=s.replace("Waiting for the final report. Resident protection depends on Windows.","Partial results update during scanning. Resident protection depends on Windows.")
    s=s.replace("Esperando el informe final. La protección residente depende de Windows.","Resultados parciales durante el análisis. La protección residente depende de Windows.")
    s=s.replace("Os resultados completos aparecem no fim da varredura. Alertas heurísticos não confirmam infecção.","Alertas parciais durante a varredura; o relatório completo é consolidado no fim. Alertas heurísticos não confirmam infecção.")
    s=s.replace("Complete results appear after the scan. Heuristic alerts do not confirm infection.","Partial alerts during scanning; final report after completion. Heuristic alerts do not confirm infection.")
    s=s.replace("Los resultados completos aparecen al finalizar el análisis. Las alertas heurísticas no confirman una infección.","Alertas parciales durante el análisis; informe final al terminar. Las alertas heurísticas no confirman una infección.")
    s=rx(s,r"el\('skipped'\)\.textContent=[^;]+;", "el('skipped').textContent=(s.running||s.done)?fmt(r.skipped)+' itens ignorados ou não examinados.':'Aguardando início da análise.';")
    s=once(s,"if(r.scan_id&&!s.running)","if(r.scan_id)")
    s=once(s,"if(!state.done||state.running){details.textContent='';return}","if(!state.done&&!state.running){details.textContent='';return}")
    s=once(s,"fmt(all.length)+' no relatório completo'","fmt(zState.finding_total??all.length)+' no relatório'+(zState.running?' parcial (exibição limitada aos últimos 500)':' completo')")
    s=once(s,"let running=s.running;zSummary(s);","let running=s.running;zSummary(s);zLivePaint(s);")
    s=once(s,'<p class="sub" id="skipped"></p></div>','<p class="sub" id="skipped"></p><p id="zLiveCoverage" class="sub"></p></div><div id="zLiveFailures" class="card"><h2 id="zLiveFailureTitle">Falhas da análise</h2><p id="zLiveFailureCount" role="status"></p><div id="zLiveErrorList"></div></div>')
    s=once(s,"</style></head>", "#zLiveFailures,#zLiveCoverage,#auditList,#scope,#skipped{min-width:0;max-width:100%;overflow-wrap:anywhere;word-break:normal}#zLiveErrorList{max-height:180px;overflow:auto}#zLiveErrorList p{margin:4px 0;font-size:12px}</style></head>")
    # Function declaration hoisted; initial refresh can safely call it.
    js=r"""
function zLivePaint(s){
 const r=s.report||{},z=id=>document.getElementById(id),lang=document.documentElement.lang;
 const t=({
 'pt-BR':['Cobertura parcial','Escopo concluído','Total elegível desconhecido','Falhas da análise','Erros','Detalhes disponíveis na tela','Relatório parcial','Relatório concluído','Aguardando varredura','Itens ignorados','Exportar relatório parcial','Exportar relatório'],
 'en-US':['Partial coverage','Scope finished','Eligible total unknown','Scan failures','Errors','Details shown','Partial report','Finished report','Waiting for scan','Skipped items','Export partial report','Export report'],
 'es-ES':['Cobertura parcial','Ámbito finalizado','Total elegible desconocido','Fallos del análisis','Errores','Detalles visibles','Informe parcial','Informe finalizado','Esperando análisis','Elementos omitidos','Exportar informe parcial','Exportar informe']
 })[lang]||['Cobertura parcial','Escopo concluído','Total elegível desconhecido','Falhas da análise','Erros','Detalhes disponíveis na tela','Relatório parcial','Relatório concluído','Aguardando varredura','Itens ignorados','Exportar relatório parcial','Exportar relatório'];
 const coverage=z('zLiveCoverage'),out=z('zLiveErrorList');
 if(!coverage||!out)return;
 const active=s.running||s.done,partial=s.running||r.partial||r.cancelled||s.error;
 coverage.textContent=active?(partial?t[0]:t[1])+' · '+fmt(s.files)+' · '+volume(s.bytes)+' · '+t[9]+': '+fmt(r.skipped)+' · '+t[2]:t[8];
 z('zLiveFailureTitle').textContent=t[3];
 const errors=r.errors||[],n=Math.max(Number(r.error_count)||0,errors.length);
 z('zLiveFailureCount').textContent=active?t[4]+': '+fmt(n)+' · '+t[5]+': '+fmt(errors.length):t[8];
 const key=JSON.stringify([errors,r.error_categories,r.scan_error,r.audit_error,s.error]);
 if(out.dataset.key!==key){out.dataset.key=key;out.replaceChildren();
 for(const [k,v] of Object.entries(r.error_categories||{}))out.append(node('p',k+': '+fmt(v)));
 for(const e of errors)out.append(node('p',String(e)));
 for(const e of [r.scan_error,r.audit_error,s.error].filter(Boolean))out.append(node('p',String(e)));
 }
 z('export').textContent=s.running?t[10]:t[11];
 const mark=(r.scan_id||String(r.started||''))+':'+s.running+':'+s.done;
 if(active&&z('auditList')&&z('auditList').dataset.scanMark!==mark){z('auditList').dataset.scanMark=mark;Promise.resolve().then(()=>auditRefresh()).catch(()=>{});}
}
;(()=>{
 const coverage=document.getElementById('scope').closest('.card'),failure=document.getElementById('zLiveFailures');
 const chart=Array.from(document.getElementById('overview').children).find(n=>n.matches('.chart'));
 if(chart&&coverage){chart.after(coverage);coverage.after(failure);}
 const history=document.getElementById('auditTitle').closest('.card');if(history&&failure)failure.after(history);
})();
"""
    s=once(s,'</script></body></html>',js+'</script></body></html>')
    # Keep history metadata and export actual saved authenticated records.
    s=once(s,"for(let v of h.Events||[])out.append(node('p',new Date(v.At).toLocaleString()+' · '+v.Kind+' · '+(v.Mode||'')+' · '+(v.Stars?v.Stars+'/5':fmt(v.Files))))","for(let v of h.Events||[]){const row=node('div','','entry');row.append(node('p',new Date(v.At).toLocaleString()+' · '+v.Kind+' · '+(v.Mode||'')+' · '+(v.Stars?v.Stars+'/5':fmt(v.Files))));if(v.Kind==='scan_finished'){row.append(node('p','Alertas: '+fmt(v.Alerts)+' · Falhas: '+fmt(v.ErrorCount)+' · Ignorados: '+fmt(v.Skipped)+(v.Cancelled?' · Cancelado':v.Partial?' · Parcial':'')));const b=node('button','Exportar relatório salvo');b.onclick=()=>action(async()=>{el('msg').textContent='Relatório salvo em: '+await window.auditExportHistory(v.ScanID)});row.append(b)}out.append(row)}")
    return s
for path,fn in [('main.go',main),('engine/scanner.go',scanner),('desktop_windows.go',desktop),('desktop.html',html)]:edit(path,fn)
