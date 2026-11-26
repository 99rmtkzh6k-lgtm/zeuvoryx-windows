package scanretry

import (
 "context"
 "errors"
 "os"
 "runtime"
 "syscall"
 "time"
)

// Open retries only Windows sharing and lock violations. Never force access or change file permissions.
func Open(ctx context.Context, path string) (*os.File,error) {
 return openWithRetry(ctx,path,os.Open,func(err error)bool{
  return runtime.GOOS=="windows" && transientWindows(err)
 },[]time.Duration{40*time.Millisecond,120*time.Millisecond})
}

func transientWindows(err error) bool { return errors.Is(err,syscall.Errno(32)) || errors.Is(err,syscall.Errno(33)) }

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
