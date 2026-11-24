package engine

import (
 "bytes"
 "context"
 "crypto/sha256"
 "encoding/hex"
 "errors"
 "testing"
)

type reader308 struct {
 r *bytes.Reader
 max int
 calls int
 cancel context.CancelFunc
}
func (r *reader308) Read(p []byte)(int,error){
 if len(p)>r.max {r.max=len(p)}
 r.calls++
 n,e:=r.r.Read(p)
 if r.cancel!=nil {r.cancel()}
 return n,e
}
func Test308CompleteHashBoundedReads(t *testing.T){
 for _,size:=range []int{0,1,262144,262145,1<<20,34<<20}{
  data:=bytes.Repeat([]byte{0x5a},size)
  for _,limit:=range []int64{64<<10,32<<20}{
   r:=&reader308{r:bytes.NewReader(data)}
   sample,h,e:=readScanContent(context.Background(),r,int64(size),limit)
   n:=int64(size);if n>limit{n=limit}
   sum:=sha256.Sum256(data)
   if e!=nil||!bytes.Equal(sample,data[:int(n)])||h!=hex.EncodeToString(sum[:])||r.max>256<<10{t.Fatalf("size=%d limit=%d max=%d hash=%s err=%v",size,limit,r.max,h,e)}
  }
 }
}
func Test308CancellationBetweenPrefixChunks(t *testing.T){
 ctx,cancel:=context.WithCancel(context.Background());defer cancel()
 data:=make([]byte,2<<20)
 r:=&reader308{r:bytes.NewReader(data),cancel:cancel}
 _,_,e:=readScanContent(ctx,r,int64(len(data)),32<<20)
 if !errors.Is(e,context.Canceled)||r.calls!=1||r.max>256<<10{t.Fatalf("calls=%d max=%d err=%v",r.calls,r.max,e)}
}
func Test308EntropyEligibility(t *testing.T){
 data:=make([]byte,1<<16);for i:=range data{data[i]=byte(i%256)}
 for _,ext:=range []string{".bin",".txt",".ps1",".js",".exe"}{
  before:=len(data)>=4096&&entropy(data)>7.8&&riskyExtensions[ext]
  after:=len(data)>=4096&&riskyExtensions[ext]&&entropy(data)>7.8
  if before!=after{t.Fatal(ext)}
 }
}
func Benchmark308EntropyLegacy(b *testing.B){
 data:=make([]byte,8<<20);for i:=range data{data[i]=byte(i%256)}
 b.ResetTimer()
 for i:=0;i<b.N;i++{if len(data)>=4096&&entropy(data)>7.8&&riskyExtensions[".bin"]{b.Fatal("unexpected")}}
}
func Benchmark308EntropyCandidate(b *testing.B){
 data:=make([]byte,8<<20);for i:=range data{data[i]=byte(i%256)}
 b.ResetTimer()
 for i:=0;i<b.N;i++{if len(data)>=4096&&riskyExtensions[".bin"]&&entropy(data)>7.8{b.Fatal("unexpected")}}
}
