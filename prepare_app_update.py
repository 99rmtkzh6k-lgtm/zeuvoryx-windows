from pathlib import Path

def once(s,a,b):
    if s.count(a)!=1:
        raise RuntimeError("Unexpected source: "+a[:100])
    return s.replace(a,b,1)

p=Path("desktop_windows.go")
s=p.read_text(encoding="utf-8")
s=once(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/updater"')
needle='\tif !bind("getScanState", func() map[string]any {'
s=once(s,needle,'''    if !bind("checkAppRelease",func()(updater.Status,error){
        return updater.Check(context.Background(),runtime.GOOS,runtime.GOARCH)
    }) { return }
    if !bind("openAppReleases",func()error{
        return exec.Command("rundll32.exe","url.dll,FileProtocolHandler",updater.ReleasesURL).Start()
    }) { return }
'''+needle)
s=s.replace("Olympus 3.0.1 Competitive","Olympus 3.0.2 Competitive")
p.write_text(s,encoding="utf-8")
p=Path("desktop.html")
s=p.read_text(encoding="utf-8")
s=once(s,'<section id="overview">','''<div class="card" style="margin-bottom:10px"><strong>Atualizações do aplicativo — 3.0.2</strong><p id="appReleaseState" class="sub" role="status">Verificação pendente.</p><div class="actions" style="margin-top:8px"><button id="appReleaseCheck">Verificar atualizações</button><button id="appReleaseOpen" hidden>Ver instalador no GitHub</button></div><p class="sub">Instale por cima da versão existente. Não desinstale. A instalação automática ainda não está habilitada.</p></div><section id="overview">''')
s=once(s,'</script></body></html>', ''' 
async function checkAppReleaseNow(){
 const b=el('appReleaseCheck'),o=el('appReleaseOpen'),s=el('appReleaseState');
 if(b.disabled)return;b.disabled=true;o.hidden=true;s.textContent='Consultando GitHub…';
 try {const r=await window.checkAppRelease();s.textContent=r.available?'Nova versão disponível: '+r.latest:r.found?'Nenhuma versão mais recente disponível.':'Nenhuma versão estável compatível publicada.';o.hidden=!r.available}
 catch(e){s.textContent='Não foi possível verificar agora. Tente novamente.'}
 finally{b.disabled=false}
}
el('appReleaseCheck').onclick=checkAppReleaseNow;
el('appReleaseOpen').onclick=()=>action(async()=>{await window.openAppReleases()});
setTimeout(checkAppReleaseNow,1500);
</script></body></html>''')
p.write_text(s,encoding="utf-8")
