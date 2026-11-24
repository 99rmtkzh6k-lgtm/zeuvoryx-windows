package main
import("context";"encoding/json";"fmt";"os";"time";"github.com/zeuvoryx/zeuvoryx/globalintel")
func main(){ctx,cancel:=context.WithTimeout(context.Background(),time.Minute);defer cancel();s,e:=globalintel.CollectSeed(ctx,"drop-seed");if e!=nil{fmt.Fprintln(os.Stderr,e);os.Exit(1)};if e=json.NewEncoder(os.Stdout).Encode(s);e!=nil{os.Exit(1)}}
