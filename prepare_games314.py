from pathlib import Path
from html.parser import HTMLParser
import re
def regex_one(s,pattern,replacement):
    s,n=re.subn(pattern,lambda m:replacement(m),s)
    if n!=1:raise RuntimeError('Games regex anchor '+pattern+' count='+str(n))
    return s
def one(s,a,b):
    if s.count(a)!=1:raise RuntimeError("Games anchor "+a[:100])
    return s.replace(a,b,1)
Path("games314_windows.go").write_text(Path("lab/games314_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
s=one(s,'if !bind("recentSignature314",recentSignature314){return}', 'if !bind("recentSignature314",recentSignature314){return}\nif !bind("allowLocalGame315",allowLocalGame315){return}')
s=one(s,'import (','import (\n "github.com/zeuvoryx/zeuvoryx/familygate"')
s=regex_one(s,r'\bdropCache\s*:=\s*globalintel\.OpenDaily\(dropPath\)',lambda m:m[0]+'''
familyPath:="";if dropDirErr==nil{familyPath=filepath.Join(dropDir,"ZEUVORYX","family314.json")}
family314=familygate.Open(familyPath)
if !bind("familyStatus314",family314.Status){return}
if !bind("familySave314",func(p familygate.Policy,password string)(bool,error){e:=family314.Save(p,password);return e==nil,e}){return}
if !bind("familyRate314",func(id string,r familygate.Rating,password string)(bool,error){if !localGame315(id){return false,fmt.Errorf("Jogo fora do catálogo local")};e:=family314.Rate(id,r,password);return e==nil,e}){return}
if !bind("familyDNSCheck314",func()familygate.DNSCheck{return familySystemDNS314(nativeCtx)}){return}
''')
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=one(s,'</nav>','<button id="zGamesNav314" type="button" aria-expanded="false">Games · Jogos</button><button id="zFamilyNav314" type="button">Controle parental</button></nav>')
s=one(s,'</main>',Path("lab/games314.html").read_text(encoding="utf-8")+Path("lab/family314.html").read_text(encoding="utf-8")+'</main>')
s=one(s,'</style></head>',Path("lab/games314.css").read_text(encoding="utf-8")+'</style></head>')
s=one(s,'</script></body></html>',Path("lab/games315_core.js").read_text(encoding="utf-8")+Path("lab/classic315_core.js").read_text(encoding="utf-8")+Path("lab/chess315_vendor.js").read_text(encoding="utf-8")+Path("lab/games314.js").read_text(encoding="utf-8")+Path("lab/family314.js").read_text(encoding="utf-8")+'</script></body></html>')
class Check(HTMLParser):
    def __init__(self):super().__init__();self.ids=[]
    def handle_starttag(self,tag,attrs):
        a=dict(attrs)
        if 'id' in a:self.ids.append(a['id'])
c=Check();c.feed(s)
for k in ('zGames314','zGamesNav314','zGameCatalog','zGameBoard','zGameBack','zGameClose315','zFamilyNav314','zFamily314','start','quick','zTitaniumIdentity'):
    if c.ids.count(k)!=1:raise RuntimeError("Games missing or duplicate "+k)
if len(s.encode("utf-8"))>1500000:raise RuntimeError("Desktop budget exceeded")
p.write_text(s,encoding="utf-8")
