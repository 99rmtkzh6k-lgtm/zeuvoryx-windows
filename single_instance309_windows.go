//go:build windows
package main
import("fmt";"os";"strings";"syscall";"unsafe")
var instanceKernel=syscall.NewLazyDLL("kernel32.dll")
var createInstanceMutex=instanceKernel.NewProc("CreateMutexW")
var instanceSession=instanceKernel.NewProc("ProcessIdToSessionId")
const instanceName="Local\\ZEUVORYX.Desktop.Lab303"
func legacyInstancePresent()(bool,error){
 var session uint32
 ok,_,e:=instanceSession.Call(uintptr(os.Getpid()),uintptr(unsafe.Pointer(&session)))
 if ok==0{return false,fmt.Errorf("session: %v",e)}
 h,e:=syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS,0);if e!=nil{return false,e};defer syscall.CloseHandle(h)
 var p syscall.ProcessEntry32;p.Size=uint32(unsafe.Sizeof(p))
 if e=syscall.Process32First(h,&p);e!=nil{return false,e}
 for {
  if p.ProcessID!=uint32(os.Getpid())&&strings.EqualFold(syscall.UTF16ToString(p.ExeFile[:]),"ZEUVORYX.exe"){
   var other uint32;ok,_,err:=instanceSession.Call(uintptr(p.ProcessID),uintptr(unsafe.Pointer(&other)))
   if ok==0{return false,fmt.Errorf("existing ZEUVORYX session: %v",err)}
   if other==session{return true,nil}
  }
  e=syscall.Process32Next(h,&p);if e==syscall.ERROR_NO_MORE_FILES{return false,nil};if e!=nil{return false,e}
 }
}
func acquireDesktopInstance()(func(),error){
 name,e:=syscall.UTF16PtrFromString(instanceName);if e!=nil{return nil,e}
 h,_,last:=createInstanceMutex.Call(0,0,uintptr(unsafe.Pointer(name)))
 if h==0{return nil,fmt.Errorf("mutex: %v",last)}
 close:=func(){_ = syscall.CloseHandle(syscall.Handle(h))}
 if last==syscall.ERROR_ALREADY_EXISTS{close();return nil,fmt.Errorf("ZEUVORYX já está aberto nesta sessão. Feche a outra versão antes de continuar.")}
 busy,e:=legacyInstancePresent();if e!=nil{close();return nil,e}
 if busy{close();return nil,fmt.Errorf("Outra versão do ZEUVORYX está aberta. Encerre a varredura e feche-a antes de continuar.")}
 return close,nil
}
