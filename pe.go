package engine

import (
	"bytes"
	"debug/pe"
	"encoding/binary"
	"fmt"
	"path/filepath"
	"strings"
)

// analyzePE is an explainable, conservative static feature detector, not a trained ML model.
// No scanned executable is executed by this analyzer.
func analyzePE(path string, data []byte) *Finding {
	if len(data) < 2 || !bytes.Equal(data[:2], []byte{'M', 'Z'}) {
		return nil
	}
	f, err := pe.NewFile(bytes.NewReader(data))
	if err != nil {
		return nil
	}
	defer f.Close()
	var score int
	var reasons []string
	imports, err := boundedImportedSymbols(f, data)
	if err == nil {
		seen := map[string]bool{}
		for _, imp := range imports {
			seen[strings.ToLower(imp)] = true
		}
		has := func(parts ...string) bool {
			for key := range seen {
				for _, p := range parts {
					if strings.Contains(key, p) {
						return true
					}
				}
			}
			return false
		}
		// These APIs have many legitimate uses. Require combinations, not isolated calls.
		if has("virtualallocex", "ntallocatevirtualmemory") && has("writeprocessmemory", "ntwritevirtualmemory") && has("createremotethread", "ntcreatethreadex") {
			score += 45
			reasons = append(reasons, "PE: combinação de APIs de injeção remota (indicador, não prova)")
		}
		if has("isdebuggerpresent") && has("checkremotedebuggerpresent") {
			score += 10
			reasons = append(reasons, "PE: verificações múltiplas contra depuração")
		}
	}
	ext := strings.ToLower(filepath.Ext(path))
	if ext == ".pdf" || ext == ".jpg" || ext == ".png" || ext == ".txt" {
		score += 35
		reasons = append(reasons, "Assinatura PE em arquivo com extensão não executável")
	}
	if score == 0 {
		return nil
	}
	return &Finding{Risk: score, Reasons: reasons}
}

// boundedImportedSymbols resolves each RVA independently against raw file data.
// It avoids repeated section copies and rejects truncated or out-of-range tables.
func boundedImportedSymbols(f *pe.File, data []byte) ([]string, error) {
 var dir pe.DataDirectory
 var width uint64
 switch h := f.OptionalHeader.(type) {
 case *pe.OptionalHeader32:
  if h.NumberOfRvaAndSizes < 2 { return nil,nil }; dir=h.DataDirectory[1]; width=4
 case *pe.OptionalHeader64:
  if h.NumberOfRvaAndSizes < 2 { return nil,nil }; dir=h.DataDirectory[1]; width=8
 default: return nil,nil
 }
 if dir.VirtualAddress==0 && dir.Size==0 { return nil,nil }
 bad := func() error { return fmt.Errorf("PE import directory incomplete or outside analysis budget") }
 if len(f.Sections)>96 || dir.VirtualAddress==0 || dir.Size<20 || uint64(dir.VirtualAddress)+uint64(dir.Size)>1<<32 { return nil,bad() }
 slice := func(rva uint64, n uint64) []byte {
  if rva>=1<<32 { return nil }
  for _,s:=range f.Sections {
   base:=uint64(s.VirtualAddress); if rva<base || s.Offset==0 {continue}; delta:=rva-base
   if delta>=uint64(s.Size) || n>uint64(s.Size)-delta {continue}
   start:=uint64(s.Offset)+delta
   if start>uint64(len(data)) || n>uint64(len(data))-start {return nil}
   return data[start:start+n]
  }; return nil
 }
 cstr := func(rva uint64) (string,bool) {
  if rva>=1<<32{return "",false}
  for _,s:=range f.Sections {
   base:=uint64(s.VirtualAddress); if rva<base || s.Offset==0 {continue}; delta:=rva-base
   if delta>=uint64(s.Size){continue}; start:=uint64(s.Offset)+delta
   if start>=uint64(len(data)){return "",false}; n:=uint64(s.Size)-delta
   if n>uint64(len(data))-start{n=uint64(len(data))-start}; if n>4096{n=4096}
   b:=data[start:start+n]; end:=bytes.IndexByte(b,0); if end<=0{return "",false}; return string(b[:end]),true
  }; return "",false
 }
 out:=make([]string,0,64); remaining:=65536; textBudget:=2<<20
 for i:=uint64(0); i<4096 && (i+1)*20<=uint64(dir.Size); i++ {
  d:=slice(uint64(dir.VirtualAddress)+i*20,20); if d==nil{return nil,bad()}
  if bytes.Equal(d,make([]byte,20)){return out,nil}
  thunk:=uint64(binary.LittleEndian.Uint32(d)); name:=uint64(binary.LittleEndian.Uint32(d[12:])); if thunk==0{thunk=uint64(binary.LittleEndian.Uint32(d[16:]))}
  dll,ok:=cstr(name); if thunk==0 || name==0 || !ok{return nil,bad()}
  for j:=uint64(0);;j++ {
   if remaining==0{return nil,bad()}; remaining--
   b:=slice(thunk+j*width,width); if b==nil{return nil,bad()}; var v uint64
   if width==8{v=binary.LittleEndian.Uint64(b)}else{v=uint64(binary.LittleEndian.Uint32(b))}
   if v==0{break}; if v&(uint64(1)<<(width*8-1))!=0{continue}; if v>0xfffffffd{return nil,bad()}
   symbol,ok:=cstr(v+2); if !ok{return nil,bad()}; cost:=len(symbol)+1+len(dll); if cost>textBudget{return nil,bad()}; textBudget-=cost; out=append(out,symbol+":"+dll)
  }
 }
 return nil,bad()
}
