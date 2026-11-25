"""Bind actual model activity, no fake progress, no autonomous training."""
from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise ValueError("AI anchor "+a[:100])
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8-sig")
if '"github.com/zeuvoryx/zeuvoryx/adaptive"' not in s:
    s=once(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/adaptive"')
s=once(s,'if !bind("getScanState",','''if !bind("zAIStatus", adaptive.Status) { return }
    if !bind("zAISelfCheck", adaptive.SelfCheck) { return }
    if !bind("getScanState",''')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'</style></head>','''#zAIActivity{margin:8px 0} .zai-dot{display:inline-block;width:10px;height:10px;border-radius:50%;background:#89929b;margin-right:6px}.zai-active .zai-dot{background:#168578;animation:zaiPulse .65s ease-in-out 2}@keyframes zaiPulse{50%{opacity:.35;transform:scale(1.35)}}@media(prefers-reduced-motion:reduce){.zai-active .zai-dot{animation:none}}</style></head>''')
panel='''<div class="card" id="zAIActivity"><h2 id="zAITitle">Inteligência artificial — atividade real</h2><p class="sub"><span class="zai-dot"></span><span id="zAIState">—</span></p><div class="grid"><div><div class="label" id="zAICountLabel">Inferências em arquivos</div><div class="metric" id="zAICount">—</div></div><div><div class="label" id="zAISuccessLabel">Concluídas</div><div class="metric" id="zAISuccess">—</div></div><div><div class="label" id="zAIErrorLabel">Falhas</div><div class="metric" id="zAIErrors">—</div></div></div><p class="sub" id="zAIDetails">—</p><p class="sub" id="zAILimit">—</p><button id="zAITest">Testar IA agora</button><p class="sub" id="zAITestResult" role="status">—</p></div>'''
s=once(s,'<section id="engine" class="hidden card">','<section id="engine" class="hidden card">'+panel)
js="""
;(()=>{
const z=id=>document.getElementById(id), panel=z('zAIActivity');
const t={
'pt-BR':['Inteligência artificial — atividade real','Inferências em arquivos','Concluídas','Falhas','Processando','Em espera por evidência heurística','Modelo','Última inferência','Índice auxiliar','Aprendizado contínuo ainda não implantado; modelo consultivo, não altera a decisão do scanner.','Testar IA agora','Autoteste sintético aprovado','Autoteste falhou','Ainda não testado','Testes sintéticos','Indisponível'],
'en-US':['Artificial intelligence — live activity','File inferences','Completed','Errors','Processing','Waiting for heuristic evidence','Model','Last inference','Advisory score','Continuous learning not implemented; advisory model, scanner verdict unchanged.','Test AI now','Synthetic self-check passed','Self-check failed','Not tested','Synthetic checks','Unavailable'],
'es-ES':['Inteligencia artificial — actividad real','Inferencias en archivos','Completadas','Errores','Procesando','Esperando evidencia heurística','Modelo','Última inferencia','Índice auxiliar','Aprendizaje continuo no implantado; modelo consultivo, no altera la decisión.','Probar IA','Autoprueba sintética aprobada','Autoprueba fallida','Sin prueba','Pruebas sintéticas','No disponible']
};
let last=null,busy=false;
const copy=()=>t[document.documentElement.lang]||t['pt-BR'];
function paint(d){const a=copy();
 ['zAITitle','zAICountLabel','zAISuccessLabel','zAIErrorLabel'].forEach((k,i)=>z(k).textContent=a[i]);
 z('zAITest').textContent=a[10];z('zAILimit').textContent=a[9];
 const active=d.running>0||(last!==null&&d.observations>last);
 if(active){panel.classList.remove('zai-active');void panel.offsetWidth;panel.classList.add('zai-active')}else panel.classList.remove('zai-active');
 last=d.observations;z('zAIState').textContent=d.running>0?a[4]:a[5];
 z('zAICount').textContent=d.observations;z('zAISuccess').textContent=d.successful;z('zAIErrors').textContent=d.errors;
 z('zAIDetails').textContent=a[6]+': '+d.version+' · '+a[7]+': '+(d.last?new Date(d.last).toLocaleTimeString(document.documentElement.lang):'—')+' · '+a[8]+': '+(typeof d.last_score==='number'?Math.round(d.last_score*100)+'/100':'—');
 z('zAITestResult').textContent=(d.self_tests?(d.self_test_ok?a[11]:a[12]):a[13])+' · '+a[14]+': '+d.self_tests;
}
async function refresh(){if(busy||document.hidden||z('engine').classList.contains('hidden'))return;busy=true;try{paint(await window.zAIStatus())}catch(_){z('zAIState').textContent=copy()[15]}finally{busy=false}}
z('zAITest').onclick=async()=>{z('zAITest').disabled=true;try{paint(await window.zAISelfCheck())}catch(_){z('zAITestResult').textContent=copy()[15]}finally{z('zAITest').disabled=false}};
setInterval(refresh,2000);refresh();
})();
"""
s=once(s,'</script></body></html>',js+'</script></body></html>')
p.write_text(s,encoding="utf-8")
