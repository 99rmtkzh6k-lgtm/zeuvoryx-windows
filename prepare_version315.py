from pathlib import Path
# Only generated sources and installer identity; preserve visual assets and all capabilities.
paths=list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]
for p in paths:
    if ".git" in p.parts: continue
    s=p.read_text(encoding="utf-8")
    n=s.replace("3.0.14","3.0.15")
    if n!=s:p.write_text(n,encoding="utf-8")
for name in ("desktop.html","setup309.iss"):
    if "3.0.15" not in Path(name).read_text(encoding="utf-8"):raise RuntimeError("Version identity missing: "+name)
