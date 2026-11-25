from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1:raise RuntimeError("313 anchor "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
for name,src in [('engine/clam313.go','lab/clam313.go.txt'),('engine/clam313_test.go','lab/clam313_test.go.txt'),('behavior313_windows.go','lab/behavior313_windows.go.txt')]:
    Path(name).write_text(Path(src).read_text(encoding="utf-8"),encoding="utf-8")
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
a=s.index('if !bind("chooseClamFolder311"')
s=s[:a]+'''if !bind("behavior313",behavior313){return}
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
    for _,r:=range s.report.SignatureReports{for _,h:=range r.Findings{if h.Path==path&&h.Signature==signature&&h.Kind=="signature_detection"{found=true}}}
    s.mu.Unlock()
    if native.status().Running||!found{return "",fmt.Errorf("Detecção não disponível para revisão")}
    tool,e:=clamExecutable();if e!=nil{return "",e}
    ctx,cancel:=context.WithTimeout(context.Background(),5*time.Minute);defer cancel()
    item,e:=s.quarantine.AddClamRechecked(ctx,tool,path,signature);if e!=nil{return "",e}
    return item.ID,nil
}){return}
'''+s[a:]
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'<div id="z312SigDetails"></div>','<button id="z313Update">Atualizar assinaturas oficiais</button><p id="z313UpdateState" role="status"></p><div id="z312SigDetails"></div><div class="card"><h2>Telemetria comportamental (Sysmon opcional)</h2><button id="z313Behavior">Consultar eventos registrados</button><p id="z313BehaviorState" role="status"></p><div id="z313BehaviorSignals"></div></div>')
old="d.append(node('p',f.kind+' · '+f.signature+' · '+f.path));"
new="""const row=node('div');row.append(node('p',f.kind+' · '+f.signature+' · '+f.path));if(!v.running&&f.kind==='signature_detection'&&!/eicar|heuristic|encrypted|pua\\.|potentially|test-signature/i.test(f.signature)){const b=node('button','Revalidar e colocar em quarentena');b.onclick=()=>action(async()=>{if(!confirm('Reexaminar e isolar reversivelmente este arquivo? '+f.path))return;b.disabled=true;try{const id=await window.isolateClam313(f.path,f.signature);row.append(node('p','Quarentena registrada: '+id));}catch(e){row.append(node('p','Não isolado: '+String(e)));b.disabled=false;}});row.append(b);}d.append(row);"""
s=once(s,old,new)
s=once(s,'</script></body></html>',"""document.getElementById('z313Update').onclick=()=>action(async()=>{const p=el('z313UpdateState');p.textContent='Atualizando...';try{p.textContent=await window.refreshClam313()}catch(e){p.textContent=String(e)}});
document.getElementById('z313Behavior').onclick=()=>action(async()=>{const p=el('z313BehaviorState'),d=el('z313BehaviorSignals');d.replaceChildren();p.textContent='Consultando eventos locais...';try{const r=await window.behavior313();p.textContent=r.coverage+' Eventos: '+fmt(r.events);for(const h of r.signals||[]){d.append(node('p',h.Image+' · '+h.Reason+' · arquivos '+h.FileCreates+' · registros '+h.StartupWrites));}}catch(e){p.textContent=String(e);}});</script></body></html>""")
p.write_text(s,encoding="utf-8")
for p in list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]:
    if ".git" in p.parts:continue
    s=p.read_text(encoding="utf-8");n=s.replace("3.0.12","3.0.13")
    if s!=n:p.write_text(n,encoding="utf-8")
