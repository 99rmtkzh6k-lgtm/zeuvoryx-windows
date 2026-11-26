from pathlib import Path
from hashlib import sha1
import json
import re

def load_retro():
    root=Path("lab/retro21")
    manifest=json.loads((root/"catalog.json").read_text(encoding="utf-8"))
    assert len(manifest)==10 and len({g["id"] for g in manifest})==10
    docs={}
    storage="""<script>const zArcadeSessionStore=(()=>{const m=new Map();return Object.freeze({getItem:k=>m.has(String(k))?m.get(String(k)):null,setItem:(k,v)=>m.set(String(k),String(v)),removeItem:k=>m.delete(String(k)),clear:()=>m.clear()});})();</script>"""
    policy="""<meta http-equiv="Content-Security-Policy" content="default-src 'none'; script-src 'unsafe-inline'; style-src 'unsafe-inline'; img-src data:; media-src data:; connect-src 'none'; object-src 'none'; frame-src 'none'; base-uri 'none'; form-action 'none'">"""
    for g in manifest:
        assert re.fullmatch(r"[a-z0-9-]+",g["id"])
        raw=(root/(g["id"]+".html")).read_bytes()
        assert sha1(b"blob "+str(len(raw)).encode()+b"\0"+raw).hexdigest()==g["sha"],g["id"]
        html=raw.decode("utf-8")
        assert not re.search(r'<(?:iframe|object|embed|base|form)\b|<script\b[^>]*\bsrc\s*=|<link\b[^>]*\bhref\s*=|(?:fetch|XMLHttpRequest|WebSocket|postMessage|eval)\s*\(|window\.parent|chrome\.webview|document\.cookie|https?://',html,re.I),g["id"]
        html=re.sub(r"\blocalStorage\b","zArcadeSessionStore",html)
        html,n=re.subn(r"<head>",lambda m:m[0]+policy+storage,html,count=1,flags=re.I)
        assert n==1 and "connect-src 'none'" in html
        docs[g["id"]]=html
    return manifest,docs
