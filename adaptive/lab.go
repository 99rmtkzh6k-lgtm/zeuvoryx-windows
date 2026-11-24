// Synthetic, defensive-only lab: no malware, network or executable samples.
package adaptive

import (
 "fmt"
 "math"
 "math/rand"
)

const LabVersion = "adaptive-0.2-synthetic-lab"
const LabTotal = 20000000

type Counts struct { TP, TN, FP, FN int }
func (c *Counts) add(score float64, label int) {
 if score>=0.5 {if label==1 {c.TP++}else{c.FP++}} else {if label==1{c.FN++}else{c.TN++}}
}
func (c Counts) total() int{return c.TP+c.TN+c.FP+c.FN}
func (c Counts) FPR() float64{return float64(c.FP)/float64(c.FP+c.TN)}
func (c Counts) FNR() float64{return float64(c.FN)/float64(c.FN+c.TP)}
type Comparison struct { Baseline, Candidate Counts }
type LabReport struct {
 Schema string
 SyntheticOnly bool
 TrainCases, ValidationCases, TestCases int
 Validation, Test Comparison
 TestFamilies [10]Comparison
 Candidate Model
 Accepted bool
}
var LabFamilies=[10]string{"ordinary","administration","packed-data","developer-tools","ambiguous-benign","disguised-execution","suspicious-script","tamper-signals","PE-signals","combined-signals"}

// Draw overlapping scenarios from explicit family assumptions. A label is
// fixed by the synthetic family, not by the current model or its prediction.
func labExample(r *rand.Rand) (Example,int) {
 family:=r.Intn(10)
 rates:=[10][5]float64{
  {.01,.08,.08,.01,.04},{.02,.70,.40,.18,.12},{.02,.08,.15,.03,.32},
  {.04,.30,.35,.08,.42},{.12,.55,.50,.22,.32},
  {.82,.40,.47,.10,.35},{.12,.88,.78,.32,.14},{.10,.65,.55,.83,.18},
  {.18,.15,.35,.28,.82},{.62,.72,.77,.72,.65},
 }
 var x Vector
 for i,p:=range rates[family] {
  // Binary and stepped features match the production feature extractor.
  if i<2 {if r.Float64()<p{x[i]=1};continue}
  steps:=2;if i==2{steps=4}
  for j:=0;j<steps;j++{if r.Float64()<p{x[i]+=1/float64(steps)}}
 }
 label:=0;if family>=5{label=1}
 return Example{X:x,Label:label},family
}
func labAccept(base,candidate Counts) bool {
 if base.total()==0||candidate.total()!=base.total()||base.TP+base.FN==0||base.TN+base.FP==0{return false}
 // Synthetic non-regression gate; no real-world efficacy claim.
 return candidate.FPR()<=base.FPR()+0.005 &&
 candidate.FNR()<=base.FNR()+0.005 &&
 candidate.FPR()+candidate.FNR()<=base.FPR()+base.FNR()
}
func RunLab(trainN,valN,testN int,progress func(string,int)) (LabReport,error){
 report:=LabReport{Schema:"zeuvoryx-synthetic-lab-v1",SyntheticOnly:true,TrainCases:trainN,ValidationCases:valN,TestCases:testN}
 if trainN<1000||valN<1000||testN<1000||trainN>LabTotal||valN>LabTotal||testN>LabTotal||trainN+valN+testN>LabTotal{return report,fmt.Errorf("invalid lab budget")}
 base,e:=Train(synthetic());if e!=nil{return report,e}
 candidate:=base
 // Streaming supervised mini-batches; bounded memory and no file collection.
 train:=rand.New(rand.NewSource(31021))
 var grad Vector
 var bias float64
 batch:=0
 for i:=0;i<trainN;i++{
  row,_:=labExample(train)
  p,e:=candidate.Score(row.X);if e!=nil{return report,e}
  err:=p-float64(row.Label);bias+=err;batch++
  for j,v:=range row.X{grad[j]+=err*v}
  if batch==256||i+1==trainN{
   lr:=0.08/math.Sqrt(1+float64(i)/100000)
   candidate.Bias-=lr*bias/float64(batch)
   for j:=range candidate.Weights{candidate.Weights[j]-=lr*(grad[j]/float64(batch)+.002*candidate.Weights[j])}
   grad=Vector{};bias=0;batch=0
  }
  if progress!=nil&&((i+1)%1000000==0||i+1==trainN){progress("train",i+1)}
 }
 evaluate:=func(seed int64,n int,phase string,cmp *Comparison,byFamily bool)error{
  rng:=rand.New(rand.NewSource(seed))
  for i:=0;i<n;i++{
   row,f:=labExample(rng)
   b,e:=base.Score(row.X);if e!=nil{return e}
   c,e:=candidate.Score(row.X);if e!=nil{return e}
   cmp.Baseline.add(b,row.Label);cmp.Candidate.add(c,row.Label)
   if byFamily{report.TestFamilies[f].Baseline.add(b,row.Label);report.TestFamilies[f].Candidate.add(c,row.Label)}
   if progress!=nil&&((i+1)%1000000==0||i+1==n){progress(phase,i+1)}
  }
  return nil
 }
 if e=evaluate(62043,valN,"validation",&report.Validation,false);e!=nil{return report,e}
 if e=evaluate(93047,testN,"test",&report.Test,true);e!=nil{return report,e}
 report.Candidate=candidate
 report.Accepted=labAccept(report.Validation.Baseline,report.Validation.Candidate)&&labAccept(report.Test.Baseline,report.Test.Candidate)
 return report,nil
}
