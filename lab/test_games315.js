const assert=require('node:assert/strict');const g=require('./games315_core');
for(let cycle=0;cycle<5;cycle++){
 let m=g.memory(()=>0.42);assert.equal(m.cards.length,12);for(const x of 'ABCDEF')assert.equal(m.cards.filter(v=>v===x).length,2);
 for(const x of 'ABCDEF'){const a=m.cards.indexOf(x),b=m.cards.lastIndexOf(x);m=g.flip(m,a);assert.strictEqual(g.flip(m,a),m);m=g.flip(m,b)}assert.equal(m.matched.length,12);assert.equal(m.moves,6);
 let n=g.memory(()=>0.42);const a=0,b=n.cards.findIndex(x=>x!==n.cards[a]);n=g.flip(g.flip(n,a),b);assert.equal(n.open.length,2);assert.equal(n.matched.length,0);assert.equal(n.moves,1);
 let d=g.drawGame(()=>0.42);assert.equal(new Set(d.cards.map(c=>c.rank+c.suit)).size,52);for(let i=0;i<51;i++){const a=d.cards[d.index].rank,b=d.cards[d.index+1].rank;d=g.guess(d,b>a?'higher':'lower')}assert.equal(d.index,51);assert.strictEqual(g.guess(d,'higher'),d);assert.ok(d.score<=51);
 let t=g.tic();for(const i of [0,3,1,4,2])t=g.move(t,i);assert.equal(t.result,'X');assert.strictEqual(g.move(t,5),t);t=g.tic();for(const i of [0,1,2,4,3,5,7,6,8])t=g.move(t,i);assert.equal(t.result,'draw');
 console.log('cycle '+(cycle+1)+': memory, cards and tic-tac-toe passed');
}
