package main

import (
 "encoding/json"
 "flag"
 "fmt"
 "os"
 "path/filepath"
 "time"

 "github.com/zeuvoryx/zeuvoryx/adaptive"
)

func run() error {
 out:=flag.String("out","lab-results/adaptive-report.json","local JSON report path")
 flag.Parse()
 started:=time.Now()
 fmt.Println("Synthetic lab: 1M supervised training + 1M validation + 18M held-out test scenarios")
 report,err:=adaptive.RunLab(1000000,1000000,18000000,func(phase string,n int){
  fmt.Printf("%s: %d scenarios processed (%s elapsed)\n",phase,n,time.Since(started).Round(time.Millisecond))
 })
 if err!=nil{return err}
 if report.TrainCases+report.ValidationCases+report.TestCases!=adaptive.LabTotal{return fmt.Errorf("incorrect scenario accounting")}
 payload:=struct{
  Report adaptive.LabReport
  Families [10]string
  ElapsedMS int64
  Deployment string
 }{report,adaptive.LabFamilies,time.Since(started).Milliseconds(),"Candidate not deployed; synthetic-only review required"}
 b,err:=json.MarshalIndent(payload,"","  ");if err!=nil{return err}
 if err=os.MkdirAll(filepath.Dir(*out),0755);err!=nil{return err}
 if err=os.WriteFile(*out,b,0600);err!=nil{return err}
 fmt.Printf("Test baseline FPR=%.4f FNR=%.4f; candidate FPR=%.4f FNR=%.4f\n",report.Test.Baseline.FPR(),report.Test.Baseline.FNR(),report.Test.Candidate.FPR(),report.Test.Candidate.FNR())
 fmt.Printf("Synthetic non-regression gate passed: %v. Candidate not deployed. Report: %s\n",report.Accepted,*out)
 return nil
}
func main(){if err:=run();err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(1)}}
