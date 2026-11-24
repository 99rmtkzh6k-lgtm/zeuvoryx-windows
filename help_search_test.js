const assert=require('node:assert/strict');
const {search}=require('./help_search.js');
const data=require('./help_articles.json');
const cases=[["pt-br","não estou conseguindo instalar","install"],["pt-br","varredura travou","fullscan"],["pt-br","restaurar quarentena","quarantine"],["pt-br","exportar relatório","report"],["pt-br","onedrive","cloud"],["en","app closes crash","startup"],["en","quarantine restore","quarantine"],["es","instalación","install"],["es","exportar informe","report"],["ja","フルスキャンが停止した","fullscan"],["ja","隔離 復元","quarantine"],["pt-br","xxxxxxxxxxxxx",""]];
for(let i=0;i<5;i++)for(const [l,q,want] of cases){assert.equal(search(data[l],q)[0]?.id||'',want,l+' '+q)}
for(const rows of Object.values(data)){assert.equal(rows.length,14);assert.equal(new Set(rows.map(a=>a.id)).size,14);assert.equal(search(rows,'').length,7);assert.equal(search(rows,'<script>alert(123)</script>')[0]?.id||'','');}
console.log('60 query checks plus dataset checks passed');
