#!/usr/bin/env python3
"""Compare two locally exported ZEUVORYX JSON reports; no upload or network."""
import argparse,json,datetime
from pathlib import Path
from collections import Counter

def dt(s):
    return datetime.datetime.fromisoformat(s.replace('Z','+00:00'))
def read(p):
    r=json.loads(Path(p).read_text(encoding='utf-8-sig')); t=(dt(r['finished'])-dt(r['started'])).total_seconds()
    if t<=0:raise ValueError('Duração inválida')
    f=r.get('findings') or []; hashes=[x['sha256'] for x in f if x.get('sha256')]
    return dict(root=r.get('root'),engine=r.get('engine_version','anterior'),seconds=t,files=r['files'],bytes=r['bytes'],files_s=r['files']/t,mb_s=r['bytes']/1e6/t,alerts=len(f),levels=dict(Counter(x.get('level','unknown') for x in f)),unique_hashes=len(set(hashes)),skipped=r.get('skipped',0),partial=bool(r.get('partial') or r.get('skipped') or r.get('errors') or r.get('cancelled')),cancelled=bool(r.get('cancelled')))
if __name__=='__main__':
    p=argparse.ArgumentParser(description=__doc__);p.add_argument('antigo');p.add_argument('novo');a=p.parse_args();x,y=read(a.antigo),read(a.novo)
    print(json.dumps({'anterior':x,'novo':y,'variacao_percentual_arquivos_s':round((y['files_s']/x['files_s']-1)*100,2) if x['files_s'] else None,'variacao_percentual_MB_s':round((y['mb_s']/x['mb_s']-1)*100,2) if x['mb_s'] else None,'escopo_e_volume_iguais':x['root']==y['root'] and x['files']==y['files'] and x['bytes']==y['bytes'],'nota':'Alertas não são malwares comprovados. Compare somente testes equivalentes; cobertura parcial e cancelamento afetam a comparação.'},indent=2,ensure_ascii=False))
