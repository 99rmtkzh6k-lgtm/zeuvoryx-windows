"""3.0.10 native performance panel after all desktop generators."""
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
if 'ZEUVORYX | 3.0.10 LAB' not in s: raise ValueError('Missing 3.0.10 identity')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8-sig")
css="""
#zPerformancePanel{margin:0 0 10px;padding:10px 12px}
.zp-head{display:flex;align-items:center;justify-content:space-between;gap:6px;flex-wrap:wrap}
.zp-head h2{font-size:15px}.zp-grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(115px,1fr));gap:6px;margin-top:6px}
.zp-item{text-align:center;border:1px solid var(--line);border-radius:9px;padding:5px 3px;background:var(--panel)}
.zp-item svg{display:block;width:100%;height:108px;max-width:176px;margin:auto}.zp-needle{stroke:#ffb760;stroke-width:2.8;stroke-linecap:round;transition:transform .35s ease}.zp-tick{stroke:#91a5b8;stroke-width:1}.zp-number{fill:#a9bed1;font-size:8px}.zp-hub{fill:#ffb760;stroke:#17263b;stroke-width:1.5}.zp-track{opacity:.6}
.zp-track{fill:none;stroke:#a6adb5;stroke-width:9;stroke-linecap:round}
.zp-arc{fill:none;stroke:#237f85;stroke-width:9;stroke-linecap:round;stroke-dasharray:0 100}
.zp-value{font-size:19px;font-weight:700;fill:currentColor}
.zp-label{font-size:12px;font-weight:600}.zp-detail{font-size:11px;min-height:28px;color:var(--muted);overflow-wrap:anywhere}
#zpClock,#zpStatus{font-size:11px;color:var(--muted)}
@media(forced-colors:active){.zp-arc{stroke:Highlight}.zp-track{stroke:GrayText}}
"""
panel='<div id="zPerformancePanel" class="card"><div class="zp-head"><h2 id="zpTitle">Desempenho do computador</h2><span id="zpClock">—</span></div><div class="zp-grid">'
for k,label in [('cpu','Processador'),('ram','Memória RAM'),('free','RAM disponível'),('disk','Disco do Windows')]:
    panel+=f'<div class="zp-item"><div class="zp-label" id="zp-{k}-label">{label}</div><svg viewBox="0 0 120 90" aria-hidden="true"><path class="zp-track" d="M12 60 A48 48 0 0 1 108 60"/><path id="zp-{k}-arc" class="zp-arc" pathLength="100" d="M12 60 A48 48 0 0 1 108 60"/><g class="zp-tick"><path d="M12 60h7 M26 26l5 5 M60 12v7 M94 26l-5 5 M108 60h-7"/></g><g class="zp-number" text-anchor="middle"><text x="12" y="70">0</text><text x="60" y="10">50</text><text x="108" y="70">100</text></g><line class="zp-needle" id="zp-{k}-needle" x1="60" y1="60" x2="23" y2="60"/><circle class="zp-hub" cx="60" cy="60" r="4"/><text class="zp-value" id="zp-{k}-value" x="60" y="84" text-anchor="middle">—</text></svg><div id="zp-{k}-detail" class="zp-detail">—</div></div>'
panel+='</div><p id="zpStatus" class="sub" role="status">Aguardando dados do Windows</p></div>'
s=once(s,'</style></head>',css+'</style></head>')
s=once(s,'<section id="overview">','<section id="overview">'+panel)
s=once(s,'<div class="metric" id="hits">—</div></div></div>','<div class="metric" id="hits">—</div></div><div class="card" id="zStorageCard"><div class="label" id="zsLabel">ARMAZENAMENTO E PROGRESSO</div><div class="metric" id="zsCapacity">—</div><p class="sub" id="zsSpace">—</p><p class="sub" id="zsProgress">—</p><p class="sub" id="z310Balance">—</p><progress id="z314ProgressBar" max="100" value="0" style="width:100%;height:12px;accent-color:#62dfd1" hidden></progress><p class="sub" id="z314Eta">—</p><p class="sub" id="z314ScopeNote"></p><p class="sub" id="z310Speed" role="status">—</p></div></div>')
s=once(s,'</style></head>','#overview .grid{grid-template-columns:repeat(auto-fit,minmax(132px,1fr))}</style></head>')
s=once(s,'3.0.10 LAB BETA</div>','3.0.10 LAB BETA</div><div class="sub" id="zpRevision">3.0.10 · Performance</div>')
js="""
;(()=>{
const z=id=>document.getElementById(id);
const dict={
'pt-BR':['Desempenho do computador','Processador','Memória RAM','RAM disponível','Disco do Windows','núcleos lógicos','em uso','livres','de','Dados locais do Windows','Aguardando dados do Windows','Medição indisponível','GiB'],
'en-US':['Computer performance','Processor','Physical RAM','Available RAM','Windows drive','logical processors','in use','available','of','Local Windows metrics','Waiting for Windows metrics','Measurement unavailable','GiB'],
'es-ES':['Rendimiento del equipo','Procesador','Memoria RAM','RAM disponible','Disco de Windows','procesadores lógicos','en uso','libres','de','Datos locales de Windows','Esperando datos de Windows','Medición no disponible','GiB']};
let busy=false,confirmed=false;
const valid=n=>typeof n==='number'&&Number.isFinite(n)&&n>=0&&n<=100;
function gauge(k,n,detail){const v=valid(n);const needle=z('zp-'+k+'-needle');needle.setAttribute('transform','rotate('+(v?n*1.8:0)+' 60 60)');needle.style.visibility=v?'visible':'hidden';z('zp-'+k+'-arc').setAttribute('stroke-dasharray',v?n+' 100':'0 100');z('zp-'+k+'-value').textContent=v?Math.round(n)+'%':'—';z('zp-'+k+'-detail').textContent=detail;}
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
  const st={
   'pt-BR':['CAPACIDADE DOS VOLUMES E PROGRESSO','ocupado para o usuário','disponível ao usuário','volumes acessíveis','omitidos','Analisado','Total elegível/restante: indeterminado'],
   'en-US':['VOLUME CAPACITY AND PROGRESS','used for caller','available to caller','accessible volumes','omitted','Scanned','Eligible total/remaining: unknown'],
   'es-ES':['CAPACIDAD DE VOLÚMENES Y PROGRESO','ocupado para el usuario','disponible al usuario','volúmenes accesibles','omitidos','Analizado','Total elegible/restante: indeterminado']
  }[lang]||['ARMAZENAMENTO E PROGRESSO','ocupado','livre','volumes acessíveis','omitidos','Analisado','Total elegível/restante: indeterminado'];
  z('zsLabel').textContent=st[0];
  const local=d.local_count>0&&d.local_total>0&&d.local_free<=d.local_total;
  z('zsCapacity').textContent=local?gib(d.local_total):'—';
  z('zsCapacity').title=local?String(d.local_total)+' bytes · GetDiskFreeSpaceExW · '+String(d.local_count)+' fixed lettered volumes':'';
  z('zsSpace').textContent=local?gib(d.local_total-d.local_free)+' '+st[1]+' / '+gib(d.local_free)+' '+st[2]+' · '+d.local_count+' '+st[3]+' · '+d.local_omitted+' '+st[4]:t[11];
  z('zsProgress').textContent=st[5]+': '+z('bytes').textContent;
  if(window.z310Storage)window.z310Storage(local?d.local_total-d.local_free:null);
  z('zpStatus').textContent=ram&&valid(d.cpu)?t[9]+' · '+new Date(d.at).toLocaleTimeString(lang):t[10];
  if(!confirmed&&ram&&valid(d.cpu)&&z('zp-cpu-value').textContent.endsWith('%')&&z('zp-ram-value').textContent.endsWith('%')&&z('zsCapacity').textContent!=='—'&&document.getElementById('zHelpOpen')&&document.getElementById('zAIActivity')) confirmed=await window.zPerformanceRendered();
 }catch(_){['cpu','ram','free','disk'].forEach(k=>gauge(k,null,t[11]));z('zpStatus').textContent=t[11];z('zpClock').textContent='—';z('zsCapacity').textContent='—';z('zsCapacity').title='';z('zsSpace').textContent=t[11];if(window.z310Storage)window.z310Storage(null);}
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
for k in ['zPerformancePanel','zpTitle','zpClock','zpRevision','zStorageCard','zsCapacity','zsProgress','start','quick','stop','export']+[f'zp-{k}-{x}' for k in ('cpu','ram','free','disk') for x in ('label','arc','value','detail')]:
    if c.ids.count(k)!=1:raise ValueError("Bad panel/control "+k)
p.write_text(s,encoding="utf-8")
print("3.0.10 native performance panel and controls verified")
