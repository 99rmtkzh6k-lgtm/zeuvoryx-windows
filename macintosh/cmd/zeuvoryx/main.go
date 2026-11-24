package main
import("context";"encoding/json";"fmt";"os";"os/signal";"github.com/zeuvoryx/zeuvoryx/macintosh/scanner";"github.com/zeuvoryx/zeuvoryx/labcheck")
func main(){
 if len(os.Args)==2&&os.Args[1]=="selftest"{r,e:=labcheck.Run();if e!=nil||!r.Update||!r.Statistics||!r.Policy{fmt.Fprintln(os.Stderr,"diagnostic failed");os.Exit(1)};if e=json.NewEncoder(os.Stdout).Encode(r);e!=nil{os.Exit(1)};return}
 if len(os.Args)!=3||os.Args[1]!="scan"{fmt.Fprintln(os.Stderr,"usage: zeuvoryx-engine scan PATH | selftest");os.Exit(2)}
 ctx,stop:=signal.NotifyContext(context.Background(),os.Interrupt);defer stop()
 r,e:=scanner.Scan(ctx,os.Args[2]);if e!=nil{fmt.Fprintln(os.Stderr,e);os.Exit(1)};if e=json.NewEncoder(os.Stdout).Encode(r);e!=nil{os.Exit(1)}
}
