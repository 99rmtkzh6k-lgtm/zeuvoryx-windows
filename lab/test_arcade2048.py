"""Check vendored game identity and desktop isolation in the generated inputs."""
from hashlib import sha256
from html.parser import HTMLParser
from pathlib import Path
import re


class Assets(HTMLParser):
    def handle_starttag(self, tag, attrs):
        for name, _ in attrs:
            assert name not in {"src", "href", "srcset", "action"}, (tag, name)


base = Path(__file__).resolve().parent
game = (base / "arcade2048/game.html").read_bytes()
approved = (base / "approved_game_hashes.json")
import json
records=json.loads(approved.read_text(encoding="utf-8"))
assert sha256(game).hexdigest() == records["2048"]["sha256"], "2048 hash mismatch; do not silently bless a modified game"
html = game.decode()
launcher = (base / "games314.js").read_text(encoding="utf-8")
shared = (base / 'gamewindow.js').read_text(encoding='utf-8')
binding = (base / "games314_windows.go.txt").read_text(encoding="utf-8")
generator = (base.parent / "prepare_games314.py").read_text(encoding="utf-8")
for cycle in range(1, 6):
    Assets().feed(html)
    assert "default-src 'none'" in html and "connect-src 'none'" in html
    assert not re.search(r"@import|\b(?:fetch|XMLHttpRequest|WebSocket)\b", html)
    assert "new GameManager" in html
    assert "frame.setAttribute('sandbox','allow-scripts')" in shared
    assert "await window.allowLocalGame315(id)" in launcher and "!==true" in launcher
    assert "board.replaceChildren()" in shared and "win.registerGame(game" in launcher and "win.open(id)" in launcher
    assert 'return id=="2048"' in binding
    assert "games315_core.js" not in generator and "chess315_vendor.js" not in generator
    print(f"Arcade integration contract {cycle}/5 passed")

import sys
sys.path.insert(0,str(base.parent))
from retro21 import load_retro
for cycle in range(1,6):
    catalog,docs=load_retro()
    for item in catalog:
        id=item["id"]
        assert ('id=="'+id+'"') in binding
        assert "connect-src 'none'" in docs[id]
        assert "zArcadeSessionStore" in docs[id]
    assert 'zGameRestart21' in launcher and 'zRetroCatalog' in launcher
    print(f"Retro catalog integrity and isolation {cycle}/5 passed (not gameplay)")
