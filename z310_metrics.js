// Pure report counters; no claims of physical disk I/O.
function z310Delta(prev,s,now){
 const r=s.report||{},id=String(r.scan_id||r.started||'');
 const n=Math.max(0,Number(s.finding_total??(r.findings||[]).length)||0);
 const files=Math.max(0,Number(s.files)||0),bytes=Math.max(0,Number(s.bytes)||0);
 const active=Boolean(s.running||s.done)&&!!id;
 const same=active&&prev&&prev.id===id;
 const dt=same?(now-prev.now)/1000:0;
 const measurable=same&&s.running&&prev.running&&dt>0&&files>=prev.files&&bytes>=prev.bytes;
 return {state:active?{id,n,files,bytes,now,running:!!s.running}:null,alerts:active?Math.max(0,n-(same?prev.n:0)):0,filesRate:measurable?(files-prev.files)/dt:null,bytesRate:measurable?(bytes-prev.bytes)/dt:null};
}
function z310Remaining(used,bytes,mode){return mode==='full'&&Number.isFinite(used)&&used>=0&&Number.isFinite(bytes)&&bytes>=0?Math.max(0,used-bytes):null}
if(typeof module!=='undefined'&&module.exports)module.exports={z310Delta,z310Remaining};
