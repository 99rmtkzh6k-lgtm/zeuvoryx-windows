package scanretry

import("context";"errors";"os";"path/filepath";"testing";"time";"syscall")
func TestRetryBusyThenSuccess(t *testing.T){
 path:=filepath.Join(t.TempDir(),"benign.txt")
 if e:=os.WriteFile(path,[]byte("benign"),0600);e!=nil{t.Fatal(e)}
 busy:=errors.New("busy");n:=0
 f,e:=openWithRetry(context.Background(),path,func(p string)(*os.File,error){n++;if n<3{return nil,busy};return os.Open(p)},func(e error)bool{return errors.Is(e,busy)},[]time.Duration{0,0})
 if e!=nil||f==nil||n!=3{t.Fatalf("attempts=%d error=%v",n,e)};defer f.Close()
}
func TestRetryBoundAndNonBusy(t *testing.T){
 for _,transient:=range []bool{false,true}{
  marker:=errors.New("failure");n:=0
  _,e:=openWithRetry(context.Background(),"unused",func(string)(*os.File,error){n++;return nil,marker},func(error)bool{return transient},[]time.Duration{0,0})
  want:=1;if transient{want=3};if !errors.Is(e,marker)||n!=want{t.Fatalf("attempts=%d want=%d err=%v",n,want,e)}
 }
}
func TestRetryCancellation(t *testing.T){
 ctx,cancel:=context.WithCancel(context.Background());defer cancel()
 n:=0;busy:=errors.New("busy")
 _,e:=openWithRetry(ctx,"unused",func(string)(*os.File,error){n++;cancel();return nil,busy},func(error)bool{return true},[]time.Duration{time.Hour})
 if !errors.Is(e,context.Canceled)||n!=1{t.Fatalf("attempts=%d error=%v",n,e)}
 _,e=Open(ctx,"unused");if !errors.Is(e,context.Canceled){t.Fatal(e)}
}

func TestWindowsRetryClassification(t *testing.T){for _,n:=range []syscall.Errno{32,33,5,2}{e:=&os.PathError{Op:"open",Path:"test",Err:n};if transientWindows(e)!=(n==32||n==33){t.Fatalf("classification %d",n)}}}
