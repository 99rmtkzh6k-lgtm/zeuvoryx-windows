"""Localize known UI messages without rewriting user data or detection logic."""
import html
import json
import re
from pathlib import Path

def once(s, a, b):
    if s.count(a) != 1:
        raise RuntimeError("Expected one anchor: "+a[:100])
    return s.replace(a,b,1)

def catalogs():
    out={"en-US":{}, "es-ES":{}}
    for line in Path("locales/ui.tsv").read_text(encoding="utf-8").splitlines():
        if not line or line.startswith("#"): continue
        fields=line.split("|")
        if len(fields)!=3 or not all(x.strip() for x in fields):
            raise ValueError("Invalid translation row")
        pt,en,es=fields
        if pt in out["en-US"]: raise ValueError("Duplicate translation")
        out["en-US"][pt]=en
        out["es-ES"][pt]=es
    return out

def transform(s, translations):
    s=once(s,'<html lang="pt-BR">','<html lang="pt-BR" data-z-mode="auto" data-z-lang="pt-BR">')
    s=once(s,'</nav>','</nav><label class="z-lang-label" for="zLanguage">Idioma</label><select id="zLanguage" aria-label="Idioma"><option value="auto">Automático (Windows)</option><option value="pt-BR">Português</option><option value="en-US">English</option><option value="es-ES">Español</option></select>')
    s=once(s,'</style></head>','.z-lang-label{display:block;margin:10px 0 4px;color:var(--muted);font-size:11px}#zLanguage{max-width:100%;width:100%;padding:6px;border-radius:7px;border:1px solid var(--line);background:var(--panel);color:#eef4f8;font:inherit;font-size:12px}</style></head>')
    s=once(s,'Motor atual: SHA-256, regras heurísticas e análise PE. Modelo treinado de malware: não instalado. Detector neural de vídeo: não instalado. C2PA: verificador de procedência digital separado.','Modelo experimental local: regressão logística treinada com dados sintéticos. Pontuação auxiliar apenas em alertas heurísticos; não altera detecções ou quarentena. Ainda não validado com amostras reais. Não aprende automaticamente.')
    s=once(s,'OLYMPUS 3.0.1 BETA</div>','OLYMPUS 3.0.2 BETA</div>')
    before,sep,js=s.partition('<script>')
    if not sep or js.count('</script>')!=1: raise ValueError("Unexpected script structure")
    # Translate only known, exact source-code string literals (never external values).
    for pt in sorted(translations["en-US"],key=len,reverse=True):
        if "'" in pt or '"' in pt or '\\' in pt: continue
        for q in ("'",'"'):
            literal=q+pt+q
            js=js.replace(literal,"zT("+literal+")")
    # Locale formatting is presentation only.
    js=js.replace("toLocaleString('pt-BR')","toLocaleString(zLocale)")
    raw=json.dumps(translations,ensure_ascii=False,separators=(",",":")).replace("<","\\u003c")
    init=r"""
const zCatalog=__CATALOG__;
let zLocale=document.documentElement.dataset.zLang || 'pt-BR';
function zT(s){return zCatalog[zLocale]?.[s] ?? s}
const zStatic=[];
function zCaptureStatic(){
  const walker=document.createTreeWalker(document.body,NodeFilter.SHOW_TEXT);
  let n;while(n=walker.nextNode()){
    if(n.parentElement.closest('script,style,textarea,pre'))continue;
    const k=n.nodeValue.trim();
    if(Object.prototype.hasOwnProperty.call(zCatalog['en-US'],k))zStatic.push([n,null,k,n.nodeValue]);
  }
  document.querySelectorAll('[placeholder],[aria-label],[alt]').forEach(e=>{
    for(const a of ['placeholder','aria-label','alt']){
      const k=e.getAttribute(a);
      if(k && Object.prototype.hasOwnProperty.call(zCatalog['en-US'],k))zStatic.push([e,a,k,k]);
    }
  });
}
function zPaintStatic(){
  document.documentElement.lang=zLocale;
  for(const [n,a,k,original] of zStatic){
    if(a)n.setAttribute(a,zT(k));
    else if(n.isConnected)n.nodeValue=original.replace(k,zT(k));
  }
}
zCaptureStatic();zPaintStatic();
const zSelect=document.getElementById('zLanguage');
zSelect.value=document.documentElement.dataset.zMode || 'auto';
zSelect.onchange=async()=>{
  const previous=document.documentElement.dataset.zMode || 'auto';
  zSelect.disabled=true;
  try{
    const result=await window.zSetLanguage(zSelect.value);
    zLocale=result;document.documentElement.dataset.zMode=zSelect.value;
    zPaintStatic();
    key='';nativeKey='';
    await refresh();await refreshNative();
    if(!el('quarantine').classList.contains('hidden'))await loadQ();
  }catch(e){zSelect.value=previous;document.getElementById('msg').textContent=String(e)}
  finally{zSelect.disabled=false}
};
""".replace("__CATALOG__",raw)
    js=once(js,"'use strict';","'use strict';\n"+init)
    return before+sep+js

def main():
    p=Path("desktop.html")
    s=transform(p.read_text(encoding="utf-8"),catalogs())
    p.write_text(s,encoding="utf-8")
    # The Windows preference is selected before the first WebView render.
    p=Path("desktop_windows.go")
    s=p.read_text(encoding="utf-8")
    s=once(s,'w.SetHtml(desktopHTML)','''mode := zReadLanguage()
    lang := zResolveLanguage(mode)
    if !bind("zSetLanguage", func(mode string) (string,error) {
        if err := zSaveLanguage(mode); err != nil { return "",err }
        return zResolveLanguage(mode),nil
    }) { return }
    localized := strings.Replace(desktopHTML, `data-z-mode="auto" data-z-lang="pt-BR"`, `data-z-mode="`+mode+`" data-z-lang="`+lang+`"`,1)
    w.SetHtml(localized)''')
    p.write_text(s,encoding="utf-8")
    print("Three-language interface integrated; unknown technical messages remain original")

if __name__=="__main__":
    main()
