//go:build windows

package main
import("log";"runtime";"sync";"syscall";"time";"unsafe")
var pKernel=syscall.NewLazyDLL("kernel32.dll")
var pMem=pKernel.NewProc("GlobalMemoryStatusEx")
var pTimes=pKernel.NewProc("GetSystemTimes")
var pDisk=pKernel.NewProc("GetDiskFreeSpaceExW")
var pWin=pKernel.NewProc("GetWindowsDirectoryW")
var pDrives=pKernel.NewProc("GetLogicalDrives")
var pDriveType=pKernel.NewProc("GetDriveTypeW")
var pVolumeID=pKernel.NewProc("GetVolumeNameForVolumeMountPointW")
// Fixed lettered volumes only; deduplicate by volume GUID, no disk traversal.
func pLocalStorage()(uint64,uint64,int,int){
 mask,_,_:=pDrives.Call();if mask==0{return 0,0,0,1}
 seen:=map[string]bool{};var total,available uint64;count,omitted:=0,0
 for i:=0;i<26;i++{if mask&(1<<i)==0{continue};root:=[4]uint16{uint16('A'+i),':','\\',0}
  kind,_,_:=pDriveType.Call(uintptr(unsafe.Pointer(&root[0])));if kind!=3{continue}
  var id [100]uint16;ok,_,_:=pVolumeID.Call(uintptr(unsafe.Pointer(&root[0])),uintptr(unsafe.Pointer(&id[0])),uintptr(len(id)));if ok==0{omitted++;continue}
  key:=syscall.UTF16ToString(id[:]);if key==""{omitted++;continue};if seen[key]{continue};seen[key]=true
  var caller,t,f uint64;ok,_,_=pDisk.Call(uintptr(unsafe.Pointer(&root[0])),uintptr(unsafe.Pointer(&caller)),uintptr(unsafe.Pointer(&t)),uintptr(unsafe.Pointer(&f)))
  if ok==0||t==0||caller>t||total>^uint64(0)-t||available>^uint64(0)-caller{omitted++;continue}
  total+=t;available+=caller;count++
 }
 return total,available,count,omitted
}
type pMemory struct{Length,Load uint32;Total,Available,PageTotal,PageAvailable,VirtualTotal,VirtualAvailable,Extended uint64}
type pTicks struct{Idle,Kernel,User uint64}
func pPercent(total,free uint64)*float64{if total==0||free>total{return nil};v:=100*float64(total-free)/float64(total);return &v}
func pCPU(a,b pTicks)*float64{if b.Idle<a.Idle||b.Kernel<a.Kernel||b.User<a.User{return nil};k,u,i:=b.Kernel-a.Kernel,b.User-a.User,b.Idle-a.Idle;if k>^uint64(0)-u{return nil};return pPercent(k+u,i)}
func pFT(t syscall.Filetime)uint64{return uint64(t.HighDateTime)<<32|uint64(t.LowDateTime)}
type pSampler struct{mu sync.Mutex;last,diskAt time.Time;cache map[string]any;prior pTicks;priorOK bool;ready bool}
func(p *pSampler) sample()map[string]any{
 p.mu.Lock();defer p.mu.Unlock();now:=time.Now()
 if !p.last.IsZero()&&now.Sub(p.last)<2*time.Second{return p.cache}
 x:=map[string]any{"at":now.Format(time.RFC3339),"logical_cpu":runtime.NumCPU(),"cpu":nil,"ram_total":uint64(0),"ram_available":uint64(0),"ram_percent":nil,"disk_root":"","disk_total":uint64(0),"disk_free":uint64(0),"disk_percent":nil}
 m:=pMemory{Length:uint32(unsafe.Sizeof(pMemory{}))}
 if ok,_,_:=pMem.Call(uintptr(unsafe.Pointer(&m)));ok!=0{if v:=pPercent(m.Total,m.Available);v!=nil{x["ram_total"],x["ram_available"],x["ram_percent"]=m.Total,m.Available,*v}}
 var idle,kernel,user syscall.Filetime
 if ok,_,_:=pTimes.Call(uintptr(unsafe.Pointer(&idle)),uintptr(unsafe.Pointer(&kernel)),uintptr(unsafe.Pointer(&user)));ok!=0{
 b:=pTicks{pFT(idle),pFT(kernel),pFT(user)};if p.priorOK{if v:=pCPU(p.prior,b);v!=nil{x["cpu"]=*v}};p.prior,p.priorOK=b,true
 }else{p.priorOK=false}
 if p.diskAt.IsZero()||now.Sub(p.diskAt)>=15*time.Second{
 var buf [260]uint16;n,_,_:=pWin.Call(uintptr(unsafe.Pointer(&buf[0])),uintptr(len(buf)))
 if n>=3&&n<uintptr(len(buf))&&buf[1]==':'&&buf[2]=='\\'{
 root:=syscall.UTF16ToString(buf[:3]);ptr,_:=syscall.UTF16PtrFromString(root);var caller,total,free uint64
 if ok,_,_:=pDisk.Call(uintptr(unsafe.Pointer(ptr)),uintptr(unsafe.Pointer(&caller)),uintptr(unsafe.Pointer(&total)),uintptr(unsafe.Pointer(&free)));ok!=0{if v:=pPercent(total,caller);v!=nil{x["disk_root"],x["disk_total"],x["disk_free"],x["disk_percent"]=root,total,caller,*v}}
 };total,available,count,omitted:=pLocalStorage();x["local_total"],x["local_free"],x["local_count"],x["local_omitted"]=total,available,count,omitted;p.diskAt=now
 }else{for _,k:=range []string{"disk_root","disk_total","disk_free","disk_percent","local_total","local_free","local_count","local_omitted"}{x[k]=p.cache[k]}}
 p.last,p.cache=now,x;p.ready=x["cpu"]!=nil&&x["ram_percent"]!=nil;return x
}
func(p *pSampler) rendered()bool{p.mu.Lock();defer p.mu.Unlock();if !p.ready{return false};log.Print("stability performance_ui_ready revision=R5");return true}
