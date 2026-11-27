package updater

import (
 "sync"
 "testing"
)

func TestDownloadGateRejectsDuplicate(t *testing.T) {
 var gate DownloadGate
 release, ok := gate.Begin("build50-sha.zip")
 if !ok { t.Fatal("first download denied") }
 if _, ok := gate.Begin("build50-sha.zip"); ok { t.Fatal("duplicate accepted") }
 other, ok := gate.Begin("another.zip")
 if !ok { t.Fatal("independent artifact denied") }
 other()
 release()
 release()
 done, ok := gate.Begin("build50-sha.zip")
 if !ok { t.Fatal("lock not released") }
 done()
}
func TestDownloadGateConcurrent(t *testing.T) {
 var gate DownloadGate
 var wg sync.WaitGroup
 start := make(chan struct{})
 successes := make(chan bool,64)
 hold := make(chan struct{})
 for i:=0;i<64;i++ {
  wg.Add(1)
  go func(){
   defer wg.Done()
   <-start
   release,ok:=gate.Begin("same")
   successes<-ok
   if ok { <-hold; release() }
  }()
 }
 close(start)
 close(hold)
 wg.Wait()
 close(successes)
 count:=0
 for ok:=range successes { if ok { count++ } }
 if count!=1 { t.Fatalf("expected one winner, got %d",count) }
}
func TestDownloadGateEmptyID(t *testing.T) {
 var gate DownloadGate
 if _,ok:=gate.Begin("");ok {t.Fatal("empty id accepted")}
}
