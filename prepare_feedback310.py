"""Final UI integration: rates, estimate and actual report-driven sound."""
from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise ValueError("310 anchor "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8-sig")
s=once(s,'if !bind("getScanState",','''if !bind("z310Maintenance",func(kind string)(bool,error){
    operationGate.Lock();defer operationGate.Unlock()
    s.mu.Lock();running:=s.running;s.mu.Unlock()
    if kind!="startup"&&(running||native.status().Running){return false,fmt.Errorf("Stop or finish the scan before disk maintenance")}
    return false,fmt.Errorf("Operação externa removida; use as funções integradas")
    }){return}
    if !bind("z310Beep",z310Beep){return}
    if !bind("getScanState",''')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'zLivePaint(s);','zLivePaint(s);z310Paint(s);')
s=once(s,'</style></head>','#z310Notice{border:2px solid #b87919;padding:10px;margin:8px 0;overflow-wrap:anywhere}#z310Speed{font-weight:600}#z310Balance{overflow-wrap:anywhere}</style></head>')
s=once(s,'</header>','</header><div id="z310Notice" class="card hidden" role="status"><p id="z310AlertText"></p><button id="z310Open">Ver alertas</button></div>')
s=once(s,'<section id="tools" class="hidden card">','<section id="tools" class="hidden card"><div class="entry" id="z310MaintenancePanel"><h2 id="z310MaintTitle">Central de otimização</h2><p class="sub" id="z310MaintInfo">Ferramentas nativas do Windows. Escolha e confirme ações no Windows.</p><button id="z310Drive">Otimizar SSD/HD</button><button id="z310Clean">Limpeza de disco</button><button id="z310Startup">Programas de inicialização</button><p class="sub" id="z310MaintStatus" role="status"></p></div>')
js=Path("z310_metrics.js").read_text(encoding="utf-8")+r"""
let z310Previous=null,z310Latest=null,z310Used=null,z310SoundAt=-10000;
function z310Words(){return ({
'pt-BR':['Velocidade de análise','arquivos/s','Aguardando medição','Parada','Saldo estimado do espaço ocupado','Não é o total real restante','Estimativa indisponível neste escopo','Novos alertas no relatório','Ver alertas','Revise os achados; isolamento exige confirmação após parar ou concluir a análise.'],
'en-US':['Scan rate','files/s','Measuring','Stopped','Estimated used-space balance','Not the actual remaining scope','Estimate unavailable for this scope','New report alerts','View alerts','Review findings; isolation requires confirmation after stopping or completing the scan.'],
'es-ES':['Velocidad de análisis','archivos/s','Midiendo','Detenido','Saldo estimado del espacio ocupado','No es el total real restante','Estimación no disponible para este ámbito','Nuevas alertas','Ver alertas','Revise los hallazgos; el aislamiento requiere confirmación después de detener o finalizar.']
})[document.documentElement.lang]||['Scan rate','files/s','Measuring','Stopped','Estimated used-space balance','Not the actual remaining scope','Estimate unavailable for this scope','New report alerts','View alerts','Review findings after stopping or completing the scan.']}
function z310Balance(){
 const s=z310Latest||{},r=s.report||{},i=r.inventory_314||{},bar=document.getElementById('z314ProgressBar'),out=document.getElementById('z310Balance'),eta=document.getElementById('z314Eta'),note=document.getElementById('z314ScopeNote');
 bar.hidden=true;eta.textContent='Previsão: indisponível';note.textContent='';
 if(i.state==='counting'){out.textContent='Contando arquivos reais: '+fmt(i.files)+' · '+volume(i.bytes);return}
 if(!['ready','partial'].includes(i.state)){out.textContent='Escopo de análise ainda não medido';return}
 const bytes=Number(s.bytes)||0,files=Number(s.files)||0,total=Number(i.bytes),count=Number(i.files);
 const remainder=Math.max(0,total-bytes),left=Math.max(0,count-files);
 out.textContent=(i.state==='partial'?'Inventário parcial: ':'Escopo inventariado: ')+volume(total)+' · '+fmt(count)+' arquivos | Saldo do inventário: '+volume(remainder)+' · '+fmt(left)+' arquivos (somente scanner próprio)';
 if(i.state==='ready'&&total>0&&bytes<=total){bar.hidden=false;bar.value=Math.min(100,100*bytes/total)}
 note.textContent='Inventário lógico de arquivos, não espaço físico ocupado. '+fmt(i.errors)+' falhas de inventário · '+fmt(i.excluded)+' itens excluídos. Arquivos podem mudar; omissões não são examinadas.';
 if(s.running){eta.textContent='Tempo total indeterminado: '+(i.scanner_done?'scanner próprio encerrado; aguardando ClamAV e consolidação.':i.state==='partial'?'inventário parcial e motor independente em andamento.':'motores independentes em andamento.');return}
 if(!s.running&&s.done){eta.textContent=r.cancelled?'Varredura cancelada':r.partial?'Varredura encerrada com cobertura parcial':'Varredura encerrada';}
}
window.z310Storage=n=>{z310Used=n;z310Balance()};
function z310Paint(s){
 const now=performance.now(),d=z310Delta(z310Previous,s,now),w=z310Words(),z=id=>document.getElementById(id);
 const prior=z310Previous;
 z310Previous=d.state;z310Latest=s;z310Balance();
 const win=z('z315WindowsSummary'),native=s.report?.windows_local_314,history=native?.defenderHistory;
 const mon=z('z315ChangeStatus'),m=s.report?.change_monitor_315;if(mon&&m)mon.textContent=(m.enabled?'Monitor de alterações ativo':'Monitor de alterações desativado')+' · '+fmt((m.active_roots||[]).length)+' pastas ativas de '+fmt((m.roots||[]).length)+' · '+fmt(m.change_signals)+' sinais de alteração (não são vírus) · '+fmt(Object.keys(m.problems||{}).length)+' falhas de cobertura.';
 if(win){
   const active=history?.activeThreats;
   const message=history?.state==='available'?'Microsoft Defender: '+fmt((history.entries||[]).length)+' registros nos últimos 7 dias'+(history.truncated?' (limite de 100)':'')+' · '+(typeof active==='number'?fmt(active)+' ameaças ativas':'ameaças ativas não verificadas')+' · '+(native.observedAt?new Date(native.observedAt).toLocaleTimeString():'horário indisponível'):'Microsoft Defender: histórico não verificado'+(s.report?.windows_problem_314?' · '+s.report.windows_problem_314:'');
   win.textContent=message+' | Resultados nativos, separados dos motores ZEUVORYX.';
 }
 z('z310Speed').textContent=w[0]+': '+(!s.running?w[3]:d.bytesRate===null?w[2]:volume(d.bytesRate)+'/s · '+Math.round(d.filesRate)+' '+w[1]);
 if(!d.state||!prior||prior.id!==d.state.id)z('z310Notice').classList.add('hidden');
 if(d.alerts>0){
   z('z310Notice').classList.remove('hidden');
   z('z310AlertText').textContent=w[7]+': '+d.alerts+' · '+w[9];
   z('z310Open').textContent=w[8];
   if(now-z310SoundAt>=1500){z310SoundAt=now;Promise.resolve(window.z310Beep()).catch(()=>{});}
 }
}
document.getElementById('z310Open').onclick=()=>document.querySelector('[data-page=alerts]').click();
;(()=>{
 const z=id=>document.getElementById(id);
 const copy={
 'pt-BR':['Central de otimização','Operações integradas: o Windows escolhe TRIM ou desfragmentação conforme a unidade. A limpeza é selecionada e confirmada nesta tela.','Otimizar SSD/HD','Limpeza de disco','Programas de inicialização','Ferramenta aberta; execute a ação no Windows.','Encerre a análise antes de manutenção de disco.'],
 'en-US':['Optimization center','Integrated operations: Windows selects TRIM or defragmentation. Review and confirm cleanup here.','Optimize SSD/HDD','Disk cleanup','Startup programs','Tool opened; complete the action in Windows.','Stop scanning before disk maintenance.'],
 'es-ES':['Centro de optimización','Operaciones integradas: Windows elige TRIM o desfragmentación. Confirme la limpieza aquí.','Optimizar SSD/HDD','Limpieza de disco','Programas de inicio','Herramienta abierta; complete la acción en Windows.','Detenga el análisis antes del mantenimiento.']};
 const words=()=>copy[document.documentElement.lang]||copy['en-US'];
 function paint(){const w=words();['z310MaintTitle','z310MaintInfo','z310Drive','z310Clean','z310Startup'].forEach((id,i)=>z(id).textContent=w[i])}
 for(const [id,kind] of [['z310Drive','drive'],['z310Clean','cleanup'],['z310Startup','startup']]){
 z(id).onclick=async()=>{const w=words();if(kind!=='startup'&&z310Latest?.running){z('z310MaintStatus').textContent=w[6];return}
 if(kind==='cleanup'){z('z310Diagnose').click();return} z('z310MaintStatus').textContent='Gerenciamento de inicialização externo removido; nenhuma alteração aplicada.'};
 }
 paint();new MutationObserver(paint).observe(document.documentElement,{attributes:true,attributeFilter:['lang']});
})();
"""
s=once(s,'</script></body></html>',js+'\n</script></body></html>')
p.write_text(s,encoding="utf-8")
print("310 feedback integration complete")
