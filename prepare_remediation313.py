from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1:raise RuntimeError("313 anchor "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
for name,src in [('engine/clam313.go','lab/clam313.go.txt'),('engine/clam313_test.go','lab/clam313_test.go.txt'),('engine/remediation316.go','lab/remediation316.go.txt'),('behavior313_windows.go','lab/behavior313_windows.go.txt')]:
    Path(name).write_text(Path(src).read_text(encoding="utf-8"),encoding="utf-8")
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/remediation"')
a=s.index('if !bind("chooseClamFolder311"')
s=s[:a]+'''if !bind("behavior313",behavior313){return}
if !bind("quarantineEICAR316",func(path string)(remediation.Outcome,error){
    operationGate.Lock();defer operationGate.Unlock()
    s.mu.Lock();defer s.mu.Unlock()
    if s.running||native.status().Running{return remediation.Outcome{},fmt.Errorf("Aguarde a varredura")}
    for _,hit:=range s.report.Findings{
        if hit.Path!=path||hit.Evidence!="test_eicar"||!strings.EqualFold(hit.SHA256,remediation.EICARSHA256){continue}
        o:=remediation.Outcome{Path:hit.Path,SHA256:hit.SHA256,Detector:"ZEUVORYX",Signature:"EICAR",ScanStarted:s.report.Started,TestArtifact:true}
        item,e:=s.quarantine.AddEICAR316(hit)
        return s.quarantine.RecordOutcome316(o,item,e)
    }
    return remediation.Outcome{},fmt.Errorf("Controle EICAR integral não encontrado no relatório atual")
}){return}
if !bind("remediationHistory316",s.quarantine.Outcomes316){return}
if !bind("remediationExport316",func()(string,error){h,e:=s.quarantine.Outcomes316();if e!=nil{return "",e};return jsonReport315("zeuvoryx-remediacao",h)}){return}
if !bind("refreshClam313",func()(string,error){
    operationGate.Lock();defer operationGate.Unlock()
    s.mu.Lock();running:=s.running;s.mu.Unlock()
    if running||native.status().Running{return "",fmt.Errorf("Aguarde a operação atual")}
    return refreshClam313()
}){return}
if !bind("isolateClam313",func(path,signature string)(string,error){
    operationGate.Lock();defer operationGate.Unlock()
    s.mu.Lock()
    if s.running {s.mu.Unlock();return "",fmt.Errorf("Aguarde a varredura")}
    found:=false
    scanStarted:=s.report.Started
    for _,r:=range s.report.SignatureReports{for _,h:=range r.Findings{if h.Path==path&&h.Signature==signature&&h.Kind=="signature_detection"{found=true}}}
    s.mu.Unlock()
    if native.status().Running||!found{return "",fmt.Errorf("Detecção não disponível para revisão")}
    tool,e:=clamExecutable();if e!=nil{return "",e}
    ctx,cancel:=context.WithTimeout(context.Background(),5*time.Minute);defer cancel()
    item,e:=s.quarantine.AddClamRechecked(ctx,tool,path,signature)
    o,err:=s.quarantine.RecordOutcome316(remediation.Outcome{Path:path,Detector:"ClamAV",Signature:signature,ScanStarted:scanStarted},item,e)
    if err!=nil{return "",err}
    return o.QuarantineID,nil
}){return}
'''+s[a:]
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'<div id="qitems" class="list"></div>','<div id="qitems" class="list"></div><div class="entry"><h2>Destino das ameaças e controles</h2><p>Detecção não comprova neutralização. Ações verificadas por hash e existência do arquivo; sem atribuir a ação ao Defender ou a outro motor sem evidência.</p><button id="z316History">Consultar histórico de remediação</button><button id="z316Export">Salvar histórico JSON</button><p id="z316State" role="status"></p><div id="z316Events"></div></div>')
s=once(s,'<div id="z312SigDetails"></div>','<button id="z313Update">Atualizar assinaturas oficiais</button><p id="z313UpdateState" role="status"></p><div id="z312SigDetails"></div><div class="card"><h2>Telemetria comportamental (Sysmon opcional)</h2><button id="z313Behavior">Consultar eventos registrados</button><p id="z313BehaviorState" role="status"></p><div id="z313BehaviorSignals"></div></div>')
old="d.append(node('p',f.kind+' · '+f.signature+' · '+f.path));"
new="""const row=node('div');row.append(node('p',f.kind+' · '+f.signature+' · '+f.path));if(!v.running&&f.kind==='signature_detection'&&!/eicar|heuristic|encrypted|pua\\.|potentially|test-signature/i.test(f.signature)){const b=node('button','Revalidar e colocar em quarentena');b.onclick=()=>action(async()=>{if(!await ask('Reexaminar e isolar reversivelmente este arquivo? '+f.path))return;b.disabled=true;try{const id=await window.isolateClam313(f.path,f.signature);row.append(node('p','Quarentena registrada: '+id));}catch(e){row.append(node('p','Não isolado: '+String(e)));b.disabled=false;}});row.append(b);}d.append(row);"""
s=once(s,old,new)
s=once(s,'</script></body></html>',r""" 
;(()=>{
const findings=el('findings');
let working=false;
function decorate(){for(const row of findings.children){if(row.querySelector('[data-eicar316]'))continue;const text=row.textContent;if(!/EICAR/i.test(text))continue;const path=row.querySelector('p')?.textContent;if(!path)continue;const b=node('button','Revalidar e isolar controle EICAR');b.dataset.eicar316='1';b.onclick=()=>action(async()=>{if(working)return;working=true;try{if(!await ask('Isolar reversivelmente apenas o controle EICAR integral deste relatório? '+path))return;const result=await window.quarantineEICAR316(path);el('msg').textContent=result.status==='quarantine_verified'?'Quarentena verificada: '+result.quarantine_id:'Destino não confirmado';await loadQ()}catch(e){el('msg').textContent='Destino não confirmado: '+String(e)}finally{working=false}});row.append(b);}}
new MutationObserver(decorate).observe(findings,{childList:true});decorate();
el('z316History').onclick=()=>action(async()=>{const h=await window.remediationHistory316(),d=el('z316Events');d.replaceChildren();const rows=(h.events||[]).slice().sort((a,b)=>String(b.at).localeCompare(String(a.at)));for(const v of rows){const row=node('div','','entry');row.append(node('strong',v.status==='quarantine_verified'?'Quarentena verificada na data da ação':'Destino não confirmado / falha'),node('p',v.original_path),node('p','Detector: '+v.detector+' | Executor: '+(v.executor||'não confirmado')+' | '+v.at),node('p','SHA-256: '+(v.sha256||'não disponível')+' | '+(v.test_artifact?'Controle de teste':'Detecção')+' | '+(v.error||'Registro histórico; o estado atual pode mudar.')));d.append(row)}el('z316State').textContent=rows.length+' registros | inválidos: '+h.invalid+(h.truncated?' | leitura limitada':'')});
el('z316Export').onclick=()=>action(async()=>{el('z316State').textContent='Relatório salvo: '+await window.remediationExport316()});
})();
</script></body></html>""")
s=once(s,'</script></body></html>',"""document.getElementById('z313Update').onclick=()=>action(async()=>{const p=el('z313UpdateState');p.textContent='Atualizando...';try{p.textContent=await window.refreshClam313()}catch(e){p.textContent=String(e)}});
document.getElementById('z313Behavior').onclick=()=>action(async()=>{const p=el('z313BehaviorState'),d=el('z313BehaviorSignals');d.replaceChildren();p.textContent='Consultando eventos locais...';try{const r=await window.behavior313();p.textContent=r.coverage+' Eventos: '+fmt(r.events);for(const h of r.signals||[]){d.append(node('p',h.Image+' · '+h.Reason+' · arquivos '+h.FileCreates+' · registros '+h.StartupWrites));}}catch(e){p.textContent=String(e);}});</script></body></html>""")
p.write_text(s,encoding="utf-8")
for p in list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]:
    if ".git" in p.parts:continue
    s=p.read_text(encoding="utf-8");n=s.replace("3.0.12","3.0.13")
    if s!=n:p.write_text(n,encoding="utf-8")
