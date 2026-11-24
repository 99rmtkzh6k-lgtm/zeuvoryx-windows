"""Integrate local diagnostics only. No production keys, network or destructive operations."""
from pathlib import Path

def one(s,a,b):
    if s.count(a)!=1:
        raise RuntimeError("Expected exactly one anchor: "+a[:100])
    return s.replace(a,b,1)

p=Path("desktop_windows.go")
s=p.read_text(encoding="utf-8")
s=one(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/labcheck"')
s=one(s,'if !bind("getScanState", func() map[string]any {','if !bind("zLabCheck", labcheck.Run) { return }\n if !bind("getScanState", func() map[string]any {')
s=one(s,'ZEUVORYX | Olympus 3.0.3 Competitive','ZEUVORYX | 3.0.3 LAB')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html")
s=p.read_text(encoding="utf-8")
s=one(s,'OLYMPUS 3.0.2 BETA</div>','3.0.3 LAB BETA</div>')
s=one(s,'<section id="engine" class="hidden card">','''<section id="engine" class="hidden card"><div class="entry"><h2 id="zLabTitle">Diagnóstico local 3.0.3</h2><p id="zLabNote">Laboratório: sem servidor mundial, chave oficial, atualização automática ou remediação automática. Os testes abaixo usam dados fictícios e não alteram arquivos.</p><button id="zLabRun">Testar componentes</button><p id="zLabResult" role="status" aria-live="polite"></p></div>''')
s=one(s,'</script></body></html>','''
const zLabCopy={
 'pt-BR':['Diagnóstico local 3.0.3','Laboratório: sem servidor mundial, chave oficial, atualização automática ou remediação automática. Testes fictícios; não alteram arquivos.','Testar componentes','Atualizações assinadas','Estatísticas assinadas','Política de remediação','Aprovado','Falhou','Testando...'],
 'en-US':['Local diagnostics 3.0.3','Lab: no global server, owner key, automatic updates or automatic remediation. Inert tests; no files changed.','Test components','Signed updates','Signed statistics','Remediation policy','Passed','Failed','Testing...'],
 'es-ES':['Diagnóstico local 3.0.3','Laboratorio: sin servidor mundial, clave del propietario, actualización automática ni remediación automática. Pruebas inertes; sin modificar archivos.','Probar componentes','Actualizaciones firmadas','Estadísticas firmadas','Política de remediación','Aprobado','Falló','Probando...']
};
function zLabText(){return zLabCopy[document.documentElement.lang]||zLabCopy['pt-BR']}
let zLabLast=null;
function zLabPaint(){
 const t=zLabText();
 el('zLabTitle').textContent=t[0];el('zLabNote').textContent=t[1];el('zLabRun').textContent=t[2];
 if(zLabLast!==null)el('zLabResult').textContent=[zLabLast.update,zLabLast.statistics,zLabLast.policy].map((v,i)=>t[i+3]+': '+t[v===true?6:7]).join(' | ');
}
new MutationObserver(zLabPaint).observe(document.documentElement,{attributes:true,attributeFilter:['lang']});
el('zLabRun').onclick=async()=>{
 const b=el('zLabRun');b.disabled=true;zLabLast=null;el('zLabResult').textContent=zLabText()[8];
 try{zLabLast=await window.zLabCheck();zLabPaint()}catch(e){el('zLabResult').textContent=zLabText()[7]+': '+String(e)}finally{b.disabled=false}
};
zLabPaint();
</script></body></html>''')
p.write_text(s,encoding="utf-8")
print("3.0.3 local diagnostics integrated; production updates and telemetry disabled")
