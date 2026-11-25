"""Final visual overlay, preserving offline approved artwork and functional controls."""
from pathlib import Path
from html.parser import HTMLParser
import re, base64, hashlib
p=Path("desktop.html")
s=p.read_text(encoding="utf-8-sig")
m=re.search(r"header\{min-height:150px;[^}]*?background:(url\('data:image/jpeg;base64,[A-Za-z0-9+/=]+'\))",s)
if not m:raise RuntimeError("Titanium: existing Olympus artwork not found")
art=m.group(1)
approved_sha256="7f5e51ee96b08e133e12b799206b5688661c7864a193ee595d66ae16cfb3155c"
raw=base64.b64decode(art.split("base64,",1)[1][:-2],validate=True)
if hashlib.sha256(raw).hexdigest()!=approved_sha256:raise RuntimeError("Titanium: artwork differs from exact user-approved Zeus Athens image")
Path("approved-titanium-art.sha256").write_text(approved_sha256+"\n",encoding="ascii")
# Store the same artwork only once in CSS, reducing the embedded document size.
s=s.replace(art,"var(--z-titanium-art)")
css=Path("lab/titanium314.css").read_text(encoding="utf-8")
if s.count("</style></head>")!=1:raise RuntimeError("Titanium style anchor")
s=s.replace("</style></head>",":root{--z-titanium-art:"+art+"}\n"+css+"</style></head>",1)
if s.count("3.0.14 LAB BETA</div>")!=1:raise RuntimeError("Titanium version anchor")
s=s.replace("3.0.14 LAB BETA</div>","3.0.14 LAB BETA</div><div id=\"zTitaniumIdentity\">OLYMPUS TITANIUM · GLOBAL THREAT DEFENSE</div>",1)
s=s.replace("3.0.14 · Performance","3.0.14 · Olympus Titanium")
if s.count("<header>")!=1:raise RuntimeError("Titanium header anchor")
s=s.replace("<header>",'<div id="zTitaniumBanner" aria-label="Identidade Olympus Titanium sobre a arte oficial de Zeus"><div class="zTitaniumBrand"><strong>ZEUVORYX</strong><span>OLYMPUS TITANIUM</span><small>GLOBAL THREAT DEFENSE · 3.0.14 LAB</small></div></div><header>',1)
# Keep wallpaper visible while using the interface. Optional distraction-free art view.
if s.count("</script></body></html>")!=1:raise RuntimeError("Titanium script anchor")
s=s.replace("</script></body></html>", """
;(()=>{const brand=document.querySelector('.zTitaniumBrand');
const open=document.createElement('button');open.type='button';open.textContent='Ver arte inteira';open.id='zArtOpen';brand.appendChild(open);
const back=document.createElement('button');back.type='button';back.id='zArtRestore';back.textContent='Voltar aos controles';back.hidden=true;document.body.appendChild(back);
function setView(active){document.body.dataset.zArtFocus=active?'1':'0';back.hidden=!active;if(active)back.focus();else open.focus();}
open.onclick=()=>setView(true);back.onclick=()=>setView(false);
document.addEventListener('keydown',e=>{if(e.key==='Escape'&&document.body.dataset.zArtFocus==='1')setView(false)});
})();
</script></body></html>""",1)
if "background-size:contain,cover,cover" not in s:raise RuntimeError("Titanium full-window wallpaper missing")
class Check(HTMLParser):
    def __init__(self):super().__init__();self.ids=[]
    def handle_starttag(self,tag,attrs):
        a=dict(attrs)
        if "id" in a:self.ids.append(a["id"])
c=Check();c.feed(s)
for k in ["start","quick","stop","export","z314Diagnostic","z314LiveHits","z314ProgressBar","zp-cpu-needle","zTitaniumIdentity","zTitaniumBanner","zDropRefresh","z310Optimizer","z314Youtube"]:
    if c.ids.count(k)!=1:raise RuntimeError("Titanium missing/duplicated control: "+k)
if len(s.encode("utf-8"))>1500000:raise RuntimeError("Titanium page budget exceeded")
p.write_text(s,encoding="utf-8")
d=Path("desktop_windows.go");x=d.read_text(encoding="utf-8")
x=x.replace("ZEUVORYX | 3.0.14 LAB","ZEUVORYX | 3.0.14 LAB - Olympus Titanium")
d.write_text(x,encoding="utf-8")
i=Path("setup309.iss");x=i.read_text(encoding="utf-8")
x=x.replace("AppName=ZEUVORYX 3.0.14 LAB","AppName=ZEUVORYX 3.0.14 LAB - Olympus Titanium").replace("UninstallDisplayName=ZEUVORYX 3.0.14 LAB","UninstallDisplayName=ZEUVORYX 3.0.14 LAB - Olympus Titanium")
i.write_text(x,encoding="utf-8")
print("Olympus Titanium: reused local artwork, global transparent layout, deep buttons, control IDs and page size verified")
