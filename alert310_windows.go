package main
import "syscall"
// System sound only for findings in the report.
func z310Beep() bool {
 r,_,_:=syscall.NewLazyDLL("user32.dll").NewProc("MessageBeep").Call(0x30)
 return r!=0
}
