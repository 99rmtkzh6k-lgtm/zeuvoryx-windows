from pathlib import Path
from html.parser import HTMLParser
def one(s,a,b):
    if s.count(a)!=1:raise RuntimeError("Games anchor "+a[:100])
    return s.replace(a,b,1)
Path("games314_windows.go").write_text(Path("lab/games314_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=one(s,'if !bind("openYouTubeVerify314",openYouTubeVerify314){return}', 'if !bind("openYouTubeVerify314",openYouTubeVerify314){return}\nif !bind("openGame314",openGame314){return}')
s=one(s,'if !bind("zDropStatus",dropCache.Status){return}', 'if !bind("zDropStatus",dropCache.Status){return}\nif !bind("reviewGame314",func(id string)(gamecatalog.Review,error){return reviewGame314(nativeCtx,id,dropCache)}){return}')
s=one(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/gamecatalog"')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=one(s,'</nav>','<button id="zGamesNav314" type="button" aria-expanded="false">Games · Jogos</button></nav>')
s=one(s,'</main>',Path("lab/games314.html").read_text(encoding="utf-8")+'</main>')
s=one(s,'</style></head>',Path("lab/games314.css").read_text(encoding="utf-8")+'</style></head>')
s=one(s,'</script></body></html>',Path("lab/games314.js").read_text(encoding="utf-8")+'</script></body></html>')
class Check(HTMLParser):
    def __init__(self):super().__init__();self.ids=[]
    def handle_starttag(self,tag,attrs):
        a=dict(attrs)
        if 'id' in a:self.ids.append(a['id'])
c=Check();c.feed(s)
for k in ('zGames314','zGamesNav314','zGameCatalog','zGameBoard','zGameBack','start','quick','zTitaniumIdentity'):
    if c.ids.count(k)!=1:raise RuntimeError("Games missing or duplicate "+k)
if len(s.encode("utf-8"))>1500000:raise RuntimeError("Desktop budget exceeded")
p.write_text(s,encoding="utf-8")
