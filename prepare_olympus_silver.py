"""Olympus Silver: real offline embedded Windows interface."""
from pathlib import Path
from PIL import Image, ImageOps, ImageDraw
import base64, hashlib, re, io
from html.parser import HTMLParser
EXPECTED="c0ae621b1784260b67f8f8953d9ad96ac6cb2a5c9b96cd66f3125f22ef9de754"
def art():
    raw=Path("zeuvoryx_brand.jpg").read_bytes()
    if hashlib.sha256(raw).hexdigest()!=EXPECTED: raise ValueError("Wrong Zeus artwork")
    with Image.open(io.BytesIO(raw)) as source: im=source.convert("RGB")
    def crop(box):
        return im.crop(tuple(round(x*d) for x,d in zip(box,(im.width,im.height,im.width,im.height))))
    face=ImageOps.fit(crop((.27,.005,.74,.225)),(256,256),method=Image.Resampling.LANCZOS,centering=(.6,.4))
    temple=ImageOps.fit(crop((.03,.635,.97,.846)),(1050,360),method=Image.Resampling.LANCZOS)
    wide=temple.copy()
    portrait=ImageOps.fit(crop((.09,.005,.97,.229)),(550,360),method=Image.Resampling.LANCZOS,centering=(.64,.42))
    mask=Image.new("L",portrait.size,255);d=ImageDraw.Draw(mask)
    for x in range(200): d.line((x,0,x,359),fill=round(255*x/199))
    wide.paste(portrait,(500,0),mask)
    def data(image):
        buf=io.BytesIO();image.save(buf,format="JPEG",quality=83,optimize=True);b=buf.getvalue()
        if len(b)>200000: raise ValueError("Theme artwork too large")
        return "data:image/jpeg;base64,"+base64.b64encode(b).decode("ascii")
    return face,data(face),data(temple),data(wide)
CSS="""
:root{color-scheme:light;--bg:#d7d9dc;--panel:#f3f4f5;--line:#969ba2;--muted:#34383e;--cyan:#30363e;--blue:#242b32}
body,.shell{color:#111315;background:#d7d9dc}
aside{color:#111315;background:linear-gradient(180deg,#f0f1f2f7,#d3d6daed 76%,#c7cbd1e0),url('__TEMPLE__') center bottom/auto 70% no-repeat;border-color:#969ba2}
main{color:#111315;background:linear-gradient(180deg,#d9dde340,#d9dde365),url('__OLYMPUS__') center top/cover no-repeat;background-attachment:local}
.brand{color:#101214;letter-spacing:.7px}.brand img{width:44px;height:44px;object-fit:cover;border:1px solid #7e838a;border-radius:10px;flex:none}
.tag,.note,.sub,.label,.muted,.speed,.hero p,.entry p,footer{color:#34383e}
header{min-height:150px;padding:16px;align-items:flex-start;border:1px solid #969ba2;border-radius:14px;background:url('__OLYMPUS__') center/cover no-repeat}
header>div:first-child{background:#f2f3f4ed;border-radius:10px;padding:10px 14px;max-width:60%;box-shadow:0 2px 7px #0002}
header h1{color:#111315}header .actions{background:#e5e7e9ed;border-radius:10px;padding:6px}
button,input,select,textarea{color:#111315;background:linear-gradient(#fafbfc,#d9dce0);border:1px solid #8c9299}
button.primary,.nav button.active{color:#111315;background:linear-gradient(#e0e3e7,#b4bac2);border:1px solid #777e86}
button:disabled{opacity:.62}button:hover:not(:disabled){filter:brightness(.96)}
button:focus-visible,input:focus-visible,select:focus-visible,textarea:focus-visible{outline:3px solid #202830;outline-offset:2px}
.nav button{color:#111315;background:#f5f6f78c;border-color:#b7bbc1}
.card,.entry,dialog{background:#f3f4f5f5;color:#111315;border-color:#969ba2;box-shadow:0 2px 8px #0001}
.hero{background:linear-gradient(90deg,#eff1f3fa,#eff1f3ef 72%,#e2e5e9e8),url('__TEMPLE__') center/cover;color:#111315;border-color:#969ba2}
.pill,.media-field,.media-report{background:#e2e5e9;color:#111315;border-color:#969ba2}
.zeus-art,.mythic{background:#d7d9dc;border-color:#8a9097}
.ring svg circle:first-child{stroke:#a4aab1}.ring svg .arc,#curve{stroke:#30363e}.chart text{fill:#34383e}.chart svg path{stroke:#a3aab1}
.risk,#msg{color:#58250c}.z-authorship strong{color:#111315;border-color:#92989e}
[hidden]{display:none!important}
@media(max-width:720px){header{min-height:105px;padding:8px;flex-wrap:wrap}header>div:first-child{max-width:100%}.brand img{width:30px;height:30px}}
@media(forced-colors:active){main,aside,header,.hero{background:Canvas!important;color:CanvasText}button{forced-color-adjust:auto}}
"""
class Check(HTMLParser):
    def __init__(self): super().__init__();self.ids=[];self.pictures=[]
    def handle_starttag(self,tag,attrs):
        a=dict(attrs)
        if "id" in a: self.ids.append(a["id"])
        if tag=="img": self.pictures.append(a)
def apply():
    face,portrait,temple,wide=art()
    p=Path("desktop.html");s=p.read_text(encoding="utf-8-sig")
    s,n=re.subn(r'(<div class="brand">)<svg\b.*?</svg>',lambda m:m[1]+'<img alt="Zeus" src="'+portrait+'">',s,count=1,flags=re.S)
    if n!=1: raise ValueError("Brand anchor missing")
    css=CSS.replace("__TEMPLE__",temple).replace("__OLYMPUS__",wide)
    if s.count("</style></head>")!=1: raise ValueError("Style anchor missing")
    s=s.replace("</style></head>",css+"</style></head>",1)
    if len(s.encode("utf-8"))>1500000: raise ValueError("Embedded page exceeds budget")
    chk=Check();chk.feed(s)
    for name in ("zUpdateCheck","start","quick","stop","export","auditLoad","auditSave","auditState","zLocalAuditState"):
        if chk.ids.count(name)!=1: raise ValueError("Invalid control: "+name)
    if not any(x.get("alt")=="Zeus" and x.get("src","").startswith("data:image/jpeg;base64,") for x in chk.pictures): raise ValueError("Missing Zeus portrait")
    p.write_text(s,encoding="utf-8")
    face.save("ZEUVORYX.png")
    face.save("ZEUVORYX.ico",sizes=[(n,n) for n in (16,24,32,48,64,128,256)])
    print("Olympus silver: embedded scenery, portrait, icon and controls verified")
if __name__=="__main__": apply()
