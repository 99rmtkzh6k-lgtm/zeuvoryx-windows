"""R2 native performance panel after all desktop generators."""
from pathlib import Path
from html.parser import HTMLParser
def once(s,a,b):
    if s.count(a)!=1: raise ValueError("Missing/duplicate performance anchor: "+a[:100])
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8-sig")
s=once(s,'if !bind("getScanState",','''var perf pSampler
    if !bind("zPerformance", perf.sample) { return }
    if !bind("zPerformanceRendered", perf.rendered) { return }
    if !bind("getScanState",''')
s=s.replace('ZEUVORYX | 3.0.8 LAB','ZEUVORYX | 3.0.8 LAB R2')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8-sig")
css="""
#zPerformancePanel{margin:0 0 10px;padding:10px 12px}
.zp-head{display:flex;align-items:center;justify-content:space-between;gap:6px;flex-wrap:wrap}
.zp-head h2{font-size:15px}.zp-grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(115px,1fr));gap:6px;margin-top:6px}
.zp-item{text-align:center;border:1px solid var(--line);border-radius:9px;padding:5px 3px;background:var(--panel)}
.zp-item svg{display:block;width:100%;height:58px;max-width:110px;margin:auto}
.zp-track{fill:none;stroke:#a6adb5;stroke-width:9;stroke-linecap:round}
.zp-arc{fill:none;stroke:#237f85;stroke-width:9;stroke-linecap:round;stroke-dasharray:0 100}
.zp-value{font-size:19px;font-weight:700;fill:currentColor}
.zp-label{font-size:12px;font-weight:600}.zp-detail{font-size:11px;min-height:28px;color:var(--muted);overflow-wrap:anywhere}
#zpClock,#zpStatus{font-size:11px;color:var(--muted)}
@media(forced-colors:active){.zp-arc{stroke:Highlight}.zp-track{stroke:GrayText}}
"""
panel='<div id="zPerformancePanel" class="card"><div class="zp-head"><h2 id="zpTitle">Desempenho do computador</h2><span id="zpClock">—</span></div><div class="zp-grid">'
for k,label in [('cpu','Processador'),('ram','Memória RAM'),('free','RAM disponível'),('disk','Disco do Windows')]:
    panel+=f'<div class="zp-item"><div class="zp-label" id="zp-{k}-label">{label}</div><svg viewBox="0 0 120 72" aria-hidden="true"><path class="zp-track" d="M12 60 A48 48 0 0 1 108 60"/><path id="zp-{k}-arc" class="zp-arc" pathLength="100" d="M12 60 A48 48 0 0 1 108 60"/><text class="zp-value" id="zp-{k}-value" x="60" y="56" text-anchor="middle">—</text></svg><div id="zp-{k}-detail" class="zp-detail">—</div></div>'
panel+='</div><p id="zpStatus" class="sub" role="status">Aguardando dados do Windows</p></div>'
s=once(s,'</style></head>',css+'</style></head>')
s=once(s,'<section id="overview">','<section id="overview">'+panel)
s=once(s,'3.0.8 LAB BETA</div>','3.0.8 LAB BETA</div><div class="sub" id="zpRevision">R2 · Performance</div>')
js="""
;(()=>{
const z=id=>document.getElementById(id);
const dict={
'pt-BR':['Desempenho do computador','Processador','Memória RAM','RAM disponível','Disco do Windows','núcleos lógicos','em uso','livres','de','Dados locais do Windows','Aguardando dados do Windows','Medição indisponível','GiB'],
'en-US':['Computer performance','Processor','Physical RAM','Available RAM','Windows drive','logical processors','in use','available','of','Local Windows metrics','Waiting for Windows metrics','Measurement unavailable','GiB'],
'es-ES':['Rendimiento del equipo','Procesador','Memoria RAM','RAM disponible','Disco de Windows','procesadores lógicos','en uso','libres','de','Datos locales de Windows','Esperando datos de Windows','Medición no disponible','GiB']};
let busy=false,confirmed=false;
const valid=n=>typeof n==='number'&&Number.isFinite(n)&&n>=0&&n<=100;
function gauge(k,n,detail){const v=valid(n);z('zp-'+k+'-arc').setAttribute('stroke-dasharray',v?n+' 100':'0 100');z('zp-'+k+'-value').textContent=v?Math.round(n)+'%':'—';z('zp-'+k+'-detail').textContent=detail;}
async function refresh(){
 if(busy||document.hidden||z('overview').classList.contains('hidden'))return;
 busy=true;
 const lang=document.documentElement.lang||'pt-BR', t=dict[lang]||dict['pt-BR'];
 const gib=n=>(Number(n)/1073741824).toLocaleString(lang,{maximumFractionDigits:1})+' '+t[12];
 z('zpTitle').textContent=t[0];['cpu','ram','free','disk'].forEach((k,i)=>z('zp-'+k+'-label').textContent=t[i+1]);
 try{
  const d=await window.zPerformance();
  z('zpClock').textContent=new Date().toLocaleTimeString(lang);
  gauge('cpu',d.cpu,d.logical_cpu>0?d.logical_cpu+' '+t[5]:t[11]);
  const ram=valid(d.ram_percent)&&d.ram_total>0&&d.ram_available<=d.ram_total;
  gauge('ram',ram?d.ram_percent:null,ram?gib(d.ram_total-d.ram_available)+' '+t[6]+' '+t[8]+' '+gib(d.ram_total):t[11]);
  gauge('free',ram?100-d.ram_percent:null,ram?gib(d.ram_available)+' '+t[7]+' '+t[8]+' '+gib(d.ram_total):t[11]);
  const disk=valid(d.disk_percent)&&d.disk_total>0&&d.disk_free<=d.disk_total;
  gauge('disk',disk?d.disk_percent:null,disk?d.disk_root+' · '+gib(d.disk_total-d.disk_free)+' '+t[6]+' / '+gib(d.disk_total):t[11]);
  z('zpStatus').textContent=ram&&valid(d.cpu)?t[9]+' · '+new Date(d.at).toLocaleTimeString(lang):t[10];
  if(!confirmed&&ram&&valid(d.cpu)&&z('zp-cpu-value').textContent.endsWith('%')&&z('zp-ram-value').textContent.endsWith('%')) confirmed=await window.zPerformanceRendered();
 }catch(_){['cpu','ram','free','disk'].forEach(k=>gauge(k,null,t[11]));z('zpStatus').textContent=t[11];z('zpClock').textContent='—';}
 finally{busy=false}
}
const overview=z('overview'),hero=Array.from(overview.children).find(n=>n.classList.contains('hero'));
if(hero)hero.after(z('zPerformancePanel'));
setInterval(refresh,3000);refresh();
})();
"""
s=once(s,'</script></body></html>',js+'</script></body></html>')
class Check(HTMLParser):
    def __init__(self):super().__init__();self.ids=[]
    def handle_starttag(self,tag,attrs):
        a=dict(attrs)
        if "id" in a:self.ids.append(a["id"])
c=Check();c.feed(s)
for k in ['zPerformancePanel','zpTitle','zpClock','zpRevision','start','quick','stop','export']+[f'zp-{k}-{x}' for k in ('cpu','ram','free','disk') for x in ('label','arc','value','detail')]:
    if c.ids.count(k)!=1:raise ValueError("Bad panel/control "+k)
p.write_text(s,encoding="utf-8")
print("R2 native performance panel and controls verified")
