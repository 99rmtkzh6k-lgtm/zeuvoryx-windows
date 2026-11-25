package main
import (
 "fmt"
 "os"
 "path/filepath"
 "strings"
 "syscall"
 "time"
 "unsafe"
)
const opt310Limit=200
const opt310Enumerate=5000
const opt310UnsafeAttrs=0x1|0x2|0x4|0x10|0x400|0x1000|0x4000
type opt310Item struct{ID int `json:"id"`;Name string `json:"name"`;Bytes uint64 `json:"bytes"`;Modified string `json:"modified"`;path string;meta syscall.ByHandleFileInformation}
type opt310Plan struct{Items []opt310Item `json:"items"`;Bytes uint64 `json:"bytes"`;Truncated bool `json:"truncated"`;created time.Time;root string}
type opt310Result struct{Removed int `json:"removed"`;Bytes uint64 `json:"bytes"`;Skipped int `json:"skipped"`}
func opt310Root() (string,error){c,e:=os.UserCacheDir();if e!=nil{return "",e};return filepath.Join(c,"Temp"),nil}
// Require ordinary directory ancestry, never traverse reparse points or network roots.
func opt310CheckRoot(root string)error{
 if !filepath.IsAbs(root)||len(filepath.VolumeName(root))!=2||filepath.VolumeName(root)[1]!=':'{return fmt.Errorf("Unsafe temporary directory")}
 for p:=filepath.Clean(root);;p=filepath.Dir(p){
  q,e:=syscall.UTF16PtrFromString(p);if e!=nil{return e};a,e:=syscall.GetFileAttributes(q);if e!=nil{return e}
  if a&0x400!=0||a&0x10==0{return fmt.Errorf("Unsafe temporary directory ancestry")}
  if filepath.Dir(p)==p{break}
 };return nil
}
func opt310Open(path string,remove bool)(syscall.Handle,syscall.ByHandleFileInformation,error){
 var m syscall.ByHandleFileInformation
 q,e:=syscall.UTF16PtrFromString(path);if e!=nil{return 0,m,e}
 access:=uint32(0x80);if remove{access|=0x10000}
 h,e:=syscall.CreateFile(q,access,1,nil,3,0x00200000,0);if e!=nil{return 0,m,e}
 e=syscall.GetFileInformationByHandle(h,&m)
 if e!=nil||m.FileAttributes&opt310UnsafeAttrs!=0||m.NumberOfLinks!=1{syscall.CloseHandle(h);return 0,m,fmt.Errorf("File unavailable or protected")}
 return h,m,nil
}
func opt310Size(m syscall.ByHandleFileInformation)uint64{return uint64(m.FileSizeHigh)<<32|uint64(m.FileSizeLow)}
func opt310Old(m syscall.ByHandleFileInformation,now time.Time)bool{return now.Sub(time.Unix(0,m.LastWriteTime.Nanoseconds()))>=7*24*time.Hour}
func opt310Same(a,b syscall.ByHandleFileInformation)bool{return a.VolumeSerialNumber==b.VolumeSerialNumber&&a.FileIndexHigh==b.FileIndexHigh&&a.FileIndexLow==b.FileIndexLow&&a.LastWriteTime==b.LastWriteTime&&a.CreationTime==b.CreationTime&&a.FileAttributes==b.FileAttributes&&a.NumberOfLinks==b.NumberOfLinks&&opt310Size(a)==opt310Size(b)}
func opt310Preview(root string,now time.Time)(opt310Plan,error){
 p:=opt310Plan{Items:[]opt310Item{},created:now,root:root}
 if e:=opt310CheckRoot(root);e!=nil{return p,e}
 f,e:=os.Open(root);if e!=nil{return p,e};defer f.Close()
 entries,e:=f.ReadDir(opt310Enumerate);if e!=nil&&len(entries)==0{return p,nil}
 p.Truncated=len(entries)==opt310Enumerate
 for _,v:=range entries{
  if len(p.Items)>=opt310Limit{p.Truncated=true;break}
  if !v.Type().IsRegular()||!strings.EqualFold(filepath.Ext(v.Name()),".tmp"){continue}
  path:=filepath.Join(root,v.Name());h,m,e:=opt310Open(path,false);if e!=nil{continue};syscall.CloseHandle(h)
  if !opt310Old(m,now){continue}
  size:=opt310Size(m);p.Items=append(p.Items,opt310Item{ID:len(p.Items)+1,Name:v.Name(),Bytes:size,Modified:time.Unix(0,m.LastWriteTime.Nanoseconds()).Format(time.RFC3339),path:path,meta:m});p.Bytes+=size
 };return p,nil
}
func opt310Apply(p opt310Plan,ids []int,now time.Time)(opt310Result,error){
 var result opt310Result
 if p.created.IsZero()||now.Before(p.created)||now.Sub(p.created)>5*time.Minute{return result,fmt.Errorf("Preview expired; examine again")}
 if len(ids)==0||len(ids)>len(p.Items){return result,fmt.Errorf("Invalid selection")}
 selected:=map[int]bool{};for _,id:=range ids{if id<1||id>len(p.Items)||selected[id]{return result,fmt.Errorf("Invalid selection")};selected[id]=true}
 if e:=opt310CheckRoot(p.root);e!=nil{return result,e}
 for _,id:=range ids{
  v:=p.Items[id-1];if filepath.Dir(v.path)!=p.root||!strings.EqualFold(filepath.Ext(v.path),".tmp"){result.Skipped++;continue}
  h,m,e:=opt310Open(v.path,true);if e!=nil{result.Skipped++;continue}
  if !opt310Same(m,v.meta)||!opt310Old(m,now){syscall.CloseHandle(h);result.Skipped++;continue}
  // Mark the already verified handle for deletion; no second path lookup.
  yes:=byte(1);ok,_,_:=syscall.NewLazyDLL("kernel32.dll").NewProc("SetFileInformationByHandle").Call(uintptr(h),4,uintptr(unsafe.Pointer(&yes)),1)
  e=syscall.CloseHandle(h);if ok==0||e!=nil{result.Skipped++;continue}
  result.Removed++;result.Bytes+=v.Bytes
 };return result,nil
}
