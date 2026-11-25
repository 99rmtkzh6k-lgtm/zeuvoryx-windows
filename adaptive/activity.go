package adaptive

import ("sync";"time")
type Activity struct {
 Version string `json:"version"`
 Running int `json:"running"`
 Observations uint64 `json:"observations"`
 Successful uint64 `json:"successful"`
 Errors uint64 `json:"errors"`
 SelfChecking bool `json:"self_checking"`
 SelfTests uint64 `json:"self_tests"`
 SelfTestOK bool `json:"self_test_ok"`
 Last string `json:"last"`
 LastScore *float64 `json:"last_score"`
}
var activityMu sync.Mutex
var activity=Activity{Version:Version}
func Status()Activity{activityMu.Lock();defer activityMu.Unlock();return activity}
func Observe(x Vector)(float64,error){
 activityMu.Lock();activity.Running++;activity.Observations++;activityMu.Unlock()
 m,e:=Baseline();var v float64;if e==nil{v,e=m.Score(x)}
 activityMu.Lock();defer activityMu.Unlock();activity.Running--;activity.Last=time.Now().Format(time.RFC3339)
 if e!=nil{activity.Errors++}else{activity.Successful++;activity.LastScore=&v}
 return v,e
}
var checkMu sync.Mutex
// Explicit synthetic self-check only; never counts as scanned file or training.
func SelfCheck()Activity{
 checkMu.Lock();defer checkMu.Unlock()
 activityMu.Lock();activity.SelfChecking=true;activityMu.Unlock()
 m,e:=Baseline();ok:=false
 if e==nil{a,e1:=m.Score(Vector{});b,e2:=m.Score(Vector{1,1,1,1,1});ok=e1==nil&&e2==nil&&b>a&&a>=0&&b<=1}
 activityMu.Lock();activity.SelfChecking=false;activity.SelfTests++;activity.SelfTestOK=ok;activityMu.Unlock()
 return Status()
}
