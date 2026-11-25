"""Final identity and startup hardening after existing source generators."""
from pathlib import Path
import re
def once(s,a,b):
    if s.count(a)!=1: raise RuntimeError("Expected exactly one anchor: "+a[:120]+" count="+str(s.count(a)))
    return s.replace(a,b,1)
def edit(p,fn):
    f=Path(p); f.write_text(fn(f.read_text(encoding="utf-8-sig")),encoding="utf-8")
def main(s):
    s=once(s,'flag.Parse()','''showVersion := flag.Bool("version",false,"show engine version")
    flag.Parse()
    if *showVersion { fmt.Println(engine.ScanEngineVersion); return }''')
    return s
def desktop(s):
    s=s.replace("3.0.5","3.0.9")
    s=once(s,'w := webview.NewWithOptions(','''profileBase,profileErr:=os.UserConfigDir()
    if profileErr!=nil { log.Print("stability webview_profile_unavailable"); desktopError("Não foi possível localizar o perfil local do Windows."); return }
    profile:=filepath.Join(profileBase,"ZEUVORYX","webview-306")
    if profileErr=os.MkdirAll(profile,0700);profileErr!=nil {log.Print("stability webview_profile_create_failed");desktopError("Não foi possível criar o perfil local da interface.");return}
    log.Print("stability webview_init")
    w := webview.NewWithOptions(''')
    s=once(s,'webview.WebViewOptions{Debug:', 'webview.WebViewOptions{DataPath:profile, Debug:')
    s=once(s,'w.SetHtml(localized)', '''if !bind("zStartupReady",func()bool{log.Print("stability ui_ready version=3.0.9-lab");return true}){return}
    w.SetHtml(localized)''')
    s=s.replace("windows-build-lab-305.yml","windows-build-lab-309.yml")
    return s
def html(s):
    s=s.replace("3.0.5","3.0.9")
    # Final labels must agree with the updater, not the historical generator version.
    match = re.search(r'const Current\s*=\s*"([0-9]+\.[0-9]+\.[0-9]+)"', Path("updater/check.go").read_text(encoding="utf-8"))
    if match is None or match.group(1) != "3.0.9":
        raise RuntimeError("Updater identity mismatch")
    version = match.group(1)
    for label in ("Diagnóstico local", "Local diagnostics", "Atualizações LAB —", "LAB updates —", "Actualizaciones LAB —"):
        pattern = re.escape(label) + r" [0-9]+\.[0-9]+\.[0-9]+"
        s, count = re.subn(pattern, lambda _: label + " " + version, s)
        if count == 0:
            raise RuntimeError("Missing localized identity: " + label)
        if any(m.group() != label + " " + version for m in re.finditer(pattern, s)):
            raise RuntimeError("Stale localized identity: " + label)
    s=once(s,"setTimeout(zUpdateRun,1500);","/* Explicit update checks only; no network call during startup. */")
    return once(s,"</script></body></html>",";window.zStartupReady().catch(()=>{});\n</script></body></html>")
edit("main.go",main)
edit("desktop_windows.go",desktop)
edit("desktop.html",html)
edit("engine/scan_upgrade.go",lambda s:once(s,'const ScanEngineVersion = "3.0.5-lab"','const ScanEngineVersion = "3.0.9-lab"'))
