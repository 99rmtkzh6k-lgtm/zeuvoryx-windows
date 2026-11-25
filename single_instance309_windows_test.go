//go:build windows
package main
import "testing"
func Test309SingleInstance(t *testing.T){
 release,e:=acquireDesktopInstance();if e!=nil{t.Fatal(e)}
 if next,e:=acquireDesktopInstance();e==nil{next();release();t.Fatal("duplicate admitted")}
 release()
 next,e:=acquireDesktopInstance();if e!=nil{t.Fatal(e)};next()
}
