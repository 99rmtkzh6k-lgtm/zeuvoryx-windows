package main
import("testing";"path/filepath")
func Test310MaintenanceAllowlist(t *testing.T){
 root:=`C:\Windows\System32`
 for k,v:=range map[string]string{"drive":"dfrgui.exe","cleanup":"cleanmgr.exe","startup":"Taskmgr.exe"}{p,e:=maintenance310Path(root,k);if e!=nil||p!=filepath.Join(root,v){t.Fatalf("%s %s %v",k,p,e)}}
 for _,v:=range []string{"","cmd","powershell","../cmd.exe","drive /x","drive;calc"}{if _,e:=maintenance310Path(root,v);e==nil{t.Fatal("unexpected action",v)}}
 if _,e:=maintenance310Path("relative","drive");e==nil{t.Fatal("relative path accepted")}
}
