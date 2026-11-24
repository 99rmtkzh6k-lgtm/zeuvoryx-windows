"""Commercial-free candidate: block unapproved remote feed at the UI backend.
Apply after source generators. Preserve authorized local import and CISA lookup.
Do not change current 3.0.6 binary or distribute a reduced-coverage build.
"""
from pathlib import Path
import re

def transform_go(s):
    pattern=r'(if\s+!bind\("syncThreatHashes",\s*func\(key string\)\s*\(map\[string\]any,\s*error\)\s*\{).*?(\}\)\s*\{return\})'
    s,n=re.subn(pattern,lambda m:m.group(1)+'\n return nil, fmt.Errorf("ZVX-LICENSE-001: consulta externa indisponível; licença comercial não verificada")\n '+m.group(2),s,count=1,flags=re.S)
    if n!=1 or 'bind("importThreatHashes"' not in s or 'bind("checkKnownCVE"' not in s:
        raise ValueError('unexpected intelligence binding structure')
    return s

def transform_html(s):
    anchor="</script></body></html>"
    if s.count(anchor)!=1 or 'zFreeFeedPolicy' in s:
        raise ValueError('unexpected desktop structure')
    payload=r"""
;(function zFreeFeedPolicy(){
const key=document.getElementById('intelKey'),button=document.getElementById('intelSync'),out=document.getElementById('intelResult');
if(!key||!button||!out)return;
key.value='';key.disabled=true;button.disabled=true;button.onclick=null;
const messages={'pt-BR':'Consulta externa indisponível: licença comercial não verificada. Importe somente indicadores autorizados.','en-US':'External lookup unavailable: commercial license not verified. Import only authorized indicators.','es-ES':'Consulta externa no disponible: licencia comercial no verificada. Importe solo indicadores autorizados.'};
function update(){out.textContent=messages[document.documentElement.lang]||messages['en-US'];}
update();
const select=document.getElementById('zLanguage');
if(select)select.addEventListener('change',()=>setTimeout(update,150));
})();
"""
    return s.replace(anchor,payload+anchor,1)

if __name__=='__main__':
    go=Path('desktop_windows.go')
    html=Path('desktop.html')
    g=transform_go(go.read_text(encoding='utf-8-sig'))
    h=transform_html(html.read_text(encoding='utf-8-sig'))
    go.write_text(g,encoding='utf-8')
    html.write_text(h,encoding='utf-8')
