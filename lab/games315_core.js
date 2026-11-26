// Original local games. No network, persistent storage, downloads, or native bridge.
const zGames315Core=(()=>{
 function shuffle(values,rng=Math.random){const a=values.slice();for(let i=a.length-1;i>0;i--){const j=Math.floor(rng()*(i+1));[a[i],a[j]]=[a[j],a[i]]}return a}
 function memory(rng){return {cards:shuffle(['A','B','C','D','E','F','A','B','C','D','E','F'],rng),open:[],matched:[],moves:0}}
 function flip(s,i){if(!Number.isInteger(i)||i<0||i>=s.cards.length||s.matched.includes(i)||s.open.includes(i))return s;let open=s.open.length===2?[]:s.open.slice();open.push(i);let matched=s.matched.slice(),moves=s.moves;if(open.length===2){moves++;if(s.cards[open[0]]===s.cards[open[1]]){matched=matched.concat(open);open=[]}}return {...s,open,matched,moves}}
 function deck(rng){return shuffle(Array.from({length:52},(_,i)=>({rank:i%13+1,suit:['♠','♥','♦','♣'][Math.floor(i/13)]})),rng)}
 function drawGame(rng){const cards=deck(rng);return {cards,index:0,score:0,last:'Escolha maior ou menor.'}}
 function guess(s,choice){if(!['higher','lower'].includes(choice)||s.index>=s.cards.length-1)return s;const a=s.cards[s.index].rank,b=s.cards[s.index+1].rank;const equal=a===b,ok=choice==='higher'?b>a:b<a;return {...s,index:s.index+1,score:s.score+Number(ok),last:equal?'Empate: sem ponto.':ok?'Acertou!':'Não foi desta vez.'}}
 const lines=[[0,1,2],[3,4,5],[6,7,8],[0,3,6],[1,4,7],[2,5,8],[0,4,8],[2,4,6]];
 function tic(){return {board:Array(9).fill(''),turn:'X',result:''}}
 function move(s,i){if(s.result||!Number.isInteger(i)||i<0||i>8||s.board[i])return s;const board=s.board.slice();board[i]=s.turn;const result=lines.some(l=>l.every(j=>board[j]===s.turn))?s.turn:board.every(Boolean)?'draw':'';return {board,turn:s.turn==='X'?'O':'X',result}}
 function ticBot(s){if(s.result||s.turn!=='O')return -1;function score(t,depth){if(t.result)return t.result==='O'?10-depth:t.result==='X'?depth-10:0;const a=t.board.flatMap((v,i)=>v?[]:[score(move(t,i),depth+1)]);return t.turn==='O'?Math.max(...a):Math.min(...a)}let best=-Infinity,idx=-1;for(const i of [4,0,2,6,8,1,3,5,7])if(!s.board[i]){const n=score(move(s,i),1);if(n>best){best=n;idx=i}}return idx}
 function chessBot(s){if(s.game_over()||s.turn()!=='b')return null;const val={p:100,n:320,b:330,r:500,q:900,k:0};const moves=s.moves({verbose:true});let best=-Infinity,choice=null;for(const m of moves){const req={from:m.from,to:m.to,promotion:m.promotion};s.move(req);const reply=s.moves({verbose:true});const risk=Math.max(0,...reply.map(x=>val[x.captured]||0));const n=(s.in_checkmate()?100000:0)+(val[m.captured]||0)+(m.promotion?800:0)-risk+(['d4','e4','d5','e5'].includes(m.to)?24:0)+(s.in_check()?12:0);s.undo();if(n>best){best=n;choice=req}}return choice?s.move(choice):null}
 return {shuffle,memory,flip,deck,drawGame,guess,tic,move,ticBot,chessBot};
})();
if(typeof module!=='undefined'&&module.exports)module.exports=zGames315Core;
