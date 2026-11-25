package changeguard
import("context";"errors";"testing";"time")
func TestRunScansPendingAndRetry(t *testing.T){
 tr:=New(nil);ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second);defer cancel()
 n:=0;events:=0
 err:=RunScans(ctx,tr,5*time.Millisecond,time.Second,func(context.Context)error{
  n++;if n==1{tr.Signal(time.Now());return errors.New("temporary")}
  if n==2{tr.Signal(time.Now())}
  return nil
 },func(c ScanCycle){events++;if events==1&&c.Error==""{t.Error("lost failure")};if events==3{if c.Signals!=2{t.Error("lost pending notification")};cancel()}})
 if !errors.Is(err,context.Canceled)||n!=3{t.Fatal(n,err)}
}
