const {chromium}=require('playwright');
const {pathToFileURL}=require('url');
const path=require('path');
const assert=require('node:assert/strict');
(async()=>{
 const browser=await chromium.launch({channel:'msedge',headless:true});
 try{
  const page=await browser.newPage();
  const errors=[],network=[],popups=[];
  page.on('pageerror',e=>errors.push(String(e)));
  page.on('request',r=>{if(/^https?:/i.test(r.url()))network.push(r.url())});
  page.on('popup',p=>{popups.push(p.url());void p.close()});
  await page.goto(pathToFileURL(path.resolve('build48-browser-harness.html')).href);
  const nav=page.locator('#zGamesNav314');
  await nav.click();
  assert.equal(await page.locator('[data-game-id]').count(),6);
  for(const cat of ['Puzzle','Retro','Ação']){await page.selectOption('#zGameFilter',cat);assert.equal(await page.locator('[data-game-id]').count(),2)}
  await page.selectOption('#zGameFilter','');
  await page.evaluate(()=>window.blockedIds.add('guerra'));
  await page.locator('#zGameBack').click();await nav.click();
  await page.locator('[data-game-id="guerra"] button:disabled').waitFor();
  assert.equal(await page.locator('#zGameBoard iframe').count(),0);
  await page.evaluate(()=>window.blockedIds.clear());
  await page.locator('#zGameBack').click();await nav.click();
  for(const id of ['2048','snake','blocks','maze','guerra','robozinho']){
   const start=Date.now();
   await page.locator('[data-game-id="'+id+'"] button:not(:disabled)').click();
   const loc=page.locator('#zGameBoard iframe');
   await loc.waitFor();
   assert.equal(await loc.getAttribute('sandbox'),'allow-scripts');
   assert.equal(await loc.getAttribute('referrerpolicy'),'no-referrer');
   const frame=await (await loc.elementHandle()).contentFrame();
   await frame.locator('body').waitFor();
   if(id==='snake')await frame.getByRole('button',{name:'START GAME',exact:true}).click();
   if(id==='guerra'){
     await frame.getByRole('button',{name:'JOGAR (PLAY)',exact:true}).click();
     await frame.getByRole('button',{name:'LEONARDO',exact:true}).click();
     assert.equal(await frame.evaluate(()=>typeof comprar==='function'&&typeof comprarArma==='function'),true);
   }
   for(let i=0;i<6;i++){
     if(id==='2048'||id==='snake'){await frame.locator('body').press(['ArrowLeft','ArrowUp','ArrowRight','ArrowDown'][i%4]);}
     else if(id==='blocks'){await frame.locator('#rotate').click();await frame.locator('#down').click();}
     else if(id==='maze'){await frame.locator(['#right','#down','#left','#up'][i%4]).click();}
     else if(id==='guerra'){await frame.evaluate(()=>{minerar();vender();comprar();});}
     else{await frame.locator('#r-catar').click();await frame.locator('#r-amassar').click();await frame.locator('#r-conteiner').click();}
     await page.waitForTimeout(10000);
   }
   assert.ok(Date.now()-start>=60000,id+' duration');
   assert.equal(await page.locator('#zGameBoard iframe').count(),1);
   await page.locator('#zGameClose315').click();
   assert.equal(await page.locator('#zGameBoard iframe').count(),0);
   assert.equal(frame.isDetached(),true);
   console.log('Build48 browser game '+id+': 60 seconds, interaction and iframe teardown confirmed');
  }
  assert.deepEqual(errors,[],'JavaScript runtime errors');
  assert.deepEqual(network,[],'unexpected HTTP request');
  assert.deepEqual(popups,[],'unexpected popup');
  console.log('Build48: six games in Edge on Windows VM; mock parent denial and exact sandbox confirmed. Native policy tested separately.');
 }finally{await browser.close()}
})().catch(e=>{console.error(e);process.exitCode=1});
