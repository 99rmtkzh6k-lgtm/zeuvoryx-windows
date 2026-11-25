package scanplan
import("testing";"time")
func TestAwaitOptionalResult(t *testing.T){ch:=make(chan int,1);ch<-42;if v,ok:=Await(ch,time.Second);!ok||v!=42{t.Fatal(v,ok)};if _,ok:=Await(ch,time.Millisecond);ok{t.Fatal("missing result reported as received")};ch<-7;close(ch);if v,ok:=Await(ch,time.Second);!ok||v!=7{t.Fatal(v,ok)};if _,ok:=Await(ch,time.Second);ok{t.Fatal("closed channel accepted")}}
