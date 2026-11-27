"""Validate final Build48 desktop and prepare a browser harness of its real launcher/assets."""
from html.parser import HTMLParser
from pathlib import Path
import json
import re
from prepare_build48 import IDS, js_json

s=Path("desktop.html").read_text(encoding="utf-8")
def extract(name):
    key="const "+name+"="
    assert s.count(key)==1,name
    return json.JSONDecoder().raw_decode(s[s.index(key)+len(key):])[0]
catalog=extract("zRetroCatalog")
docs=extract("zRetroDocs")
arcade=extract("zArcade2048")
assert ["2048"]+[g["id"] for g in catalog]==IDS
assert set(docs)==set(IDS[1:])
assert len(IDS)==6
class Check(HTMLParser):
    def __init__(self):
        super().__init__()
        self.ids=[]
    def handle_starttag(self,tag,attrs):
        a=dict(attrs)
        if "id" in a:self.ids.append(a["id"])
        assert tag not in {"iframe","object","embed","base"},tag
        for attr in ["src","href","srcset","action"]:
            assert not a.get(attr,"").lower().startswith(("http:","https:","//")),(tag,attr)
for id,doc in {"2048":arcade,**docs}.items():
    p=Check();p.feed(doc);p.close()
    assert "connect-src 'none'" in doc and "default-src 'none'" in doc,id
    assert "pagehide" in doc and "ctx.close()" in doc,id
    assert "allow-same-origin" not in doc and "allow-top-navigation" not in doc and "allow-popups" not in doc,id
    assert not re.search(r'<script\b[^>]*\bsrc\s*=|<link\b[^>]*\bhref\s*=',doc,re.I),id
assert "function comprar(){comprarArma();}" in docs["guerra"]
assert "const z48Mode=1;" in docs["guerra"] and "const z48Mode=2;" in docs["robozinho"]
assert '<button onclick="abrirJogo(2)">' not in docs["guerra"]
assert '<button onclick="abrirJogo(1)">' not in docs["robozinho"]
assert 'id="z310Startup"' not in s
assert "refresh();setInterval(refresh,500);" in s
assert "Varredura cancelada por você" in s and "Motor parado" in s
assert "chart(s.files,running&&!r.inventory_314?.scanner_done);" in s
assert "measureSpeed(s.bytes,running&&!r.inventory_314?.scanner_done);" in s
assert "Date.now()-z48CancelAt>2000" in s
assert "Volume indisponível" in s and "Volume unavailable" in s and "Volumen no disponible" in s
assert "local?gib(d.local_total):zT('Volume indisponível')" in s
for label in ["Scanner complementar. Não substitui o Microsoft Defender.","Verificar procedência C2PA","Não detecta todos os vídeos de IA.","IA auxiliar, não é proteção principal."]:
    assert label in s,label
launcher=Path("lab/build48/launcher.js").read_text(encoding="utf-8")
assert launcher in s
assert "frame.setAttribute('sandbox','allow-scripts')" in launcher
assert "frame.setAttribute('referrerpolicy','no-referrer')" in launcher
assert not any(x in launcher for x in ["allow-same-origin","allow-top-navigation","allow-popups","window.open"])
binding=Path("games314_windows.go").read_text(encoding="utf-8")
m=re.search(r'func localGame315\(id string\)\s*bool\s*\{([^}]+)\}',binding)
assert m
assert re.findall(r'id\s*==\s*"([^"]+)"',m[1])==IDS
for fn in ["familyRate314","familySave314","familyStatus314"]:
    assert fn in s and fn in Path("desktop_windows.go").read_text(encoding="utf-8")
section=re.search(r'<section id="zGames314"[\s\S]*?</section>',s)
assert section
prefix="""<!doctype html><html lang="pt-BR"><head><meta charset="utf-8"><style>.hidden,[hidden]{display:none!important}</style></head><body>
<button id="zGamesNav314">Games</button><button id="zFamilyNav314">Controle parental</button><button data-page="overview">Visão geral</button><button data-page="alerts">Alertas</button><section id="overview"></section><section id="alerts"></section>"""
bridge="""<script>
window.blockedIds=new Set();window.getScanState=async()=>({running:false,files:0,bytes:0,report:{}});
window.allowLocalGame315=async id=>!window.blockedIds.has(id);
function node(tag,text){let e=document.createElement(tag);e.textContent=text;return e}
function volume(n){return n+' B'}
"""
payload="const zArcade2048="+js_json(arcade)+";const zRetroCatalog="+js_json(catalog)+";const zRetroDocs="+js_json(docs)+";"
Path("build48-browser-harness.html").write_text(prefix+section[0]+bridge+payload+launcher+"</script></body></html>",encoding="utf-8")
print("Build48: exact catalog, generated bindings, UI text and isolation contract verified. Browser harness uses mock host authorization only.")
