from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1:raise RuntimeError("314 anchor "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
def edit(path,fn):
    p=Path(path);p.write_text(fn(p.read_text(encoding="utf-8-sig")),encoding="utf-8")
Path("bundle314_windows.go").write_text(Path("lab/bundle314_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
Path("bundle314_windows_test.go").write_text(Path("lab/bundle314_windows_test.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
Path("optional_signatures315_windows.go").write_text(Path("lab/optional_signatures315_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
Path("optional_signatures315_windows_test.go").write_text(Path("lab/optional_signatures315_windows_test.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
def desktop(s):
    anchor='if !bind("chooseClamFolder311"'
    if s.count(anchor)!=1:raise RuntimeError("Optional providers anchor")
    return s.replace(anchor,'if !bind("openOptionalSignature315",openOptionalSignature315){return}\n'+anchor,1)
def main(s):
    return once(s,'func main() {','func main() {\n if len(os.Args)==3 && os.Args[1]=="-signature-check314" {os.Exit(signatureCheck314(os.Args[2]))}\n')
def html(s):
    s=s.replace("Requer instalação oficial separada do ClamAV e assinaturas atualizadas pelo FreshClam.","ClamAV oficial integrado; usa assinaturas locais sem internet, com aviso se antigas. Internet necessária para adquirir a primeira base quando ausente.")
    s=s.replace("quando ClamAV estiver instalado.","quando a base oficial estiver disponível.")
    s=s.replace("ClamAV indisponível","ClamAV ou base indisponível")
    s=once(s,'<div id="z312SigDetails"></div>','<div id="z312SigDetails"></div><div class="card"><h2>Assinaturas complementares opcionais</h2><p>ZEUVORYX e atualizações oficiais ClamAV permanecem gratuitos. Serviços independentes podem cobrar por bases adicionais. A contratação, pagamento, renovação e cancelamento são tratados diretamente pelo usuário com o fornecedor. Nenhuma contratação é obrigatória. Nesta versão o catálogo apenas consulta fornecedores; a ativação das bases está em validação. Não altere manualmente o diretório oficial. A falha ou cancelamento de extensão não deverá interromper a proteção gratuita. Os direitos legais do consumidor e as licenças originais permanecem preservados. Consulte TERMOS_EXTENSOES_ZEUVORYX.txt incluído na instalação.</p><button id="zOptionalSecurite315">Conhecer SecuriteInfo</button><button id="zOptionalPatrol315">Conhecer Malware Patrol</button><p id="zOptionalSignatureState315" role="status"></p></div>')
    s=once(s,'</script></body></html>',"""for(const [id,provider] of [['zOptionalSecurite315','securiteinfo'],['zOptionalPatrol315','malwarepatrol']]){el(id).onclick=()=>action(async()=>{const p=el('zOptionalSignatureState315');try{const ok=await window.openOptionalSignature315(provider);p.textContent=ok?'Solicitação de abertura enviada ao navegador. Nenhuma assinatura foi ativada.':'Não foi possível abrir a página.'}catch(e){p.textContent=String(e)}});}</script></body></html>""")
    return s
def installer(s):
    s=once(s,'[Setup]','[Setup]\nInfoBeforeFile=TERMOS_EXTENSOES_ZEUVORYX.txt')
    s=once(s,'[Files]','[Files]\nSource: "TERMOS_EXTENSOES_ZEUVORYX.txt"; DestDir: "{app}"; Flags: ignoreversion')
    return once(s,'[Files]','[Files]\nSource: "dist\\clamav\\*"; DestDir: "{app}\\clamav"; Flags: ignoreversion recursesubdirs createallsubdirs\nSource: "dist\\third-party\\*"; DestDir: "{app}\\third-party"; Flags: ignoreversion recursesubdirs createallsubdirs\nSource: "dist\\media-verifier\\*"; DestDir: "{app}\\media-verifier"; Flags: ignoreversion recursesubdirs createallsubdirs')
edit("main.go",main);edit("desktop_windows.go",desktop);edit("desktop.html",html);edit("setup309.iss",installer)
for p in list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]:
    if ".git" in p.parts:continue
    s=p.read_text(encoding="utf-8");n=s.replace("3.0.13","3.0.14")
    if s!=n:p.write_text(n,encoding="utf-8")
