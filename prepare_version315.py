from pathlib import Path
# Only generated sources and installer identity; preserve visual assets and all capabilities.
paths=list(Path(".").rglob("*.go"))+[Path("desktop.html"),Path("setup309.iss")]
for p in paths:
    if ".git" in p.parts: continue
    s=p.read_text(encoding="utf-8")
    n=s.replace("3.0.14","3.0.18")
    if n!=s:p.write_text(n,encoding="utf-8")
for name in ("desktop.html","setup309.iss"):
    if "3.0.18" not in Path(name).read_text(encoding="utf-8"):raise RuntimeError("Version identity missing: "+name)


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


# Preserve actual archive coverage; the physical ZIP hash remains independent.
p=Path("engine/scanner.go");s=p.read_text(encoding="utf-8")
def own316(a,b):
    global s
    if s.count(a)!=1:raise RuntimeError("Own motor 316 anchor: "+a[:100])
    s=s.replace(a,b,1)
def ownrx316(a,b,n=1):
    global s
    s,c=re.subn(a,lambda m:b(m),s)
    if c!=n:raise RuntimeError("Own motor 316 regex: "+a+" count "+str(c))
own316('type Report struct {','type Report struct {\n ArchiveCoverage316 ArchiveCoverage316 '+chr(96)+'json:"archive_coverage_316"'+chr(96))
ownrx316(r'func \(s \*Scanner\) examine\(ctx context\.Context,\s*path string,\s*size int64\)',lambda m:'func (s *Scanner) examine(ctx context.Context, path string, size int64, coverage *ArchiveCoverage316)')
ownrx316(r's\.examine\(ctx,\s*abs,\s*info\.Size\(\)\)',lambda m:'s.examine(ctx, abs, info.Size(), &report.ArchiveCoverage316)')
ownrx316(r'(?m)^([ \t]*)hits,\s*e\s*:=\s*s\.examine\(ctx,\s*path,\s*st\.Size\(\)\)',lambda m:m[1]+'var archive316 ArchiveCoverage316\n'+m[1]+'hits,e:=s.examine(ctx,path,st.Size(),&archive316)')
ownrx316(r'(hits,e:=s\.examine\(ctx,path,st\.Size\(\),&archive316\)\s*mu\.Lock\(\))',lambda m:m[1]+'\n report.ArchiveCoverage316.Merge(archive316)')
own316('if strings.EqualFold(filepath.Ext(path), ".zip") {','if zipCandidate316(path,b) {')
ownrx316(r'(archive,\s*e\s*:=\s*zip\.NewReader\(f,\s*size\))',lambda m:m[1]+'\n coverage.Containers++\n if e!=nil{coverage.OpenFailures++}')
ownrx316(r'if idx >= maxEntries \{\s*break\s*\}',lambda m:'if idx >= maxEntries { coverage.EntryBudgetStops++; break }')
ownrx316(r'if z\.FileInfo\(\)\.IsDir\(\) \|\| z\.UncompressedSize64 > uint64\(entryMax\) \|\| z\.UncompressedSize64 > uint64\(totalLimit\)-totalUncompressed \{\s*continue\s*\}',lambda m:'if z.FileInfo().IsDir(){continue}\n if z.UncompressedSize64>uint64(entryMax)||z.UncompressedSize64>uint64(totalLimit)-totalUncompressed{coverage.EntriesLimited++;continue}')
ownrx316(r'(r,\s*e\s*:=\s*z\.Open\(\)\s*if e != nil \{)\s*continue',lambda m:m[1]+'\n coverage.EntryReadFailures++;continue')
ownrx316(r'if int64\(len\(data\)\)\s*>\s*allowed\s*\{\s*break\s*\}',lambda m:'if int64(len(data))>allowed{coverage.EntriesLimited++;coverage.EntryBudgetStops++;break}')
ownrx316(r'(if e != nil \|\| uint64\(len\(data\)\)\s*!=\s*z\.UncompressedSize64 \{)\s*continue',lambda m:m[1]+'\n coverage.EntryReadFailures++;continue')
own316('sum := sha256.Sum256(data)','coverage.EntriesExamined++\n sum := sha256.Sum256(data)')
p.write_text(s,encoding="utf-8")
p=Path("engine/scan_upgrade.go");s=p.read_text(encoding="utf-8")
ownrx316(r'(?m)^([ \t]*r\.Partial\s*=.*)$',lambda m:m[1]+' || r.ArchiveCoverage316.Partial()')
p.write_text(s,encoding="utf-8")
p=Path("main.go");s=p.read_text(encoding="utf-8")
ownrx316(r'(r\.Findings\s*=\s*append\(r\.Findings,\s*one\.Findings\.\.\.\))',lambda m:m[1]+'\n r.ArchiveCoverage316.Merge(one.ArchiveCoverage316)')
p.write_text(s,encoding="utf-8")
for name in ("archive316.go","archive316_test.go","ioc316_test.go"):
    Path("engine",name).write_text(Path("lab",name+".txt").read_text(encoding="utf-8"),encoding="utf-8")
