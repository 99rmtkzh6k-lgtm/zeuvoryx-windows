package auditmetrics

import("encoding/json";"strings";"testing")

func TestSignatureEvidence(t *testing.T){
 e:=Signature(20,0,100,-1,nil,false);if !e.SummaryUsable||e.Partial||e.EngineErrors!=nil||*e.Scanned!=20{t.Fatal(e)}
 e=Signature(-1,-1,-1,-1,[]string{"heuristic_or_limit"},true);if e.SummaryUsable||!e.Partial||e.Scanned!=nil||e.EngineInfected!=nil||e.SignatureRecords!=0||e.HeuristicOrLimitRecords!=1{t.Fatal(e)}
 b,err:=json.Marshal(e);if err!=nil||!strings.Contains(string(b),`"scanned_files":null`){t.Fatal(string(b),err)}
 e=Signature(10,2,100,0,[]string{"signature_detection","test_eicar"},false);if !e.SummaryUsable||e.SignatureRecords!=1||e.TestRecords!=1||e.ObservedRecords!=2{t.Fatal(e)}
 for _,v:=range []SignatureEvidence{Signature(10,2,100,0,[]string{"signature_detection"},false),Signature(0,0,100,0,nil,false),Signature(1,2,100,0,[]string{"signature_detection","signature_detection"},false),Signature(1,1,100,0,[]string{"unknown"},false),Signature(1,1,100,0,[]string{"heuristic_or_limit"},false),Signature(1,0,100,1,nil,false),Signature(1,0,0,0,nil,false)}{if !v.Partial||v.SummaryUsable{t.Fatal(v)}}
}
