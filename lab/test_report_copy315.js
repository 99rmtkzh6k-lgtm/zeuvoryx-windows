// Exercise the actual generated report UI handler with a minimal DOM double.
const fs=require('node:fs'),vm=require('node:vm'),assert=require('node:assert/strict');
const html=fs.readFileSync(process.argv[2],'utf8');
const m=html.match(/\/\* Z315_REPORT_COPY_BEGIN \*\/([\s\S]*?)\/\* Z315_REPORT_COPY_END \*\//);
assert.ok(m,'missing generated report copy code');
async function run(fallback){
 const nodes=new Map();
 const el=id=>{if(!nodes.has(id))nodes.set(id,{value:'',textContent:'',hidden:true,disabled:false,addEventListener(){},focus(){},select(){this.selected=true}});return nodes.get(id)};
 el('z315ReportSource').value='current';
 const expected=JSON.stringify({partial:true,files:42});
 let copied='',calls=0,release;
 const ctx={el,z315ReportBusy:false,z315ReportPaint:async()=>{},
 window:{reportText315:async()=>{calls++;return expected}},
 navigator:{clipboard:{writeText:async text=>{if(fallback)throw Error('denied');copied=text}}},
 document:{execCommand(cmd){assert.equal(cmd,'copy');copied=el('z315Text').value;return true}}};
 vm.runInNewContext(m[1],ctx);
 await el('z315ShowText').onclick();
 assert.equal(el('z315TextPanel').hidden,false);assert.equal(el('z315Text').value,expected);
 assert.equal(ctx.z315ReportBusy,false);assert.equal(calls,1);
 await el('z315CopyText').onclick();assert.equal(copied,expected);
 ctx.window.reportText315=()=>new Promise(resolve=>{release=resolve});
 const pending=el('z315ShowText').onclick();assert.equal(ctx.z315ReportBusy,true);
 assert.equal(el('z315ReportSave').disabled,true);
 el('z315ReportSource').value='diag';release(expected);await pending;
 assert.equal(el('z315Text').value,'');assert.equal(ctx.z315ReportBusy,false);
}
(async()=>{for(let i=0;i<5;i++){await run(false);await run(true)}console.log('Generated report copy UI: 5 normal and 5 fallback cycles passed')})().catch(e=>{console.error(e);process.exitCode=1});
