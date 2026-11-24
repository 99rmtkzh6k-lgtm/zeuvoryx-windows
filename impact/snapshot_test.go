package impact
import("crypto/ed25519";"crypto/rand";"encoding/json";"math";"testing";"time")
func TestCompactSignedStatistics(t *testing.T){
 pub,priv,e:=ed25519.GenerateKey(rand.Reader);if e!=nil{t.Fatal(e)}
 now:=time.Unix(1800000000,0)
 sign:=func(s Snapshot)([]byte,[]byte){b,e:=json.Marshal(s);if e!=nil{t.Fatal(e)};return b,ed25519.Sign(priv,append([]byte(Domain),b...))}
 good:=Snapshot{Sequence:1,Timestamp:now.Unix(),Removed:3000000000,Virus:1000000000,Spyware:1000000000,Ransomware:1000000000,Devices:100,Stars:[5]uint64{1,0,0,0,1}}
 b,s:=sign(good)
 if len(b)>=MaxPayload{t.Fatal("too large")}
 v,e:=Verify(b,s,pub,0,now);if e!=nil{t.Fatal(e)}
 if rating,count:=v.Rating();rating!=3||count!=2{t.Fatal(rating,count)}
 if _,e:=Verify(b,s,pub,1,now);e==nil{t.Fatal("replay")}
 b[0]='[';if _,e:=Verify(b,s,pub,0,now);e==nil{t.Fatal("tampered")}
 for _,change:=range []func(*Snapshot){func(s *Snapshot){s.Removed++},func(s *Snapshot){s.Virus=math.MaxUint64},func(s *Snapshot){s.Devices=s.Removed+1},func(s *Snapshot){s.Stars=[5]uint64{math.MaxUint64,1}},func(s *Snapshot){s.Timestamp=now.Add(time.Hour).Unix()},func(s *Snapshot){s.Timestamp=now.Add(-MaxAge-time.Second).Unix()},func(s *Snapshot){s.Sequence=0}}{bad:=good;change(&bad);b,s:=sign(bad);if _,e:=Verify(b,s,pub,0,now);e==nil{t.Fatal("invalid counters accepted")}}
}

func TestRatingBounds(t *testing.T){
 if r,n:=(Snapshot{}).Rating();r!=0||n!=0{t.Fatal(r,n)}
 if r,n:=(Snapshot{Stars:[5]uint64{0,0,0,0,math.MaxUint64}}).Rating();r!=5||n!=math.MaxUint64{t.Fatal(r,n)}
 if r,n:=(Snapshot{Stars:[5]uint64{math.MaxUint64,1}}).Rating();r!=0||n!=0{t.Fatal(r,n)}
}
func FuzzSignedSnapshot(f *testing.F){
 pub,priv,e:=ed25519.GenerateKey(rand.Reader);if e!=nil{f.Fatal(e)}
 f.Add([]byte("{}"));f.Add([]byte("null"));f.Add([]byte("[]"))
 f.Fuzz(func(t *testing.T,b []byte){
  if len(b)>MaxPayload{return}
  sig:=ed25519.Sign(priv,append([]byte(Domain),b...))
  s,e:=Verify(b,sig,pub,0,time.Unix(1800000000,0))
  if e==nil{r,n:=s.Rating();if n>0&&(r<1||r>5){t.Fatal("invalid rating")}}
 })
}
