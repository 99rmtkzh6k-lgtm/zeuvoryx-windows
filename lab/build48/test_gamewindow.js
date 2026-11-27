const {chromium}=require('playwright');
const {pathToFileURL}=require('url');
const path=require('path');
const assert=require('node:assert/strict');

const keys=['ArrowLeft','ArrowUp','ArrowRight','ArrowDown'];
const vectors=[['ArrowRight',1,0],['ArrowDown',0,1],['ArrowLeft',-1,0],['ArrowUp',0,-1]];
const pause=ms=>new Promise(resolve=>setTimeout(resolve,ms));
async function poll(read,accept,label,timeout=5000){
 const deadline=Date.now()+timeout;
 let last;
 do{
  last=await read();
  if(accept(last))return last;
  await pause(50);
 }while(Date.now()<deadline);
 assert.fail(label+'; last='+JSON.stringify(last));
}
async function focusFrame(page,frame){
 await page.locator('#zGameBoard iframe').focus();
 await poll(()=>frame.evaluate(()=>document.hasFocus()),v=>v===true,'iframe keyboard focus');
}
async function focusHost(page,id){
 const b=page.locator(id);
 await b.focus();
 await poll(()=>b.evaluate(el=>el===document.activeElement&&document.hasFocus()),v=>v===true,'host keyboard focus '+id);
}
async function counterEnter(page,frame,button,label,delta){
 await button.waitFor({state:'visible'});
 assert.equal(await button.isEnabled(),true,label+' button disabled');
 const read=async()=>{
  const text=await frame.locator('#topo-recursos').textContent();
  const match=text?.match(label==='Robozinho'?/Lixo\s*\(Waste\)\s*(\d+)/:/Ouro\s*\(Gold\)\s*(\d+)/);
  assert.ok(match,label+' visible resource counter not found: '+text);
  return Number(match[1]);
 };
 const before=await read();
 await button.focus();
 assert.equal(await button.evaluate(el=>el===document.activeElement&&document.hasFocus()),true,label+' keyboard focus missing');
 await page.keyboard.press('Enter');
 const after=await poll(read,v=>v!==before,label+' Enter did not change resources',2000);
 assert.equal(after,before+delta,label+' Enter exact resource delta');
 console.log(label+' keyboard Enter: visible resource '+before+' -> '+after);
}
async function tileState(frame){
 return frame.evaluate(()=>({
  board:Array.from(document.querySelectorAll('.tile-container .tile'),t=>JSON.stringify([t.textContent,t.className.match(/tile-position-\d+-\d+/)?.[0]])).sort().join('|'),
  count:document.querySelectorAll('.tile-container .tile').length,
  score:document.querySelector('.score-container')?.textContent.trim()
 }));
}
async function reset2048(frame){
 await frame.getByText('New Game',{exact:true}).click();
 await poll(()=>tileState(frame),v=>v.count===2&&v.score==='0','2048 restart: two tiles and zero score');
}
async function move2048(page,frame){
 await frame.locator('.tile-container .tile').first().waitFor({state:'visible'});
 await poll(()=>tileState(frame),v=>v.count===2&&v.score==='0','2048 initial board');
 await frame.locator('.game-container').click();
 assert.equal(await frame.evaluate(()=>document.hasFocus()),true,'2048 iframe keyboard focus');
 const before=(await tileState(frame)).board;
 let moved=false;
 for(const key of keys){
  await page.keyboard.press(key);
  const deadline=Date.now()+1000;
  do{
   if((await tileState(frame)).board!==before){moved=true;break}
   await pause(50);
  }while(Date.now()<deadline);
  if(moved)break;
 }
 assert.ok(moved,'2048 real keyboard must change tile values/positions');
 console.log('2048 keyboard movement verified');
}
// The following are script-classic lexical bindings, verified against the
// game sources. No private IIFE state is exposed. These reads inspect actual
// coordinates/direction used by the renderer rather than arbitrary pixels.
async function snakeMove(page,frame){
 await frame.getByRole('button',{name:'START GAME',exact:true}).click();
 await focusFrame(page,frame);
 const read=()=>frame.evaluate(()=>({running:game.gameRunning,head:{...game.snake[0]},dir:{...game.direction},next:{...game.nextDirection},size:GRID_SIZE,body:game.snake.map(s=>({...s})),obstacles:game.obstacles.map(s=>({...s}))}));
 const before=await poll(read,v=>v.running,'Snake start');
 const choice=vectors.find(([,dx,dy])=>{
  const a=before.head.x+dx,b=before.head.y+dy;
  return dx*before.dir.x+dy*before.dir.y===0&&a>=0&&b>=0&&a<before.size&&b<before.size&&![...before.body,...before.obstacles].some(p=>p.x===a&&p.y===b);
 });
 assert.ok(choice,'Snake no legal perpendicular direction');
 const [key,dx,dy]=choice;
 await page.keyboard.press(key);
 const after=await poll(read,v=>v.running&&v.dir.x===dx&&v.dir.y===dy&&v.next.x===dx&&v.next.y===dy&&((v.head.x-before.head.x)*dx+(v.head.y-before.head.y)*dy)>0,'Snake head did not move in requested direction',2000);
 console.log('Snake directional keyboard: '+key+' '+JSON.stringify(before.head)+' -> '+JSON.stringify(after.head));
}
async function blocksMove(page,frame){
 await focusFrame(page,frame);
 const read=()=>frame.evaluate(()=>({x,y,over,left:fits(p,x-1,y),right:fits(p,x+1,y)}));
 const before=await poll(read,v=>!v.over&&(v.left||v.right),'Blocks movable piece');
 const dx=before.right?1:-1,key=dx===1?'ArrowRight':'ArrowLeft';
 await page.keyboard.press(key);
 const after=await poll(read,v=>v.x===before.x+dx&&!v.over,'Blocks horizontal movement, not automatic fall',1500);
 console.log('Blocks keyboard: '+key+' x='+before.x+' -> '+after.x);
}
async function mazeMove(page,frame,original){
 await focusFrame(page,frame);
 const read=original?
  ()=>frame.evaluate(()=>({x:player.x,y:player.y,grid:maze.map(row=>row.slice()),over:game.gameOver})):
  ()=>frame.evaluate(()=>({x:p.x,y:p.y,grid:grid.slice(),dir:dir.slice(),over}));
 const before=await read();
 assert.equal(before.over,false,'Maze must be active');
 if(!original)assert.deepEqual(before.dir,[0,0],'integrated Maze must start stationary');
 const free=(x,y)=>before.grid[y]?.[x]===(original?0:'.');
 const choice=vectors.find(([,dx,dy])=>free(before.x+dx,before.y+dy));
 assert.ok(choice,'Maze no legal adjacent cell');
 const [key,dx,dy]=choice;
 await page.keyboard.down(key);
 let after;
 try{
  after=await poll(read,v=>{
   const ax=v.x-before.x,ay=v.y-before.y;
   return ax*dx+ay*dy>0&&ax*dy-ay*dx===0;
  },'Maze keyboard did not move along legal edge',1500);
 }finally{await page.keyboard.up(key)}
 const distance=(after.x-before.x)*dx+(after.y-before.y)*dy;
 for(let i=1;i<=distance;i++)assert.ok(free(before.x+dx*i,before.y+dy*i),'Maze crossed wall');
 console.log((original?'Digital Maze':'Maze')+' keyboard '+key+' ('+before.x+','+before.y+') -> ('+after.x+','+after.y+')');
}
const retroControls={
 'retro-racer':{start:'START RACE',key:'ArrowRight',read:f=>f.evaluate(()=>game.player.x)},
 'cartridge-dash':{start:'START DASH',key:'ArrowRight',read:f=>f.evaluate(()=>game.player.x)},
 'jump-quest':{start:'START GAME',key:'ArrowRight',read:f=>f.evaluate(()=>game.player.x)},
 'laser-defender':{start:'START',key:'ArrowRight',read:f=>f.evaluate(()=>game.base.x)},
 'meteor-strike':{key:'ArrowRight',read:f=>f.evaluate(()=>game.player.x)},
 'ping-pong-2p':{key:'ArrowDown',read:f=>f.evaluate(()=>game.p2.y)},
 'breakout':{start:'START GAME',key:'ArrowRight',read:f=>f.evaluate(()=>game.paddle.x)},
 'asteroids':{start:'START',key:'ArrowRight',read:f=>f.evaluate(()=>game.ship.angle)}
};
async function semantic(page,frame,id){
 switch(id){
  case '2048':return move2048(page,frame);
  case 'robozinho':return counterEnter(page,frame,frame.locator('#r-catar'),'Robozinho',10);
  case 'guerra':
   await frame.getByRole('button',{name:'JOGAR (PLAY)',exact:true}).click();
   await frame.getByRole('button',{name:'LEONARDO',exact:true}).click();
   return counterEnter(page,frame,frame.getByRole('button',{name:/MINERAR/}),'Guerra',20);
  case 'snake':return snakeMove(page,frame);
  case 'blocks':return blocksMove(page,frame);
  case 'maze':return mazeMove(page,frame,false);
  case 'digital-maze':return mazeMove(page,frame,true);
  default:{
   const m=retroControls[id];assert.ok(m,'No semantic contract '+id);
   if(m.start)await frame.getByRole('button',{name:m.start,exact:true}).click();
   await focusFrame(page,frame);
   if(id==='ping-pong-2p'){
    await page.keyboard.press('Space');
    await poll(()=>frame.evaluate(()=>game.gameRunning),v=>v===true,'Ping Pong start');
   }
   const before=await m.read(frame);assert.ok(Number.isFinite(before),id+' invalid initial coordinate');
   await page.keyboard.down(m.key);
   let after;
   try{after=await poll(()=>m.read(frame),v=>Number.isFinite(v)&&v>before,id+' expected positive directional movement',2000)}
   finally{await page.keyboard.up(m.key)}
   console.log(id+' keyboard '+m.key+' coordinate '+before+' -> '+after);
  }
 }
}
async function interact(page,frame,id){
 const end=Date.now()+30000;
 let actions=0;
 while(Date.now()<end){
  if(id==='robozinho')await frame.locator('#r-catar').click();
  else if(id==='guerra')await frame.getByRole('button',{name:/MINERAR/}).click();
  else if(id==='blocks'){await frame.locator('#rotate').click();await frame.locator('#down').click()}
  else if(id==='maze')await frame.locator(['#right','#down','#left','#up'][actions%4]).click();
  else{
   const key=id==='ping-pong-2p'?['ArrowUp','ArrowDown'][actions%2]:keys[actions%4];
   await page.keyboard.down(key);try{await pause(120)}finally{await page.keyboard.up(key)}
  }
  actions++;await pause(380);
 }
 assert.ok(actions>0,id+' 30 second interaction absent');
 if(id==='2048')await reset2048(frame);
 console.log(id+' 30s interaction: '+actions+' input rounds; no claim of full game completion');
}
function monitor(page){
 const errors=[],network=[],popups=[];
 page.on('pageerror',e=>errors.push(String(e)));
 page.on('request',r=>{if(/^https?:/i.test(r.url()))network.push(r.url())});
 page.on('popup',p=>{popups.push(p.url());void p.close()});
 return ()=>{assert.deepEqual(errors,[],'page runtime errors');assert.deepEqual(network,[],'unexpected HTTP requests');assert.deepEqual(popups,[],'unexpected popups')};
}
const mode=page=>page.locator('#zGames314').getAttribute('data-window');
const zoom=page=>page.locator('#zGameZoomReset').textContent();
const full=page=>page.evaluate(()=>document.fullscreenElement?.id==='zGameLocal');
const waitMode=(page,m)=>poll(()=>mode(page),v=>v===m,'window mode '+m);
const waitZoom=(page,z)=>poll(()=>zoom(page),v=>v===z,'zoom '+z);
const waitFull=(page,f)=>poll(()=>full(page),v=>v===f,'fullscreen '+f);
async function bounds(page,label,scaled=false){
 const box=await page.evaluate(()=>{
  const r=s=>{const e=document.querySelector(s),b=e?.getBoundingClientRect();return b?{left:b.left,top:b.top,right:b.right,bottom:b.bottom,width:b.width,height:b.height}:null};
  const v=document.querySelector('.zGameViewport');
  return {viewport:r('.zGameViewport'),local:r('#zGameLocal'),frame:r('#zGameBoard iframe'),width:innerWidth,height:innerHeight,overflow:v?getComputedStyle(v).overflow:null};
 });
 const inside=(a,b)=>a&&b&&a.left>=b.left-2&&a.top>=b.top-2&&a.right<=b.right+2&&a.bottom<=b.bottom+2;
 const host={left:0,top:0,right:box.width,bottom:box.height};
 assert.ok(box.local?.width>0&&box.local.height>0&&inside(box.local,host),label+' window outside host: '+JSON.stringify(box));
 assert.ok(box.viewport?.width>0&&box.viewport.height>0&&inside(box.viewport,box.local),label+' viewport outside window: '+JSON.stringify(box));
 if(scaled)assert.ok(/auto|scroll|hidden|clip/.test(box.overflow),label+' zoom not clipped by viewport');
 else assert.ok(box.frame?.width>0&&box.frame.height>0&&inside(box.frame,box.viewport),label+' iframe outside usable viewport: '+JSON.stringify(box));
}
async function closeGame(page,frame){
 const detached=page.waitForEvent('framedetached',{predicate:f=>f===frame,timeout:2000});
 await Promise.all([detached,page.locator('#zGameClose315').click()]);
 assert.equal(await page.locator('#zGameBoard iframe').count(),0);
 assert.equal(frame.isDetached(),true);
 await waitMode(page,'closed');
}
async function openGame(page,id){
 await page.locator('[data-game-id="'+id+'"] button:not(:disabled)').click();
 const loc=page.locator('#zGameBoard iframe');
 await loc.waitFor({state:'visible'});
 assert.equal(await loc.getAttribute('sandbox'),'allow-scripts');
 assert.equal(await loc.getAttribute('referrerpolicy'),'no-referrer');
 const frame=await (await loc.elementHandle()).contentFrame();
 await frame.locator('body').waitFor({state:'visible'});
 await waitMode(page,'normal');
 await page.locator('#zGameLocal').scrollIntoViewIfNeeded();
 await poll(()=>frame.evaluate(()=>document.readyState),v=>v==='complete','game document ready');
 await bounds(page,id+' normal');
 return frame;
}
async function maximize(page,id){
 await page.locator('#zGameMax').click();await waitMode(page,'maximized');
 assert.equal(await page.locator('.shell > aside').isVisible(),false);
 await bounds(page,id+' maximized');
 const b=await page.locator('#zGameBoard iframe').boundingBox();
 assert.ok(b.width>page.viewportSize().width*.8&&b.height>page.viewportSize().height*.4,id+' collapsed maximized viewport');
 console.log(id+' maximized within usable host');
}
async function exerciseWindow(page,frame,id){
 const loc=page.locator('#zGameBoard iframe');
 await page.locator('#zGameZoomPlus').click();await waitZoom(page,'110%');await bounds(page,id+' zoom',true);
 await focusFrame(page,frame);await page.keyboard.press('Control+-');await waitZoom(page,'100%');
 await focusHost(page,'#zGameZoomReset');await page.keyboard.press('Control+=');await waitZoom(page,'110%');
 await focusHost(page,'#zGameZoomReset');await page.keyboard.press('Control+0');await waitZoom(page,'100%');
 await bounds(page,id+' zoom reset');
 await page.locator('#zGameFull').click();await waitFull(page,true);await bounds(page,id+' fullscreen button');
 await page.locator('#zGameFull').click();await waitFull(page,false);
 await focusHost(page,'#zGameFull');await page.keyboard.press('Alt+Enter');await waitFull(page,true);
 await focusFrame(page,frame);await page.keyboard.press('F11');await waitFull(page,false);
 await focusFrame(page,frame);await page.keyboard.press('Alt+Enter');await waitFull(page,true);
 await focusHost(page,'#zGameFull');await page.keyboard.press('F11');await waitFull(page,false);
 await bounds(page,id+' fullscreen exit');
 await page.locator('#zGameMin').click();await waitMode(page,'minimized');
 assert.equal(await loc.count(),1);assert.equal(await loc.isVisible(),false);assert.equal(frame.isDetached(),false);
 const dock=page.locator('#zGameDockRestore');await dock.waitFor({state:'visible'});await dock.scrollIntoViewIfNeeded();
 const db=await dock.boundingBox(),vp=page.viewportSize();
 assert.ok(db&&db.x>=-2&&db.y>=-2&&db.x+db.width<=vp.width+2&&db.y+db.height<=vp.height+2,id+' minimized dock outside host');
 await dock.click();await waitMode(page,'maximized');
 assert.equal(frame.isDetached(),false);await bounds(page,id+' restored');
 await page.locator('#zGameMax').click();await waitMode(page,'normal');
 await page.locator('#zGameLocal').scrollIntoViewIfNeeded();await bounds(page,id+' normal restored');
 await closeGame(page,frame);
 console.log(id+' window controls passed: zoom, host/iframe Alt+Enter and F11, fullscreen, minimize, restore, normal, teardown');
}
async function denialAfterSwitch(page,next){
 const frame=await openGame(page,next);
 await page.locator('#zGameMin').click();await waitMode(page,'minimized');
 await page.evaluate(id=>window.blockedIds.add(id),next);
 const detached=page.waitForEvent('framedetached',{predicate:f=>f===frame,timeout:5000});
 await page.locator('#zGameDockRestore').click();await detached;await waitMode(page,'closed');
 assert.equal(await page.locator('#zGameBoard iframe').count(),0);
 assert.equal(frame.isDetached(),true);
 console.log(next+' game switch and simulated parental revalidation denial passed');
}
(async()=>{
 const browser=await chromium.launch({channel:'msedge',headless:true});
 const ids=['robozinho','2048','snake','blocks','maze','guerra'];
 try{
  for(let cycle=0;cycle<20;cycle++){
   const size=[[1366,768],[1600,900],[1920,1080]][cycle%3],dpi=[1,1.25,1.5][Math.floor(cycle/3)%3];
   const page=await browser.newPage({viewport:{width:Math.floor(size[0]/dpi),height:Math.floor(size[1]/dpi)},deviceScaleFactor:dpi});
   const verify=monitor(page),id=ids[cycle%ids.length],next=ids[(cycle+1)%ids.length];
   try{
    await page.goto(pathToFileURL(path.resolve('build48-browser-harness.html')).href);
    await page.locator('#zGamesNav314').click();
    const frame=await openGame(page,id);
    await maximize(page,id);
    await semantic(page,frame,id);await interact(page,frame,id);
    await exerciseWindow(page,frame,id);await denialAfterSwitch(page,next);verify();
    console.log('GameWindow browser cycle '+(cycle+1)+'/20 '+id+' '+size.join('x')+' DPI '+dpi+' : passed; simulated native denial, not physical');
   }catch(e){console.error('GameWindow integrated cycle '+(cycle+1)+'/20 game '+id);throw e}
   finally{await page.close()}
  }
  const original=await browser.newPage({viewport:{width:1366,height:768}});
  const verify=monitor(original);
  const originalIds=['2048','retro-racer','cartridge-dash','jump-quest','laser-defender','digital-maze','meteor-strike','ping-pong-2p','breakout','snake','asteroids'];
  try{
   await original.goto(pathToFileURL(path.resolve('gamewindow-original-harness.html')).href);
   await original.locator('#zGamesNav314').click();
   const tiles=original.locator('#zGameCatalog [data-game-id]');
   assert.deepEqual((await tiles.evaluateAll(a=>a.map(t=>t.dataset.gameId))).sort(),originalIds.slice().sort(),'original catalog mismatch');
   for(const id of originalIds){
    try{
     const frame=await openGame(original,id);await maximize(original,id);
     await semantic(original,frame,id);await interact(original,frame,id);await exerciseWindow(original,frame,id);
     verify();
     console.log('Original catalog: '+id+' semantic input, 30s interaction, shared controls verified in browser; NOT physical');
    }catch(e){console.error('GameWindow original catalog game '+id);throw e}
   }
   await original.evaluate(()=>window.blockedIds.add('2048'));
   await original.locator('[data-game-id="2048"] button').click();
   await poll(()=>original.locator('#zGameStatus').textContent(),v=>v?.includes('Jogo indisponível'),'original parental denial');
   assert.equal(await original.locator('#zGameBoard iframe').count(),0);verify();
   console.log('GameWindow original catalog 11/11 and simulated parental denial passed');
  }finally{await original.close()}
 }finally{await browser.close()}
})().catch(e=>{console.error(e);process.exitCode=1});
