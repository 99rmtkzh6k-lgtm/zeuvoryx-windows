package main

import (
 "fmt"
 "path/filepath"
 "runtime"
 "syscall"
 "unsafe"
)
type reportGUID315 struct{a uint32;b,c uint16;d [8]byte}
type reportDataObject315 struct{vtable *[3]uintptr}
var dataIID315=reportGUID315{0x0000010e,0,0,[8]byte{0xc0,0,0,0,0,0,0,0x46}}
func hresult315(op string,r uintptr)error{if int32(r)<0{return fmt.Errorf("%s: HRESULT 0x%08X",op,uint32(r))};return nil}
func reportShell315(drag bool)(bool,error){
 path,e:=lastReport315();if e!=nil{return false,e}
 // A dedicated STA owns the OLE objects; no external window or process is opened.
 runtime.LockOSThread();defer runtime.UnlockOSThread()
 ole:=syscall.NewLazyDLL("ole32.dll");shell:=syscall.NewLazyDLL("shell32.dll");user:=syscall.NewLazyDLL("user32.dll")
 r,_,_:=ole.NewProc("OleInitialize").Call(0);if e=hresult315("OleInitialize",r);e!=nil{return false,e};defer ole.NewProc("OleUninitialize").Call()
 parse:=shell.NewProc("SHParseDisplayName")
 fullName,e:=syscall.UTF16PtrFromString(path);if e!=nil{return false,e}
 parentName,e:=syscall.UTF16PtrFromString(filepath.Dir(path));if e!=nil{return false,e}
 var full,parent uintptr
 var obj *reportDataObject315
 r,_,_ = parse.Call(uintptr(unsafe.Pointer(fullName)),0,uintptr(unsafe.Pointer(&full)),0,0);if e=hresult315("Arquivo",r);e!=nil{return false,e};if full==0{return false,fmt.Errorf("Arquivo indisponível")};defer ole.NewProc("CoTaskMemFree").Call(full)
 r,_,_ = parse.Call(uintptr(unsafe.Pointer(parentName)),0,uintptr(unsafe.Pointer(&parent)),0,0);if e=hresult315("Pasta",r);e!=nil{return false,e};if parent==0{return false,fmt.Errorf("Pasta indisponível")};defer ole.NewProc("CoTaskMemFree").Call(parent)
 child,_,_:=shell.NewProc("ILFindLastID").Call(full);if child==0{return false,fmt.Errorf("Identificador de arquivo indisponível")}
 r,_,_=shell.NewProc("SHCreateDataObject").Call(parent,1,uintptr(unsafe.Pointer(&child)),0,uintptr(unsafe.Pointer(&dataIID315)),uintptr(unsafe.Pointer(&obj)))
 if e=hresult315("Objeto de transferência",r);e!=nil{return false,e};if obj==nil{return false,fmt.Errorf("Objeto de transferência indisponível")}
 defer func(){syscall.SyscallN(obj.vtable[2],uintptr(unsafe.Pointer(obj)));runtime.KeepAlive(obj)}()
 if drag {
  pressed,_,_:=user.NewProc("GetAsyncKeyState").Call(1);if uint16(pressed)&0x8000==0{return false,fmt.Errorf("Mantenha o botão esquerdo pressionado e arraste")}
  var effect uint32
  r,_,_=shell.NewProc("SHDoDragDrop").Call(0,uintptr(unsafe.Pointer(obj)),0,1,uintptr(unsafe.Pointer(&effect)))
  if e=hresult315("Arrastar arquivo",r);e!=nil{return false,e}
  return uint32(r)==0x00040100&&effect==1,nil
 }
 r,_,_=ole.NewProc("OleSetClipboard").Call(uintptr(unsafe.Pointer(obj)));if e=hresult315("Copiar arquivo",r);e!=nil{return false,e}
 r,_,_=ole.NewProc("OleFlushClipboard").Call();if e=hresult315("Confirmar cópia",r);e!=nil{return false,e}
 return true,nil
}
