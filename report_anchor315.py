"""Structural CSS insertion for the generated desktop; independent of Arcade."""
from html.parser import HTMLParser


class _Boundaries(HTMLParser):
    def __init__(self, html):
        super().__init__(convert_charrefs=False)
        self.starts = [0]
        for line in html.split('\n')[:-1]:
            self.starts.append(self.starts[-1] + len(line) + 1)
        self.in_head = False
        self.styles = []
        self.head_close = None
        self.body_open = None
        self.html_open_end = None
        self.decl_end = None
        self.feed(html)
        self.close()

    def offset(self):
        line, col = self.getpos()
        return self.starts[line - 1] + col

    def handle_decl(self, decl):
        if self.decl_end is None:
            self.decl_end = self.offset() + len(decl) + 3

    def handle_starttag(self, tag, attrs):
        if tag == "html" and self.html_open_end is None:
            self.html_open_end = self.offset() + len(self.get_starttag_text())
        elif tag == "head":
            self.in_head = True
        elif tag == "body":
            self.in_head = False
            if self.body_open is None:
                self.body_open = self.offset()

    def handle_endtag(self, tag):
        if tag == "style" and self.in_head:
            self.styles.append(self.offset())
        elif tag == "head":
            self.in_head = False
            if self.head_close is None:
                self.head_close = self.offset()


def insert_report_css(html, css):
    """Return modified HTML and the actual anchor used; no dependence on games."""
    if not css or css in html:
        raise ValueError("empty or already integrated report CSS")
    tags = _Boundaries(html)
    if tags.styles:
        pos, insertion, anchor = tags.styles[-1], css, "</style> in head"
    else:
        insertion = "<style>" + css + "</style>"
        if tags.head_close is not None:
            pos, anchor = tags.head_close, "</head>"
        elif tags.body_open is not None:
            pos, anchor = tags.body_open, "<body>"
        else:
            pos = tags.html_open_end if tags.html_open_end is not None else (tags.decl_end or 0)
            anchor = "document start fallback"
    return html[:pos] + insertion + html[pos:], anchor
