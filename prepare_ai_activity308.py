"""Bind actual model activity, no fake progress, no autonomous training."""
from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise ValueError("AI anchor "+a[:100])
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8-sig")
if '"github.com/zeuvoryx/zeuvoryx/adaptive"' not in s:
    s=once(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/adaptive"')
s=once(s,'if !bind("getScanState",','''go adaptive.SelfCheck() // local startup check, not a file inference
    if !bind("zAIStatus", adaptive.Status) { return }
    if !bind("zAISelfCheck", adaptive.SelfCheck) { return }
    if !bind("getScanState",''')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'</style></head>','''.zRobotWrap{display:flex;align-items:center;gap:12px;min-height:94px;margin:5px 0}.zRobot{flex:0 0 108px;overflow:visible}.zRobotBody{transform-origin:60px 55px}.zRobotLift{transform-origin:60px 54px}.zRobotBed{opacity:0}.zRobotWrap[data-mode="working"] .zRobotLift{animation:zLift .75s ease-in-out 2}.zRobotWrap[data-mode="idle"] .zRobotBody{transform:translate(0px,19px) rotate(-90deg) scale(.80)}.zRobotWrap[data-mode="idle"] .zRobotBed{opacity:1}.zRobotWrap[data-mode="error"] .zRobotBody{opacity:.55}@keyframes zLift{0%,100%{transform:translateY(0)}50%{transform:translateY(-12px)}}@media(prefers-reduced-motion:reduce){.zRobotLift{animation:none!important}}#zAIActivity{margin:8px 0} .zai-dot{display:inline-block;width:10px;height:10px;border-radius:50%;background:#89929b;margin-right:6px}.zai-active .zai-dot{background:#168578;animation:zaiPulse .65s ease-in-out 2}@keyframes zaiPulse{50%{opacity:.35;transform:scale(1.35)}}@media(prefers-reduced-motion:reduce){.zai-active .zai-dot{animation:none}}</style></head>''')
panel='''<div class="card" id="zAIActivity"><h2 id="zAITitle">Inteligência artificial — atividade real</h2><div class="zRobotWrap" id="zAIRobot" data-mode="idle"><svg class="zRobot" viewBox="0 0 120 100" width="108" height="90" aria-hidden="true"><g class="zRobotBody" stroke="#50718f" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><path d="M60 15v-7" fill="none"/><circle cx="60" cy="6" r="3" fill="#18aabb"/><rect x="40" y="15" width="40" height="30" rx="9" fill="#dbe5ee"/><circle cx="51" cy="28" r="3" fill="#147b97" stroke="none"/><circle cx="69" cy="28" r="3" fill="#147b97" stroke="none"/><path d="M53 36h14" fill="none"/><rect x="43" y="49" width="34" height="27" rx="7" fill="#9bb8ca"/><path d="M52 76v15h-9m25-15v15h9" fill="none"/><g class="zRobotLift"><path d="M43 54 29 44 21 54m56 0 14-10 8 10" fill="none"/><path d="M12 55h96" stroke="#42566a"/><path d="M12 48v14m7-16v18m82-18v18m7-16v14" stroke="#42566a" stroke-width="5"/></g></g><path class="zRobotBed" d="M10 88h101M15 88v8m90-8v8" stroke="#6e8191" stroke-width="3" fill="none"/></svg><div id="zRobotText" class="sub" role="status">Inicializando modelo</div></div><p class="sub"><span class="zai-dot"></span><span id="zAIState">—</span></p><div class="grid"><div><div class="label" id="zAICountLabel">Avaliações do modelo</div><div class="metric" id="zAICount">—</div></div><div><div class="label" id="zAISuccessLabel">Concluídas</div><div class="metric" id="zAISuccess">—</div></div><div><div class="label" id="zAIErrorLabel">Falhas</div><div class="metric" id="zAIErrors">—</div></div></div><p class="sub" id="zAIDetails">—</p><p class="sub" id="zAILimit">—</p><button id="zAITest">Testar IA agora</button><p class="sub" id="zAITestResult" role="status">—</p></div>'''
s=once(s,'<section id="engine" class="hidden card">','<section id="engine" class="hidden card">'+panel)
s=once(s,'<section id="overview">','<section id="overview"><div class="zRobotWrap" id="zAIOverviewRobot" data-mode="idle"><svg class="zRobot" viewBox="0 0 120 100" width="108" height="90" aria-hidden="true"><g class="zRobotBody" stroke="#50718f" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><path d="M60 15v-7" fill="none"/><circle cx="60" cy="6" r="3" fill="#18aabb"/><rect x="40" y="15" width="40" height="30" rx="9" fill="#dbe5ee"/><circle cx="51" cy="28" r="3" fill="#147b97" stroke="none"/><circle cx="69" cy="28" r="3" fill="#147b97" stroke="none"/><path d="M53 36h14" fill="none"/><rect x="43" y="49" width="34" height="27" rx="7" fill="#9bb8ca"/><path d="M52 76v15h-9m25-15v15h9" fill="none"/><g class="zRobotLift"><path d="M43 54 29 44 21 54m56 0 14-10 8 10" fill="none"/><path d="M12 55h96" stroke="#42566a"/><path d="M12 48v14m7-16v18m82-18v18m7-16v14" stroke="#42566a" stroke-width="5"/></g></g><path class="zRobotBed" d="M10 88h101M15 88v8m90-8v8" stroke="#6e8191" stroke-width="3" fill="none"/></svg><div id="zAIOverview" class="sub" role="status">IA: inicializando modelo</div></div>')
js="""
;(()=>{
const z=id=>document.getElementById(id), panel=z('zAIActivity');
const t={
'pt-BR':['Inteligência artificial — atividade real','Avaliações do modelo','Concluídas','Falhas','Processando','Em espera por arquivos','Modelo','Última inferência','Índice auxiliar','Modelo consultivo: sinaliza indícios correlacionados para revisão; aprendizado validado pendente.','Testar IA agora','Autoteste sintético aprovado','Autoteste falhou','Ainda não testado','Testes sintéticos','Indisponível'],
'en-US':['Artificial intelligence — live activity','Model evaluations','Completed','Errors','Processing','Waiting for files','Model','Last inference','Advisory score','Advisory model flags correlated indicators; validated learning pending.','Test AI now','Synthetic self-check passed','Self-check failed','Not tested','Synthetic checks','Unavailable'],
'es-ES':['Inteligencia artificial — actividad real','Evaluaciones del modelo','Completadas','Errores','Procesando','Esperando archivos','Modelo','Última inferencia','Índice auxiliar','Modelo consultivo: indicadores correlacionados para revisión; aprendizaje validado pendiente.','Probar IA','Autoprueba sintética aprobada','Autoprueba fallida','Sin prueba','Pruebas sintéticas','No disponible']
};
let last=null,lastErrors=null,busy=false;
const copy=()=>t[document.documentElement.lang]||t['pt-BR'];
function paint(d){const a=copy();
 ['zAITitle','zAICountLabel','zAISuccessLabel','zAIErrorLabel'].forEach((k,i)=>z(k).textContent=a[i]);
 z('zAITest').textContent=a[10];z('zAILimit').textContent=a[9];
 const active=d.running>0||(last!==null&&d.successful>last);
 if(active){panel.classList.remove('zai-active');void panel.offsetWidth;panel.classList.add('zai-active')}else panel.classList.remove('zai-active');
 const lang=document.documentElement.lang||'pt-BR';
 const words={
 'pt-BR':['Modelo pronto — descansando; sem novas avaliações','Trabalhando — avaliação consultiva de arquivos','Falha na IA — verificar diagnóstico','Inicializando modelo'],
 'en-US':['Model ready — resting; no new evaluations','Working — advisory file evaluations','AI error — review diagnostics','Initializing model'],
 'es-ES':['Modelo listo — descansando; sin nuevas evaluaciones','Trabajando — evaluaciones consultivas de archivos','Fallo de IA — revisar diagnóstico','Inicializando modelo']
 }[lang]||['Modelo pronto — descansando','Trabalhando — avaliações reais','Falha na IA','Inicializando modelo'];
 const failed=(d.self_tests&&!d.self_test_ok)||(lastErrors!==null&&d.errors>lastErrors);
 const mode=failed?'error':active?'working':'idle';
 const message=failed?words[2]:active?words[1]:d.self_test_ok?words[0]:words[3];
 for(const id of ['zAIRobot','zAIOverviewRobot']){const robot=z(id);robot.dataset.mode=mode;if(mode==='working'){const lift=robot.querySelector('.zRobotLift');lift.style.animation='none';void lift.getBoundingClientRect();lift.style.animation='';}}
 z('zRobotText').textContent=message;z('zAIOverview').textContent=message+' · '+a[1]+': '+d.observations;
 last=d.successful;lastErrors=d.errors;z('zAIState').textContent=message;
 z('zAICount').textContent=d.observations;z('zAISuccess').textContent=d.successful;z('zAIErrors').textContent=d.errors;
 z('zAIDetails').textContent=a[6]+': '+d.version+' · '+a[7]+': '+(d.last?new Date(d.last).toLocaleTimeString(document.documentElement.lang):'—')+' · '+a[8]+': '+(typeof d.last_score==='number'?Math.round(d.last_score*100)+'/100':'—');
 z('zAITestResult').textContent=(d.self_tests?(d.self_test_ok?a[11]:a[12]):a[13])+' · '+a[14]+': '+d.self_tests;
}
async function refresh(){if(busy||document.hidden||(z('engine').classList.contains('hidden')&&z('overview').classList.contains('hidden')))return;busy=true;try{paint(await window.zAIStatus())}catch(_){z('zAIState').textContent=copy()[15];z('zRobotText').textContent=copy()[15];z('zAIOverview').textContent=copy()[15];for(const id of ['zAIRobot','zAIOverviewRobot'])z(id).dataset.mode='error'}finally{busy=false}}
z('zAITest').onclick=async()=>{z('zAITest').disabled=true;z('zAITestResult').textContent=({'pt-BR':'Autoteste sintético em andamento','en-US':'Synthetic self-check running','es-ES':'Autoprueba sintética en curso'})[document.documentElement.lang]||'Synthetic self-check running';try{paint(await window.zAISelfCheck())}catch(_){z('zAITestResult').textContent=copy()[15]}finally{z('zAITest').disabled=false}};
setInterval(refresh,2000);refresh();
})();
"""
s=once(s,'</script></body></html>',js+'</script></body></html>')
p.write_text(s,encoding="utf-8")
