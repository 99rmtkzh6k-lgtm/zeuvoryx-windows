// Package labcheck runs local, inert diagnostic checks of the three foundations.
// It does not contact a server, use production keys or operate on user files.
package labcheck

import (
 "bytes"
 "crypto/ed25519"
 "crypto/rand"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "time"
 "github.com/zeuvoryx/zeuvoryx/releaseguard"
 "github.com/zeuvoryx/zeuvoryx/impact"
 "github.com/zeuvoryx/zeuvoryx/remediation"
)

type Result struct {
 Update bool `json:"update"`
 Statistics bool `json:"statistics"`
 Policy bool `json:"policy"`
}
func Run() (Result,error) {
 var r Result
 pub,priv,e:=ed25519.GenerateKey(rand.Reader)
 if e!=nil{return r,e}
 now:=time.Now()
 data:=[]byte("ZEUVORYX inert laboratory package")
 h:=sha256.Sum256(data)
 m:=releaseguard.Manifest{Sequence:2,Version:"3.0.3",Channel:"beta",Platform:"darwin-universal",URL:"https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/releases/download/mac-v3.0.3/lab.zip",SHA256:hex.EncodeToString(h[:]),Size:int64(len(data)),Expires:now.Add(time.Hour).Unix()}
 raw,e:=json.Marshal(m);if e!=nil{return r,e}
 sig:=ed25519.Sign(priv,append([]byte(releaseguard.Domain),raw...))
 verified,e:=releaseguard.Verify(raw,sig,pub,1,"beta",now)
 if e!=nil{return r,e}
 _,replay:=releaseguard.Verify(raw,sig,pub,2,"beta",now)
 r.Update=releaseguard.VerifyPackage(bytes.NewReader(data),verified)==nil && releaseguard.VerifyPackage(bytes.NewReader([]byte("wrong")),verified)!=nil && replay!=nil
 s:=impact.Snapshot{Sequence:1,Timestamp:now.Unix(),Removed:1,Virus:1,Devices:1,Stars:[5]uint64{0,0,0,0,1}}
 raw,e=json.Marshal(s);if e!=nil{return r,e}
 sig=ed25519.Sign(priv,append([]byte(impact.Domain),raw...))
 v,e:=impact.Verify(raw,sig,pub,0,now);if e!=nil{return r,e}
 _,replay=impact.Verify(raw,sig,pub,1,now)
 rating,count:=v.Rating()
 r.Statistics=rating==5&&count==1&&replay!=nil
 evidence:=remediation.Evidence{Verified:true,CurrentIdentityVerified:true,UserConsent:true,RecoveryReady:true,TestArtifact:true}
 r.Policy=remediation.Allowed(remediation.Quarantine,evidence)&&!remediation.EligibleForProductionMetric(evidence)&&!remediation.Allowed(remediation.StopProcess,evidence)&&!remediation.Allowed(remediation.CleanPersistence,evidence)&&!remediation.Allowed(remediation.Repair,evidence)
 return r,nil
}
