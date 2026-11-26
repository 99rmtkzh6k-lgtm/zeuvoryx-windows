package scopeprogress
import ("context"; "testing"; "time")
func TestAwaitResult(t *testing.T) {
 ch:=make(chan int,1); ch<-42
 if v,ok:=AwaitResult(ch,time.Second); !ok||v!=42 {t.Fatal(v,ok)}
 if v,ok:=AwaitResult(ch,time.Millisecond);ok||v!=0{t.Fatal(v,ok)}
 close(ch)
 if _,ok:=AwaitResult(ch,time.Second);ok{t.Fatal("closed channel is not a result")}
}

func TestAwaitContextSharedDeadline(t *testing.T){
 ctx,cancel:=context.WithCancel(context.Background());cancel();empty:=make(chan int);ready:=make(chan int,1);ready<-7
 if _,ok:=AwaitContext(ctx,empty);ok{t.Fatal("unexpected result")};if v,ok:=AwaitContext(ctx,ready);!ok||v!=7{t.Fatal(v,ok)}
}
