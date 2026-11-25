"""Apply after the existing main.go generators, before Windows build."""
from pathlib import Path
import re
def transform(s):
    pattern=r'(func \(s \*State\) fullScan\s*\(.*?\)\s*\(engine\.Report,\s*error\)\s*\{)'
    if 'beginFullScanPower()' in s:
        raise ValueError('power request already integrated')
    s,n=re.subn(pattern,lambda m:m.group(1)+'\n\tif ctx.Err()==nil { releasePower:=beginFullScanPower(); defer releasePower() }\n',s,count=1,flags=re.S)
    if n!=1:
        raise ValueError('fullScan anchor missing')
    return s
if __name__=='__main__':
    p=Path('main.go')
    p.write_text(transform(p.read_text(encoding='utf-8-sig')),encoding='utf-8')
