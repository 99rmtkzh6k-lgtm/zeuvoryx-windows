from pathlib import Path
def once(s,a,b):
    if s.count(a)!=1:raise RuntimeError("314 alert anchor "+a[:100])
    return s.replace(a,b,1)
Path("expanded314_windows.go").write_text(Path("lab/expanded314_windows.go.txt").read_text(encoding="utf-8"),encoding="utf-8")
p=Path("desktop_windows.go");s=p.read_text(encoding="utf-8")
a=s.index('if !bind("behavior313"')
s=s[:a]+'''if !bind("recentSignature314",recentSignature314){return}
if !bind("openYouTubeVerify314",openYouTubeVerify314){return}
'''+s[a:]
p.write_text(s,encoding="utf-8")
p=Path("desktop.html");s=p.read_text(encoding="utf-8")
s=once(s,'<section id="media" class="hidden card">','<section id="media" class="hidden card"><div class="entry"><h2>Verificar vídeo do YouTube</h2><button id="z314Youtube">Abrir verificador oficial do YouTube</button><p class="sub">Cole o link do vídeo no serviço oficial. Sem sinais, o resultado é inconclusivo. Abre no navegador, sem integração direta ao player.</p></div>')
s=once(s,'<div id="z312SigDetails"></div>','<div id="z312SigDetails"></div><div class="entry"><h2>Alertas de assinatura durante a análise</h2><p id="z314LiveSummary" role="status">Nenhum sinal recebido nesta sessão.</p><p id="z315WindowsSummary" role="status">Microsoft Defender: aguardando consulta.</p><div id="z314LiveHits" role="status"></div></div>')
s=once(s,'</script></body></html>',"""document.getElementById('z314Youtube').onclick=()=>action(async()=>{await window.openYouTubeVerify314()});
;(()=>{let last=0;async function refresh(){try{const r=await window.recentSignature314();const d=el('z314LiveHits'),summary=el('z314LiveSummary');if(!r||!r.length){last=0;d.replaceChildren();summary.textContent='Nenhum sinal recebido nesta varredura; isso não comprova segurança.';return}if(r.at(-1).Sequence===last)return;last=r.at(-1).Sequence;d.replaceChildren();const limits=r.filter(x=>x.Finding.kind==='heuristic_or_limit').length,tests=r.filter(x=>x.Finding.kind==='test_eicar').length,detections=r.filter(x=>x.Finding.kind==='signature_detection').length;summary.textContent='Sinais recentes (máx. 100): '+detections+' detecções por assinatura · '+limits+' heurísticas/limites · '+tests+' testes.';for(const x of r.slice(-20).reverse()){const f=x.Finding;d.append(node('p',f.kind+' · '+f.signature+' · '+f.path+' · '+new Date(x.At).toLocaleTimeString()));}el('z310Notice').classList.remove('hidden');el('z310AlertText').textContent=r.at(-1).Finding.kind==='heuristic_or_limit'?'ClamAV: alerta heurístico ou limite; não confirma malware.':r.at(-1).Finding.kind==='test_eicar'?'ClamAV: arquivo de teste detectado.':'ClamAV: detecção por assinatura; revise antes de isolar.';}catch(e){el('z314LiveHits').textContent=String(e)}}setInterval(refresh,1000);})();
</script></body></html>""")
p.write_text(s,encoding="utf-8")
