"""Test report CSS independently of games, and inspect real generated desktop."""
import sys
from pathlib import Path
from html.parser import HTMLParser

sys.path.insert(0, str(Path(__file__).resolve().parent.parent))
from report_anchor315 import insert_report_css

CSS = "#reports315{gap:10px}"
cases = [
    ("<html><head><style>body{color:red}</style></head><body></body></html>", "</style> in head"),
    ("<HTML><HEAD><STYLE>body{color:red}</STYLE >\n</HEAD><BODY></BODY></HTML>", "</style> in head"),
    ("<html><head><script>let x='</style></head>';</script></head><body></body></html>", "</head>"),
    ("<html><head></head><body><style>p{color:red}</style></body></html>", "</head>"),
    ("<html><body><main></main></body></html>", "<body>"),
    ("<!DOCTYPE html><html><main></main></html>", "document start fallback"),
    ("<main></main>", "document start fallback"),
]
for source, expected in cases:
    actual, anchor = insert_report_css(source, CSS)
    assert anchor == expected, (anchor, expected)
    assert actual.count(CSS) == 1
    assert actual.replace(CSS, "", 1) == source or actual.replace("<style>" + CSS + "</style>", "", 1) == source
    try:
        insert_report_css(actual, CSS)
    except ValueError:
        pass
    else:
        raise AssertionError("duplicate report CSS was accepted")
print(f"Report CSS boundary fixtures: {len(cases)}/{len(cases)} passed (no games)")

if len(sys.argv) == 2:
    class Check(HTMLParser):
        def __init__(self):
            super().__init__()
            self.ids = []
            self.in_head = False
            self.in_style = False
            self.css = 0

        def handle_starttag(self, tag, attrs):
            if tag == "head":
                self.in_head = True
            if tag == "style" and self.in_head:
                self.in_style = True
            attrs = dict(attrs)
            if "id" in attrs:
                self.ids.append(attrs["id"])

        def handle_endtag(self, tag):
            if tag == "style":
                self.in_style = False
            if tag == "head":
                self.in_head = False

        def handle_data(self, data):
            if self.in_style:
                self.css += data.count(CSS)

    p = Check()
    p.feed(Path(sys.argv[1]).read_text(encoding="utf-8"))
    p.close()
    assert p.css == 1, "report CSS missing or duplicated in head"
    for id_ in ("reports315", "z315ReportSave", "z315ReportSource", "z315ShowText", "z315ReportShortcut"):
        assert p.ids.count(id_) == 1, ("report control missing or duplicated", id_)
    print("Generated desktop: report stylesheet and five controls verified")
