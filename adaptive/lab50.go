package adaptive
import("fmt";"math";"math/rand";"sort")
const Lab50Train=5000000
const Lab50Validation=5000000
const Lab50Nominal=45000000
const Lab50Shift=5000000
const patternCount=180
// Smoothed joint-frequency classifier, trained on synthetic labels. Never auto-deployed.
type PatternModel struct{ Scores [patternCount]float64;Threshold float64 }
type Lab50Report struct{Schema string;SyntheticOnly bool;TrainCases,ValidationCases,NominalTestCases,ShiftTestCases,UniqueTrainingPatterns int;Validation,NominalTest,ShiftTest Comparison;NominalFamilies,ShiftFamilies [10]Comparison;Candidate PatternModel;ValidationGate,NominalGate,ShiftGate,FamilyGate,Accepted bool; FailedFamilies []string}
func patternIndex(x Vector)int{return (((int(math.Round(x[0]))*2+int(math.Round(x[1])))*5+int(math.Round(x[2]*4)))*3+int(math.Round(x[3]*2)))*3+int(math.Round(x[4]*2))}
func strictGate(b,c Counts)bool{return b.total()>0&&b.total()==c.total()&&b.TP+b.FN>0&&b.TN+b.FP>0&&b.TP+b.FN==c.TP+c.FN&&b.TN+b.FP==c.TN+c.FP&&c.FN<=b.FN&&c.FP<=b.FP}
func countAt(h [patternCount][2]int,m PatternModel)Counts{var c Counts;for i,v:=range h{if m.Scores[i]>=m.Threshold{c.FP+=v[0];c.TP+=v[1]}else{c.TN+=v[0];c.FN+=v[1]}};return c}
func RunLab50(trainN,valN,nominalN,shiftN int,progress func(string,int))(Lab50Report,error){
 r:=Lab50Report{Schema:"zeuvoryx-synthetic-lab-v2",SyntheticOnly:true,TrainCases:trainN,ValidationCases:valN,NominalTestCases:nominalN,ShiftTestCases:shiftN}
 if trainN<1000||trainN>Lab50Train||valN<1000||valN>Lab50Validation||nominalN<1000||nominalN>Lab50Nominal||shiftN<1000||shiftN>Lab50Shift{return r,fmt.Errorf("invalid lab budget")}
 base,e:=Train(synthetic());if e!=nil{return r,e};var training,validation [patternCount][2]int
 tick:=func(phase string,i,n int){if progress!=nil&&((i+1)%1000000==0||i+1==n){progress(phase,i+1)}}
 rng:=rand.New(rand.NewSource(481771));for i:=0;i<trainN;i++{row,_:=labExample(rng);training[patternIndex(row.X)][row.Label]++;tick("train",i,trainN)}
 m:=PatternModel{};for i,v:=range training{if v[0]+v[1]>0{r.UniqueTrainingPatterns++};m.Scores[i]=(float64(v[1])+1)/(float64(v[0]+v[1])+2)}
 rng=rand.New(rand.NewSource(582883));for i:=0;i<valN;i++{row,_:=labExample(rng);validation[patternIndex(row.X)][row.Label]++;b,e:=base.Score(row.X);if e!=nil{return r,e};r.Validation.Baseline.add(b,row.Label);tick("validation",i,valN)}
 // Threshold calibrated on validation only; minimize misses without adding false alarms.
 thresholds:=append([]float64{0.5,1.0},m.Scores[:]...);sort.Float64s(thresholds);m.Threshold=0.5;best:=countAt(validation,m);found:=false
 for _,t:=range thresholds{candidate:=m;candidate.Threshold=t;c:=countAt(validation,candidate);if strictGate(r.Validation.Baseline,c)&&(!found||c.FN<best.FN||(c.FN==best.FN&&c.FP<best.FP)){m.Threshold=t;best=c;found=true}}
 r.Candidate=m;r.Validation.Candidate=best;r.ValidationGate=strictGate(r.Validation.Baseline,best)
 evaluate:=func(seed int64,n int,shift bool,cmp *Comparison,families *[10]Comparison)error{
  rng:=rand.New(rand.NewSource(seed));phase:="nominal-test";if shift{phase="shift-test"}
  for i:=0;i<n;i++{row,f:=labExample(rng);if shift{for j:=range row.X{if rng.Float64()<0.15{steps:=1;if j==2{steps=4}else if j>2{steps=2};row.X[j]=float64(rng.Intn(steps+1))/float64(steps)}}};b,e:=base.Score(row.X);if e!=nil{return e};v:=0.0;if m.Scores[patternIndex(row.X)]>=m.Threshold{v=1};cmp.Baseline.add(b,row.Label);cmp.Candidate.add(v,row.Label);families[f].Baseline.add(b,row.Label);families[f].Candidate.add(v,row.Label);tick(phase,i,n)};return nil}
 if e=evaluate(683997,nominalN,false,&r.NominalTest,&r.NominalFamilies);e!=nil{return r,e};if e=evaluate(785103,shiftN,true,&r.ShiftTest,&r.ShiftFamilies);e!=nil{return r,e}
 r.NominalGate=strictGate(r.NominalTest.Baseline,r.NominalTest.Candidate);r.ShiftGate=strictGate(r.ShiftTest.Baseline,r.ShiftTest.Candidate);r.FamilyGate=true
 for i,name:=range LabFamilies {
  if !familyGate(r.NominalFamilies[i]) {r.FamilyGate=false;r.FailedFamilies=append(r.FailedFamilies,"nominal:"+name)}
  if !familyGate(r.ShiftFamilies[i]) {r.FamilyGate=false;r.FailedFamilies=append(r.FailedFamilies,"shift:"+name)}
 }
 r.Accepted=r.ValidationGate&&r.NominalGate&&r.ShiftGate&&r.FamilyGate;return r,nil
}

// Unlike strictGate, this works for single-label scenario families.
// Require at least 100 observations, identical class totals and no regression.
func familyGate(c Comparison)bool{
 b,n:=c.Baseline,c.Candidate
 return b.total()>=100 && b.total()==n.total() &&
 b.TP+b.FN==n.TP+n.FN && b.TN+b.FP==n.TN+n.FP &&
 n.FP<=b.FP && n.FN<=b.FN
}
