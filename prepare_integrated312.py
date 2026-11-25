from pathlib import Path
import re
def once(s,a,b):
    if s.count(a)!=1: raise RuntimeError("312 anchor "+a[:90]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
def rx(s,a,b):
    s,n=re.subn(a,lambda m:b(m) if callable(b) else b,s)
    if n!=1: raise RuntimeError("312 regex "+a+" count="+str(n))
    return s
def edit(path,fn):
    p=Path(path);p.write_text(fn(p.read_text(encoding="utf-8-sig")),encoding="utf-8")
def report(s):
    s=once(s,"import (","import (\n \"github.com/zeuvoryx/zeuvoryx/signatureav\"")
    return once(s,"type Report struct {",'''type Report struct {
 SignatureRoots []string `json:"signature_roots,omitempty"`
 SignatureState string `json:"signature_state,omitempty"`
 SignatureProblem string `json:"signature_problem,omitempty"`
 SignatureReports []signatureav.Report `json:"signature_reports,omitempty"`
''')
def summary(s):
    return rx(s,r'(r\.Partial\s*=\s*r\.Cancelled[^\n]+)',lambda m:m[1]+' || r.SignatureState == "partial" || r.SignatureState == "unavailable" || r.SignatureState == "running"')
def main(s):
    a=s.index("func (s *State) scanRoots(");b=s.index("func (s *State) quickScan(",a)
    x=s[a:b]
    x=rx(x,r'(r\s*:=\s*engine\.Report\s*\{\s*Root:\s*strings\.Join\(roots,\s*"; "\),)',lambda m:m[1]+' SignatureRoots:append([]string(nil),roots...),')
    s=s[:a]+x+s[b:]
    a=s.index("func (s *State) startScan(");b=s.index("func main(",a)
    x=s[a:b]
    x=rx(x,r'(s\.mu\.Lock\(\)\s*defer s\.mu\.Unlock\(\)\s*rep\.ScanID\s*=)',lambda m:'s.signaturePhase312(ctx,&rep,path,automatic)\n'+m[1])
    return s[:a]+x+s[b:]
def html(s):
    s=once(s,'<section id="overview">','''<section id="overview"><div class="card" id="z312Signature"><h2>Motor de assinaturas ClamAV</h2><p id="z312SigState" role="status">Aguardando varredura.</p><div id="z312SigDetails"></div><p class="sub">Contagem independente; não somar arquivos dos dois motores. Sem detecção não significa computador seguro.</p></div>''')
    s=s.replace("os botões de varredura ZEUVORYX continuam heurísticos.","as varreduras normais agora incluem uma segunda etapa independente, quando ClamAV estiver instalado.")
    js=r"""
;(()=>{let last='';async function draw(){try{
const v=await window.getScanState(),r=v.report||{},p=document.getElementById('z312SigState'),d=document.getElementById('z312SigDetails');
const state=r.signature_state||'',rr=r.signature_reports||[],key=JSON.stringify([state,r.signature_problem,rr,v.running]);if(key===last)return;last=key;
p.textContent=({running:'ClamAV: análise em andamento',finished:'ClamAV: etapa finalizada',partial:'ClamAV: cobertura parcial',unavailable:'ClamAV indisponível',pending:'ClamAV aguardando'})[state]||(v.running?'Primeira etapa: scanner ZEUVORYX':'Aguardando varredura.');
d.replaceChildren();if(r.signature_problem)d.append(node('p',r.signature_problem));
for(const q of rr){d.append(node('p',q.root+' · '+q.engine_database_version+' · '+fmt(q.scanned_files)+' arquivos · '+fmt(q.known_signatures)+' assinaturas · '+q.status));for(const f of (q.findings||[]).slice(0,100)){d.append(node('p',f.kind+' · '+f.signature+' · '+f.path));}if((q.findings||[]).length>100)d.append(node('p','Mais registros no JSON completo.'));if(q.problem)d.append(node('p',q.problem));}
}catch(e){document.getElementById('z312SigState').textContent='Estado do motor indisponível';}}
setInterval(draw,3000);draw();})();
"""
    return once(s,'</script></body></html>',js+'</script></body></html>')
for path,fn in [('engine/scanner.go',report),('engine/scan_upgrade.go',summary),('main.go',main),('desktop.html',html)]:edit(path,fn)
Path("signature312_windows.go").write_text(Path("lab/signature312_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
# Version promotion only after the previous generator (which validates 3.0.10 then produces 3.0.11).
for p in list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]:
    if ".git" in p.parts:continue
    s=p.read_text(encoding="utf-8");n=s.replace("3.0.11","3.0.12")
    if s!=n:p.write_text(n,encoding="utf-8")
