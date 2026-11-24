// Package adaptive implements an experimental supervised classifier. Its score
// is advisory, not a calibrated probability of malware or a security verdict.
package adaptive

import (
    "errors"
    "math"
    "sync"
)

const Version = "adaptive-0.1-synthetic"
const Dimensions = 5
// Features: execution disguise, script type, execution signals,
// defense/backup tampering signals, PE indicators. All in [0,1].
type Vector [Dimensions]float64
type Example struct {
    X Vector `json:"features"`
    Label int `json:"label"`
}
type Model struct {
    Weights Vector `json:"weights"`
    Bias float64 `json:"bias"`
}
func valid(x Vector) bool {
    for _, v := range x { if math.IsNaN(v)||math.IsInf(v,0)||v<0||v>1 {return false} }
    return true
}
func sigmoid(z float64) float64 { return 1/(1+math.Exp(-z)) }
func (m Model) Score(x Vector) (float64,error) {
    if !valid(x)||math.IsNaN(m.Bias)||math.IsInf(m.Bias,0)||math.Abs(m.Bias)>100 {return 0,errors.New("invalid model or features")}
    z:=m.Bias
    for i,w:=range m.Weights {
        if math.IsNaN(w)||math.IsInf(w,0)||math.Abs(w)>100 {return 0,errors.New("invalid model weights")}
        z+=w*x[i]
    }
    return sigmoid(z),nil
}
// Train fits logistic regression by bounded batch gradient descent. Labels
// must be independently reviewed; this API never collects or opens user files.
func Train(rows []Example) (Model,error) {
    var m Model
    if len(rows)<12||len(rows)>10000{return m,errors.New("dataset must contain 12 to 10000 reviewed examples")}
    pos:=0
    seen:=map[Vector]int{}
    for _,r:=range rows{
        if !valid(r.X)||(r.Label!=0&&r.Label!=1){return m,errors.New("invalid training row")}
        if old,ok:=seen[r.X];ok&&old!=r.Label{return m,errors.New("conflicting labels")}
        seen[r.X]=r.Label;pos+=r.Label
    }
    if pos<3||len(rows)-pos<3{return m,errors.New("both classes required")}
    // Balance classes to limit majority-class domination.
    for epoch:=0;epoch<250;epoch++{
        var grad Vector
        var bias float64
        for _,r:=range rows{
            p,_:=m.Score(r.X)
            weight:=0.5/float64(len(rows)-pos)
            if r.Label==1 {weight=0.5/float64(pos)}
            e:=(p-float64(r.Label))*weight
            bias+=e
            for i,v:=range r.X{grad[i]+=e*v}
        }
        m.Bias-=0.2*bias
        for i:=range m.Weights {m.Weights[i]-=0.2*(grad[i]+0.02*m.Weights[i])}
    }
    return m,nil
}
// Synthetic feature vectors only. Not malware samples and not evidence of
// real-world accuracy. The model is intentionally advisory.
func synthetic() []Example {
    return []Example{
        {Vector{0,0,0,0,0},0},{Vector{0,1,0,0,0},0},
        {Vector{0,0,0,0,.2},0},{Vector{0,1,.15,0,0},0},
        {Vector{0,0,0,0,.5},0},{Vector{0,1,.3,0,0},0},
        {Vector{0,0,.1,0,.3},0},{Vector{0,1,0,0,.3},0},
        {Vector{1,1,.8,0,0},1},{Vector{0,1,.9,1,0},1},
        {Vector{1,0,0,0,.8},1},{Vector{0,1,.7,1,.2},1},
        {Vector{1,1,.6,.5,0},1},{Vector{0,0,.8,.7,1},1},
        {Vector{0,1,1,1,0},1},{Vector{1,0,.5,.5,.8},1},
    }
}
var once sync.Once
var initial Model
var initialErr error
func Baseline() (Model,error) {
    once.Do(func(){initial,initialErr=Train(synthetic())})
    return initial,initialErr
}
