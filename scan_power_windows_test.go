//go:build windows

package main

import (
 "reflect"
 "testing"
)
func TestFullScanPowerLifecycle(t *testing.T){
 for i:=0;i<5;i++{
  var calls []uintptr
  end:=beginScanPowerRequest(func()bool{return true},func(v uintptr)bool{calls=append(calls,v);return true})
  end();end()
  if !reflect.DeepEqual(calls,[]uintptr{scanESContinuous|scanESSystemRequired,scanESContinuous}){t.Fatalf("unexpected request lifecycle: %v",calls)}
 }
}
func TestFullScanPowerBattery(t *testing.T){
 called:=false
 end:=beginScanPowerRequest(func()bool{return false},func(v uintptr)bool{called=true;return true})
 end()
 if called{t.Fatal("request on battery")}
}
func TestFullScanPowerFailure(t *testing.T){
 var calls []uintptr
 end:=beginScanPowerRequest(func()bool{return true},func(v uintptr)bool{calls=append(calls,v);return false})
 end()
 if !reflect.DeepEqual(calls,[]uintptr{scanESContinuous|scanESSystemRequired}){t.Fatalf("unexpected failed request: %v",calls)}
}
func TestFullScanPowerClearFailure(t *testing.T){
 var calls []uintptr
 end:=beginScanPowerRequest(func()bool{return true},func(v uintptr)bool{calls=append(calls,v);return v!=scanESContinuous})
 end();end()
 if len(calls)!=2{t.Fatalf("unexpected clear attempts: %v",calls)}
}
