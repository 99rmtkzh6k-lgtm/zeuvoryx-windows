//go:build windows

package main

import (
 "log"
 "runtime"
 "sync"
 "syscall"
 "unsafe"
)

const scanESContinuous = 0x80000000
const scanESSystemRequired = 0x00000001

var scanPowerKernel = syscall.NewLazyDLL("kernel32.dll")
var scanPowerState = scanPowerKernel.NewProc("SetThreadExecutionState")
var scanPowerStatus = scanPowerKernel.NewProc("GetSystemPowerStatus")

type scanSystemPowerStatus struct {
 ACLineStatus byte
 BatteryFlag byte
 BatteryLifePercent byte
 SystemStatusFlag byte
 BatteryLifeTime uint32
 BatteryFullLifeTime uint32
}

// Avoid keeping laptops awake on battery or when the power source is unknown.
func scanOnAC() bool {
 var s scanSystemPowerStatus
 if scanPowerStatus.Find()!=nil{return false}
 ok,_,_:=scanPowerStatus.Call(uintptr(unsafe.Pointer(&s)))
 runtime.KeepAlive(&s)
 return ok!=0 && s.ACLineStatus==1
}
func setScanPowerState(flags uintptr) bool {
 if scanPowerState.Find()!=nil{return false}
 v,_,_:=scanPowerState.Call(flags)
 return v!=0
}

// The native request is thread-local. Acquire and clear on the same
// dedicated OS thread. No periodic polling or display wake request.
func beginScanPowerRequest(onAC func()bool, set func(uintptr)bool) func() {
 if !onAC(){log.Print("power full_scan request skipped: battery_or_unknown");return func(){}}
 ready:=make(chan bool,1)
 stop:=make(chan struct{})
 done:=make(chan struct{})
 go func(){
  defer close(done)
  runtime.LockOSThread()
  // If clearing fails, leave this goroutine locked: its OS thread is retired.
  if !set(scanESContinuous|scanESSystemRequired){
   runtime.UnlockOSThread()
   ready<-false
   return
  }
  ready<-true
  <-stop
  if set(scanESContinuous){runtime.UnlockOSThread();log.Print("power full_scan request released")}else{log.Print("power full_scan request clear failed; retire thread")}
 }()
 if !<-ready {<-done;log.Print("power full_scan request unavailable");return func(){}}
 log.Print("power full_scan system_required active; display unchanged")
 var once sync.Once
 return func(){once.Do(func(){close(stop);<-done})}
}

func beginFullScanPower() func() {
 return beginScanPowerRequest(scanOnAC,setScanPowerState)
}
