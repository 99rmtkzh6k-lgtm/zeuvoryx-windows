from pathlib import Path
import re
def once(s,a,b):
    if s.count(a)!=1:raise RuntimeError("parallel314 anchor "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
def edit(path,fn):
    p=Path(path);p.write_text(fn(p.read_text(encoding="utf-8-sig")),encoding="utf-8")
Path("parallel314_windows.go").write_text(Path("lab/parallel314_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
def main(s):
    s=once(s,'func main() {','func main() {\n if len(os.Args)==2 && os.Args[1]=="-diagnostic-check314" {os.Exit(diagnosticCLI314())}\n')
    a=s.index('func (s *State) scanRoots(');b=s.index('func (s *State) quickScan(',a)
    x=s[a:b];x=once(x,'engine.Report {','''engine.Report {
 if enabled,_:=ctx.Value(parallelKey314{}).(bool);enabled {
   inner:=context.WithValue(ctx,parallelKey314{},false)
   r,_:=s.concurrent314(ctx,roots,func()(engine.Report,error){return s.scanRoots(inner,roots,self,progress),nil})
   return r
 }
''')
    s=s[:a]+x+s[b:]
    a=s.index('func (s *State) startScan(');b=s.index('func main(',a);x=s[a:b]
    x=once(x,'var rep engine.Report','ctx=context.WithValue(ctx,parallelKey314{},true)\n var rep engine.Report')
    x=once(x,'s.scanner.Scan(engine.WithLiveReport(ctx,s.publishLive),path,progress)','s.parallelCustom314(ctx,path,progress)')
    return s[:a]+x+s[b:]
def report(s):
    s=once(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/scopeprogress"')
    return once(s,'type Report struct {','''type Report struct {
 Inventory314 scopeprogress.Inventory `json:"inventory_314"`
 BehaviorLocal map[string]any `json:"behavior_local,omitempty"`
 BehaviorProblem string `json:"behavior_problem,omitempty"`
 WindowsLocal314 map[string]any `json:"windows_local_314,omitempty"`
 WindowsProblem314 string `json:"windows_problem_314,omitempty"`
''')
def desktop(s):
    a=s.index('if !bind("behavior313"')
    return s[:a]+'''if !bind("realDiagnostic314",func()(diagnostic314,error){
        operationGate.Lock();defer operationGate.Unlock()
        s.mu.Lock();running:=s.running;s.mu.Unlock()
        if running||native.status().Running{return diagnostic314{},fmt.Errorf("Aguarde a operação atual")}
        return realDiagnostic314()
    }){return}
'''+s[a:]
def html(s):
    s=once(s,'<div id="z312SigDetails"></div>','<div id="z312SigDetails"></div><div class="card"><h2>Diagnóstico real dos motores 3.0.14</h2><button id="z314Diagnostic">Testar motores instalados agora</button><p id="z314DiagnosticState" role="status"></p><pre id="z314DiagnosticDetails" style="white-space:pre-wrap;overflow-wrap:anywhere"></pre></div>')
    s=s.replace("Primeira etapa: scanner ZEUVORYX","Scanner ZEUVORYX e preparação do motor de assinaturas")
    s=s.replace("uma segunda etapa independente","uma análise independente em paralelo")
    s=s.replace("Testar componentes","Autoteste sintético de componentes")
    s=once(s,'</script></body></html>',"""document.getElementById('z314Diagnostic').onclick=()=>action(async()=>{const b=el('z314Diagnostic'),p=el('z314DiagnosticState'),d=el('z314DiagnosticDetails');b.disabled=true;p.textContent='Executando leitura real, motores e consulta do Windows...';d.textContent='';try{const r=await window.realDiagnostic314();const ok=r.LocalFile&&r.ScannerFiles===1&&r.ScannerAlerts===0&&r.ClamState==='finished'&&r.ClamFiles===1&&r.ClamSignatures>0;p.textContent=ok?'Motores responderam ao teste benigno; confira cobertura e Windows.':'Diagnóstico incompleto: verifique falhas abaixo.';d.textContent=JSON.stringify(r,null,2)}catch(e){p.textContent=String(e)}finally{b.disabled=false}});</script></body></html>""")
    return s
edit("main.go",main);edit("engine/scanner.go",report);edit("desktop_windows.go",desktop);edit("desktop.html",html)
