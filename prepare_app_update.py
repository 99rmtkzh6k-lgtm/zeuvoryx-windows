"""3.0.4 LAB: release discovery only; no untrusted package execution."""
from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise RuntimeError("Unexpected source: "+a)
    return s.replace(a,b,1)
m=Path("main.go"); ms=m.read_text(encoding="utf-8")
ms=once(ms,'func main() {','var updateInstalledVersion = flag.String("updated-version","","Installer completion version")\n\nfunc main() {')
m.write_text(ms,encoding="utf-8")
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=once(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/updater"')
needle='if !bind("getScanState", func() map[string]any {'
s=once(s,needle,'''if !bind("checkLabRelease",func()(updater.Status,error){
    return updater.Check(context.Background(),runtime.GOOS,runtime.GOARCH,true)
}) { return }
if !bind("openLabReleases",func()error{
    s.mu.Lock(); running:=s.running; s.mu.Unlock()
    if running{return fmt.Errorf("Aguarde a conclusão da varredura")}
    status,e:=updater.Check(context.Background(),runtime.GOOS,runtime.GOARCH,true);if e!=nil{return e}
    if !status.Available||status.DownloadURL==""{return fmt.Errorf("Nenhum pacote LAB mais recente encontrado")}
    return exec.Command("rundll32.exe","url.dll,FileProtocolHandler",status.DownloadURL).Start()
}) { return }
if !bind("signedUpdateEnabled", updater.SignedEnabled) { return }
if !bind("completedUpdateVersion",func()string{return updater.CompletedVersion(*updateInstalledVersion)}) { return }
if !bind("installSignedUpdate",func()error{
    operationGate.Lock();defer operationGate.Unlock()
    s.mu.Lock(); running:=s.running; s.mu.Unlock()
    if running||native.status().Running{return fmt.Errorf("Aguarde as operações em andamento")}
    status,e:=updater.Check(nativeCtx,runtime.GOOS,runtime.GOARCH,true);if e!=nil{return e}
    if !status.Available{return fmt.Errorf("Nenhuma atualização mais recente")}
    if e=updater.StartVerifiedInstall(nativeCtx,status.Latest);e!=nil{return e}
    monitorEnabled.Store(false)
    w.Dispatch(func(){w.Terminate()})
    return nil
}) { return }
'''+needle)
s=once(s,'ZEUVORYX | 3.0.3 LAB','ZEUVORYX | 3.0.4 LAB')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'3.0.3 LAB BETA</div>','3.0.4 LAB BETA</div>')
s=once(s,'<section id="overview">','''<section id="overview"><div class="card" style="margin-bottom:10px"><h2 id="zUpdateTitle">Atualizações LAB — 3.0.4</h2><p id="zUpdateState" class="sub" role="status"></p><p id="zUpdateCompleted" class="sub" role="status" hidden></p><div class="actions"><button id="zUpdateCheck">Verificar atualizações</button><button id="zUpdateOpen" hidden>Baixar atualização LAB</button></div><p id="zUpdateNote" class="sub"></p></div>''')
s=once(s,'</script></body></html>','''
const zUpdateCopy={
'pt-BR':['Atualizações LAB — 3.0.4','Verificar atualizações','Baixar atualização LAB','Download pelo GitHub; instalação manual em Sandbox. Confira a origem do pacote.','Consultando...','Nova versão LAB: ','Nenhuma atualização LAB encontrada.','Falha de consulta. Tente novamente.'],
'en-US':['LAB updates — 3.0.4','Check for updates','Download LAB update','GitHub download; manual Sandbox installation. Verify package origin.','Checking...','New LAB version: ','No LAB update found.','Check failed. Try again.'],
'es-ES':['Actualizaciones LAB — 3.0.4','Buscar actualizaciones','Descargar actualización LAB','Descarga de GitHub; instalación manual en Sandbox. Verifique el origen.','Consultando...','Nueva versión LAB: ','Sin actualizaciones LAB.','Error de consulta. Inténtelo de nuevo.']
};
let zUpdateBusy=false,zUpdateResult=null,zUpdateError=false,zSigned=false,zCompleted="";
function zUpdatePaint(){
 const t=zUpdateCopy[document.documentElement.lang]||zUpdateCopy['pt-BR'];
 const completion=el('zUpdateCompleted');completion.hidden=!zCompleted;completion.textContent=zCompleted?(({'pt-BR':'Atualização concluída. Versão em execução: ','en-US':'Update completed. Running version: ','es-ES':'Actualización completada. Versión en ejecución: '}[document.documentElement.lang]||'Atualização concluída. Versão em execução: ')+zCompleted):'';
 el('zUpdateTitle').textContent=t[0];el('zUpdateCheck').textContent=t[1];el('zUpdateOpen').textContent=zSigned?({'pt-BR':'Atualizar agora','en-US':'Update now','es-ES':'Actualizar ahora'}[document.documentElement.lang]||'Atualizar agora'):t[2];el('zUpdateNote').textContent=zSigned?({'pt-BR':'Instalação assinada sobre a versão atual; finalize as varreduras.','en-US':'Signed in-place update; finish scans first.','es-ES':'Actualización firmada; finalice los análisis.'}[document.documentElement.lang]||t[3]):t[3];
 el('zUpdateState').textContent=zUpdateBusy?t[4]:zUpdateError?t[7]:zUpdateResult?(zUpdateResult.available?t[5]+zUpdateResult.latest:t[6]):'';
 el('zUpdateOpen').hidden=!(zUpdateResult&&zUpdateResult.available&&!zUpdateError);el('zUpdateCheck').disabled=zUpdateBusy;
}
async function zUpdateRun(){
 if(zUpdateBusy)return;zUpdateBusy=true;zUpdateError=false;zUpdateResult=null;zUpdatePaint();
 try{zSigned=await window.signedUpdateEnabled();zUpdateResult=await window.checkLabRelease()}catch(e){zUpdateError=true}finally{zUpdateBusy=false;zUpdatePaint()}
}
el('zUpdateCheck').onclick=zUpdateRun;el('zUpdateOpen').onclick=()=>action(async()=>{if(zUpdateBusy)return;zUpdateBusy=true;el('zUpdateOpen').disabled=true;try{if(zSigned){const lang=document.documentElement.lang;const text={'pt-BR':'Instalar a atualização assinada e fechar o ZEUVORYX?','en-US':'Install the signed update and close ZEUVORYX?','es-ES':'¿Instalar la actualización firmada y cerrar ZEUVORYX?'};if(await ask(text[lang]||text['pt-BR']))await window.installSignedUpdate()}else await window.openLabReleases()}finally{zUpdateBusy=false;el('zUpdateOpen').disabled=false;zUpdatePaint()}});
new MutationObserver(zUpdatePaint).observe(document.documentElement,{attributes:true,attributeFilter:['lang']});zUpdatePaint();window.completedUpdateVersion().then(v=>{zCompleted=v;zUpdatePaint()}).catch(()=>{});setTimeout(zUpdateRun,1500);
</script></body></html>''')
p.write_text(s,encoding="utf-8")
