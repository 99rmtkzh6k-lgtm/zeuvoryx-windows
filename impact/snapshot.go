// Package impact validates compact, signed aggregate statistics from a trusted service.
// A valid signature establishes origin, not the truth of individual client reports.
package impact
import("bytes";"crypto/ed25519";"encoding/json";"errors";"math";"time")

const Domain="ZEUVORYX/IMPACT/v1\x00"
const MaxPayload=1024
type Snapshot struct{
 Sequence uint64 `json:"seq"`
 Timestamp int64 `json:"ts"`
 Removed uint64 `json:"removed"`
 Virus uint64 `json:"virus"`
 Spyware uint64 `json:"spyware"`
 Ransomware uint64 `json:"ransomware"`
 Other uint64 `json:"other"`
 Devices uint64 `json:"devices"`
 Defender uint64 `json:"defender"`
 Stars [5]uint64 `json:"stars"`
}
func sum(nums ...uint64)(uint64,bool){
 var n uint64
 for _,v:=range nums{if math.MaxUint64-n<v{return 0,false};n+=v}
 return n,true
}
func Verify(raw,sig []byte,key ed25519.PublicKey,previous uint64,now time.Time)(Snapshot,error){
 var s Snapshot
 fail:=func()(Snapshot,error){return Snapshot{},errors.New("invalid or unauthenticated statistics")}
 if len(raw)==0||len(raw)>=MaxPayload||len(key)!=ed25519.PublicKeySize||len(sig)!=ed25519.SignatureSize||!ed25519.Verify(key,append([]byte(Domain),raw...),sig){return fail()}
 if json.Unmarshal(raw,&s)!=nil{return fail()}
 canonical,e:=json.Marshal(s)
 if e!=nil||!bytes.Equal(canonical,raw)||s.Sequence==0||s.Sequence<=previous||s.Timestamp<=0||s.Timestamp>now.Add(5*time.Minute).Unix(){return fail()}
 n,ok:=sum(s.Virus,s.Spyware,s.Ransomware,s.Other)
 if !ok||n!=s.Removed||s.Devices>s.Removed{return fail()}
 if _,ok=sum(s.Stars[:]...);!ok{return fail()}
 return s,nil
}
func (s Snapshot) Rating()(float64,uint64){
 total,ok:=sum(s.Stars[:]...)
 if !ok||total==0{return 0,0}
 var weighted float64
 for i,n:=range s.Stars{weighted+=float64(i+1)*float64(n)}
 return weighted/float64(total),total
}
