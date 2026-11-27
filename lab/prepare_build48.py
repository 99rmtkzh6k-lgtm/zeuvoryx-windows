"""Build 48 only. Apply to generated desktop after the preserved Build 47 generators."""
from pathlib import Path
from html import escape
from html.parser import HTMLParser
import json
import re
import sys

ROOT = Path(__file__).resolve().parent.parent
sys.path.insert(0, str(ROOT))
from retro21 import load_retro

DIR = ROOT / "lab" / "build48"
CATALOG = [
    {"id":"snake","category":"Retro","title":"Snake","description":"Coleção retro licenciada MIT."},
    {"id":"blocks","category":"Puzzle","title":"Tetris-like","description":"Encaixe blocos e complete linhas. Implementação original."},
    {"id":"maze","category":"Retro","title":"Pacman-like","description":"Colete pontos no labirinto. Implementação original."},
    {"id":"guerra","category":"Ação","title":"Guerra","description":"Família Botelho vs Horda Zumbi; estratégia fictícia."},
    {"id":"robozinho","category":"Ação","title":"Robozinho","description":"Recicle, plante e restaure o planeta."},
]
IDS = ["2048"] + [item["id"] for item in CATALOG]

def one(s, a, b):
    n = s.count(a)
    if n != 1:
        raise RuntimeError("Build48 anchor count="+str(n)+": "+a[:150])
    return s.replace(a,b,1)

def regex_one(s, a, b):
    out,n=re.subn(a,lambda m:b(m) if callable(b) else b,s)
    if n!=1:
        raise RuntimeError("Build48 regex count="+str(n)+": "+a[:150])
    return out

def get(name):
    return (DIR/name).read_text(encoding="utf-8")

def js_json(value):
    return json.dumps(value,ensure_ascii=True).replace("<",r"\u003c")

def replace_json(s, name, value=None, transform=None):
    marker="const "+name+"="
    if s.count(marker)!=1:
        raise RuntimeError("Build48 JSON anchor: "+name)
    start=s.index(marker)+len(marker)
    old,end=json.JSONDecoder().raw_decode(s[start:])
    if transform is not None:
        value=transform(old)
    return s[:start]+js_json(value)+s[start+end:]

LIFECYCLE = """<script>
;(()=>{const intervals=new Set(),timeouts=new Set(),frames=new Set(),audio=new Set();
const si=window.setInterval.bind(window),ci=window.clearInterval.bind(window),
st=window.setTimeout.bind(window),ct=window.clearTimeout.bind(window),
ra=window.requestAnimationFrame.bind(window),ca=window.cancelAnimationFrame.bind(window);
window.setInterval=(fn,ms,...args)=>{let id=si(fn,ms,...args);intervals.add(id);return id};
window.clearInterval=id=>{intervals.delete(id);timeouts.delete(id);ci(id)};
window.setTimeout=(fn,ms,...args)=>{let id;const f=typeof fn==='function'?function(...a){timeouts.delete(id);return fn.apply(this,a)}:fn;id=st(f,ms,...args);timeouts.add(id);return id};
window.clearTimeout=id=>{timeouts.delete(id);intervals.delete(id);ct(id)};
window.requestAnimationFrame=fn=>{let id=ra(t=>{frames.delete(id);fn(t)});frames.add(id);return id};
window.cancelAnimationFrame=id=>{frames.delete(id);ca(id)};
for(const name of ['AudioContext','webkitAudioContext']){const Original=window[name];if(Original){try{window[name]=class extends Original{constructor(...args){super(...args);audio.add(this)}}}catch(_){}}}
addEventListener('pagehide',()=>{for(const id of intervals)ci(id);for(const id of timeouts)ct(id);for(const id of frames)ca(id);for(const ctx of audio)try{ctx.close().catch(()=>{})}catch(_){};try{if(window.speechSynthesis)window.speechSynthesis.cancel()}catch(_){}},{once:true});
})();
</script>"""

def lifecycle(doc):
    # Keep CSP ahead of every inline script; insert after actual opening head and CSP.
    pat=r'(?is)(<meta\b[^>]*http-equiv=["\']Content-Security-Policy["\'][^>]*>)'
    return regex_one(doc,pat,lambda m:m[0]+LIFECYCLE)

def family(mode):
    doc=get("universo/index.html")
    script=get("universo/script.js")
    css=get("universo/style.css")
    if "</script" in script.lower() or "</style" in css.lower():
        raise RuntimeError("Unexpected inline document boundary")
    script=one(script,"function abrirJogo(n){if(n!==1&&n!==2)return;",
               "function abrirJogo(n){if(n!==z48Mode)return;")
    script=one(script,'function carregarJogo(codigo){if(typeof codigo==="string"&&codigo.trim()){try{let o=JSON.parse(codigo);',
               'function carregarJogo(codigo){if(typeof codigo==="string"&&codigo.trim()){try{let o=JSON.parse(codigo);if((o.game===2?2:1)!==z48Mode)throw Error("Jogo fora do modo autorizado");')
    script=one(script,"function iniciarJogo2(){jogo2=novo2();","function iniciarJogo2(){if(z48Mode!==2)return;jogo2=novo2();")
    other=2 if mode==1 else 1
    doc=regex_one(doc,r'<button onclick="abrirJogo\('+str(other)+r'\)">[\s\S]*?</button>',"")
    doc=one(doc,'<link rel="stylesheet" href="style.css">',"<style>"+css+"</style>")
    doc=one(doc,'<script src="script.js"></script>',"<script>const z48Mode="+str(mode)+";\n"+script+"\nabrirJogo(z48Mode);</script>")
    return lifecycle(doc)

def docs():
    _,old=load_retro()
    return {"snake":lifecycle(old["snake"]),"blocks":lifecycle(get("blocks.html")),
            "maze":lifecycle(get("maze.html")),"guerra":family(1),"robozinho":family(2)}

def patch_html(s):
    # Use actual scanner counts, never fabricate a heartbeat or declare completion on request.
    s=one(s,"refresh();setInterval(refresh,2000);","refresh();setInterval(refresh,500);")
    s=one(s,"let last=null,series=[],key='',busy=false,confirmResolve=null;",
          "let last=null,series=[],key='',busy=false,confirmResolve=null,z48CancelAt=0,z48LastScan='';")
    s=one(s,"function volume(n){","function volume(n){if(n===null||n===undefined||n===''||!Number.isFinite(Number(n))||Number(n)<0)return zT('Volume indisponível');n=Number(n);")
    s=one(s,"let running=s.running;zSummary(s);","let running=s.running;if(running&&r.scan_id&&z48LastScan&&r.scan_id!==z48LastScan)z48CancelAt=0;if(r.scan_id)z48LastScan=r.scan_id;zSummary(s);")
    s=one(s,"chart(s.files,running);","chart(s.files,running&&!r.inventory_314?.scanner_done);")
    s=one(s,"measureSpeed(s.bytes,running);","measureSpeed(s.bytes,running&&!r.inventory_314?.scanner_done);")
    s=one(s,"el('activity').textContent=series.length?",
          "el('activity').textContent=!running?zT('Motor parado'):series.length?")
    s=one(s,"el('speed').textContent=zT('Velocidade: varredura parada.')",
          "el('speed').textContent=zT('Motor parado')")
    s=one(s,"el('stateTitle').textContent=running?zT('Varredura em andamento'):s.error?",
          "el('stateTitle').textContent=running?(z48CancelAt?zT('Cancelando varredura'):zT('Varredura em andamento')):r.cancelled&&z48CancelAt?zT('Varredura cancelada por você'):s.error?")
    s=one(s,"el('stateText').textContent=s.error||",
          "el('stateText').textContent=running&&z48CancelAt?(Date.now()-z48CancelAt>2000?zT('Cancelamento solicitado; motor ainda encerrando. Limite de 2 segundos excedido.'):zT('Aguardando confirmação de parada do motor.')):s.error||")
    s=one(s,"el('stop').disabled=!running;","el('stop').disabled=!running||!!z48CancelAt;")
    s=one(s,"if(await window.startQuickScan())resetChart();","if(await window.startQuickScan()){z48CancelAt=0;resetChart();}")
    s=one(s,"if(await window.startScan())resetChart();","if(await window.startScan()){z48CancelAt=0;resetChart();}")
    s=one(s,"await window.stopScan();await refresh()","z48CancelAt=Date.now();try{if(await window.stopScan()!==true)z48CancelAt=0;}catch(e){z48CancelAt=0;throw e}await refresh()")
    s=one(s,"const gib=n=>(Number(n)/1073741824).toLocaleString(lang,{maximumFractionDigits:1})+' '+t[12];",
          "const gib=n=>n===null||n===undefined||n===''||!Number.isFinite(Number(n))||Number(n)<0?zT('Volume indisponível'):(Number(n)/1073741824).toLocaleString(lang,{maximumFractionDigits:1})+' '+t[12];")
    s=one(s,"z('zsCapacity').textContent=local?gib(d.local_total):'—';","z('zsCapacity').textContent=local?gib(d.local_total):zT('Volume indisponível');")
    s=one(s,'<button id="z310Startup">Programas de inicialização</button>',"")
    s=one(s,"'z310MaintTitle','z310MaintInfo','z310Drive','z310Clean','z310Startup'",
          "'z310MaintTitle','z310MaintInfo','z310Drive','z310Clean'")
    s=one(s,"[['z310Drive','drive'],['z310Clean','cleanup'],['z310Startup','startup']]",
          "[['z310Drive','drive'],['z310Clean','cleanup']]")
    s=s.replace("Análise local de arquivos. Esta versão não bloqueia a execução de malware nem substitui o Microsoft Defender.",
                "Scanner complementar. Não substitui o Microsoft Defender.")
    s=s.replace("Arquivos examinados por intervalo de atualização de aproximadamente 2 segundos.",
                "Arquivos realmente examinados pelo scanner próprio; atualização a cada 500 ms.")
    s=s.replace("Verificar procedência</button>","Verificar procedência C2PA</button>")
    s=s.replace("A falta de credenciais é inconclusiva.","A falta de credenciais é inconclusiva. Não detecta todos os vídeos de IA.")
    s=one(s,'<h2 id="zAITitle">Inteligência artificial — atividade real</h2>',
          '<h2 id="zAITitle">Inteligência artificial — atividade real</h2><p class="sub">IA auxiliar, não é proteção principal.</p>')
    s=one(s,'<div id="zAIOverview" class="sub" role="status">IA: inicializando modelo</div>',
          '<div id="zAIOverview" class="sub" role="status">IA: inicializando modelo</div><p class="sub">IA auxiliar, não é proteção principal.</p>')
    # Match generated attribute order, not a reserialized DOM.
    translation={
        "en-US":{"Motor parado":"Scanner stopped","Cancelando varredura":"Canceling scan","Varredura cancelada por você":"Scan canceled by you","Cancelamento solicitado; motor ainda encerrando. Limite de 2 segundos excedido.":"Cancellation requested; scanner still shutting down. Two-second limit exceeded.","Aguardando confirmação de parada do motor.":"Waiting for scanner shutdown confirmation.","IA auxiliar, não é proteção principal.":"Auxiliary AI, not primary protection.","Verificar procedência C2PA":"Verify C2PA provenance"},
        "es-ES":{"Motor parado":"Motor detenido","Cancelando varredura":"Cancelando análisis","Varredura cancelada por você":"Análisis cancelado por usted","Cancelamento solicitado; motor ainda encerrando. Limite de 2 segundos excedido.":"Cancelación solicitada; motor aún cerrándose. Límite de dos segundos superado.","Aguardando confirmação de parada do motor.":"Esperando confirmación de detención del motor.","IA auxiliar, não é proteção principal.":"IA auxiliar, no es protección principal.","Verificar procedência C2PA":"Verificar procedencia C2PA"}}
    for lang,values in translation.items():
        values["Volume indisponível"]="Volume unavailable" if lang=="en-US" else "Volumen no disponible"
        values["Scanner complementar. Não substitui o Microsoft Defender."]="Complementary scanner. Does not replace Microsoft Defender." if lang=="en-US" else "Escáner complementario. No sustituye a Microsoft Defender."
        values["Arquivos realmente examinados pelo scanner próprio; atualização a cada 500 ms."]="Actual files examined by the local scanner; refresh every 500 ms." if lang=="en-US" else "Archivos realmente examinados por el escáner propio; actualización cada 500 ms."
    s=one(s,"zCaptureStatic();zPaintStatic();","for(const [lang,v] of Object.entries("+js_json(translation)+"))Object.assign(zCatalog[lang],v);zCaptureStatic();zPaintStatic();")
    s=one(s,"Fechar partida</button>","Fechar jogo</button>")
    s=one(s,"w[0]+': '+(!s.running?w[3]:", "w[0]+': '+(!s.running||s.report?.inventory_314?.scanner_done?zT('Motor parado'):")
    s=one(s,"Proteção durante os jogos</strong>","Estado da varredura durante os jogos</strong>")
    s=regex_one(s,r'(<select id="zGameFilter">)[\s\S]*?(</select>)',lambda m:m[1]+'<option value="">Todos</option><option>Puzzle</option><option>Retro</option><option>Ação</option>'+m[2])
    opts='<option value="2048">2048</option>'+''.join('<option value="'+escape(g["id"],quote=True)+'">'+escape(g["title"])+'</option>' for g in CATALOG)
    s=regex_one(s,r'(<select id="zFamilyGame">)[\s\S]*?(</select>)',lambda m:m[1]+opts+m[2])
    s=replace_json(s,"zArcade2048",transform=lifecycle)
    s=replace_json(s,"zRetroCatalog",CATALOG)
    s=replace_json(s,"zRetroDocs",docs())
    old=(ROOT/"lab/games314.js").read_text(encoding="utf-8")
    s=one(s,old,get("launcher.js"))
    if len(s.encode("utf-8"))>1500000:
        raise RuntimeError("Build48 desktop budget exceeded")
    return s

def patch_binding(s):
    return regex_one(s,r'func localGame315\(id string\)bool\{return [^}]+\}',
                     'func localGame315(id string)bool{return '+'||'.join('id=="'+i+'"' for i in IDS)+'}')

def main():
    html=ROOT/"desktop.html"
    native=ROOT/"games314_windows.go"
    result=patch_html(html.read_text(encoding="utf-8"))
    binding=patch_binding(native.read_text(encoding="utf-8"))
    html.write_text(result,encoding="utf-8")
    native.write_text(binding,encoding="utf-8")
    print("Build48: six offline catalog entries, exact native allowlist, scanner UI and maintenance text applied")

if __name__=="__main__":
    main()
