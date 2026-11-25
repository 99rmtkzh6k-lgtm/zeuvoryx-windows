//go:build windows

// Package nativeamsi provides a bounded, read-only Windows AMSI client.
package nativeamsi

import (
 "fmt"
 "path/filepath"
 "runtime"
 "syscall"
 "unsafe"
)

const MaxContent = 8 * 1024 * 1024

type Verdict struct {
 Raw uint32
 Status string
}

// Interpret preserves provider scores; administrator policy is not malware.
func Interpret(raw uint32) Verdict {
 status := "provider_score"
 switch {
 case raw >= 32768: status = "malware"
 case raw >= 16384 && raw <= 20479: status = "policy_block"
 case raw == 0: status = "clean_by_provider"
 case raw == 1: status = "not_detected_by_provider"
 }
 return Verdict{Raw:raw, Status:status}
}

// Scan submits in-memory content to the configured Windows provider.
// It does not execute content, change security policies or guarantee coverage.
// An error means inconclusive, never clean. No truncation is allowed.
func Scan(content []byte, name string) (Verdict, error) {
 if len(content)==0 || len(content)>MaxContent { return Verdict{},fmt.Errorf("AMSI: content outside 1..%d bytes",MaxContent) }
 source,err:=syscall.UTF16PtrFromString(name);if err!=nil{return Verdict{},err}
 app,_:=syscall.UTF16PtrFromString("ZEUVORYX LAB")
 kernel:=syscall.NewLazyDLL("kernel32.dll")
 getdir:=kernel.NewProc("GetSystemDirectoryW")
 if err:=getdir.Find();err!=nil{return Verdict{},err}
 var path [32768]uint16
 n,_,e:=getdir.Call(uintptr(unsafe.Pointer(&path[0])),uintptr(len(path)))
 if n==0||n>=uintptr(len(path)){return Verdict{},fmt.Errorf("AMSI: system directory: %v",e)}
 dll:=syscall.NewLazyDLL(filepath.Join(syscall.UTF16ToString(path[:n]),"amsi.dll"))
 init:=dll.NewProc("AmsiInitialize")
 scan:=dll.NewProc("AmsiScanBuffer")
 finish:=dll.NewProc("AmsiUninitialize")
 for _,p:=range []*syscall.LazyProc{init,scan,finish}{if err:=p.Find();err!=nil{return Verdict{},err}}
 var ctx uintptr
 hr,_,_:=init.Call(uintptr(unsafe.Pointer(app)),uintptr(unsafe.Pointer(&ctx)))
 runtime.KeepAlive(app)
 if int32(hr)<0||ctx==0{return Verdict{},fmt.Errorf("AMSI initialize HRESULT %#08x",uint32(hr))}
 defer finish.Call(ctx)
 var raw uint32
 hr,_,_=scan.Call(ctx,uintptr(unsafe.Pointer(&content[0])),uintptr(len(content)),uintptr(unsafe.Pointer(source)),0,uintptr(unsafe.Pointer(&raw)))
 runtime.KeepAlive(content);runtime.KeepAlive(source)
 return scanResult(hr,raw)
}

// A provider that did not return S_OK has not supplied an authoritative verdict.
// In particular ERROR_NOT_READY is a coverage failure, not a clean result.
func scanResult(hr uintptr, raw uint32) (Verdict,error) {
 if uint32(hr)==0x80070015 {return Verdict{},fmt.Errorf("AMSI provider not ready (HRESULT 0x80070015); result inconclusive")}
 if uint32(hr)!=0 {return Verdict{},fmt.Errorf("AMSI scan HRESULT %#08x; result inconclusive",uint32(hr))}
 return Interpret(raw),nil
}
