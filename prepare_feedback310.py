"""Final UI integration: rates, estimate and actual report-driven sound."""
from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise ValueError("310 anchor "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8-sig")
s=once(s,'if !bind("getScanState",','if !bind("z310Beep",z310Beep){return}\n    if !bind("getScanState",')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'zLivePaint(s);','zLivePaint(s);z310Paint(s);')
s=once(s,'</style></head>','#z310Notice{border:2px solid #b87919;padding:10px;margin:8px 0;overflow-wrap:anywhere}#z310Speed{font-weight:600}#z310Balance{overflow-wrap:anywhere}</style></head>')
s=once(s,'</header>','</header><div id="z310Notice" class="card hidden" role="status"><p id="z310AlertText"></p><button id="z310Open">Ver alertas</button></div>')
js=Path("z310_metrics.js").read_text(encoding="utf-8")+r"""
let z310Previous=null,z310Latest=null,z310Used=null,z310SoundAt=-10000;
function z310Words(){return ({
'pt-BR':['Velocidade de análise','arquivos/s','Aguardando medição','Parada','Saldo estimado do espaço ocupado','Não é o total real restante','Estimativa indisponível neste escopo','Novos alertas no relatório','Ver alertas','Revise os achados; isolamento exige confirmação após parar ou concluir a análise.'],
'en-US':['Scan rate','files/s','Measuring','Stopped','Estimated used-space balance','Not the actual remaining scope','Estimate unavailable for this scope','New report alerts','View alerts','Review findings; isolation requires confirmation after stopping or completing the scan.'],
'es-ES':['Velocidad de análisis','archivos/s','Midiendo','Detenido','Saldo estimado del espacio ocupado','No es el total real restante','Estimación no disponible para este ámbito','Nuevas alertas','Ver alertas','Revise los hallazgos; el aislamiento requiere confirmación después de detener o finalizar.']
})[document.documentElement.lang]||['Scan rate','files/s','Measuring','Stopped','Estimated used-space balance','Not the actual remaining scope','Estimate unavailable for this scope','New report alerts','View alerts','Review findings after stopping or completing the scan.']}
function z310Balance(){
 const s=z310Latest||{},w=z310Words(),v=z310Remaining(z310Used,Number(s.bytes),s.report?.scan_mode);
 document.getElementById('z310Balance').textContent=v===null?w[6]:w[4]+': '+(v/1073741824).toLocaleString(document.documentElement.lang,{maximumFractionDigits:2})+' GiB · '+w[5];
}
window.z310Storage=n=>{z310Used=n;z310Balance()};
function z310Paint(s){
 const now=performance.now(),d=z310Delta(z310Previous,s,now),w=z310Words(),z=id=>document.getElementById(id);
 const prior=z310Previous;
 z310Previous=d.state;z310Latest=s;z310Balance();
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
"""
s=once(s,'</script></body></html>',js+'\n</script></body></html>')
p.write_text(s,encoding="utf-8")
print("310 feedback integration complete")
