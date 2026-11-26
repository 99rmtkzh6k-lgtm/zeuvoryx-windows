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
 return {shuffle,memory,flip,deck,drawGame,guess,tic,move};
})();
if(typeof module!=='undefined'&&module.exports)module.exports=zGames315Core;
