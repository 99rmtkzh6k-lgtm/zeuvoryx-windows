//go:build windows

// Standalone 2048 LAB host. Separate process and native top-level window.
// Build from repository root: go build -o ZEUVORYX-2048-LAB.exe ./lab/arcade2048
package main

import (
 "embed"
 "runtime"
 "syscall"
 "unsafe"

 webview "github.com/jchv/go-webview2"
)

//go:embed game.html
var gameHTML string

var user32 = syscall.NewLazyDLL("user32.dll")
var sendMessageW = user32.NewProc("SendMessageW")
var setForegroundWindow = user32.NewProc("SetForegroundWindow")
var showWindow = user32.NewProc("ShowWindow")
var setProcessDPIAware = user32.NewProc("SetProcessDPIAware")

const (
 wmSysCommand = 0x0112
 scMinimize = 0xF020
 scMaximize = 0xF030
 scRestore = 0xF120
 scClose = 0xF060
 swShow = 5
)

func main() {
 runtime.LockOSThread()
 defer runtime.UnlockOSThread()
 // DPI awareness must be set before creating the HWND.
 setProcessDPIAware.Call()
 w := webview.NewWithOptions(webview.WebViewOptions{
  Debug: false, AutoFocus: true,
  WindowOptions: webview.WindowOptions{
   Title: "ZEUVORYX LAB | 2048", Width: 1024, Height: 640, Center: true,
  },
 })
 if w == nil { return }
 defer w.Destroy()
 w.SetSize(900, 600, webview.HintMin)
 hwnd := uintptr(w.Window())
 // Native window system commands; not CSS-only imitation.
 command := func(code uintptr) bool {
  w.Dispatch(func(){ sendMessageW.Call(hwnd, wmSysCommand, code, 0) })
  return true
 }
 if err := w.Bind("z2048Minimize", func() bool { return command(scMinimize) }); err != nil { return }
 if err := w.Bind("z2048Maximize", func() bool { return command(scMaximize) }); err != nil { return }
 if err := w.Bind("z2048Restore", func() bool { return command(scRestore) }); err != nil { return }
 if err := w.Bind("z2048Close", func() bool { return command(scClose) }); err != nil { return }
 // Standard Windows titlebar retains native minimize, maximize and close.
 // The embedded game is local, offline, and unchanged in gameplay behavior.
 showWindow.Call(hwnd, swShow)
 setForegroundWindow.Call(hwnd)
 w.SetHtml(gameHTML)
 w.Run()
 // The host exits without cancelling or changing the main antivirus process.
 _ = unsafe.Pointer(nil)
}
