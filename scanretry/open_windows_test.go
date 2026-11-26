//go:build windows

package scanretry
import("context";"errors";"os";"path/filepath";"syscall";"testing")
func TestWindowsExclusiveHandleRemainsProtected(t *testing.T){
 p:=filepath.Join(t.TempDir(),"locked.txt")
 if e:=os.WriteFile(p,[]byte("benign"),0600);e!=nil{t.Fatal(e)}
 name,e:=syscall.UTF16PtrFromString(p);if e!=nil{t.Fatal(e)}
 h,e:=syscall.CreateFile(name,syscall.GENERIC_READ,0,nil,syscall.OPEN_EXISTING,syscall.FILE_ATTRIBUTE_NORMAL,0)
 if e!=nil{t.Fatal(e)}
 f,e:=Open(context.Background(),p)
 if f!=nil{f.Close()};syscall.CloseHandle(h)
 if !errors.Is(e,syscall.Errno(32)){t.Fatalf("expected sharing violation, got %v",e)}
 var skipped *SkippedError
 if !errors.As(e,&skipped)||skipped.Path!=p||skipped.CodigoWindows!=32||skipped.Hora.IsZero(){t.Fatalf("missing structured skip: %v",e)}
 f,e=Open(context.Background(),p);if e!=nil{t.Fatal(e)};f.Close()
}
