from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1:raise RuntimeError("314 anchor "+a[:100]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
def edit(path,fn):
    p=Path(path);p.write_text(fn(p.read_text(encoding="utf-8-sig")),encoding="utf-8")
Path("bundle314_windows.go").write_text(Path("lab/bundle314_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
Path("bundle314_windows_test.go").write_text(Path("lab/bundle314_windows_test.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
def main(s):
    return once(s,'func main() {','func main() {\n if len(os.Args)==3 && os.Args[1]=="-signature-check314" {os.Exit(signatureCheck314(os.Args[2]))}\n')
def html(s):
    s=s.replace("Requer instalação oficial separada do ClamAV e assinaturas atualizadas pelo FreshClam.","ClamAV oficial incluído no instalador; atualização automática da base pelo FreshClam antes da análise. Internet necessária para a primeira atualização.")
    s=s.replace("quando ClamAV estiver instalado.","quando a base oficial estiver disponível.")
    s=s.replace("ClamAV indisponível","ClamAV ou base indisponível")
    return s
def installer(s):
    return once(s,'[Files]','[Files]\nSource: "dist\\clamav\\*"; DestDir: "{app}\\clamav"; Flags: ignoreversion recursesubdirs createallsubdirs\nSource: "dist\\third-party\\*"; DestDir: "{app}\\third-party"; Flags: ignoreversion recursesubdirs createallsubdirs\nSource: "dist\\media-verifier\\*"; DestDir: "{app}\\media-verifier"; Flags: ignoreversion recursesubdirs createallsubdirs')
edit("main.go",main);edit("desktop.html",html);edit("setup309.iss",installer)
for p in list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]:
    if ".git" in p.parts:continue
    s=p.read_text(encoding="utf-8");n=s.replace("3.0.13","3.0.14")
    if s!=n:p.write_text(n,encoding="utf-8")
