"""Advisory-only network database. Existing malware/hash and quarantine policies remain unchanged."""
from pathlib import Path
def one(s,a,b):
    if s.count(a)!=1: raise RuntimeError("Global intelligence anchor: "+a)
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=one(s,"import (",'import (\n "github.com/zeuvoryx/zeuvoryx/globalintel"')
needle='if !bind("getScanState", func() map[string]any {'
s=one(s,needle,'''dropBase:=globalintel.Bundled(time.Now())
var dropLastAttempt time.Time
if t:=dropBase.Status().Updated;t!=""{dropLastAttempt,_=time.Parse(time.RFC3339,t)}
if !bind("zDropRefresh",func()(globalintel.Status,error){
    operationGate.Lock();defer operationGate.Unlock()
    s.mu.Lock();running:=s.running;s.mu.Unlock()
    if running||native.status().Running{return dropBase.Status(),fmt.Errorf("Aguarde o fim das operações")}
    if time.Since(dropLastAttempt)<24*time.Hour{return dropBase.Status(),fmt.Errorf("Intervalo mínimo: 24 h")}
    dropLastAttempt=time.Now()
    next,e:=globalintel.Refresh(nativeCtx);if e!=nil{return dropBase.Status(),e}
    dropBase=next;return dropBase.Status(),nil
}){return}
if !bind("zDropStatus",func()globalintel.Status{operationGate.Lock();defer operationGate.Unlock();return dropBase.Status()}){return}
if !bind("zDropMatch",func(ip string)(bool,error){
    operationGate.Lock();defer operationGate.Unlock()
    return dropBase.Match(ip,time.Now())
}){return}
'''+needle)
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=one(s,'<section id="intel" class="hidden card">','''<section id="intel" class="hidden card"><div class="entry"><h2>The Spamhaus Project — DROP</h2><p id="zDropHelp" class="sub"></p><button id="zDropRefresh">Atualizar DROP</button><p id="zDropStatus" role="status"></p><input id="zDropIP" class="media-field" maxlength="80" placeholder="IPv4 / IPv6"><button id="zDropMatch">Consultar IP</button><p id="zDropResult" role="status"></p></div>''')
s=one(s,'</script></body></html>','''
;(()=>{const d=id=>document.getElementById(id);const copy={
'pt-BR':['Base consultiva em memória, atualização manual a cada 24 h no máximo. Não bloqueia conexões.','Atualizar DROP','Consultar IP','Fonte indisponível ou desatualizada','Correspondência na lista DROP — revisar conexão','Sem correspondência; não comprova segurança','Consultando...'],
'en-US':['In-memory advisory feed, manual refresh at most every 24 h. Does not block connections.','Refresh DROP','Check IP','Feed unavailable or stale','DROP match — review connection','No match; does not prove safety','Checking...'],
'es-ES':['Base consultiva en memoria, actualización manual como máximo cada 24 h. No bloquea conexiones.','Actualizar DROP','Consultar IP','Fuente no disponible o desactualizada','Coincidencia DROP — revisar conexión','Sin coincidencia; no demuestra seguridad','Consultando...']};
const lang=()=>copy[document.documentElement.lang]||copy['pt-BR'];
function labels(){const t=lang();d('zDropHelp').textContent=t[0];d('zDropRefresh').textContent=t[1];d('zDropMatch').textContent=t[2]}
function show(r){d('zDropStatus').textContent=r.updated?'IPv4: '+r.ipv4+' | IPv6: '+r.ipv6+' | '+r.updated+' | '+r.source+' | '+r.copyright4+' | '+r.copyright6:lang()[3]}
labels();window.zDropStatus().then(show).catch(()=>{});new MutationObserver(labels).observe(document.documentElement,{attributes:true,attributeFilter:['lang']});
d('zDropRefresh').onclick=async()=>{const b=d('zDropRefresh');b.disabled=true;d('zDropStatus').textContent=lang()[6];try{const r=await window.zDropRefresh();show(r)}catch(_){d('zDropStatus').textContent=lang()[3]}finally{b.disabled=false}};
d('zDropMatch').onclick=async()=>{try{const yes=await window.zDropMatch(d('zDropIP').value);d('zDropResult').textContent=lang()[yes?4:5]}catch(_){d('zDropResult').textContent=lang()[3]}};
})();
</script></body></html>''')
p.write_text(s,encoding="utf-8")
