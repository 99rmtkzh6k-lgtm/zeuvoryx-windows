package adaptive

import (
    "math"
    "testing"
)

func TestBaseline(t *testing.T) {
    m,err:=Baseline();if err!=nil{t.Fatal(err)}
    quiet,_:=m.Score(Vector{0,0,0,0,0})
    suspicious,_:=m.Score(Vector{1,1,1,1,1})
    if quiet>=suspicious||quiet<0||suspicious>1{t.Fatal(quiet,suspicious)}
    again,_:=Baseline();if again!=m{t.Fatal("non-deterministic model")}
}
func TestTrainRejectsInvalid(t *testing.T) {
    rows:=synthetic()
    for _,bad:=range []Example{{Vector{math.NaN(),0,0,0,0},0},{Vector{2,0,0,0,0},0},{Vector{},2},{Vector{},1}}{
        copyRows:=append(append([]Example{},rows...),bad)
        if _,err:=Train(copyRows);err==nil{t.Fatal("accepted invalid or conflicting training data")}
    }
    if _,err:=Train(rows[:8]);err==nil{t.Fatal("accepted insufficient dataset")}
}
func TestTrainLearnsFromLabels(t *testing.T) {
    rows:=synthetic()
    m,err:=Train(rows);if err!=nil{t.Fatal(err)}
    for i:=range rows{rows[i].Label=1-rows[i].Label}
    opposite,err:=Train(rows);if err!=nil{t.Fatal(err)}
    x:=Vector{1,1,1,1,1}
    a,_:=m.Score(x);b,_:=opposite.Score(x)
    if a<=b{t.Fatal("model did not respond to changed labels",a,b)}
}
func TestConcurrentBaseline(t *testing.T) {
    t.Parallel()
    for i:=0;i<20;i++{t.Run("score",func(t *testing.T){t.Parallel();m,e:=Baseline();if e!=nil{t.Fatal(e)};if _,e=m.Score(Vector{0,1,.2,0,0});e!=nil{t.Fatal(e)}})}
}
func TestInvalidModel(t *testing.T){
    m:=Model{Bias:math.NaN()}
    if _,e:=m.Score(Vector{});e==nil{t.Fatal("invalid model accepted")}
}
