"""One Reports destination, no extra windows or repeated export buttons."""
from pathlib import Path
import re
def once(s,a,b):
    if s.count(a)!=1:raise RuntimeError("reports anchor "+a[:100])
    return s.replace(a,b,1)
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'</nav>','<button data-page="reports315">Relatórios</button></nav>')
s=once(s,'</main>','''<section id="reports315" class="hidden card"><h2>Central de relatórios</h2><p class="sub">Dados reais do ZEUVORYX e do Windows, salvos localmente. Nenhum envio automático.</p><label for="z315ReportSource">Relatório</label><select id="z315ReportSource"><option value="current">Varredura atual ou última</option><option value="history">Varredura do histórico</option><option value="diag">Diagnóstico gravado</option></select><label for="z315ReportSearch">Pesquisar histórico</label><input id="z315ReportSearch" type="search" placeholder="Data, modo ou identificador"><select id="z315ReportHistory" aria-label="Relatório histórico"><option value="">Selecione uma varredura</option></select><p id="z315ReportPreview" role="status"></p><button id="z315ReportSave" class="primary">Gerar relatório JSON</button><div id="z315Transfer" class="entry" hidden><strong id="z315ReportFile" role="button" tabindex="0" style="cursor:grab;touch-action:none;user-select:none;overflow-wrap:anywhere">Arquivo JSON</strong><p class="sub">Mantenha o botão esquerdo pressionado sobre o arquivo e arraste. Ou copie e cole no destino compatível.</p><button id="z315CopyFile">Copiar arquivo</button></div><p id="z315ReportStatus" role="status" style="overflow-wrap:anywhere"></p><details><summary>Histórico e avaliação</summary><div id="z315ReportAudit"></div></details><details><summary>Diagnóstico técnico</summary><div id="z315ReportDiag"></div></details></section></main>''')
s=once(s,'</style></head>','#reports315{gap:10px}#reports315 #z315ReportHistory,#reports315 #z315ReportSearch{max-width:100%}#reports315 [hidden]{display:none!important}#reports315 details{margin-top:10px}#reports315 #auditList{max-height:220px;overflow:auto}</style></head>')
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
