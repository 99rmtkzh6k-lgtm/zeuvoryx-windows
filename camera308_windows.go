//go:build windows

package main

import (
 "encoding/binary"
 "sync"
 "syscall"
 "time"
)

type cameraEvidence struct { Start, Stop uint64; Valid bool }
func cameraClassify(rows []cameraEvidence, incomplete bool) string {
 usable:=0
 for _,r:=range rows{
  if !r.Valid || r.Start==0 || (r.Stop!=0 && r.Stop<r.Start){incomplete=true;continue}
  usable++
  if r.Stop==0 {return "activity"}
 }
 if incomplete||usable==0{return "unknown"}
 return "no_activity_recorded"
}
func cameraOpen(parent syscall.Handle,path string)(syscall.Handle,error){
 p,e:=syscall.UTF16PtrFromString(path);if e!=nil{return 0,e}
 var h syscall.Handle
 e=syscall.RegOpenKeyEx(parent,p,0,syscall.KEY_READ,&h);return h,e
}
func cameraQword(h syscall.Handle,name string)(uint64,bool){
 p,e:=syscall.UTF16PtrFromString(name);if e!=nil{return 0,false}
 var b [8]byte;var kind uint32;n:=uint32(len(b))
 e=syscall.RegQueryValueEx(h,p,nil,&kind,&b[0],&n)
 if e!=nil||kind!=syscall.REG_QWORD||n!=8{return 0,false}
 return binary.LittleEndian.Uint64(b[:]),true
}
// Bounded current-user ConsentStore metadata only. Undocumented format;
// cannot prove webcam state, device presence, malware or protection.
func cameraRead() (rows []cameraEvidence,incomplete bool) {
 root,e:=cameraOpen(syscall.HKEY_CURRENT_USER,`SOFTWARE\Microsoft\Windows\CurrentVersion\CapabilityAccessManager\ConsentStore\webcam`)
 if e!=nil{return nil,true};defer syscall.RegCloseKey(root)
 budget:=512
 var visit func(syscall.Handle,int)
 visit=func(h syscall.Handle,depth int){
  for idx:=uint32(0);;idx++{
   if budget<=0{incomplete=true;return}
   var name [1024]uint16;n:=uint32(len(name))
   e:=syscall.RegEnumKeyEx(h,idx,&name[0],&n,nil,nil,nil,nil)
   if e==syscall.Errno(259){return};if e!=nil{incomplete=true;return}
   budget--
   key:=syscall.UTF16ToString(name[:n]);child,e:=cameraOpen(h,key)
   if e!=nil{incomplete=true;continue}
   if depth==0&&key=="NonPackaged"{visit(child,1)}else{
    start,a:=cameraQword(child,"LastUsedTimeStart");stop,b:=cameraQword(child,"LastUsedTimeStop")
    rows=append(rows,cameraEvidence{start,stop,a&&b})
   }
   syscall.RegCloseKey(child)
  }
 }
 visit(root,0);return
}
type cameraSampler struct {mu sync.Mutex;last time.Time;state string}
func (c *cameraSampler) sample()string{
 c.mu.Lock();defer c.mu.Unlock()
 now:=time.Now();if !c.last.IsZero()&&now.Sub(c.last)<10*time.Second{return c.state}
 rows,incomplete:=cameraRead();c.state=cameraClassify(rows,incomplete);c.last=now
 return c.state
}
