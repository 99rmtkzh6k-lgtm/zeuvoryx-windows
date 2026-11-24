"""3.0.4 LAB: release discovery only; no untrusted package execution."""
from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise RuntimeError("Unexpected source: "+a)
    return s.replace(a,b,1)
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/updater"')
needle='if !bind("getScanState", func() map[string]any {'
s=once(s,needle,'''if !bind("checkLabRelease",func()(updater.Status,error){
    return updater.Check(context.Background(),runtime.GOOS,runtime.GOARCH,true)
}) { return }
if !bind("openLabReleases",func()error{
    return exec.Command("rundll32.exe","url.dll,FileProtocolHandler",updater.ReleasesURL).Start()
}) { return }
'''+needle)
s=once(s,'ZEUVORYX | 3.0.3 LAB','ZEUVORYX | 3.0.4 LAB')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'3.0.3 LAB BETA</div>','3.0.4 LAB BETA</div>')
s=once(s,'<section id="overview">','''<section id="overview"><div class="card" style="margin-bottom:10px"><h2 id="zUpdateTitle">Atualizações LAB — 3.0.4</h2><p id="zUpdateState" class="sub" role="status"></p><div class="actions"><button id="zUpdateCheck">Verificar atualizações</button><button id="zUpdateOpen" hidden>Ver versões no GitHub</button></div><p id="zUpdateNote" class="sub"></p></div>''')
s=once(s,'</script></body></html>','''
const zUpdateCopy={
'pt-BR':['Atualizações LAB — 3.0.4','Verificar atualizações','Ver versões no GitHub','Consulta apenas. Não instala automaticamente. Teste em Sandbox.','Consultando...','Nova versão LAB: ','Nenhuma atualização LAB encontrada.','Falha de consulta. Tente novamente.'],
'en-US':['LAB updates — 3.0.4','Check for updates','View GitHub releases','Discovery only. No automatic installation. Test in Sandbox.','Checking...','New LAB version: ','No LAB update found.','Check failed. Try again.'],
'es-ES':['Actualizaciones LAB — 3.0.4','Buscar actualizaciones','Ver versiones en GitHub','Solo consulta. Sin instalación automática. Pruebe en Sandbox.','Consultando...','Nueva versión LAB: ','Sin actualizaciones LAB.','Error de consulta. Inténtelo de nuevo.']
};
let zUpdateBusy=false,zUpdateResult=null,zUpdateError=false;
function zUpdatePaint(){
 const t=zUpdateCopy[document.documentElement.lang]||zUpdateCopy['pt-BR'];
 el('zUpdateTitle').textContent=t[0];el('zUpdateCheck').textContent=t[1];el('zUpdateOpen').textContent=t[2];el('zUpdateNote').textContent=t[3];
 el('zUpdateState').textContent=zUpdateBusy?t[4]:zUpdateError?t[7]:zUpdateResult?(zUpdateResult.available?t[5]+zUpdateResult.latest:t[6]):'';
 el('zUpdateOpen').hidden=!(zUpdateResult&&zUpdateResult.available&&!zUpdateError);el('zUpdateCheck').disabled=zUpdateBusy;
}
async function zUpdateRun(){
 if(zUpdateBusy)return;zUpdateBusy=true;zUpdateError=false;zUpdateResult=null;zUpdatePaint();
 try{zUpdateResult=await window.checkLabRelease()}catch(e){zUpdateError=true}finally{zUpdateBusy=false;zUpdatePaint()}
}
el('zUpdateCheck').onclick=zUpdateRun;el('zUpdateOpen').onclick=()=>action(async()=>{await window.openLabReleases()});
new MutationObserver(zUpdatePaint).observe(document.documentElement,{attributes:true,attributeFilter:['lang']});zUpdatePaint();setTimeout(zUpdateRun,1500);
</script></body></html>''')
p.write_text(s,encoding="utf-8")
