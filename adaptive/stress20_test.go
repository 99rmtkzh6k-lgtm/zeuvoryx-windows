package adaptive
import("math";"math/rand";"os";"testing";"time")
func TestZStressModel5M(t *testing.T){
 if os.Getenv("ZEUVORYX_QA20")!="1"{t.Skip("opt-in workload")}
 m,e:=Baseline();if e!=nil{t.Fatal(e)}
 rng:=rand.New(rand.NewSource(30520261124));var counts Counts;start:=time.Now()
 for i:=0;i<5000000;i++{x,_:=labExample(rng);v,e:=m.Score(x.X);if e!=nil||math.IsNaN(v)||v<0||v>1{t.Fatalf("case %d score %v %v",i,v,e)};counts.add(v,x.Label)}
 if counts.total()!=5000000{t.Fatal(counts)}
 t.Logf("QA20 model: %d completed, synthetic confusion matrix=%+v elapsed=%s; not real malware accuracy",counts.total(),counts,time.Since(start))
}
