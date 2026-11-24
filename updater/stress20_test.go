package updater
import("os";"testing";"time";"strconv";"net/url")
func TestZStressUpdate5M(t *testing.T){
 if os.Getenv("ZEUVORYX_QA20")!="1"{t.Skip("opt-in workload")}
 start:=time.Now()
 for i:=0;i<5000000;i++{
  a:="3."+strconv.Itoa(i/1000000)+"."+strconv.Itoa(i%1000000)
  x,e:=Sequence(a);if e!=nil{t.Fatalf("valid sequence %s: %v",a,e)}
  b:="3."+strconv.Itoa((i+1)/1000000)+"."+strconv.Itoa((i+1)%1000000)
  y,e:=Sequence(b);if e!=nil||y!=x+1{t.Fatalf("ordering %s -> %s",a,b)}
  if _,e=Sequence(a+"-lab");e==nil{t.Fatal("accepted invalid version")}
  u,e:=url.Parse("https://github.com.evil.invalid/"+a);if e!=nil||allowedURL(u){t.Fatal("unapproved hostname")}
 }
 t.Logf("QA20 updater: 5000000 iterations of ordering and URL rejection; elapsed=%s",time.Since(start))
}
