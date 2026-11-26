;(()=>{
const games=[['2048','Quebra-cabeça','2048','Combine os blocos numerados. Jogo de Gabriele Cirulli, licença MIT.']];
const z=id=>document.getElementById(id),panel=z('zGames314'),nav=z('zGamesNav314'),status=z('zGameStatus'),board=z('zGameBoard');
let gameTimer=null,visible=false,querying=false,launch=0;
function closeGame(){launch++;board.replaceChildren();z('zGameLocal').hidden=true}
function close(){if(gameTimer!==null){clearInterval(gameTimer);gameTimer=null}visible=false;closeGame();panel.classList.add('hidden');nav.classList.remove('active');nav.setAttribute('aria-expanded','false')}
nav.onclick=()=>{visible=true;if(gameTimer!==null)clearInterval(gameTimer);refreshGameScan();gameTimer=setInterval(refreshGameScan,3000);for(const b of document.querySelectorAll('[data-page]')){b.classList.remove('active');const p=z(b.dataset.page);if(p)p.classList.add('hidden')}panel.classList.remove('hidden');nav.classList.add('active');nav.setAttribute('aria-expanded','true');render()};
for(const b of document.querySelectorAll('[data-page]'))b.addEventListener('click',close);
z('zGameBack').onclick=()=>{close();document.querySelector('[data-page="overview"]').click()};
z('zGameAlerts314').onclick=()=>{close();document.querySelector('[data-page="alerts"]').click()};
z('zFamilyNav314').addEventListener('click',close);
async function refreshGameScan(){if(!visible||querying)return;querying=true;try{const v=await window.getScanState(),r=v.report||{};if(!visible)return;const count=Number(v.files??r.files??0),bytes=Number(v.bytes??r.bytes??0),hits=Number(v.finding_total??0);z('zGameScanState314').textContent=v.running?'Varredura em andamento · '+(r.signature_state||'motor de assinaturas aguardando'):(v.done?'Varredura encerrada':'Nenhuma varredura ativa');z('zGameScanMetrics314').textContent='Arquivos: '+count.toLocaleString('pt-BR')+' · Volume: '+volume(bytes)+' · Indícios do scanner: '+hits;}catch(e){if(visible){z('zGameScanState314').textContent='Progresso indisponível';z('zGameScanMetrics314').textContent=''}}finally{querying=false}}
function render(){if(!visible)return;const f=z('zGameFilter').value,q=z('zGameSearch').value.trim().toLocaleLowerCase('pt-BR'),out=z('zGameCatalog');out.replaceChildren();let n=0;for(const [id,category,title,desc] of games){if((f&&f!==category)||(q&&!(title+' '+category+' '+desc).toLocaleLowerCase('pt-BR').includes(q)))continue;n++;const tile=document.createElement('article');tile.className='zGameTile';tile.dataset.platform=category;const b=node('button','Jogar aqui');b.type='button';b.onclick=()=>start(id);tile.append(node('small',category+' · offline'),node('strong',title),node('small',desc),b);out.append(tile)}status.textContent=n+' jogos locais no filtro. Nenhum abre o navegador.'}
async function start(id){
 const game=games.find(entry=>entry[0]===id);
 if(!game||!visible)return;
 const ticket=++launch;
 try{await window.allowLocalGame315(id)}catch(e){if(ticket===launch)status.textContent='Jogo indisponível: '+String(e);return}
 if(!visible||ticket!==launch)return;
 closeGame();
 const frame=document.createElement('iframe');
 frame.title='2048 — Gabriele Cirulli';
 frame.setAttribute('sandbox','allow-scripts');
 frame.style.cssText='width:100%;height:680px;border:0;background:#faf8ef;grid-column:1/-1';
 frame.srcdoc=zArcade2048;
 z('zGameTitle').textContent=game[2];
 z('zGameInstructions').textContent=game[3];
 z('zGameLocal').hidden=false;
 board.append(frame);
 z('zGameLocal').scrollIntoView({block:'nearest'});
}
z('zGameFilter').onchange=render;
z('zGameSearch').oninput=render;
z('zGameClose315').onclick=closeGame;
})();
