package scanplan
import "testing"
func TestSerial(t *testing.T){
 for _,v:=range []struct{known bool;total,available uint64;want bool}{
 {false,16*GiB,10*GiB,true},{true,0,0,true},{true,4*GiB,4*GiB,true},{true,6*GiB,4*GiB,true},{true,8*GiB,2*GiB,true},{true,8*GiB,3*GiB,false},{true,16*GiB,8*GiB,false},{true,8*GiB,9*GiB,true},
 }{if got:=Serial(v.known,v.total,v.available);got!=v.want{t.Fatalf("%+v got %t",v,got)}}
}
