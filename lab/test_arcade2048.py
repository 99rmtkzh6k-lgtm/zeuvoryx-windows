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
    assert 'id=="2048"' in binding and 'classicGame315(id)' in binding
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


# Browser proof exercises the exact launcher and generated iframe source.
# Native WebView callbacks are mocked; these screenshots are not installed-app proof.
if "--browser" in sys.argv:
    from playwright.sync_api import sync_playwright
    out = base.parent / "arcade-proof"
    out.mkdir(exist_ok=True)
    shell = """<!doctype html><html><head><meta charset="utf-8"><style>body{margin:0;background:#101b30;color:#fff;font:16px system-ui}main{padding:12px;max-width:1200px;margin:auto}.hidden{display:none!important}.card,.entry{padding:12px}button,input,select{font:inherit;padding:8px}button{cursor:pointer} [hidden]{display:none!important}"""+(base/"games314.css").read_text(encoding="utf-8")+"</style></head><body><button id='zGamesNav314'>Jogos</button><button id='zFamilyNav314'>Família</button><button id='start' data-page='overview'>Escanear</button><button id='quick'>Rápido</button><button data-page='alerts'>Alertas</button><main><section id='overview'></section><section id='alerts'></section>"+(base/"games314.html").read_text(encoding="utf-8")+"</main></body></html>"
    bridge = """window.busy=false;window.records={};window.getScanState=async()=>({running:window.busy,report:{}});window.allowLocalGame315=async()=>true;window.getGameRecord315=async id=>window.records[id]||0;window.saveGameRecord315=async(id,n)=>{window.records[id]=Math.max(window.records[id]||0,n);return true};const zRetroCatalog=[],zRetroDocs={},zArcade2048='';const volume=n=>String(n);const node=(tag,value)=>{const e=document.createElement(tag);e.textContent=value;return e};"""
    with sync_playwright() as pw:
        try:
            browser=pw.chromium.launch(channel="msedge",headless=True)
        except Exception:
            browser=pw.chromium.launch(channel="chrome",headless=True)
        for width,height in [(1366,768),(1920,1080)]:
            page=browser.new_page(viewport={"width":width,"height":height},device_scale_factor=1)
            page.set_content(shell)
            page.add_script_tag(content=bridge)
            page.add_script_tag(content=launcher)
            page.locator("#zGamesNav314").click()
            def launch_game(title):
                page.locator("#zGameSearch").fill(title)
                page.locator("#zGameCatalog article").filter(has_text=title).get_by_role("button",name="Jogar aqui").click()
                frame=page.frame_locator("#zGameBoard iframe")
                frame.locator("#info").wait_for(state="visible")
                assert frame.locator("#msg").inner_text()!="Jogo indisponível, tente novamente"
                page.locator("#zGameFullscreen315").click()
                return frame
            def back():
                page.locator("#zGameClose315").click()
                assert page.locator("#zGameBoard iframe").count()==0
            f=launch_game("Tetris")
            assert f.locator("canvas.tetris").get_attribute("width")=="300"
            assert f.locator("canvas.tetris").get_attribute("height")=="600"
            f.get_by_role("button",name="Girar / Espaço").click()
            f.get_by_role("button",name="Pausar / P").click()
            assert "Pausado" in f.locator("#info").inner_text()
            f.get_by_role("button",name="Pausar / P").click()
            page.screenshot(path=str(out/f"tetris-{width}x{height}.png"))
            page.evaluate("window.busy=true")
            f.locator("#pause").wait_for(state="visible")
            page.evaluate("window.busy=false")
            f.locator("#pause").wait_for(state="hidden")
            back()
            f=launch_game("Poker")
            assert f.locator(".playing").count()==8
            f.get_by_role("button",name="Check",exact=True).click()
            assert f.locator(".playing").count()==11
            page.screenshot(path=str(out/f"poker-{width}x{height}.png"))
            for _ in range(3):
                f.get_by_role("button",name="Check",exact=True).click()
            assert "venceu" in f.locator("#msg").inner_text()
            assert f.locator(".playing").count()==13
            chips=[int(re.search(r"(\d+) fichas",v).group(1)) for v in f.locator(".players strong").all_inner_texts()]
            assert sum(chips)==4000,chips
            f.get_by_role("button",name="Nova Rodada").click()
            assert f.locator(".playing").count()==8
            back()
            f=launch_game("Campo minado")
            assert f.locator(".mines button").count()==81
            f.locator(".mines button").nth(40).click()
            assert f.locator("#msg").inner_text()!="Derrota"
            available=f.locator(".mines button:not(:disabled)")
            if available.count():
                available.first.click(button="right")
                assert "Bombas restantes: 9" in f.locator("#info").inner_text()
            page.screenshot(path=str(out/f"campo-minado-{width}x{height}.png"))
            back()
            f=launch_game("Jogo da memória")
            assert f.locator(".memory button").count()==16
            f.locator(".memory button").nth(0).click()
            f.locator(".memory button").nth(1).click()
            assert "Tentativas: 1" in f.locator("#info").inner_text()
            page.screenshot(path=str(out/f"memoria-{width}x{height}.png"))
            back()
            assert page.evaluate("document.documentElement.scrollWidth <= innerWidth")
            print(f"PASS gameplay browser harness {width}x{height}: four games, pause, fullscreen, poker chip conservation")
            page.close()
        browser.close()
