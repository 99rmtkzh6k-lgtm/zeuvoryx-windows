from pathlib import Path
# Only generated sources and installer identity; preserve visual assets and all capabilities.
paths=list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]
for p in paths:
    if ".git" in p.parts: continue
    s=p.read_text(encoding="utf-8")
    n=s.replace("3.0.14","3.0.21")
    if n!=s:p.write_text(n,encoding="utf-8")
for name in ("desktop.html","setup309.iss"):
    if "3.0.21" not in Path(name).read_text(encoding="utf-8"):raise RuntimeError("Version identity missing: "+name)


# Replace the existing generated Unicode normalizer, preserving the scanner
# call site and guardedReader. Earlier CI steps already normalize the sample.
import re
p=Path("engine/readguard.go")
s=p.read_text(encoding="utf-8")
anchor=r'(?ms)^func normalizeScriptSample\(data \[\]byte\) \[\]byte \{.*?^\}'
replacement='''func normalizeScriptSample(data []byte) []byte {
 if decoded, ok := scriptguard.DecodeBOM(data); ok {
  return []byte(strings.TrimPrefix(decoded, "\\ufeff"))
 }
 return data
}'''
s,n=re.subn(anchor,lambda m:replacement,s)
if n!=1: raise RuntimeError("Unicode normalizer integration anchor: "+str(n))
for package in ("encoding/binary","unicode/utf16"):
    s,n=re.subn(r'(?m)^[ \t]*"'+re.escape(package)+r'"[ \t]*$',"",s)
    if n!=1: raise RuntimeError("Unicode import removal: "+package)
if s.count('import (')!=1: raise RuntimeError("Unicode import anchor")
s=s.replace('import (','import (\n "github.com/zeuvoryx/zeuvoryx/scriptguard"',1)
p.write_text(s,encoding="utf-8")
scanner=Path("engine/scanner.go").read_text(encoding="utf-8")
if not re.search(r'normalizeScriptSample\(data\)',scanner): raise RuntimeError("Unicode scanner call missing")
Path("engine/unicode315_test.go").write_text(r'''package engine
import("encoding/binary";"testing";"unicode/utf16")
func TestUnicode315ScannerIntegration(t *testing.T){
 for _,order:=range []binary.ByteOrder{binary.LittleEndian,binary.BigEndian}{
  for _,text:=range []string{"Write-Output 'hello'","Documentation example: Invoke-Expression and FromBase64String("}{
   b:=[]byte{0xff,0xfe};if order==binary.BigEndian{b=[]byte{0xfe,0xff}}
   for _,v:=range utf16.Encode([]rune(text)){var a [2]byte;order.PutUint16(a[:],v);b=append(b,a[:]...)}
   plain:=analyze("sample.ps1",[]byte(text),"",nil);wide:=analyze("sample.ps1",b,"",nil)
   if text!="Write-Output 'hello'" && plain==nil {t.Fatal("expected advisory finding")}
   if (plain==nil)!=(wide==nil){t.Fatalf("different results for equivalent script: %v %v",plain,wide)}
   if plain!=nil&&plain.Risk!=wide.Risk{t.Fatalf("different risk: %d %d",plain.Risk,wide.Risk)}
  }
 }
}
''',encoding="utf-8")

# Bind the same measured resource policy to both actual generated file scanners.
for name,pattern in (("main.go",r"stabilityWorkers\(runtime\.NumCPU\(\)\)"),("engine/scanner.go",r"(?<=:=)\s*runtime\.NumCPU\(\)")):
    p=Path(name);s=p.read_text(encoding="utf-8")
    replacement=" scanplan.CurrentWorkers(runtime.NumCPU())" if name=="engine/scanner.go" else "scanplan.CurrentWorkers(runtime.NumCPU())"
    s,n=re.subn(pattern,lambda m:replacement,s)
    if n!=1:raise RuntimeError("adaptive worker anchor "+name+": "+str(n))
    if s.count("import (")!=1:raise RuntimeError("adaptive worker import "+name)
    s=s.replace("import (",'import (\n "github.com/zeuvoryx/zeuvoryx/scanplan"',1)
    p.write_text(s,encoding="utf-8")
