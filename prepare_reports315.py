"""One Reports destination, no extra windows or repeated export buttons."""
from pathlib import Path
import re
from report_anchor315 import insert_report_css
def once(s,a,b):
    if s.count(a)!=1:raise RuntimeError("reports anchor "+a[:100])
    return s.replace(a,b,1)
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'</nav>','<button data-page="reports315">Relatórios</button></nav>')
# The generated desktop uses an explicit navigation registry. Without this entry,
# selecting Reports hides the other sections but leaves the Reports panel hidden.
nav_pattern=r"(\[[^\[\]\r\n]*'overview'[^\[\]\r\n]*'intel'[^\[\]\r\n]*\])(?=\.forEach\(id=>el\(id\)\.classList\.toggle\('hidden',id!==b\.dataset\.page\))"
def reports_nav(m):
    items=m.group(1)
    if "'reports315'" in items:raise RuntimeError("reports navigation already registered")
    # Diagnostics is now nested inside Reports, not a top-level destination.
    items=items.replace(",'diagnostics'","")
    return items[:-1]+",'reports315']"
s,n=re.subn(nav_pattern,reports_nav,s)
if n!=1:raise RuntimeError("reports navigation registry count "+str(n))
if len(re.findall(r"'reports315'\]\.forEach\(id=>el\(id\)\.classList\.toggle\('hidden',id!==b\.dataset\.page\)",s))!=1:
    raise RuntimeError("reports navigation registration not verified")

s=once(s,'</main>','''<section id="reports315" class="hidden card"><h2>Central de relatórios</h2><p class="sub">Dados reais do ZEUVORYX e do Windows, salvos localmente. Nenhum envio automático.</p><label for="z315ReportSource">Relatório</label><select id="z315ReportSource"><option value="current">Varredura atual ou última</option><option value="history">Varredura do histórico</option><option value="diag">Diagnóstico gravado</option></select><label for="z315ReportSearch">Pesquisar histórico</label><input id="z315ReportSearch" type="search" placeholder="Data, modo ou identificador"><select id="z315ReportHistory" aria-label="Relatório histórico"><option value="">Selecione uma varredura</option></select><p id="z315ReportPreview" role="status"></p><button id="z315ReportSave" class="primary">Gerar relatório JSON</button><div id="z315Transfer" class="entry" hidden><strong id="z315ReportFile" role="button" tabindex="0" style="cursor:grab;touch-action:none;user-select:none;overflow-wrap:anywhere">Arquivo JSON</strong><p class="sub">Mantenha o botão esquerdo pressionado sobre o arquivo e arraste. Ou copie e cole no destino compatível.</p><button id="z315CopyFile">Copiar arquivo</button></div><p id="z315ReportStatus" role="status" style="overflow-wrap:anywhere"></p><details><summary>Histórico e avaliação</summary><div id="z315ReportAudit"></div></details><details><summary>Diagnóstico técnico</summary><div id="z315ReportDiag"></div></details></section></main>''')
# Insert CSS only at an actual structural HTML boundary (not a string in JS).
report_css = '#reports315{gap:10px}#reports315 #z315ReportHistory,#reports315 #z315ReportSearch{max-width:100%}#reports315 [hidden]{display:none!important}#reports315 details{margin-top:10px}#reports315 #auditList{max-height:220px;overflow:auto}'
s, anchor = insert_report_css(s, report_css)
print("reports CSS anchor:", anchor)
s=once(s,'let h=await window.auditHistory();out.replaceChildren();','let h=await window.auditHistory();z315ReportHistory(h.Events||[]);out.replaceChildren();')
old="const b=node('button','Exportar relatório salvo');b.onclick=()=>action(async()=>{el('msg').textContent='Relatório salvo em: '+await window.auditExportHistory(v.ScanID)});row.append(b)"
s=once(s,old,"")
s=once(s,'</script></body></html>',r'''
let z315Rows=[],z315ReportBusy=false;
function z315ReportHistory(rows){z315Rows=rows.filter(v=>v.Kind==='scan_finished'&&v.ScanID);z315FilterReports()}
function z315FilterReports(){const sel=el('z315ReportHistory'),prev=sel.value,q=el('z315ReportSearch').value.trim().toLowerCase();sel.replaceChildren(new Option('Selecione uma varredura',''));for(const v of z315Rows){const label=new Date(v.At).toLocaleString()+' · '+(v.Mode||'')+' · '+(v.Cancelled?'Cancelada':v.Partial?'Parcial':'Encerrada')+' · '+v.ScanID;if(!q||label.toLowerCase().includes(q))sel.add(new Option(label,v.ScanID))}if(Array.from(sel.options).some(x=>x.value===prev))sel.value=prev}
async function z315ReportPaint(){const kind=el('z315ReportSource').value,preview=el('z315ReportPreview'),btn=el('z315ReportSave');try{if(kind==='history'){const v=z315Rows.find(x=>x.ScanID===el('z315ReportHistory').value);preview.textContent=v?'Arquivos: '+fmt(v.Files)+' · Alertas: '+fmt(v.Alerts)+' · Falhas: '+fmt(v.ErrorCount):'Selecione uma varredura salva.';btn.disabled=z315ReportBusy||!v;return}if(kind==='diag'){const d=await window.diagReport();if(el('z315ReportSource').value!==kind)return;const n=(d.samples||[]).length;preview.textContent='Amostras de diagnóstico: '+fmt(n);btn.disabled=z315ReportBusy||!n;return}const v=await window.getScanState(),r=v.report||{};if(el('z315ReportSource').value!==kind)return;preview.textContent=(v.running?'Relatório parcial':v.done?(r.cancelled?'Varredura cancelada':'Última varredura'):'Nenhuma varredura iniciada')+' · Arquivos: '+fmt(r.files??v.files)+' · Volume: '+volume(r.bytes??v.bytes??0)+' · Alertas do scanner: '+fmt(v.finding_total??0);btn.disabled=z315ReportBusy||!(v.running||v.done)}catch(e){preview.textContent=String(e);btn.disabled=true}}
;(()=>{const source=el('z315ReportSource'),status=el('z315ReportStatus'),save=el('z315ReportSave'),section=el('reports315');const legacy=el('export');legacy.hidden=true;section.append(legacy);const transfer=el('z315Transfer'),file=el('z315ReportFile'),copy=el('z315CopyFile');let transferring=false;async function handoff(drag){if(transferring||transfer.hidden)return;transferring=true;copy.disabled=true;try{const ok=drag?await window.dragReport315():await window.copyReport315();status.textContent=ok?(drag?'Arquivo entregue ao destino pelo Windows. Confirme o envio no aplicativo de destino.':'Arquivo copiado para a área de transferência. Cole no destino compatível.'):'Transferência cancelada ou não aceita; tente Copiar arquivo.'}catch(e){status.textContent='Transferência indisponível: '+String(e)}finally{transferring=false;copy.disabled=false}}file.onpointerdown=e=>{if(e.button===0){e.preventDefault();handoff(true)}};file.onkeydown=e=>{if(e.key==='Enter'||e.key===' '){e.preventDefault();handoff(false)}};copy.onclick=()=>handoff(false);const audit=el('auditTitle').closest('.card');el('z315ReportAudit').append(audit);const failures=el('zLiveFailures');if(failures)el('z315ReportAudit').prepend(failures);const diag=el('diagnostics');if(diag){diag.classList.remove('hidden');diag.removeAttribute('id');el('z315ReportDiag').append(diag);const old=el('diagSave');old.hidden=true;const nav=document.querySelector('[data-page="diagnostics"]');if(nav){nav.hidden=true;nav.style.display='none'}/* Preserve legacy polling ID on the nested section. */diag.id='diagnostics';}
function showHistory(){const visible=source.value==='history';for(const x of [el('z315ReportHistory'),el('z315ReportSearch'),section.querySelector('label[for="z315ReportSearch"]')])if(x)x.hidden=!visible}showHistory();source.onchange=()=>{status.textContent='';showHistory();z315ReportPaint()};el('z315ReportSearch').oninput=()=>{z315FilterReports();z315ReportPaint()};el('z315ReportHistory').onchange=z315ReportPaint;
document.querySelector('[data-page="reports315"]').addEventListener('click',()=>{auditRefresh();z315ReportPaint()});
save.onclick=async()=>{if(z315ReportBusy||transferring)return;z315ReportBusy=true;transfer.hidden=true;save.disabled=true;status.textContent='Gravando o arquivo JSON no computador...';try{const kind=source.value;let path;if(kind==='history'){const id=el('z315ReportHistory').value;if(!id)throw Error('Selecione um relatório');path=await window.auditExportHistory(id)}else if(kind==='diag')path=await window.diagExport();else path=await window.exportReport();if(typeof path!=='string'||!path.trim())throw Error('O aplicativo não confirmou o arquivo');file.textContent='📄 '+path.split(/[\\\\/]/).pop();transfer.hidden=false;status.textContent='Arquivo salvo: '+path+' — arraste ou copie somente se desejar.'}catch(e){status.textContent='Falha ao salvar: '+String(e)}finally{z315ReportBusy=false;await z315ReportPaint()}};
setInterval(()=>{if(!section.classList.contains('hidden')){const d=el('diagnostics');if(d)d.classList.remove('hidden');z315ReportPaint()}},2000);
})();
</script></body></html>''')
p.write_text(s,encoding="utf-8")
# Replace both report export handlers; all three report types share a verified native writer.
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
if '"io"' not in s:s=once(s,'import (','import (\n "io"')
a=s.index('if !bind("auditExportHistory"');b=s.index('if !bind("exportReport",',a)
s=s[:a]+'''if !bind("auditExportHistory",func(id string)(string,error){
 a,e:=audit.Default();if e!=nil{return "",e};data,e:=a.ReadReport(id);if e!=nil{return "",e}
 if !json.Valid(data){return "",fmt.Errorf("Relatório histórico inválido")}
 return saveReport315("zeuvoryx-history",func(w io.Writer)error{n,e:=w.Write(data);if e==nil&&n!=len(data){return io.ErrShortWrite};return e})
 }){return}
 '''+s[b:]
s,n=re.subn(r'if !bind\("exportReport",\s*func\(\)\s*\(string,\s*error\)\s*\{[\s\S]*?\n\s*\}\)\s*\{\s*return\s*\}',lambda m:'''if !bind("exportReport",func()(string,error){
 r,ok:=s.snapshotForExport();if !ok{return "",fmt.Errorf("Nenhuma varredura iniciada")}
 return jsonReport315("zeuvoryx-report",r)
 }){return}
 if !bind("copyReport315",func()(bool,error){return reportShell315(false)}){return}
 if !bind("dragReport315",func()(bool,error){return reportShell315(true)}){return}''',s)
if n!=1:raise RuntimeError("current export anchor "+str(n))
s=once(s,'return exportDiag(diag.report())','r:=diag.report();if len(r.Samples)==0{return "",fmt.Errorf("Nenhuma amostra de diagnóstico")};return jsonReport315("zeuvoryx-diagnostico",r)')
p.write_text(s,encoding="utf-8")

# Multi-engine learning report (explicit opt-in, local aggregate evidence only).
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'if !bind("copyReport315",','''if !bind("fieldStart",fieldStart){return}
 if !bind("fieldActive",fieldActive){return}
 if !bind("fieldCancel",fieldCancel){return}
 if !bind("fieldExport",func()(string,error){return fieldExport(s)}){return}
 if !bind("copyReport315",''')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'<option value="diag">Diagnóstico gravado</option>','<option value="diag">Diagnóstico gravado</option><option value="learning">Aprendizado dos motores</option>')
s=once(s,'<p id="z315ReportPreview" role="status"></p>','<p id="z315ReportPreview" role="status"></p><div id="fieldPanel" hidden><p>Inicie a observação, execute as varreduras e gere o JSON. Somente eventos e contagens; nenhum envio automático.</p><button id="fieldStart">Iniciar observação</button><button id="fieldCancel">Cancelar observação</button><p id="fieldStatus" role="status"></p></div>')
s=once(s,"try{if(kind==='history')","try{el('fieldPanel').hidden=kind!=='learning';if(kind==='learning'){const at=await window.fieldActive();preview.textContent=at?'Observando desde '+at:'Inicie a observação';btn.disabled=z315ReportBusy||!at;return}if(kind==='history')")
s=once(s,"else if(kind==='diag')path=await window.diagExport();","else if(kind==='diag')path=await window.diagExport();else if(kind==='learning')path=await window.fieldExport();")
s=once(s,'</script></body></html>',"""el('fieldStart').onclick=async()=>{try{el('fieldStatus').textContent='Iniciada: '+await window.fieldStart();await z315ReportPaint()}catch(e){el('fieldStatus').textContent=String(e)}};
el('fieldCancel').onclick=async()=>{try{await window.fieldCancel();el('fieldStatus').textContent='Observação cancelada';await z315ReportPaint()}catch(e){el('fieldStatus').textContent=String(e)}};
</script></body></html>""")
p.write_text(s,encoding="utf-8")

# Recover local scan reports independently of WebView and manual export.
Path("autoreport315_windows.go").write_text(Path("lab/autoreport315_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
p=Path("main.go");s=p.read_text(encoding="utf-8")
a=s.index("func (s *State) startScan(");b=s.index("func main(",a)
x=s[a:b];x=once(x,"var rep engine.Report","var rep engine.Report\n stopAuto315:=s.startAutoReport315(scanID)\n defer func(){stopAuto315(rep)}()")
s=s[:a]+x+s[b:];p.write_text(s,encoding="utf-8")
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'if !bind("copyReport315",','''if !bind("openReports315",openReports315){return}
 if !bind("autoReportStatus315",autoReportStatus315){return}
 if !bind("copyReport315",''')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'<button id="z315ReportSave" class="primary">Gerar relatório JSON</button>','<button id="z315ReportSave" class="primary">Gerar relatório JSON</button><button id="z315OpenFolder">Abrir pasta dos relatórios</button><p id="z315AutoStatus" role="status">Salvamento automático local a cada 30 segundos e ao encerrar.</p>')
s=once(s,'</script></body></html>',r'''
el('z315OpenFolder').onclick=async()=>{try{await window.openReports315()}catch(e){el('z315AutoStatus').textContent='Falha ao abrir pasta: '+String(e)}};
setInterval(async()=>{try{if(!el('reports315').classList.contains('hidden')){const v=await window.autoReportStatus315();if(v)el('z315AutoStatus').textContent=v}}catch(e){el('z315AutoStatus').textContent=String(e)}},2000);
</script></body></html>''')
p.write_text(s,encoding="utf-8")

# Display copyable JSON directly, even when disk export is unavailable.
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'if !bind("openReports315",','''if !bind("reportText315",func(kind,id string)(string,error){
 switch kind {
 case "current":
  r,ok:=s.snapshotForExport();if !ok{return "",fmt.Errorf("Nenhuma varredura iniciada")}
  return formatReportText315(r)
 case "history":
  a,e:=audit.Default();if e!=nil{return "",e};data,e:=a.ReadReport(id);if e!=nil{return "",e}
  return formatRawReportText315(data)
 case "diag":
  r:=diag.report();if len(r.Samples)==0{return "",fmt.Errorf("Nenhuma amostra de diagnóstico")}
  return formatReportText315(r)
 default:return "",fmt.Errorf("Tipo de relatório inválido")
 }
 }){return}
 if !bind("savedReportText315",readLastReportText315){return}
 if !bind("openReports315",''')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'<button id="z315ReportSave" class="primary">Gerar relatório JSON</button>','''<button id="z315ShowText" class="primary">Exibir relatório JSON para copiar</button><button id="z315ReportSave">Salvar arquivo JSON</button><div id="z315TextPanel" hidden><label for="z315Text">Relatório JSON — revise antes de compartilhar</label><textarea id="z315Text" rows="12" readonly spellcheck="false" aria-label="Relatório JSON selecionável" style="width:100%;box-sizing:border-box;resize:vertical;font-family:monospace"></textarea><button id="z315SelectText">Selecionar tudo</button><button id="z315CopyText">Copiar texto JSON</button><p id="z315TextStatus" role="status"></p></div>''')
s=once(s,'</script></body></html>',r'''
/* Z315_REPORT_COPY_BEGIN */
;(()=>{
 const source=el('z315ReportSource'),panel=el('z315TextPanel'),area=el('z315Text'),status=el('z315TextStatus'),show=el('z315ShowText');
 let busy=false;
 function clear(){area.value='';panel.hidden=true;status.textContent=''}
 source.addEventListener('change',clear);el('z315ReportHistory').addEventListener('change',clear);
 show.onclick=async()=>{if(busy||z315ReportBusy)return;busy=true;z315ReportBusy=true;show.disabled=true;el('z315ReportSave').disabled=true;clear();panel.hidden=false;status.textContent='Obtendo dados reais do relatório...';
 const kind=source.value,id=el('z315ReportHistory').value;
 try{let v;if(kind==='learning'){await window.fieldExport();v=await window.savedReportText315()}else{v=await window.reportText315(kind,id)}
 if(source.value!==kind||(kind==='history'&&id!==el('z315ReportHistory').value)){clear();return}
 if(typeof v!=='string'||!v.trim())throw Error('Relatório vazio');JSON.parse(v);area.value=v;
 const snapshot=JSON.parse(v),at=snapshot.snapshot_at||snapshot.finished||snapshot.started;status.textContent='Snapshot '+(at?new Date(at).toLocaleString():'sem horário')+' · '+(snapshot.in_progress?'PARCIAL — clique em Exibir novamente para atualizar':snapshot.cancelled?'CANCELADO':snapshot.partial?'FINAL PARCIAL':'Relatório disponível')+'. Copiar texto JSON copia exatamente este snapshot.';area.focus();area.select();
 }catch(e){status.textContent=String(e);area.value=''}finally{busy=false;z315ReportBusy=false;show.disabled=false;await z315ReportPaint()}};
 el('z315SelectText').onclick=()=>{area.focus();area.select()};
 el('z315CopyText').onclick=async()=>{if(!area.value)return;area.focus();area.select();try{
 if(navigator.clipboard&&navigator.clipboard.writeText){try{await navigator.clipboard.writeText(area.value);status.textContent='Texto JSON copiado.';return}catch(_){}}
 if(document.execCommand('copy'))status.textContent='Texto JSON copiado.';else status.textContent='Texto selecionado. Pressione Ctrl+C para copiar.';
 }catch(_){status.textContent='Texto selecionado. Pressione Ctrl+C para copiar.'}};
})();
/* Z315_REPORT_COPY_END */
</script></body></html>''')
# Keep a direct shortcut on the main screen; no additional page registry.
s=once(s,'<section id="overview">','<section id="overview"><div class="card"><button id="z315ReportShortcut">Gerar relatório para copiar</button></div>')
s=once(s,'</script></body></html>',r'''el('z315ReportShortcut').onclick=()=>{document.querySelector('[data-page="reports315"]').click();el('z315ReportSource').value='current';el('z315ReportSource').dispatchEvent(new Event('change'));el('z315ShowText').click()};</script></body></html>''')
for id in ("z315ShowText","z315Text","z315CopyText","z315ReportShortcut"):
    if s.count('id="'+id+'"')!=1:raise RuntimeError("missing report text control "+id)
p.write_text(s,encoding="utf-8")

# P0: never hold the scan-state lock during final report serialization or disk IO.
# State remains running until the audit result is known; the UI and cancel command
# can still acquire the lock. A storage failure is a separate audit error.
p=Path("main.go");s=p.read_text(encoding="utf-8")
a=s.index("func (s *State) startScan(");b=s.index("func main(",a)
x=s[a:b]
x,n=re.subn(r's\.mu\.Lock\(\)\s*defer s\.mu\.Unlock\(\)\s*(rep\.ScanID\s*=)',lambda m:m[1],x,count=1)
if n!=1:raise RuntimeError("P0 final scan mutex anchor "+str(n))
x=once(x,"s.report = rep","s.mu.Lock()\n defer s.mu.Unlock()\n s.report = rep")
s=s[:a]+x+s[b:];p.write_text(s,encoding="utf-8")

# Build the technical diagnosis from a reviewed snapshot using derived facts only.
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'if !bind("savedReportText315",', '''if !bind("saveTechnicalDiagnosis315",func(data string)(string,error){
 if len(data)>64<<10 || !json.Valid([]byte(data)){return "",fmt.Errorf("Diagnóstico JSON inválido")}
 return jsonReport315("zeuvoryx-diagnostico-tecnico",json.RawMessage(data))
 }){return}
 if !bind("savedReportText315",''')
if '"encoding/json"' not in s:
    s=once(s,"import (",'import (\n "encoding/json"')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'<div id="z315TextPanel" hidden>', '''<div class="card"><h2>Diagnóstico de Testes</h2><button id="z315BuildDiagnosis">Analisar falha / Atualizar diagnóstico</button><div id="z315DiagnosisPanel" hidden><label for="z315Diagnosis">Causas possíveis e soluções técnicas</label><textarea id="z315Diagnosis" readonly rows="12" style="width:100%;box-sizing:border-box"></textarea><button id="z315CopyDiagnosis">Copiar diagnóstico completo para IA</button><button id="z315SaveDiagnosis">Exportar diagnóstico JSON</button></div><p id="z315DiagnosisStatus" role="status"></p></div><div id="z315TextPanel" hidden>''')
s=once(s,'</script></body></html>',Path("lab/test_diagnosis.js").read_text(encoding="utf-8")+r'''
;(()=>{
 const source=el('z315ReportSource'),history=el('z315ReportHistory'),panel=el('z315DiagnosisPanel'),area=el('z315Diagnosis'),status=el('z315DiagnosisStatus'),build=el('z315BuildDiagnosis');
 let data='',revision=0;
 function clear(){revision++;data='';area.value='';panel.hidden=true;status.textContent=''}
 source.addEventListener('change',clear);history.addEventListener('change',clear);
 build.onclick=async()=>{
  clear();const attempt=revision,kind=source.value,id=history.value;build.disabled=true;
  try{
   if(!['current','history','diag'].includes(kind))throw Error('Selecione a varredura atual, o histórico ou o diagnóstico de testes.');
   const raw=await window.reportText315(kind,id);
   if(attempt!==revision)return;
   const diagnosis=buildTestDiagnosis(JSON.parse(raw),kind);
   data=JSON.stringify(diagnosis,null,2);area.value=data;panel.hidden=false;
   status.textContent='Diagnóstico atualizado a partir do relatório selecionado.';
  }catch(e){if(attempt===revision)status.textContent=String(e)}
  finally{build.disabled=false}
 };
 el('z315CopyDiagnosis').onclick=async()=>{
  if(!data)return;
  area.focus();area.select();
  try{await navigator.clipboard.writeText(data);status.textContent='Diagnóstico copiado.'}
  catch(_){status.textContent=document.execCommand('copy')?'Diagnóstico copiado.':'Pressione Ctrl+C para copiar o texto selecionado.'}
 };
 const save=el('z315SaveDiagnosis');
 save.onclick=async()=>{
  if(!data)return;
  save.disabled=true;
  try{status.textContent='Diagnóstico salvo: '+await window.saveTechnicalDiagnosis315(data)}
  catch(e){status.textContent=String(e)}
  finally{save.disabled=false}
 };
})();
</script></body></html>''')
p.write_text(s,encoding="utf-8")
