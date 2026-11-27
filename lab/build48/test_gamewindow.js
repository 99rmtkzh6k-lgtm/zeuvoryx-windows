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
   await page.locator('#zGameMax').click();assert.equal(await mode(),'maximized');
   const box=await loc.boundingBox();assert.ok(box.width>page.viewportSize().width*.8&&box.height>page.viewportSize().height*.4,'collapsed viewport '+JSON.stringify(box));
   const until=Date.now()+30000;
   if(id==='2048'){
    const snapshot=()=>frame.evaluate(()=>Array.from(document.querySelectorAll('.tile-container .tile'),t=>t.className+':'+t.textContent).sort().join('|'));
    await frame.locator('.game-container').click();const before=await snapshot();await page.keyboard.press('ArrowRight');
    await page.waitForTimeout(250);assert.notEqual(await snapshot(),before,'2048 right arrow must change board');
   }else if(id==='robozinho'){
    const b=frame.locator('#r-catar');await b.focus();await page.keyboard.press('Enter');
   }else if(id==='snake'){await frame.getByRole('button',{name:'START GAME',exact:true}).click()}
   else if(id==='guerra'){await frame.getByRole('button',{name:'JOGAR (PLAY)',exact:true}).click();await frame.getByRole('button',{name:'LEONARDO',exact:true}).click()}
   while(Date.now()<until){
    if(id==='2048'||id==='snake')await page.keyboard.press(['ArrowLeft','ArrowUp','ArrowRight','ArrowDown'][Math.floor(Date.now()/500)%4]);
    else if(id==='blocks')await frame.locator('#rotate').click();
    else if(id==='maze')await frame.locator('#right').click();
    else if(id==='robozinho')await frame.locator('#r-catar').click();
    else await frame.evaluate(()=>minerar());
    await page.waitForTimeout(500)
   }
   await page.locator('#zGameZoomPlus').click();assert.equal(await page.locator('#zGameZoomReset').textContent(),'110%');
   await page.keyboard.press('Control+-');await page.waitForFunction(()=>document.getElementById('zGameZoomReset').textContent==='100%');
   await page.locator('#zGameFull').click();await page.waitForFunction(()=>document.fullscreenElement?.id==='zGameLocal');
   await page.locator('#zGameFull').click();await page.waitForFunction(()=>!document.fullscreenElement);
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
 }finally{await browser.close()}
})().catch(e=>{console.error(e);process.exitCode=1});
