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
 owner,ok:=gate.Begin("same")
 if !ok {t.Fatal("initial reservation failed")}
 var wg sync.WaitGroup
 rejected:=make(chan bool,64)
 for i:=0;i<64;i++ {
  wg.Add(1)
  go func(){
   defer wg.Done()
   _,ok:=gate.Begin("same")
   rejected<-!ok
  }()
 }
 wg.Wait()
 close(rejected)
 for denied:=range rejected {if !denied {t.Fatal("duplicate reservation succeeded")}}
 owner()
 next,ok:=gate.Begin("same")
 if !ok {t.Fatal("reservation not released")}
 next()
}
func TestDownloadGateEmptyID(t *testing.T) {
 var gate DownloadGate
 if _,ok:=gate.Begin("");ok {t.Fatal("empty id accepted")}
}
