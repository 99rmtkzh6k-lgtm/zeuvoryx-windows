const assert=require('node:assert/strict');const {z310Delta:d,z310Remaining:r}=require('./z310_metrics');
for(let i=0;i<5;i++){
 const s={running:true,files:100,bytes:1048576,finding_total:0,report:{scan_id:'a'}};
 let a=d(null,s,1000);assert.equal(a.alerts,0);assert.equal(a.bytesRate,null);
 let b=d(a.state,{...s,files:120,bytes:3145728,finding_total:2},3000);assert.equal(b.filesRate,10);assert.equal(b.bytesRate,1048576);assert.equal(b.alerts,2);
 let c=d(b.state,{...s,files:120,bytes:3145728,finding_total:2},5000);assert.equal(c.alerts,0);assert.equal(c.bytesRate,0);
 assert.equal(d(c.state,{...s,report:{scan_id:'b'}},6000).bytesRate,null);
 assert.equal(d(c.state,{...s,running:false,done:true,finding_total:3},6000).alerts,1);
 assert.equal(d(c.state,{...s,files:1,bytes:1},6000).bytesRate,null);
 assert.equal(r(100,30,'full'),70);assert.equal(r(100,200,'full'),0);assert.equal(r(100,30,'quick'),null);assert.equal(r(null,30,'full'),null);
}
console.log('310: five cycles of measured rates, alert deduplication, scan identity and approximate remaining passed');
