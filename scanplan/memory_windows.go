//go:build windows

package scanplan
import("syscall";"unsafe")
type memoryStatus struct{Length,Load uint32;Total,Available,PageTotal,PageAvailable,VirtualTotal,VirtualAvailable,Extended uint64}
var memoryProc=syscall.NewLazyDLL("kernel32.dll").NewProc("GlobalMemoryStatusEx")
func Memory()(bool,uint64,uint64){
 m:=memoryStatus{Length:uint32(unsafe.Sizeof(memoryStatus{}))}
 ok,_,_:=memoryProc.Call(uintptr(unsafe.Pointer(&m)))
 return ok!=0,m.Total,m.Available
}
