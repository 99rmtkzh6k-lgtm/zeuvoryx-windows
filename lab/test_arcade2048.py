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
from retro21 import load_universo
family = load_universo()
for cycle in range(1, 6):
    Assets().feed(family)
    assert "default-src 'none'" in family and "connect-src 'none'" in family
    assert not re.search(r"@import|\b(?:fetch|XMLHttpRequest|WebSocket|postMessage)\s*\(|window\.parent|chrome\.webview|https?://", family, re.I)
    assert 'id=="familia-botelho-zumbis"' in binding
    assert 'family_game_id = "familia-botelho-zumbis"' in generator
    assert 'retro_docs[family_game_id] = family_game' in generator
    for action in ("minerar", "vender", "comprarArma", "atacar", "invasao", "catarLixo", "amassarLixo", "colocarNoConteiner", "regarArvore", "trocarCupom", "construirFoguete", "buscarPovo", "darBeijo"):
        assert 'onclick="'+action+'()"' in family and 'function '+action+'()' in family
    print(f"Family game integration contract {cycle}/5 passed (not gameplay)")

# Complete single-file game contract and syntax check.
import shutil
import subprocess
import tempfile
assert not (base / "familia-botelho/game.html").exists()
assert not (base / "familia-botelho/botelho.html").exists()
assert set(p.name for p in (base.parent / "universo-botelho").iterdir()) == {"index.html", "style.css", "script.js", "README.md"}
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
assert 'onclick="abrirJogo(1)"' in family and 'onclick="abrirJogo(2)"' in family
print("Unified Universo HTML JavaScript syntax and function contract passed")

# Five complete campaigns per game; simulated DOM, not browser/audio validation.
with tempfile.TemporaryDirectory() as tmp:
    harness = Path(tmp) / "universo-logic.js"
    harness.write_text("const source = " + __import__("json").dumps(scripts[0]) + ";\n" + "\nconst ids = new Set([\"universo-header\",\"topo-recursos\",\"botoes-musica\",\"musica\",\"vozes\",\"universo-voltar\",\"universo-menu\",\"universo-g1\",\"tela-inicial\",\"ajuda\",\"tela-escolha\",\"tela-jogo\",\"fase\",\"clima\",\"objetivo\",\"progresso\",\"recompensa\",\"mapa\",\"base-hp\",\"map-icons\",\"horda\",\"status\",\"heroi\",\"botoes-acao\",\"invasao-btn\",\"ataque-btn\",\"ninho-btn\",\"next\",\"botoes-construcao\",\"b-centroTreinamento\",\"b-fabricaArmas\",\"b-fabricaJipes\",\"b-barBotelho\",\"b-torreVigia\",\"b-muro\",\"tela-loja\",\"loja-saldo\",\"super-comida\",\"tela-fim\",\"resultado\",\"universo-g2\",\"tela-robo\",\"robo-titulo\",\"robo-cor\",\"robo-fase\",\"robo-progresso\",\"robo-mapa\",\"robo-avatar\",\"robo-nome\",\"robo-arvores\",\"robo-foguete\",\"robo-status\",\"robo-familia\",\"r-catar\",\"r-amassar\",\"r-conteiner\",\"r-regar\",\"r-cupom\",\"r-foguete\",\"r-buscar\",\"r-beijo\",\"r-next\",\"robo-final\",\"universo-comum\",\"log-vozes\",\"save-code\"]);\nlet checks=0;function ck(v,msg){if(!v)throw Error(msg);checks++}\nfor(let cycle=0;cycle<5;cycle++){\n  let els=new Map(),store={};const document={getElementById(id){if(!ids.has(id))throw Error(\"unknown DOM id \"+id);if(!els.has(id))els.set(id,{hidden:false,textContent:\"\",value:\"\",disabled:false,style:{}});return els.get(id)},addEventListener(){}};\n  const localStorage={setItem(k,v){store[k]=v},getItem(k){return store[k]}},window={document,localStorage};\n  new Function(\"window\",\"document\",\"localStorage\",source)(window,document,localStorage);\n  let a=window;a.iniciarJogo2();\n  function recycle(n){for(let i=0;i<n;i++){a.catarLixo();a.amassarLixo();a.colocarNoConteiner()}}\n  function next2(){ck(a.verificarObjetivoFase2(),\"robot phase objective\");a.proximaFase2()}\n  a.catarLixo();a.catarLixo();next2();a.amassarLixo();a.amassarLixo();next2();a.colocarNoConteiner();next2();a.colocarNoConteiner();recycle(3);next2();recycle(5);for(let i=0;i<4;i++)a.regarArvore();next2();recycle(10);for(let i=0;i<4;i++)a.regarArvore();next2();recycle(30);a.construirFoguete();next2();a.buscarPovo();for(let i=0;i<5;i++)a.darBeijo();next2();recycle(30);for(let i=0;i<6;i++)a.regarArvore();for(let i=0;i<3;i++){a.buscarPovo();for(let b=0;b<5;b++)a.darBeijo()}next2();recycle(20);for(let i=0;i<4;i++)a.regarArvore();for(let i=0;i<2;i++){a.buscarPovo();ck(!a.verificarVitoriaFinal2(),\"no victory before kisses\");for(let b=0;b<5;b++)a.darBeijo()}ck(a.verificarVitoriaFinal2(),\"robot victory\");a.salvarJogo();let s=JSON.parse(document.getElementById(\"save-code\").value).jogo2;ck(s.pessoasVoltaram===30&&s.arvores===100&&s.familiaBeijos===5&&s.beijos===30,\"robot values\");ck(a.carregarJogo()===true,\"robot save\");\n  a.menuUniverso();a.abrirJogo(1);a.iniciarJogo();a.escolherPersonagem(\"LEONARDO\");\n  for(let p=1;p<=10;p++){for(let i=0;i<20;i++)a.minerar();for(let i=0;i<8;i++)a.vender();for(let i=0;i<3;i++)a.comprarArma();if(p===2){a.construirPredio(\"barBotelho\");a.construirPredio(\"centroTreinamento\");a.treinarSoldado()}if(p===3){a.construirPredio(\"fabricaJipes\");a.fabricarJipe()}if(p===4)a.construirPredio(\"torreVigia\");if(p===5)a.escolherPersonagem(\"CRISZINHA\");for(let w=0;w<(p===9?5:1);w++){a.invasao();for(let k=0;k<100&&!document.getElementById(\"ataque-btn\").disabled;k++){a.atacar();a.comprarComida()}ck(document.getElementById(\"ataque-btn\").disabled,\"wave clear \"+p)}if(p===6||p===7)a.destruirNinho();if(p<10){ck(a.verificarObjetivoFase(),\"zombie phase \"+p);a.proximaFase()}else ck(a.verificarVitoriaFinal(),\"zombie victory\")}\n  a.salvarJogo();ck(a.carregarJogo()===true,\"zombie save\");\n}\n\nconsole.log(checks+' simulated logic assertions passed');", encoding="utf-8")
    subprocess.run([node, str(harness)], check=True)
