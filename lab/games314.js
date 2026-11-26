
// First-party offline games. No network, no real currency, no privileged bridge in frames.
function zClassicGame(id,initialRecord){
 'use strict';
 const root=document.getElementById('game'),info=document.getElementById('info'),msg=document.getElementById('msg');
 const btn=(name,fn)=>{const b=document.createElement('button');b.textContent=name;b.type='button';b.onclick=()=>{if(!paused){try{fn()}catch(e){fail()}}};return b};
 let paused=false,stopped=false,cleanup=()=>{},draw=()=>{};
 const record=n=>{if(Number.isSafeInteger(n)&&n>=0&&n<=100000000&&n>initialRecord){initialRecord=n;parent.postMessage({type:'zeuvoryx-record',score:n},'*')}};
 function fail(){stopped=true;cleanup();root.replaceChildren();msg.textContent='Jogo indisponível, tente novamente';}
 window.addEventListener('error',fail);window.addEventListener('unhandledrejection',fail);
 window.addEventListener('message',e=>{if(e.source!==parent)return;if(e.data?.type==='zeuvoryx-pause'){paused=!!e.data.paused;document.getElementById('pause').hidden=!paused;}});
 function status(s){info.textContent=s+' · Recorde: '+initialRecord}
 function shuffle(a){for(let i=a.length-1;i>0;i--){const j=Math.floor(Math.random()*(i+1));[a[i],a[j]]=[a[j],a[i]]}return a}
 function grid(cols){const d=document.createElement('div');d.className='grid';d.style.setProperty('--cols',cols);root.append(d);return d}
 function start(){
 cleanup();cleanup=()=>{};root.replaceChildren();msg.textContent='';stopped=false;
 if(id==='zvx-tetris')tetris();else if(id==='zvx-poker')poker();else if(id==='zvx-mines')mines();else if(id==='zvx-memory')memory();else fail();
 }
 function tetris(){
 const shapes=[[[1,1,1,1]],[[1,1],[1,1]],[[0,1,0],[1,1,1]],[[1,0,0],[1,1,1]],[[0,0,1],[1,1,1]],[[0,1,1],[1,1,0]],[[1,1,0],[0,1,1]]], colors=['#172335','#37baf4','#ffe45d','#b47cff','#f49b39','#5b82ff','#5edf88','#ff6677'];
 const canvas=document.createElement('canvas');canvas.width=300;canvas.height=600;canvas.className='tetris';root.append(canvas);const c=canvas.getContext('2d');
 const panel=document.createElement('div');panel.className='controls';root.append(panel);
 const next=document.createElement('canvas');next.width=120;next.height=90;next.className='next';panel.append(next);const nc=next.getContext('2d');
 const nextLabel=document.createElement('div');nextLabel.textContent='Próximo bloco';panel.prepend(nextLabel);
 let b=Array.from({length:20},()=>Array(10).fill(0)),score=0,lines=0,level=1,bag=[],piece,upcoming=take(),last=0,elapsed=0,raf=0,manualPause=false;
 function take(){if(!bag.length)bag=shuffle([0,1,2,3,4,5,6]);return bag.pop()}
 function collision(p){return p.m.some((row,y)=>row.some((v,x)=>v&&(p.x+x<0||p.x+x>=10||p.y+y>=20||(p.y+y>=0&&b[p.y+y][p.x+x]))))}
 function spawn(){const k=upcoming;upcoming=take();piece={m:shapes[k].map(r=>r.slice()),x:3,y:0,k:k+1};if(collision(piece)){stopped=true;record(score);msg.textContent='Game Over';}}
 function paint(){c.fillStyle=colors[0];c.fillRect(0,0,300,600);const square=(x,y,k)=>{c.fillStyle=colors[k];c.fillRect(x*30+1,y*30+1,28,28);c.strokeStyle='#ffffff18';c.strokeRect(x*30+1,y*30+1,28,28)};b.forEach((r,y)=>r.forEach((v,x)=>{if(v)square(x,y,v)}));if(piece)piece.m.forEach((r,y)=>r.forEach((v,x)=>{if(v&&piece.y+y>=0)square(piece.x+x,piece.y+y,piece.k)}));nc.fillStyle=colors[0];nc.fillRect(0,0,120,90);nc.fillStyle=colors[upcoming+1];shapes[upcoming].forEach((r,y)=>r.forEach((v,x)=>{if(v)nc.fillRect(x*25+4,y*25+4,23,23)}));status('Pontos: '+score+' · Nível: '+level+' · Linhas: '+lines+(manualPause?' · Pausado':''));}
 function move(dx){if(stopped||manualPause)return;const p={...piece,x:piece.x+dx};if(!collision(p))piece=p;paint()}
 function rotate(){if(stopped||manualPause)return;const m=piece.m[0].map((_,x)=>piece.m.map(r=>r[x]).reverse());for(const dx of [0,-1,1,-2,2]){const p={...piece,m,x:piece.x+dx};if(!collision(p)){piece=p;break}}paint()}
 function down(soft){if(stopped||manualPause)return;const p={...piece,y:piece.y+1};if(!collision(p)){piece=p;if(soft)score++;}else{piece.m.forEach((r,y)=>r.forEach((v,x)=>{if(v)b[piece.y+y][piece.x+x]=piece.k}));const keep=b.filter(r=>r.some(v=>!v));const n=20-keep.length;b=Array.from({length:n},()=>Array(10).fill(0)).concat(keep);score+=[0,100,300,500,800][n]*level;lines+=n;level=1+Math.floor(lines/10);record(score);spawn()}paint()}
 function pause(){if(!stopped){manualPause=!manualPause;paint()}}
 function key(e){if(paused)return;const k=e.key.toLowerCase();if(!['arrowleft','arrowright','arrowdown','arrowup',' ','p'].includes(k))return;e.preventDefault();if(k==='p')pause();else if(k==='arrowleft')move(-1);else if(k==='arrowright')move(1);else if(k==='arrowdown')down(true);else rotate()}
 panel.append(btn('←',()=>move(-1)),btn('→',()=>move(1)),btn('Girar / Espaço',rotate),btn('↓',()=>down(true)),btn('Pausar / P',pause));
 function tick(t){const dt=last?Math.min(t-last,100):0;last=t;if(!paused&&!manualPause&&!stopped){elapsed+=dt;if(elapsed>=Math.max(80,800-(level-1)*65)){elapsed=0;down(false)}}if(!stopped)raf=requestAnimationFrame(tick)}
 document.addEventListener('keydown',key);cleanup=()=>{cancelAnimationFrame(raf);document.removeEventListener('keydown',key)};spawn();paint();raf=requestAnimationFrame(tick);
 }
 function mines(){
 let cells=Array.from({length:81},()=>({mine:false,open:false,flag:false,n:0})),started=false,opened=0,flags=0,seconds=0;const g=grid(9);g.classList.add('mines');const bs=[];
 const near=i=>{let a=[];for(let dy=-1;dy<=1;dy++)for(let dx=-1;dx<=1;dx++){let y=Math.floor(i/9)+dy,x=i%9+dx;if((dx||dy)&&y>=0&&y<9&&x>=0&&x<9)a.push(y*9+x)}return a};
 function seed(i){for(const k of shuffle(cells.map((_,j)=>j).filter(j=>j!==i)).slice(0,10))cells[k].mine=true;cells.forEach((c,j)=>c.n=near(j).filter(k=>cells[k].mine).length);started=true}
 function paint(){cells.forEach((c,i)=>{const b=bs[i];b.textContent=c.open?(c.mine?'💣':c.n||''):c.flag?'⚑':'·';b.classList.toggle('revealed',c.open);b.setAttribute('aria-label','Casa '+(i+1)+(c.flag?' marcada':''));b.disabled=stopped||c.open});status('Bombas restantes: '+(10-flags)+' · Tempo: '+seconds+' s')}
 function reveal(i){if(stopped||cells[i].flag||cells[i].open)return;if(!started)seed(i);if(cells[i].mine){cells.forEach(c=>{if(c.mine)c.open=true});stopped=true;msg.textContent='Derrota';paint();return}const q=[i];while(q.length){const k=q.pop(),c=cells[k];if(c.open||c.flag||c.mine)continue;c.open=true;opened++;if(!c.n)q.push(...near(k))}if(opened===71){stopped=true;record(Math.max(1,100000-seconds));msg.textContent='Vitória!'}paint()}
 for(let i=0;i<81;i++){const b=btn('·',()=>reveal(i));b.oncontextmenu=e=>{e.preventDefault();if(paused||stopped||cells[i].open)return;cells[i].flag=!cells[i].flag;flags+=cells[i].flag?1:-1;paint()};bs.push(b);g.append(b)}
 const timer=setInterval(()=>{if(started&&!paused&&!stopped){seconds++;paint()}},1000);cleanup=()=>clearInterval(timer);paint();
 }
 function memory(){
 const symbols=shuffle(['☀','☾','♣','♦','★','♥','⚓','♫'].flatMap(s=>[s,s]));const g=grid(4);g.classList.add('memory');let first=-1,second=-1,attempts=0,matched=new Set(),wait=0;const bs=[];
 function paint(){bs.forEach((b,i)=>{b.textContent=matched.has(i)||i===first||i===second?symbols[i]:'?';b.disabled=matched.has(i)||i===first||second>=0||stopped});status('Tentativas: '+attempts+' · Pares: '+matched.size/2+'/8')}
 function choose(i){if(stopped||matched.has(i)||i===first||second>=0)return;if(first<0){first=i;paint();return}second=i;attempts++;paint();if(symbols[first]===symbols[second]){matched.add(first);matched.add(second);first=second=-1;if(matched.size===16){stopped=true;record(Math.max(1,10000-attempts));msg.textContent='Vitória!'}paint()}else{let left=8;wait=setInterval(()=>{if(paused)return;if(--left<=0){clearInterval(wait);first=second=-1;paint()}},100)}}
 symbols.forEach((_,i)=>{const b=btn('?',()=>choose(i));bs.push(b);g.append(b)});cleanup=()=>clearInterval(wait);paint();
 }
 function poker(){
 const ranks=['','', '2','3','4','5','6','7','8','9','10','J','Q','K','A'],names=['','','Dois','Três','Quatro','Cinco','Seis','Sete','Oito','Nove','Dez','Valetes','Damas','Reis','Ases'],suits=['♠','♥','♦','♣'];
 const cmp=(a,b)=>{for(let i=0;i<Math.max(a.length,b.length);i++){const d=(a[i]||0)-(b[i]||0);if(d)return d}return 0};
 function five(cs){let r=cs.map(c=>c.r).sort((a,b)=>b-a),u=[...new Set(r)],groups=u.map(v=>[r.filter(x=>x===v).length,v]).sort((a,b)=>b[0]-a[0]||b[1]-a[1]),flush=cs.every(c=>c.s===cs[0].s),straight=u.length===5?(u[0]-u[4]===4?u[0]:u.join(',')==='14,5,4,3,2'?5:0):0;
 if(flush&&straight)return [8,straight];if(groups[0][0]===4)return [7,groups[0][1],groups[1][1]];if(groups[0][0]===3&&groups[1][0]===2)return [6,groups[0][1],groups[1][1]];if(flush)return [5,...r];if(straight)return [4,straight];if(groups[0][0]===3)return [3,...groups.map(g=>g[1])];if(groups[0][0]===2&&groups[1][0]===2)return [2,...groups.map(g=>g[1])];if(groups[0][0]===2)return [1,...groups.map(g=>g[1])];return [0,...r]}
 function best(cs){let a=[-1];for(let i=0;i<cs.length;i++)for(let j=i+1;j<cs.length;j++){const v=five(cs.filter((_,k)=>k!==i&&k!==j));if(cmp(v,a)>0)a=v}return a}
 const label=v=>['Carta alta '+ranks[v[1]],'Par de '+names[v[1]],'Dois pares','Trinca de '+names[v[1]],'Sequência','Flush','Full house','Quadra','Straight flush'][v[0]];
 let chips=[1000,1000,1000,1000],hands,table,pot,street,finished,folded,round=0;const titles=['Você','Bot 1','Bot 2','Bot 3'],view=document.createElement('div');view.className='poker';root.append(view);const actions=document.createElement('div');actions.className='actions';root.append(actions);
 function cards(cs,hide){const d=document.createElement('div');d.className='cards';for(const c of cs){const span=document.createElement('span');span.className='playing'+(!hide&&(c.s==='♥'||c.s==='♦')?' red':'');span.textContent=hide?'▧':ranks[c.r]+c.s;d.append(span)}return d}
 function paint(){view.replaceChildren();const players=document.createElement('div');players.className='players';for(let i=0;i<4;i++){const x=document.createElement('div');const h=document.createElement('strong');h.textContent=titles[i]+' · '+chips[i]+' fichas';x.append(h,cards(hands[i],i!==0&&!finished));players.append(x)}view.append(players,cards(table.slice(0,[0,3,4,5][street]),false));status('Rodada '+round+' · Pote '+pot+' · '+['Pré-flop','Flop','Turn','River'][street]+' · Fichas fictícias');actions.replaceChildren();if(finished){actions.append(btn('Nova Rodada',deal));return}actions.append(btn('Check',()=>step(false)),btn('Apostar 25',()=>step(true)),btn('Desistir',()=>{folded=true;street=3;settle()}))}
 function deal(){if(chips.some(c=>c===0)){chips=[1000,1000,1000,1000];msg.textContent='Nova partida: fichas fictícias reiniciadas'}else msg.textContent='';const deck=shuffle(suits.flatMap(s=>Array.from({length:13},(_,i)=>({s,r:i+2}))));hands=Array.from({length:4},()=>[deck.pop(),deck.pop()]);table=Array.from({length:5},()=>deck.pop());pot=0;street=0;finished=false;folded=false;round++;for(let i=0;i<4;i++){const n=Math.min(10,chips[i]);chips[i]-=n;pot+=n}paint()}
 function step(bet){if(finished)return;if(bet){const n=Math.min(25,...chips);if(n===0){msg.textContent='Sem fichas para nova aposta; use Check.';return}for(let i=0;i<4;i++){chips[i]-=n;pot+=n}}if(street===3)settle();else{street++;paint()}}
 function settle(){finished=true;street=3;const scores=hands.map(h=>best(h.concat(table)));let winners=[];for(let i=folded?1:0;i<4;i++){if(!winners.length||cmp(scores[i],scores[winners[0]])>0)winners=[i];else if(cmp(scores[i],scores[winners[0]])===0)winners.push(i)}const share=Math.floor(pot/winners.length),rem=pot%winners.length;winners.forEach((i,j)=>chips[i]+=share+(j===0?rem:0));pot=0;record(chips[0]);msg.textContent=winners.map(i=>titles[i]).join(' e ')+' venceu'+(winners.length>1?'ram':'')+' — '+label(scores[winners[0]]);paint()}
 deal();
 }
 document.getElementById('restart').onclick=()=>{if(!paused){try{start()}catch(e){fail()}}};try{start()}catch(e){fail()}
}
function zClassicDoc(id,score){
 const style='*{box-sizing:border-box}html,body{margin:0;background:#101a27;color:#f3f7ff;font:16px system-ui,sans-serif}body{padding:12px}h2,p{margin:6px 0}button{font:inherit;padding:8px 12px;background:#1d3d52;color:white;border:1px solid #8ba9c4;border-radius:7px;cursor:pointer}button:disabled{opacity:.65;cursor:default}#game{display:flex;gap:12px;flex-wrap:wrap;align-items:flex-start;justify-content:center;margin:10px auto}#info,#msg{text-align:center;font-weight:600}.tetris{width:min(300px,46vw,calc((100vh - 120px)/2));height:auto;aspect-ratio:1/2;border:1px solid #7894a8}.next{width:120px;height:90px}.controls{display:flex;flex-direction:column;gap:6px}.grid{display:grid;grid-template-columns:repeat(var(--cols),minmax(0,1fr));gap:4px;width:min(100%,440px);margin:auto}.grid button{padding:0;aspect-ratio:1;font-size:clamp(12px,3vw,24px);min-width:0}.grid .revealed{background:#496174}.memory{max-width:320px}.poker{width:100%;max-width:740px;background:#12583d;border:2px solid #ad9767;border-radius:14px;padding:12px}.players{display:grid;grid-template-columns:repeat(2,minmax(0,1fr));gap:12px}.players>div{text-align:center;min-width:0}.cards{display:flex;flex-wrap:wrap;gap:5px;justify-content:center;margin:8px auto}.playing{display:flex;align-items:center;justify-content:center;background:#fff;color:#172335;border-radius:6px;border:1px solid #aaa;min-width:38px;width:clamp(38px,9vw,72px);height:clamp(52px,12vw,98px);font-size:clamp(16px,3vw,30px);font-weight:700}.red{color:#b41d24}.actions{display:flex;gap:8px;flex-wrap:wrap;justify-content:center}#pause{padding:12px;background:#ac6a15;color:white;text-align:center}#pause[hidden]{display:none}@media(max-width:480px){body{padding:6px}.controls button{padding:5px}.players{gap:5px}.poker{padding:6px}}';
 return '<!doctype html><html lang="pt-BR"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><meta http-equiv="Content-Security-Policy" content="default-src &#39;none&#39;;script-src &#39;unsafe-inline&#39;;style-src &#39;unsafe-inline&#39;;connect-src &#39;none&#39;;base-uri &#39;none&#39;;form-action &#39;none&#39;"><style>'+style+'</style></head><body><div id="pause" hidden>Jogo pausado pela varredura</div><p id="info"></p><div id="game"></div><p id="msg" role="status"></p><div class="actions"><button id="restart">Jogar novamente</button></div><script>('+zClassicGame.toString()+')('+JSON.stringify(id)+','+Number(score||0)+');<\/script></body></html>';
}

;(()=>{
const classics=[['zvx-tetris','Clássicos','Tetris — blocos','Setas, espaço para girar, P para pausar.'],['zvx-poker','Clássicos','Poker Texas Hold’em','Mesa com três bots, apenas fichas fictícias, sem dinheiro.'],['zvx-mines','Clássicos','Campo minado','9 × 9, 10 bombas. Clique direito marca bandeira.'],['zvx-memory','Clássicos','Jogo da memória','4 × 4, oito pares.']];
const own=id=>classics.some(g=>g[0]===id);
const games=[...classics,['2048','Quebra-cabeça','2048','Combine os blocos numerados. Jogo de Gabriele Cirulli, licença MIT.'],...zRetroCatalog.map(g=>[g.id,g.category,g.title,g.description])];
const z=id=>document.getElementById(id),panel=z('zGames314'),nav=z('zGamesNav314'),status=z('zGameStatus'),board=z('zGameBoard');
let gameTimer=null,visible=false,querying=false,launch=0,currentGame=null,scanBusy=false;
function blockGame(busy){scanBusy=busy;const frame=board.querySelector('iframe');if(frame&&currentGame){if(own(currentGame)){frame.contentWindow.postMessage({type:'zeuvoryx-pause',paused:busy},'*');frame.style.pointerEvents=busy?'none':'';}else if(busy){closeGame();status.textContent='Partida encerrada para priorizar a varredura.'}}}
window.addEventListener('message',e=>{const frame=board.querySelector('iframe');if(!frame||e.source!==frame.contentWindow||!own(currentGame)||e.data?.type!=='zeuvoryx-record'||!Number.isSafeInteger(e.data.score))return;window.saveGameRecord315(currentGame,e.data.score).catch(()=>{status.textContent='Não foi possível salvar o recorde.'})});
for(const id of ['start','quick']){const b=z(id);if(b)b.addEventListener('click',()=>blockGame(true),true)}
function closeGame(){z('zGameLocal').classList.remove('zGameFull315');z('zGameFullscreen315').textContent='Tela Cheia';launch++;currentGame=null;board.replaceChildren();z('zGameLocal').hidden=true}
function close(){if(gameTimer!==null){clearInterval(gameTimer);gameTimer=null}visible=false;closeGame();panel.classList.add('hidden');nav.classList.remove('active');nav.setAttribute('aria-expanded','false')}
nav.onclick=()=>{visible=true;if(gameTimer!==null)clearInterval(gameTimer);refreshGameScan();gameTimer=setInterval(refreshGameScan,350);for(const b of document.querySelectorAll('[data-page]')){b.classList.remove('active');const p=z(b.dataset.page);if(p)p.classList.add('hidden')}panel.classList.remove('hidden');nav.classList.add('active');nav.setAttribute('aria-expanded','true');render()};
for(const b of document.querySelectorAll('[data-page]'))b.addEventListener('click',close);
z('zGameBack').onclick=()=>{close();document.querySelector('[data-page="overview"]').click()};
z('zGameAlerts314').onclick=()=>{close();document.querySelector('[data-page="alerts"]').click()};
z('zFamilyNav314').addEventListener('click',close);
async function refreshGameScan(){if(!visible||querying)return;querying=true;try{const v=await window.getScanState(),r=v.report||{};if(!visible)return;blockGame(!!v.running);const count=Number(v.files??r.files??0),bytes=Number(v.bytes??r.bytes??0),hits=Number(v.finding_total??0);z('zGameScanState314').textContent=v.running?'Varredura em andamento · '+(r.signature_state||'motor de assinaturas aguardando'):(v.done?'Varredura encerrada':'Nenhuma varredura ativa');z('zGameScanMetrics314').textContent='Arquivos: '+count.toLocaleString('pt-BR')+' · Volume: '+volume(bytes)+' · Indícios do scanner: '+hits;}catch(e){if(visible){z('zGameScanState314').textContent='Progresso indisponível';z('zGameScanMetrics314').textContent=''}}finally{querying=false}}
function render(){if(!visible)return;const f=z('zGameFilter').value,q=z('zGameSearch').value.trim().toLocaleLowerCase('pt-BR'),out=z('zGameCatalog');out.replaceChildren();let n=0;for(const [id,category,title,desc] of games){if((f&&f!==category)||(q&&!(title+' '+category+' '+desc).toLocaleLowerCase('pt-BR').includes(q)))continue;n++;const tile=document.createElement('article');tile.className='zGameTile';tile.dataset.platform=category;const b=node('button','Jogar aqui');b.type='button';b.onclick=()=>start(id);tile.append(node('small',category+' · offline'),node('strong',title),node('small',desc),b);out.append(tile)}status.textContent=n+' jogos locais no filtro. Nenhum abre o navegador.'}
async function start(id){
 const game=games.find(entry=>entry[0]===id);
 if(!game||!visible)return;
 try{const v=await window.getScanState();if(v.running){status.textContent='Aguarde o fim da varredura para jogar.';return}}catch(e){status.textContent='Jogo indisponível, tente novamente';return}
 const ticket=++launch;
 try{if(await window.allowLocalGame315(id)!==true)throw Error('Acesso não autorizado')}catch(e){if(ticket===launch)status.textContent='Jogo indisponível: '+String(e);return}
 if(!visible||ticket!==launch)return;
 let score=0;try{if(own(id))score=await window.getGameRecord315(id);const v=await window.getScanState();if(v.running){status.textContent='Aguarde o fim da varredura para jogar.';return}}catch(e){status.textContent='Jogo indisponível, tente novamente';return}
 if(!visible||ticket!==launch)return;
 closeGame();
 const frame=document.createElement('iframe');
 frame.title=game[2];
 frame.setAttribute('referrerpolicy','no-referrer');
 frame.setAttribute('sandbox','allow-scripts');
 frame.style.cssText='width:100%;height:clamp(370px,70vh,760px);border:0;background:#101a27;grid-column:1/-1';
 frame.addEventListener('load',()=>{if(own(id)&&scanBusy)frame.contentWindow.postMessage({type:'zeuvoryx-pause',paused:true},'*')});
 frame.onerror=()=>{closeGame();status.textContent='Jogo indisponível, tente novamente'};
 frame.srcdoc=own(id)?zClassicDoc(id,score):id==='2048'?zArcade2048:zRetroDocs[id];
 currentGame=id;
 z('zGameTitle').textContent=game[2];
 z('zGameInstructions').textContent=game[3];
 z('zGameLocal').hidden=false;
 board.append(frame);
 z('zGameLocal').scrollIntoView({block:'nearest'});
}
z('zGameFilter').onchange=render;
z('zGameSearch').oninput=render;
z('zGameClose315').onclick=closeGame;
z('zGameFullscreen315').onclick=()=>{const e=z('zGameLocal');e.classList.toggle('zGameFull315');z('zGameFullscreen315').textContent=e.classList.contains('zGameFull315')?'Sair da tela cheia':'Tela Cheia'};
z('zGameRestart21').onclick=()=>{if(currentGame)start(currentGame)};
})();
