package main
import("fmt";"os";"path/filepath";"syscall";"unsafe")
func drive310Args(root string)(string,error){
 if len(root)!=3||root[1]!=':'||root[2]!='\\'||!((root[0]>='A'&&root[0]<='Z')||(root[0]>='a'&&root[0]<='z')){return "",fmt.Errorf("Invalid local system volume")}
 return root[:2]+" /O /U /V",nil
}
type drive310Shell struct{
 Size,Mask uint32
 Window uintptr
 Verb,File,Parameters,Directory *uint16
 Show int32
 Instance,IDList uintptr
 Class *uint16
 ClassKey uintptr
 HotKey uint32
 Icon,Process uintptr
}
type drive310Power struct{AC,Flag,Percent,Reserved byte;Life,Full uint32}
// Explicit UAC for one fixed Microsoft tool. No elevation of ZEUVORYX itself.
func drive310Optimize()(uint32,error){
 k:=syscall.NewLazyDLL("kernel32.dll")
 var power drive310Power
 ok,_,_:=k.NewProc("GetSystemPowerStatus").Call(uintptr(unsafe.Pointer(&power)))
 if ok==0||power.AC!=1{return 0,fmt.Errorf("Connect the computer to AC power before optimization")}
 var win [32768]uint16
 n,_,_:=k.NewProc("GetWindowsDirectoryW").Call(uintptr(unsafe.Pointer(&win[0])),uintptr(len(win)))
 if n<3||n>=uintptr(len(win)){return 0,fmt.Errorf("Windows drive unavailable")}
 root:=syscall.UTF16ToString(win[:3]);args,e:=drive310Args(root);if e!=nil{return 0,e}
 drive,_:=syscall.UTF16PtrFromString(root)
 typ,_,_:=k.NewProc("GetDriveTypeW").Call(uintptr(unsafe.Pointer(drive)));if typ!=3{return 0,fmt.Errorf("System drive is not a fixed local volume")}
 var sys [32768]uint16
 n,_,_ = k.NewProc("GetSystemDirectoryW").Call(uintptr(unsafe.Pointer(&sys[0])),uintptr(len(sys)))
 if n==0||n>=uintptr(len(sys)){return 0,fmt.Errorf("Windows system directory unavailable")}
 dir:=syscall.UTF16ToString(sys[:int(n)]);path:=filepath.Join(dir,"defrag.exe")
 st,e:=os.Stat(path);if e!=nil||!st.Mode().IsRegular(){return 0,fmt.Errorf("Native optimizer unavailable")}
 verb,_:=syscall.UTF16PtrFromString("runas");file,_:=syscall.UTF16PtrFromString(path);param,_:=syscall.UTF16PtrFromString(args);working,_:=syscall.UTF16PtrFromString(dir)
 x:=drive310Shell{Mask:0x40|0x100,Verb:verb,File:file,Parameters:param,Directory:working,Show:0};x.Size=uint32(unsafe.Sizeof(x))
 ok,_,callErr:=syscall.NewLazyDLL("shell32.dll").NewProc("ShellExecuteExW").Call(uintptr(unsafe.Pointer(&x)))
 if ok==0||x.Process==0{return 0,fmt.Errorf("Windows did not authorize or start optimization: %v",callErr)}
 defer syscall.CloseHandle(syscall.Handle(x.Process))
 wait,_,_:=k.NewProc("WaitForSingleObject").Call(x.Process,0xffffffff)
 if wait!=0{return 0,fmt.Errorf("Optimization completion not verified")}
 var code uint32;ok,_,_ = k.NewProc("GetExitCodeProcess").Call(x.Process,uintptr(unsafe.Pointer(&code)))
 if ok==0{return 0,fmt.Errorf("Optimization result unavailable")}
 if code!=0{return code,fmt.Errorf("Windows optimizer returned code %d",code)}
 return code,nil
}
