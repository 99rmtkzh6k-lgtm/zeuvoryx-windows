;const ZGameWindow314=(()=>{
const error="Jogo sem GameWindow - correção obrigatória";
const command=e=>{const k=e.key;if(k==='F11'||(e.altKey&&k==='Enter'))return 'fullscreen';if(e.ctrlKey&&!e.altKey){if(k==='+'||k==='='||e.code==='NumpadAdd')return 'plus';if(k==='-'||e.code==='NumpadSubtract')return 'minus';if(k==='0')return 'reset'}return null};
class GameWindow{
 constructor(panel,board,status){
  this.panel=panel;this.board=board;this.status=status;this.local=document.getElementById('zGameLocal');this.dock=document.getElementById('zGameDock');this.catalog=new Map();this.frame=null;this.id=null;this.mode='closed';this.previous='normal';this.fullscreen=false;this.zoom=1;this.generation=0;
  const z=id=>document.getElementById(id);
  z('zGameMin').onclick=()=>this.minimize();z('zGameMax').onclick=()=>this.maximize();z('zGameFull').onclick=()=>this.toggleFullscreen();
  z('zGameZoomPlus').onclick=()=>this.apply('plus');z('zGameZoomMinus').onclick=()=>this.apply('minus');z('zGameZoomReset').onclick=()=>this.apply('reset');z('zGameDockRestore').onclick=()=>this.restore();
  this.local.addEventListener('keydown',e=>{const c=command(e);if(this.frame&&c){e.preventDefault();e.stopPropagation();this.apply(c)}},true);
  window.addEventListener('message',e=>{if(!this.frame||this.mode==='minimized'||e.source!==this.frame.contentWindow||e.origin!=='null')return;const d=e.data;if(d&&d.channel==='zGameWindow314'&&['fullscreen','plus','minus','reset'].includes(d.command))this.apply(d.command)});
  document.addEventListener('fullscreenchange',()=>{this.fullscreen=document.fullscreenElement===this.local;z('zGameFull').setAttribute('aria-pressed',String(this.fullscreen));z('zGameFull').textContent=this.fullscreen?'⛶ Sair da tela cheia':'⛶ Tela cheia';if(this.frame&&this.mode!=='minimized')this.focus()});
 }
 registerGame(game,doc){
  if(!Array.isArray(game)||!game[0]||this.catalog.has(game[0])||typeof doc!=='string'||!/<\/head\s*>/i.test(doc)||!doc.includes("connect-src 'none'"))throw Error(error);
  this.catalog.set(game[0],{game,doc});
 }
 bridge(doc){
  const js=';(()=>{const command='+command.toString()+';document.addEventListener("keydown",e=>{const c=command(e);if(!c)return;e.preventDefault();e.stopImmediatePropagation();parent.postMessage({channel:"zGameWindow314",command:c},"*")},true);window.addEventListener("load",()=>{if(document.body&&document.activeElement===document.body){document.body.tabIndex=-1;document.body.focus({preventScroll:true})}},{once:true})})();';
  return doc.replace(/<\/head\s*>/i,()=>'<script>'+js+'<'+'/script></head>');
 }
 setMode(mode){
  this.mode=mode;this.panel.dataset.window=mode;this.local.hidden=mode==='closed'||mode==='minimized';this.dock.hidden=mode!=='minimized';
  const b=document.getElementById('zGameMax');b.textContent=mode==='maximized'?'□ Restaurar':'□ Maximizar';b.setAttribute('aria-pressed',String(mode==='maximized'));
 }
 open(id){
  const item=this.catalog.get(id);if(!item)throw Error(error);
  this.close();this.id=id;this.zoom=1;const frame=document.createElement('iframe');
  frame.title=item.game[2];frame.tabIndex=0;frame.setAttribute('referrerpolicy','no-referrer');frame.setAttribute('sandbox','allow-scripts');frame.setAttribute('allow','fullscreen');frame.setAttribute('allowfullscreen','');
  frame.addEventListener('load',()=>{if(this.frame===frame&&this.mode!=='minimized')this.focus()});
  frame.srcdoc=this.bridge(item.doc);this.frame=frame;
  document.getElementById('zGameTitle').textContent=item.game[2];document.getElementById('zGameInstructions').textContent=item.game[3];document.getElementById('zGameDockRestore').textContent='Restaurar '+item.game[2];
  this.board.replaceChildren(frame);this.setMode('normal');this.setZoom(1);this.local.scrollIntoView({block:'start'});this.focus();
 }
 focus(){if(!this.frame||this.mode==='minimized'||this.local.hidden)return;try{this.frame.focus({preventScroll:true});this.frame.contentWindow.focus()}catch(_){}}
 setZoom(v){this.zoom=Math.max(.5,Math.min(2,Math.round(v*10)/10));if(this.frame)this.frame.style.transform='scale('+this.zoom+')';document.getElementById('zGameZoomReset').textContent=Math.round(this.zoom*100)+'%'}
 apply(c){if(!this.frame||this.mode==='minimized')return;if(c==='fullscreen'){void this.toggleFullscreen();return}this.setZoom(c==='reset'?1:this.zoom+(c==='plus'?.1:-.1));this.focus()}
 maximize(){if(!this.frame)return;this.setMode(this.mode==='maximized'?'normal':'maximized');this.focus()}
 async toggleFullscreen(){if(!this.frame||this.mode==='minimized')return;try{if(document.fullscreenElement===this.local)await document.exitFullscreen();else if(this.local.requestFullscreen)await this.local.requestFullscreen();else throw Error('Fullscreen indisponível neste WebView2')}catch(e){this.status.textContent='Tela cheia indisponível: '+String(e);document.getElementById('zGameTurn').textContent=this.status.textContent}this.focus()}
 minimize(){if(!this.frame||this.mode==='minimized')return;this.previous=this.mode;if(document.fullscreenElement===this.local)void document.exitFullscreen().catch(()=>{});this.setMode('minimized');document.getElementById('zGameDockRestore').focus()}
 async restore(){if(!this.frame||this.mode!=='minimized')return;const ticket=this.generation,id=this.id;try{if(await window.allowLocalGame315(id)!==true)throw Error('Acesso não autorizado');if(ticket!==this.generation||this.mode!=='minimized')return;this.setMode(this.previous);this.focus()}catch(e){if(ticket===this.generation){this.close();this.status.textContent='Jogo indisponível: '+String(e)}}}
 close(){this.generation++;if(document.fullscreenElement===this.local)void document.exitFullscreen().catch(()=>{});this.frame=null;this.id=null;this.board.replaceChildren();this.setMode('closed');this.fullscreen=false;this.zoom=1;document.getElementById('zGameTurn').textContent=''}
}
return GameWindow;
})();
