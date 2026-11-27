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
assert sha256(game).hexdigest() == "d0a8b940891c03d845d5c79d939a187232a7a75741ed45b5d9b9910fb8822ed3"
html = game.decode()
launcher = (base / "games314.js").read_text(encoding="utf-8")
binding = (base / "games314_windows.go.txt").read_text(encoding="utf-8")
generator = (base.parent / "prepare_games314.py").read_text(encoding="utf-8")
for cycle in range(1, 6):
    Assets().feed(html)
    assert "default-src 'none'" in html and "connect-src 'none'" in html
    assert not re.search(r"@import|\b(?:fetch|XMLHttpRequest|WebSocket)\b", html)
    assert "new GameManager" in html
    assert "frame.setAttribute('sandbox','allow-scripts')" in launcher
    assert "await window.allowLocalGame315(id)" in launcher and "!==true" in launcher
    assert "board.replaceChildren()" in launcher
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

# Family game: source, offline boundary and catalog registration (not gameplay).
family = (base / "familia-botelho/botelho.html").read_text(encoding="utf-8")
for cycle in range(1, 6):
    Assets().feed(family)
    assert "default-src 'none'" in family and "connect-src 'none'" in family
    assert not re.search(r"@import|\b(?:fetch|XMLHttpRequest|WebSocket|postMessage)\s*\(|window\.parent|chrome\.webview|https?://", family, re.I)
    assert 'id=="familia-botelho-zumbis"' in binding
    assert 'family_game_id = "familia-botelho-zumbis"' in generator
    assert 'retro_docs[family_game_id] = family_game' in generator
    for action in ("minerar", "vender", "comprar", "atacar", "invasao"):
        assert 'onclick="'+action+'()"' in family and 'function '+action+'()' in family
    print(f"Family game integration contract {cycle}/5 passed (not gameplay)")

# Complete single-file game contract and syntax check.
import shutil
import subprocess
import tempfile
assert not (base / "familia-botelho/game.html").exists()
for name in ("iniciarJogo","escolherPersonagem","atualizarTela","mostrarVoz","falarPersonagem","falarZumbi","minerar","vender","comprarArma","comprarComida","treinarSoldado","construirPredio","invasaoZumbi","atacarZumbi","verificarObjetivoFase","proximaFase","verificarDerrota","verificarVitoriaFinal","tocarMusica","tocarSom","ligarDesligarMusica","salvarJogo","carregarJogo"):
    assert re.search(r"function\s+" + name + r"\(", family), name
scripts = re.findall(r"<script>(.*?)</script>", family, re.S)
assert len(scripts) == 1
node = shutil.which("node")
if not node:
    raise RuntimeError("Node required for family game JavaScript syntax check")
with tempfile.TemporaryDirectory() as tmp:
    js = Path(tmp) / "botelho.js"
    js.write_text(scripts[0], encoding="utf-8")
    subprocess.run([node, "--check", str(js)], check=True)
print("Unified Botelho HTML JavaScript syntax and function contract passed")
