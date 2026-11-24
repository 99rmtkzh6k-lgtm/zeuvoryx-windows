"""Embed the lightweight offline help after the final desktop HTML generator.
Run only in the next authorized version pipeline; no runtime Python dependency.
"""
from pathlib import Path
import json

def transform(html,script,articles):
    anchor="</script></body></html>"
    if html.count(anchor)!=1 or "zHelpPanel" in html:
        raise ValueError("unexpected or previously patched desktop")
    if script.count("__Z_HELP_DATA__")!=1:
        raise ValueError("invalid help script placeholder")
    data=json.dumps(articles,ensure_ascii=False,separators=(",",":")).replace("<","\\u003c")
    payload=script.replace("__Z_HELP_DATA__",data,1)
    if "</script" in payload.lower():
        raise ValueError("unsafe script termination")
    return html.replace(anchor,";\n"+payload+"\n"+anchor,1)

if __name__=="__main__":
    p=Path("desktop.html")
    s=Path("help_search.js").read_text(encoding="utf-8")
    data=json.loads(Path("help_articles.json").read_text(encoding="utf-8"))
    assert set(data)=={"pt-br","en","es","ja"} and all(len(v)==31 for v in data.values())
    p.write_text(transform(p.read_text(encoding="utf-8"),s,data),encoding="utf-8")
