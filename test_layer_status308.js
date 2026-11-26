const assert=require('node:assert/strict'),fs=require('node:fs'),vm=require('node:vm');
const html=fs.readFileSync('desktop.html','utf8');
const start=html.indexOf('function zLayerStatus('),end=html.indexOf('function auditRow(',start);
assert.ok(start>=0&&end>start);
const classify=vm.runInNewContext(html.slice(start,end)+'; zLayerStatus');
const cases=[
 ['Microsoft Defender','Ativado','green'],
 ['Proteção em tempo real do Defender','Desativado','red'],
 ['BitLocker (disco do sistema)','Não verificado','gray'],
 ['BitLocker (disco do sistema)','On','green'],
 ['BitLocker (disco do sistema)','Off','yellow'],
 ['Conexões RDP permitidas pela configuração','Desativado','green'],
 ['Conexões RDP permitidas pela configuração','Ativado','yellow'],
 ['ASR 1','Bloqueio','green'],
 ['ASR 2','Auditoria','yellow'],
 ['Proteção de rede','Modo de auditoria','yellow'],
 ['Windows','Windows 11 | build 26100','blue'],
 ['Consulta indisponível','Acesso negado','gray'],
 ['Firewall','Não verificado','gray'],
 ['Firewall Público','Desativado','red'],
 ['SmartScreen','Gerenciado: habilitado pela política','yellow'],
 ['Secure Boot','Incompatível','gray'],
];
const cycles=[];
for(let i=0;i<5;i++){
 for(const [title,state,expected] of cases)assert.equal(classify(title,state),expected,title+' '+state);
 cycles.push({cycle:i+1,classification_checks:cases.length,status:'passed'});
}
for(const op of ['realtime-on','network-audit','folder-audit','asr-audit','remediate','update','updates']){
 assert.match(html,new RegExp('data-native="'+op+'"'));
}
for(const id of ['auditWindows','auditReport','zLayerLegend'])assert.equal(html.split('id="'+id+'"').length,2,id);
for(const id of ['openSecurity','openUpdates'])assert.equal(html.split('id="'+id+'"').length,1,id+' must remain absent from internal interface');
assert.ok(!html.includes('zCameraPanel'));
fs.writeFileSync('security-center-regressions.json',JSON.stringify({
 suite:'Security Center classification and generated controls',
 platform:process.platform,
 native_policy_mutations_tested:false,cycles
},null,2)+'\n');
console.log('Security Center: five classification cycles and generated control checks');
