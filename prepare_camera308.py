"""Read-only webcam status and explicit link to Windows privacy settings."""
from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise ValueError("Camera anchor: "+a[:90])
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'if !bind("getScanState",','''var cam cameraSampler
    if !bind("zCameraStatus", cam.sample) { return }
    if !bind("getScanState",''')
p.write_text(s,encoding="utf-8")
p=Path("security_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'"updates": "ms-settings:windowsupdate"','"updates": "ms-settings:windowsupdate", "camera": "ms-settings:privacy-webcam"')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
panel='''<div class="card" id="zCameraPanel" style="margin:8px 0"><h2 id="zcTitle">Segurança da câmera</h2><p class="sub" role="status" id="zcStatus">—</p><div class="zcLegend"><span id="zcBlue"><i style="background:#2684ff"></i><span id="zcBlueText">Sem uso identificado</span></span><span id="zcYellow"><i style="background:#eab308"></i><span id="zcYellowText">Indeterminado</span></span><span id="zcRed"><i style="background:#dc2626"></i><span id="zcRedText">Indício de uso</span></span></div><p class="sub" id="zcLimit">—</p><button id="zcSettings">Privacidade da câmera no Windows</button></div>'''
s=once(s,'</style></head>','.zcLegend{display:flex;flex-wrap:wrap;gap:6px;margin:8px 0}.zcLegend>span{display:inline-flex;align-items:center;gap:4px;padding:4px 7px;border:1px solid var(--line);border-radius:7px;font-size:11px}.zcLegend i{display:inline-block;width:10px;height:10px;border-radius:50%;border:1px solid #555}.zcLegend .zcSelected{outline:2px solid currentColor;font-weight:700}</style></head>')
s=once(s,'<section id="overview">','<section id="overview">'+panel)
js="""
;(()=>{
const z=id=>document.getElementById(id);
const texts={
'pt-BR':['Segurança da câmera','Sem uso identificado','Indeterminado','Indício de uso','Apenas registros locais do Windows; não confirma câmera desligada, invasão ou bloqueio.','Privacidade da câmera no Windows','Estado: '],
'en-US':['Camera security','No use recorded','Unknown','Possible use','Local Windows records only; does not confirm camera off, intrusion or blocking.','Windows camera privacy','Status: '],
'es-ES':['Seguridad de la cámara','Sin uso registrado','Indeterminado','Indicio de uso','Solo registros locales de Windows; no confirma cámara apagada, intrusión ni bloqueo.','Privacidad de la cámara en Windows','Estado: ']
};
let busy=false;
function draw(state){const t=texts[document.documentElement.lang]||texts['pt-BR'];const current=state==='activity'?3:state==='no_activity_recorded'?1:2;
['zcTitle','zcBlueText','zcYellowText','zcRedText','zcLimit','zcSettings'].forEach((k,i)=>z(k).textContent=t[i]);
['zcBlue','zcYellow','zcRed'].forEach((k,i)=>{z(k).classList.toggle('zcSelected',current===i+1);z(k).setAttribute('aria-current',current===i+1?'true':'false')});
z('zcStatus').textContent=t[6]+t[current];}
async function refresh(){if(busy||document.hidden||z('overview').classList.contains('hidden'))return;busy=true;try{draw(await window.zCameraStatus())}catch(_){draw('unknown')}finally{busy=false}}
z('zcSettings').onclick=()=>action(()=>window.openSecuritySettings('camera'));
const p=z('zPerformancePanel');if(p)p.after(z('zCameraPanel'));
draw('unknown');setInterval(refresh,10000);refresh();
})();
"""
s=once(s,'</script></body></html>',js+'</script></body></html>')
p.write_text(s,encoding="utf-8")
