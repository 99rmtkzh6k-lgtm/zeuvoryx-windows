package adaptive

import "testing"

func TestLabAccounting(t *testing.T){
 r,e:=RunLab(1000,1000,2000,nil);if e!=nil{t.Fatal(e)}
 if !r.SyntheticOnly||r.TrainCases!=1000||r.Validation.Baseline.total()!=1000||r.Validation.Candidate.total()!=1000||r.Test.Baseline.total()!=2000||r.Test.Candidate.total()!=2000{t.Fatal("incorrect case counts")}
 n:=0;for _,f:=range r.TestFamilies{n+=f.Candidate.total()}
 if n!=2000{t.Fatal("incorrect family counts")}
 r2,e:=RunLab(1000,1000,2000,nil);if e!=nil||r!=r2{t.Fatal("non-reproducible result",e)}
}
func TestLabRejectsRegression(t *testing.T){
 base:=Counts{TP:90,TN:90,FP:10,FN:10}
 if labAccept(base,Counts{TP:80,TN:100,FP:0,FN:20}){t.Fatal("accepted higher false negative rate")}
 if !labAccept(base,Counts{TP:92,TN:92,FP:8,FN:8}){t.Fatal("rejected non-regression")}
 if labAccept(base,Counts{}){t.Fatal("accepted missing results")}
}
func TestLabBudget(t *testing.T){if _,e:=RunLab(1,1000,1000,nil);e==nil{t.Fatal("accepted invalid budget")}}
