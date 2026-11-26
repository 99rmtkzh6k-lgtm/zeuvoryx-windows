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

let z316History=[],z316Chart=null,z316Session='',z316Last=0,z316Ended=false;
function z316PaintChart(d,s,now){
 const canvas=document.getElementById('z316ScanChart'),msg=document.getElementById('z316ChartStatus');
 if(!canvas||!msg||!d.state)return;
 if(z316Session!==d.state.id){z316History=[];z316Session=d.state.id;z316Last=0;z316Ended=false;}
 if(z316Ended)return;
 const finished=!s.running&&s.done;
 if(finished){z316Ended=true;msg.textContent=s.report?.partial?'Encerrada com cobertura parcial':'Sessão encerrada';}
 else if(!s.running)return;
 else if(now-z316Last<1000)return;
 else{z316Last=now;z316History.push({at:new Date().toLocaleTimeString(),mb:d.bytesRate===null?null:d.bytesRate/1048576,files:d.filesRate});if(z316History.length>120)z316History.shift();
 msg.textContent=s.report?.inventory_314?.scanner_done?'Scanner próprio concluído; etapa complementar ou finalização.':d.bytesRate===null?'Aguardando medição':d.bytesRate===0&&d.filesRate===0?'Sem avanço na última amostra; verificando etapa.':'Taxas medidas do scanner próprio.';}
 if(typeof Chart!=='function'){msg.textContent+=' Biblioteca gráfica indisponível.';return;}
 try{
 if(!z316Chart)z316Chart=new Chart(canvas,{type:'line',data:{labels:[],datasets:[{label:'MiB/s',data:[],borderColor:'#62dfd1',yAxisID:'y',pointRadius:0},{label:'Arquivos/s',data:[],borderColor:'#ffb760',yAxisID:'y1',pointRadius:0}]},options:{responsive:true,maintainAspectRatio:false,animation:false,spanGaps:false,plugins:{legend:{labels:{color:'#b7c6d8'}}},scales:{x:{ticks:{maxTicksLimit:5,color:'#b7c6d8'}},y:{beginAtZero:true,title:{display:true,text:'MiB/s'}},y1:{position:'right',beginAtZero:true,grid:{drawOnChartArea:false},title:{display:true,text:'Arquivos/s'}}}}});
 z316Chart.data.labels=z316History.map(v=>v.at);z316Chart.data.datasets[0].data=z316History.map(v=>v.mb);z316Chart.data.datasets[1].data=z316History.map(v=>v.files);z316Chart.update('none');
 }catch(e){msg.textContent+=' Gráfico indisponível: '+String(e).slice(0,150);}
}

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
 z310Previous=d.state;z310Latest=s;z310Balance();z316PaintChart(d,s,now);
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

# Embed pinned Chart.js directly, with licenses, no runtime CDN/network.
from urllib.request import urlopen, Request
import hashlib, base64, tarfile, io, html
url='https://registry.npmjs.org/chart.js/-/chart.js-4.5.1.tgz'
with urlopen(Request(url,headers={'User-Agent':'ZEUVORYX-build/1'}),timeout=30) as response:
    blob=response.read(8*1024*1024+1)
expected='GIjfiT9dbmHRiYi6Nl2yFCq7kkwdkp1W/lp2J99rX0yo9tgJGn3lKQATztIjb5tVtevcBtIdICNWqlq5+E8/Pw=='
if len(blob)>8*1024*1024 or base64.b64encode(hashlib.sha512(blob).digest()).decode()!=expected:raise RuntimeError('Chart.js integrity failure')
with tarfile.open(fileobj=io.BytesIO(blob),mode='r:gz') as tar:
    def read_member(name,limit):
        member=tar.getmember(name)
        if not member.isfile() or member.size>limit:raise RuntimeError('Chart.js invalid member '+name)
        return tar.extractfile(member).read().decode('utf-8')
    vendor=read_member('package/dist/chart.umd.min.js',500000)
    license=read_member('package/LICENSE.md',30000)
if '</script' in vendor.lower():raise RuntimeError('Unsafe inline chart script')
# Keep one script block: existing desktop validation extracts that block.
p=Path('desktop.html');s=p.read_text(encoding='utf-8')
import re
s,n=re.subn(r'(<p id="z310Speed"[^>]*>[^<]*</p>)',lambda match:match.group(1)+'<div style="height:150px;position:relative"><canvas id="z316ScanChart" role="img" aria-label="Histórico real de velocidade da varredura"></canvas></div><p id="z316ChartStatus" class="sub" role="status">Aguardando dados da sessão.</p><details><summary>Licenças dos gráficos</summary><pre style="white-space:pre-wrap">'+html.escape(license)+'</pre></details>',s,count=1)
if n!=1:raise RuntimeError('Chart speed anchor missing')
# Inline vendor before first existing app script, not a remote script.
s=once(s,'<script>','<script>\n'+vendor+'\n;\n')
p.write_text(s,encoding='utf-8')
