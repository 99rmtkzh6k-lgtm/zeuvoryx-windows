"""Integrated conservative maintenance after desktop generators."""
from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise ValueError("optimizer anchor "+a[:120]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8-sig")
s=once(s,'if !bind("getScanState",','''var optimizePlan opt310Plan
    if !bind("z310DriveOptimize",func()(uint32,error){
     if !operationGate.TryLock(){return 0,fmt.Errorf("Another operation is running")};defer operationGate.Unlock()
     s.mu.Lock();running:=s.running;s.mu.Unlock()
     if running||native.status().Running{return 0,fmt.Errorf("Finish the scan before optimization")}
     return drive310Optimize()
    }){return}
    if !bind("z310OptimizePreview",func()(opt310Plan,error){
     operationGate.Lock();defer operationGate.Unlock()
     s.mu.Lock();running:=s.running;s.mu.Unlock()
     optimizePlan=opt310Plan{}
     if running||native.status().Running{return optimizePlan,fmt.Errorf("Finish the scan before maintenance")}
     root,e:=opt310Root();if e!=nil{return optimizePlan,e}
     optimizePlan,e=opt310Preview(root,time.Now());return optimizePlan,e
    }){return}
    if !bind("z310OptimizeApply",func(ids []int)(opt310Result,error){
     operationGate.Lock();defer operationGate.Unlock()
     s.mu.Lock();running:=s.running;s.mu.Unlock()
     plan:=optimizePlan;optimizePlan=opt310Plan{}
     if running||native.status().Running{return opt310Result{},fmt.Errorf("Finish the scan before maintenance")}
     return opt310Apply(plan,ids,time.Now())
    }){return}
    if !bind("getScanState",''')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'<p class="sub" id="z310MaintStatus" role="status"></p>','''<p class="sub" id="z310MaintStatus" role="status"></p><div id="z310Optimizer"><button id="z310Diagnose">Diagnosticar e otimizar</button><p id="z310Diagnosis" role="status" class="sub"></p><p id="z310PlanSummary" class="sub"></p><div id="z310FileList"></div><button id="z310Apply" disabled>Limpar selecionados</button><p id="z310Result" class="sub" role="status"></p></div>''')
s=once(s,'</style></head>','#z310FileList{max-height:240px;overflow:auto}#z310FileList label{display:block;overflow-wrap:anywhere;padding:5px 0}#z310Optimizer{margin-top:12px;border-top:1px solid var(--line);padding-top:12px}</style></head>')
s=once(s,'<div id="zPerformancePanel" class="card">','<div id="zPerformancePanel" class="card"><button id="z310GoOptimize">Diagnosticar e otimizar</button>')
js=r"""
;(()=>{
 const z=id=>document.getElementById(id);
 const texts={
 'pt-BR':['Diagnosticar e otimizar','Limpar selecionados','Medição do sistema (não é diagnóstico de defeito)','CPU','RAM','Uso elevado nas duas amostras. Avalie programas em execução.','Sem uso elevado persistente nas duas amostras.','Medição insuficiente.','Temporários antigos candidatos à limpeza','Lista limitada; examine novamente após limpar.','Excluir definitivamente os temporários selecionados? Não é quarentena.','Removidos','Ignorados/alterados','Tamanho lógico removido','Examine novamente para medir o computador; não há ganho de CPU garantido.','Encerre a varredura antes da manutenção.','Diagnosticando…','Selecione apenas temporários que deseja excluir.'],
 'en-US':['Diagnose and optimize','Clean selected','System measurement (not a fault diagnosis)','CPU','RAM','High use in both samples. Review running programs.','No sustained high use in both samples.','Insufficient measurement.','Old temporary cleanup candidates','Limited list; examine again after cleanup.','Permanently delete selected temporary files? This is not quarantine.','Removed','Skipped/changed','Logical size removed','Measure again; CPU gain is not guaranteed.','Finish the scan before maintenance.','Diagnosing…','Select only temporary files you want to delete.'],
 'es-ES':['Diagnosticar y optimizar','Limpiar seleccionados','Medición del sistema (no diagnóstico de avería)','CPU','RAM','Uso elevado en ambas muestras. Revise los programas.','Sin uso elevado persistente en ambas muestras.','Medición insuficiente.','Temporales antiguos candidatos','Lista limitada; examine nuevamente.','¿Eliminar definitivamente los temporales seleccionados? No es cuarentena.','Eliminados','Omitidos/cambiados','Tamaño lógico eliminado','Mida de nuevo; no se garantiza mejora de CPU.','Finalice el análisis antes del mantenimiento.','Diagnosticando…','Seleccione solo los temporales que desea eliminar.']};
 const w=()=>texts[document.documentElement.lang]||texts['en-US'];
 let busy=false,plan=null;
 const valid=n=>typeof n==='number'&&Number.isFinite(n)&&n>=0&&n<=100;
 const percent=n=>valid(n)?Math.round(n)+'%':'—';
 function sync(){z('z310Apply').disabled=busy||!plan||!z('z310FileList').querySelector('input:checked')}
 function lang(){const a=w();z('z310GoOptimize').textContent=z('z310Diagnose').textContent=a[0];z('z310Apply').textContent=a[1]}
 function clear(){plan=null;z('z310FileList').replaceChildren();z('z310PlanSummary').textContent='';sync()}
 z('z310GoOptimize').onclick=()=>{document.querySelector('[data-page=tools]').click();z('z310Diagnose').click()};
 z('z310Drive').onclick=async()=>{
  const lang=document.documentElement.lang;
  const a=({'pt-BR':['Otimizar a unidade do Windows agora? O Windows escolhe a operação adequada ao disco. Mantenha o computador na tomada e aguarde.','Otimizando — aguarde o Windows. Não inicie varreduras.','Windows informou conclusão (código 0). Ganho de desempenho não medido.'],
  'en-US':['Optimize the Windows drive now? Windows selects the appropriate operation. Keep AC connected.','Optimizing — wait for Windows. Do not start scans.','Windows reported completion (exit 0). Performance gain not measured.'],
  'es-ES':['¿Optimizar la unidad de Windows? Mantenga el equipo conectado a la corriente.','Optimizando — espere a Windows.','Windows informó finalización (código 0). Mejora no medida.']})[lang]||['Optimize Windows drive?','Optimizing…','Windows reported completion.'];
  if(busy||z310Latest?.running){z('z310MaintStatus').textContent=w()[15];return}
  busy=true;z('z310Drive').disabled=true;
  try{if(!await ask(a[0]))return;z('z310MaintStatus').textContent=a[1];await window.z310DriveOptimize();z('z310MaintStatus').textContent=a[2]}catch(e){z('z310MaintStatus').textContent=String(e)}finally{busy=false;z('z310Drive').disabled=false;sync()}
 };

 z('z310Diagnose').onclick=async()=>{
  if(busy)return;const a=w();if(z310Latest?.running){z('z310Diagnosis').textContent=a[15];return}
  busy=true;z('z310Diagnose').disabled=true;clear();z('z310Result').textContent='';z('z310Diagnosis').textContent=a[16];
  try{
    const first=await window.zPerformance();await new Promise(r=>setTimeout(r,2200));const second=await window.zPerformance();
    const vals=[first.cpu,second.cpu,first.ram_percent,second.ram_percent],known=vals.every(valid);
    const high=known&&((first.cpu>=80&&second.cpu>=80)||(first.ram_percent>=85&&second.ram_percent>=85));
    z('z310Diagnosis').textContent=a[2]+': '+a[3]+' '+percent(first.cpu)+' / '+percent(second.cpu)+' · '+a[4]+' '+percent(first.ram_percent)+' / '+percent(second.ram_percent)+' · '+a[known?(high?5:6):7];
    plan=await window.z310OptimizePreview();
    z('z310PlanSummary').textContent=a[8]+': '+plan.items.length+' · '+volume(plan.bytes)+(plan.truncated?' · '+a[9]:'')+' · '+a[17];
    for(const v of plan.items){const label=document.createElement('label'),c=document.createElement('input');c.type='checkbox';c.value=String(v.id);c.onchange=sync;label.append(c,document.createTextNode(' '+v.name+' · '+volume(v.bytes)+' · '+new Date(v.modified).toLocaleString(document.documentElement.lang)));z('z310FileList').append(label)}
  }catch(e){z('z310Result').textContent=String(e);clear()}finally{busy=false;z('z310Diagnose').disabled=false;sync()}
 };
 z('z310Apply').onclick=async()=>{
  if(busy||!plan)return;const a=w(),ids=Array.from(z('z310FileList').querySelectorAll('input:checked'),c=>Number(c.value));if(!ids.length)return;
  busy=true;sync();z('z310Diagnose').disabled=true;
  try{if(!await ask(a[10]+' ('+ids.length+')'))return;
   const r=await window.z310OptimizeApply(ids);z('z310Result').textContent=a[11]+': '+r.removed+' · '+a[12]+': '+r.skipped+' · '+a[13]+': '+volume(r.bytes)+' · '+a[14];clear();
  }catch(e){z('z310Result').textContent=String(e);clear()}finally{busy=false;z('z310Diagnose').disabled=false;sync()}
 };
 lang();new MutationObserver(lang).observe(document.documentElement,{attributes:true,attributeFilter:['lang']});
})();
"""
s=once(s,'</script></body></html>',js+'\n</script></body></html>')
p.write_text(s,encoding="utf-8")
