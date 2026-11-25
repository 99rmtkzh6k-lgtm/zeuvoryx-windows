from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1: raise RuntimeError("signature integration anchor: "+a[:100])
    return s.replace(a,b,1)
p=Path("clam_windows.go")
p.write_text(Path("lab/clam311_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
p=Path("desktop_windows.go")
s=p.read_text(encoding="utf-8")
import re
s,n=re.subn(r'if !bind\("clamStatus",\s*func\(\)string\{[^\n]*\}\)\s*\{return\}', 'if !bind("clamStatus",clamStatus311){return}',s)
# Previous gofmt can expand the bind: replace only by structural anchors if necessary.
if n!=1:
    a=s.index('if !bind("clamStatus"')
    b=s.index('if !bind("chooseClamFile"',a)
    s=s[:a]+'if !bind("clamStatus",clamStatus311){return}\n'+s[b:]
a=s.index('if !bind("chooseClamFile"')
s=s[:a]+'if !bind("chooseClamFolder311",chooseClamFolder311){return}\nif !bind("cancelClam311",cancelClam311){return}\n'+s[a:]
p.write_text(s,encoding="utf-8")
p=Path("desktop.html")
s=p.read_text(encoding="utf-8")
s=once(s,'id="clamChoose">Examinar arquivo com ClamAV</button>','id="clamChoose">Examinar arquivo com ClamAV</button><button id="clamFolder311">Examinar pasta ou unidade com ClamAV</button><button id="clamCancel311">Cancelar ClamAV</button>')
s=s.replace("Examina um arquivo por vez; não remove automaticamente e não é proteção residente.","Modo independente por arquivo, pasta ou unidade; os botões de varredura ZEUVORYX continuam heurísticos. JSON local com nomes das assinaturas, cobertura e versão da base. Não remove automaticamente e não é proteção residente.")
anchor="el('clamChoose').onclick="
a=s.index(anchor)
s=s[:a]+'''el('clamFolder311').onclick=()=>action(async()=>{const b=el('clamFolder311');b.disabled=true;try{const p=await window.chooseClamFolder311();if(!p)return;el('clamState').textContent='ClamAV examinando pasta/unidade; pode demorar...';const r=await window.clamScanFile(p);el('clamState').textContent=r.status;el('clamReport').textContent=r.details||''}catch(e){el('clamState').textContent='Operação não confirmada: '+String(e)}finally{b.disabled=false}});
el('clamCancel311').onclick=()=>action(async()=>{await window.cancelClam311()});
'''+s[a:]
p.write_text(s,encoding="utf-8")
# Change current identity only after the historical generator scripts have run.
for p in list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]:
    if ".git" in p.parts: continue
    s=p.read_text(encoding="utf-8")
    n=s.replace("3.0.10","3.0.11")
    if n!=s:p.write_text(n,encoding="utf-8")
print("3.0.11 signature mode integrated; official ClamAV installation still required")
