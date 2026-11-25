"""Add truthful base coverage warning to generated desktop. Apply after final UI generation."""
from pathlib import Path
def transform(s):
    anchor="</script></body></html>"
    if s.count(anchor)!=1 or "zSeedCoverage" in s: raise ValueError("unexpected UI structure")
    return s.replace(anchor,r"""
;(function zSeedCoverage(){
const panel=document.getElementById('overview');if(!panel)return;
const out=document.createElement('p');out.className='sub';out.setAttribute('role','status');
panel.append(out);
const copy={
'pt-BR':['Assinaturas SHA-256 ativas: ',' | Base histórica/limitada; zero alertas não comprova segurança.','Base vazia ou indisponível: cobertura por assinaturas ausente. Mantenha o Defender ativo.'],
'en-US':['Active SHA-256 indicators: ',' | Historical/limited base; zero alerts does not prove safety.','Empty or unavailable base: no signature coverage. Keep Defender enabled.'],
'es-ES':['Indicadores SHA-256 activos: ',' | Base histórica/limitada; cero alertas no demuestra seguridad.','Base vacía o no disponible: sin cobertura de firmas. Mantenga Defender activo.']};
let count=null;
function paint(){const t=copy[document.documentElement.lang]||copy['en-US'];out.textContent=count>0?t[0]+count+t[1]:t[2]}
async function refresh(){try{const r=await window.intelStatus();count=Number(r.total)||0}catch(_){count=null}paint()}
refresh();new MutationObserver(paint).observe(document.documentElement,{attributes:true,attributeFilter:['lang']});
const b=document.getElementById('intelCount');if(b)b.addEventListener('click',refresh);
})();
"""+anchor,1)
if __name__=="__main__":
    p=Path("desktop.html")
    p.write_text(transform(p.read_text(encoding="utf-8")),encoding="utf-8")
