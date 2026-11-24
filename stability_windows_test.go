//go:build windows

package main
import("bytes";"sync";"testing")
func TestStabilityWorkers(t *testing.T){
 for n,w:=range map[int]int{0:1,1:1,2:1,3:3,4:4,8:4,32:4}{if got:=stabilityWorkers(n);got!=w{t.Fatalf("%d: %d != %d",n,got,w)}}
}
func TestStabilityLogBound(t *testing.T){
 var b bytes.Buffer;w:=&stabilityWriter{out:&b,remaining:128}
 var wg sync.WaitGroup
 for i:=0;i<16;i++{wg.Add(1);go func(){defer wg.Done();p:=bytes.Repeat([]byte("x"),32);if n,e:=w.Write(p);e!=nil||n!=len(p){t.Errorf("%d %v",n,e)}}()}
 wg.Wait()
 if b.Len()!=128||w.remaining!=0{t.Fatal(b.Len(),w.remaining)}
}
