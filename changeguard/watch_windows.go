package changeguard

import("context";"fmt";"syscall";"time";"unsafe")
var kernel= syscall.NewLazyDLL("kernel32.dll")
var first=kernel.NewProc("FindFirstChangeNotificationW")
var next=kernel.NewProc("FindNextChangeNotification")
var closeWatch=kernel.NewProc("FindCloseChangeNotification")

// Watch requests native notifications, with bounded waits for cancellation. No file access or modification.
func Watch(ctx context.Context,root string,t *Tracker)error{
 p,e:=syscall.UTF16PtrFromString(root);if e!=nil{return e}
 // FILE_NOTIFY_CHANGE_FILE_NAME | DIR_NAME | SIZE | LAST_WRITE
 h,_,e:=first.Call(uintptr(unsafe.Pointer(p)),1,0x1|0x2|0x8|0x10)
 if h==0||h==^uintptr(0){return fmt.Errorf("watch unavailable: %v",e)}
 defer closeWatch.Call(h);t.Active(root,true);defer t.Active(root,false)
 for{if ctx.Err()!=nil{return nil}
  result,e:=syscall.WaitForSingleObject(syscall.Handle(h),500);if e!=nil{return e}
  switch result{case 0: t.Signal(time.Now());r,_,e:=next.Call(h);if r==0{return fmt.Errorf("watch rearm: %v",e)}
  case 258:continue
  default:return fmt.Errorf("unexpected watch wait: %d",result)}
 }
}
