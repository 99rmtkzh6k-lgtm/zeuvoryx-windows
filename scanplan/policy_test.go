package scanplan
import (
	"testing"
	"time"
)
func TestSerial(t *testing.T){
 for _,v:=range []struct{known bool;total,available uint64;want bool}{
 {false,16*GiB,10*GiB,true},{true,0,0,true},{true,4*GiB,4*GiB,true},{true,6*GiB,4*GiB,true},{true,8*GiB,2*GiB,true},{true,8*GiB,3*GiB,false},{true,16*GiB,8*GiB,false},{true,8*GiB,9*GiB,true},
 }{if got:=Serial(v.known,v.total,v.available);got!=v.want{t.Fatalf("%+v got %t",v,got)}}
}

func TestWorkers(t *testing.T){
 for _,v:=range []struct{cpu int;known bool;total,available uint64;want int}{
 {2,true,2*GiB,GiB,1},{4,true,4*GiB,GiB,2},{4,true,4*GiB,GiB/4,1},
 {4,true,8*GiB,4*GiB,3},{8,true,16*GiB,8*GiB,4},
 {2,true,16*GiB,8*GiB,2},{8,false,16*GiB,8*GiB,1},
 {8,true,4*GiB,8*GiB,1},{0,true,0,0,1},{1,true,16*GiB,8*GiB,1},
 }{if got:=Workers(v.cpu,v.known,v.total,v.available);got!=v.want{t.Fatalf("%+v got %d",v,got)}}
 if ClamCapacity(true,2*GiB)||!ClamCapacity(true,4*GiB)||!ClamCapacity(false,0){t.Fatal("incorrect signature capacity policy")}
}

func TestSignatureBudget(t *testing.T) {
	for mode, want := range map[string]time.Duration{
		"quick":  5 * time.Minute,
		"full":   time.Hour,
		"custom": 20 * time.Minute,
	} {
		if got := SignatureBudget(mode); got != want {
			t.Errorf("%s: got %s, want %s", mode, got, want)
		}
	}
}
