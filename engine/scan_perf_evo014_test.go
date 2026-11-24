package engine

import (
 "context"
 "crypto/sha256"
 "encoding/hex"
 "math/rand"
 "os"
 "path/filepath"
 "testing"
)

func zLegacyEntropy(ext string, data []byte) bool {
 return len(data) >= 4096 && entropy(data) > 7.8 && riskyExtensions[ext]
}
func TestZMotorEntropyEquivalence(t *testing.T) {
 rng := rand.New(rand.NewSource(14))
 for _, n := range []int{0,128,4095,4096,65536,1048576} {
  p := make([]byte,n); rng.Read(p)
  q := make([]byte,n)
  for _, ext := range []string{".pdf",".exe",".zip",".tmp",".ps1",".js",".bat",".scr",""} {
   for _, data := range [][]byte{p,q} {
    if zLegacyEntropy(ext,data) != highEntropyScript(ext,data) {t.Fatalf("decision changed: %s %d",ext,n)}
   }
  }
 }
}
func TestZMotorFullHashAndScript(t *testing.T) {
 sc:=New()
 p:=filepath.Join(t.TempDir(),"sample.pdf")
 data:=make([]byte,65536)
 rand.New(rand.NewSource(19)).Read(data)
 if err:=os.WriteFile(p,data,0600);err!=nil{t.Fatal(err)}
 sum:=sha256.Sum256(data);hash:=hex.EncodeToString(sum[:]);sc.IOC[hash]="synthetic harmless test"
 hits,err:=sc.ExamineRegular(context.Background(),p,int64(len(data)))
 if err!=nil||len(hits)!=1||hits[0].SHA256!=hash||hits[0].Evidence!="ioc_match"{t.Fatal(hits,err)}
 script:=[]byte("Invoke-Expression ( harmless test fixture )")
 s:=analyze("test.ps1",script,"",map[string]string{})
 if s==nil||s.Risk<30{t.Fatal("script heuristic missing")}
}
var zEntropySink bool
var zFindingSink *Finding
func BenchmarkZMotorEntropyLegacy(b *testing.B) {
 p:=make([]byte,1<<20);rand.New(rand.NewSource(14)).Read(p)
 b.SetBytes(int64(len(p)));b.ReportAllocs();b.ResetTimer()
 for i:=0;i<b.N;i++{zEntropySink=zLegacyEntropy(".pdf",p)}
}
func BenchmarkZMotorEntropyOptimized(b *testing.B) {
 p:=make([]byte,1<<20);rand.New(rand.NewSource(14)).Read(p)
 b.ReportAllocs();b.ResetTimer()
 for i:=0;i<b.N;i++{zEntropySink=highEntropyScript(".pdf",p)}
}
func BenchmarkZMotorAnalyze(b *testing.B) {
 p:=make([]byte,1<<20);rand.New(rand.NewSource(14)).Read(p)
 ioc:=map[string]string{}
 b.SetBytes(int64(len(p)));b.ReportAllocs();b.ResetTimer()
 for i:=0;i<b.N;i++{zFindingSink=analyze("sample.pdf",p,"",ioc)}
}
