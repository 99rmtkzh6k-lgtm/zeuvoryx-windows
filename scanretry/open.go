package scanretry

import (
 "context"
 "errors"
 "fmt"
 "os"
 "runtime"
 "syscall"
 "time"
)

// SkippedError is a final access failure, not evidence of a clean file.
type SkippedError struct {
 Path string `json:"path"`
 Motivo string `json:"motivo"`
 Hora time.Time `json:"hora"`
 CodigoWindows uint32 `json:"codigo_windows"`
 cause error
}
func(e *SkippedError) Error() string {return fmt.Sprintf("SKIPPED: %q | %s | %s | Windows %d",e.Path,e.Hora.Format(time.RFC3339Nano),e.Motivo,e.CodigoWindows)}
func(e *SkippedError) Unwrap() error {return e.cause}

func windowsAccessCode(err error)(syscall.Errno,bool){
 var code syscall.Errno
 if !errors.As(err,&code){return 0,false}
 return code,code==5||code==32||code==33
}

// Access errors describe scan coverage, never evidence of infection.
// Code 5 alone cannot establish that a path is a system folder.
func accessReason(code syscall.Errno) string {
 switch code {
 case 5: return "Sem permissão de acesso — não é evidência de vírus"
 case 32,33: return "Arquivo em uso ou bloqueado por outro programa — não é evidência de vírus"
 default: return "Falha de acesso — não é evidência de vírus"
 }
}

// WindowsSkipped preserves an existing timestamp; it does not override ACLs.
func WindowsSkipped(path string,err error)*SkippedError {
 var skipped *SkippedError
 if errors.As(err,&skipped){return skipped}
 if runtime.GOOS!="windows"{return nil}
 code,ok:=windowsAccessCode(err);if !ok{return nil}
 return &SkippedError{Path:path,Motivo:accessReason(code),Hora:time.Now().UTC(),CodigoWindows:uint32(code),cause:err}
}

// os.Open on Windows requests read access with FILE_SHARE_READ|FILE_SHARE_WRITE.
// Exclusive handles and access controls remain enforced by the OS.
func Open(ctx context.Context,path string)(*os.File,error){
 return openScan(ctx,path,os.Open,runtime.GOOS=="windows")
}

func openScan(ctx context.Context,path string,open func(string)(*os.File,error),windows bool)(*os.File,error){
 f,err:=openWithRetry(ctx,path,open,func(err error)bool{
  _,ok:=windowsAccessCode(err);return windows&&ok
 },[]time.Duration{500*time.Millisecond})
 if err==nil||!windows{return f,err}
 code,ok:=windowsAccessCode(err);if !ok{return f,err}
 return nil,&SkippedError{Path:path,Motivo:accessReason(code),Hora:time.Now().UTC(),CodigoWindows:uint32(code),cause:err}
}

func transientWindows(err error) bool {return errors.Is(err,syscall.Errno(32))||errors.Is(err,syscall.Errno(33))}

func openWithRetry(ctx context.Context,path string,open func(string)(*os.File,error),busy func(error)bool,delays []time.Duration)(*os.File,error){
 for i:=0;;i++{
  if err:=ctx.Err();err!=nil{return nil,err}
  f,err:=open(path)
  if err==nil{return f,nil}
  if !busy(err)||i>=len(delays){return nil,err}
  timer:=time.NewTimer(delays[i])
  select{
  case <-ctx.Done():timer.Stop();return nil,ctx.Err()
  case <-timer.C:
  }
 }
}
