const {chromium}=require('playwright');
const {pathToFileURL}=require('url');
const path=require('path');
const assert=require('node:assert/strict');
(async()=>{
 const browser=await chromium.launch({channel:'msedge',headless:true});
 const ids=['robozinho','2048','snake','blocks','maze','guerra'];
 try{
 for(let cycle=0;cycle<20;cycle++){
  const size=[[1366,768],[1600,900],[1920,1080]][cycle%3],dpi=[1,1.25,1.5][Math.floor(cycle/3)%3];
  const page=await browser.newPage({viewport:{width:Math.floor(size[0]/dpi),height:Math.floor(size[1]/dpi)},deviceScaleFactor:dpi});
  const errors=[];page.on('pageerror',e=>errors.push(String(e)));
  const id=ids[cycle%ids.length],next=ids[(cycle+1)%ids.length];
  try{
   await page.goto(pathToFileURL(path.resolve('build48-browser-harness.html')).href);
   await page.locator('#zGamesNav314').click();
   const tile=id=>page.locator('[data-game-id="'+id+'"] button:not(:disabled)');
   await tile(id).click();
   const loc=page.locator('#zGameBoard iframe');await loc.waitFor();
   const frame=await (await loc.elementHandle()).contentFrame();await frame.locator('body').waitFor();
   const mode=()=>page.locator('#zGames314').getAttribute('data-window');
   assert.equal(await mode(),'normal');
   await page.locator('#zGameMax').click();assert.equal(await mode(),'maximized');assert.equal(await page.locator('.shell > aside').isVisible(),false);
   const box=await loc.boundingBox();assert.ok(box.width>page.viewportSize().width*.8&&box.height>page.viewportSize().height*.4,'collapsed viewport '+JSON.stringify(box));
   const until=Date.now()+30000;
   if(id==='2048'){
    const snapshot=()=>frame.evaluate(()=>Array.from(document.querySelectorAll('.tile-container .tile'),t=>t.className+':'+t.textContent).sort().join('|'));
    await frame.locator('.game-container').click();await page.keyboard.press('ArrowLeft');await page.waitForTimeout(250);const before=await snapshot();await page.keyboard.press('ArrowRight');
    await page.waitForTimeout(250);assert.notEqual(await snapshot(),before,'2048 right arrow must change board');
   }else if(id==='robozinho'){
    const b=frame.locator('#r-catar');const before=await frame.evaluate(()=>jogo2.coletado);await b.focus();await page.keyboard.press('Enter');assert.ok(await frame.evaluate(()=>jogo2.coletado)>before,'Robozinho keyboard action did not collect waste');
   }else if(id==='snake'){await frame.getByRole('button',{name:'START GAME',exact:true}).click();await page.keyboard.press('ArrowUp');await page.waitForTimeout(160);await page.keyboard.press('ArrowRight');assert.equal(await frame.evaluate(()=>game.nextDirection.x),1,'Snake right direction failed')}
   else if(id==='guerra'){await frame.getByRole('button',{name:'JOGAR (PLAY)',exact:true}).click();await frame.getByRole('button',{name:'LEONARDO',exact:true}).click();const before=await frame.evaluate(()=>jogo.minerado);await frame.getByRole('button',{name:/MINERAR/}).focus();await page.keyboard.press('Enter');assert.ok(await frame.evaluate(()=>jogo.minerado)>before,'Guerra keyboard mining did not change state')}
   else if(id==='blocks'){const before=await frame.evaluate(()=>x);await page.keyboard.press('ArrowRight');assert.ok(await frame.evaluate(()=>x)>before,'Blocks right movement failed')}
   else if(id==='maze'){const before=await frame.evaluate(()=>p.x);await page.keyboard.press('ArrowRight');assert.ok(await frame.evaluate(()=>p.x)>before,'Maze right movement failed')}
   while(Date.now()<until){
    if(id==='2048'||id==='snake')await page.keyboard.press(['ArrowLeft','ArrowUp','ArrowRight','ArrowDown'][Math.floor(Date.now()/500)%4]);
    else if(id==='blocks')await frame.locator('#rotate').click();
    else if(id==='maze')await frame.locator('#right').click();
    else if(id==='robozinho')await frame.locator('#r-catar').click();
    else await frame.getByRole('button',{name:/MINERAR/}).click();
    await page.waitForTimeout(500)
   }
   await page.locator('#zGameZoomPlus').click();assert.equal(await page.locator('#zGameZoomReset').textContent(),'110%');
   await page.keyboard.press('Control+-');await page.waitForFunction(()=>document.getElementById('zGameZoomReset').textContent==='100%');
   await page.locator('#zGameFull').click();await page.waitForFunction(()=>document.fullscreenElement?.id==='zGameLocal');
   await page.locator('#zGameFull').click();await page.waitForFunction(()=>!document.fullscreenElement);
   await page.keyboard.press('Alt+Enter');await page.waitForFunction(()=>document.fullscreenElement?.id==='zGameLocal');await page.keyboard.press('F11');await page.waitForFunction(()=>!document.fullscreenElement);
   await page.locator('#zGameMin').click();assert.equal(await mode(),'minimized');assert.equal(await loc.count(),1);assert.equal(frame.isDetached(),false);
   await page.locator('#zGameDockRestore').click();await page.waitForFunction(()=>document.getElementById('zGames314').dataset.window==='maximized');
   assert.equal(frame.isDetached(),false);
   await page.locator('#zGameMax').click();assert.equal(await mode(),'normal');
   await page.locator('#zGameClose315').click();assert.equal(await loc.count(),0);
   await tile(next).click();await loc.waitFor();await page.locator('#zGameMin').click();
   await page.evaluate(id=>window.blockedIds.add(id),next);await page.locator('#zGameDockRestore').click();await page.waitForFunction(()=>document.getElementById('zGames314').dataset.window==='closed');
   assert.equal(await loc.count(),0);assert.deepEqual(errors,[]);
   console.log('GameWindow browser cycle '+(cycle+1)+'/20 '+id+' '+size.join('x')+' DPI '+dpi+' : passed; simulated native denial, not physical');
  }finally{await page.close()}
 }
 // Original catalog is distinct; every entry needs an explicit semantic input test.
 const original=await browser.newPage({viewport:{width:1366,height:768}});
 const originalIds=['2048','retro-racer','cartridge-dash','jump-quest','laser-defender','digital-maze','meteor-strike','ping-pong-2p','breakout','snake','asteroids'];
 const movement={
  'retro-racer':{start:'START RACE',state:'game.player.x',key:'ArrowRight'},
  'cartridge-dash':{start:'START DASH',state:'game.player.x',key:'ArrowRight'},
  'jump-quest':{start:'START GAME',state:'game.player.x',key:'ArrowRight'},
  'laser-defender':{start:'START',state:'game.base.x',key:'ArrowRight'},
  'meteor-strike':{state:'game.player.x',key:'ArrowRight'},
  'ping-pong-2p':{state:'game.p2.y',key:'ArrowDown'},
  'breakout':{start:'START GAME',state:'game.paddleTargetX',key:'ArrowRight'},
  'snake':{start:'START GAME',state:'game.nextDirection.x',key:'ArrowRight'},
  'asteroids':{start:'START',state:'game.ship.angle',key:'ArrowRight'}
 };
 const fail=[];original.on('pageerror',e=>fail.push(String(e)));
 try{
  await original.goto(pathToFileURL(path.resolve('gamewindow-original-harness.html')).href);
  await original.locator('#zGamesNav314').click();
  const tiles=original.locator('#zGameCatalog [data-game-id]');
  assert.deepEqual((await tiles.evaluateAll(a=>a.map(t=>t.dataset.gameId))).sort(),originalIds.slice().sort(),'Jogo sem GameWindow - correção obrigatória: semantic test missing');
  for(const id of originalIds){
   await original.locator('[data-game-id="'+id+'"] button').click();
   const loc=original.locator('#zGameBoard iframe');await loc.waitFor();
   const frame=await (await loc.elementHandle()).contentFrame();await frame.locator('body').waitFor();
   await original.locator('#zGameMax').click();
   const bounds=await loc.boundingBox();assert.ok(bounds.width>1000&&bounds.height>300,'Original catalog window collapsed: '+id);
   const m=movement[id];
   if(id==='2048'){
    const snap=()=>frame.evaluate(()=>Array.from(document.querySelectorAll('.tile-container .tile'),t=>t.className+':'+t.textContent).sort().join('|'));
    await frame.locator('.game-container').click();await original.keyboard.press('ArrowLeft');await original.waitForTimeout(250);const before=await snap();await original.keyboard.press('ArrowRight');await original.waitForTimeout(250);assert.notEqual(await snap(),before,id+' right movement failed');
   }else if(id==='digital-maze'){
    // Respect actual maze walls; choose a legal edge and prove coordinate change.
    const edge=await frame.evaluate(()=>{const steps=[['ArrowRight',1,0],['ArrowDown',0,1],['ArrowUp',0,-1],['ArrowLeft',-1,0]];const hit=steps.find(([,dx,dy])=>maze[player.y+dy]?.[player.x+dx]===0);return {key:hit?.[0],x:player.x,y:player.y}});
    assert.ok(edge.key,'maze has no valid move');await original.keyboard.down(edge.key);await original.waitForTimeout(110);await original.keyboard.up(edge.key);
    assert.notDeepEqual(await frame.evaluate(()=>[player.x,player.y]),[edge.x,edge.y],id+' movement failed');
   }else{
    assert.ok(m,'Jogo sem GameWindow - correção obrigatória: '+id);
    if(m.start)await frame.getByRole('button',{name:m.start,exact:true}).click();
    if(id==='snake'){await original.keyboard.press('ArrowUp');await original.waitForTimeout(160)}
    // Expressions are fixed test constants, evaluated only in the isolated game context.
    const before=await frame.evaluate(m.state);
    await original.keyboard.down(m.key);await original.waitForTimeout(180);await original.keyboard.up(m.key);
    assert.notEqual(await frame.evaluate(m.state),before,id+' semantic movement failed: '+m.key);
   }
   const end=Date.now()+30000;
   while(Date.now()<end){
    await original.keyboard.press(['ArrowLeft','ArrowUp','ArrowRight','ArrowDown'][Math.floor(Date.now()/500)%4]);
    await original.waitForTimeout(500);
   }
   await original.locator('#zGameZoomPlus').click();assert.equal(await original.locator('#zGameZoomReset').textContent(),'110%');
   await original.locator('#zGameZoomReset').click();
   await original.locator('#zGameFull').click();await original.waitForFunction(()=>document.fullscreenElement?.id==='zGameLocal');
   await original.locator('#zGameFull').click();await original.waitForFunction(()=>!document.fullscreenElement);
   await original.locator('#zGameMin').click();await original.locator('#zGameDockRestore').click();await original.waitForFunction(()=>document.getElementById('zGames314').dataset.window==='maximized');
   assert.equal(frame.isDetached(),false);
   await original.locator('#zGameClose315').click();assert.equal(await loc.count(),0);
   console.log('Original catalog: '+id+' semantic input, 30s interaction, shared controls verified in browser; NOT physical');
  }
  await original.evaluate(()=>window.blockedIds.add('2048'));await original.locator('[data-game-id="2048"] button').click();
  await original.waitForFunction(()=>document.getElementById('zGameStatus').textContent.includes('Jogo indisponível'));
  assert.equal(await original.locator('#zGameBoard iframe').count(),0);assert.deepEqual(fail,[]);
 }finally{await original.close()}
 }finally{await browser.close()}
})().catch(e=>{console.error(e);process.exitCode=1});
