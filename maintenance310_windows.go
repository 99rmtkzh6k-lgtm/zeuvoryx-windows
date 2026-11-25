package main
import("fmt";"os";"os/exec";"path/filepath";"syscall";"unsafe")
// Fixed native tool allowlist. No user-supplied paths, shell, arguments or silent actions.
func maintenance310Path(root,kind string)(string,error){
 if !filepath.IsAbs(root){return "",fmt.Errorf("Windows system directory unavailable")}
 var name string
 switch kind{case "drive":name="dfrgui.exe";case "cleanup":name="cleanmgr.exe";case "startup":name="Taskmgr.exe";default:return "",fmt.Errorf("Unknown maintenance tool")}
 return filepath.Join(root,name),nil
}
func openMaintenance310(kind string)(bool,error){
 var buf [32768]uint16
 n,_,_:=syscall.NewLazyDLL("kernel32.dll").NewProc("GetSystemDirectoryW").Call(uintptr(unsafe.Pointer(&buf[0])),uintptr(len(buf)))
 if n==0||n>=uintptr(len(buf)){return false,fmt.Errorf("Windows system directory unavailable")}
 root:=syscall.UTF16ToString(buf[:int(n)])
 path,e:=maintenance310Path(root,kind);if e!=nil{return false,e}
 st,e:=os.Stat(path);if e!=nil||!st.Mode().IsRegular(){return false,fmt.Errorf("Windows tool unavailable: %s",filepath.Base(path))}
 cmd:=exec.Command(path);cmd.Dir=root
 if e=cmd.Start();e!=nil{return false,e};go func(){_ = cmd.Wait()}()
 return true,nil
}
