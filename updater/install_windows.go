//go:build windows

package updater
import("context";"errors";"os";"os/exec";"path/filepath";"strings";"syscall";"github.com/zeuvoryx/zeuvoryx/releaseguard")
// StartVerifiedInstall only runs for a default installed LAB copy; portable and custom
// paths must use the manual installer. The file is reverified under a Windows read lock.
func StartVerifiedInstall(ctx context.Context,target string)error{
 exe,e:=os.Executable();if e!=nil{return e}
 base:=os.Getenv("LOCALAPPDATA");if base==""||!strings.EqualFold(filepath.Clean(exe),filepath.Join(base,"Programs","ZEUVORYX-Lab303","ZEUVORYX.exe")){return errors.New("instalação portátil ou personalizada: use o instalador manual")}
 v,e:=StageAuthorized(ctx,target);if e!=nil{return e}
 ok:=false;defer func(){if !ok{os.RemoveAll(filepath.Dir(v.Path))}}()
 p,e:=syscall.UTF16PtrFromString(v.Path);if e!=nil{return e}
 h,e:=syscall.CreateFile(p,syscall.GENERIC_READ,syscall.FILE_SHARE_READ,nil,syscall.OPEN_EXISTING,syscall.FILE_ATTRIBUTE_NORMAL,0);if e!=nil{return e}
 f:=os.NewFile(uintptr(h),v.Path);defer f.Close()
 if e=releaseguard.VerifyPackage(f,v.Manifest);e!=nil{return e}
 if e=ctx.Err();e!=nil{return e}
 cmd:=exec.Command(v.Path,"/SILENT","/NORESTART","/CLOSEAPPLICATIONS","/RESTARTAPPLICATIONS","/ZEUVORYXUPDATE=1","/ZEUVORYXTARGET="+target)
 if e=cmd.Start();e!=nil{return e};ok=true
 _=cmd.Process.Release()
 return nil
}
