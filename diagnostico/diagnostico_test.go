package diagnostico
import (
 "encoding/json"
 "strings"
 "testing"
)
func TestMultipleHypothesesAndDedup(t *testing.T){
 r:=Analyze([]Evidence{{Code:"sharing_violation",Detail:"C:\\Users\\private"},{"Code":"sharing_violation"}})
 if len(r.Suggestions)!=2 {t.Fatalf("want 2 distinct hypotheses, got %d",len(r.Suggestions))}
 if strings.Contains(r.Prompt,"private") {t.Fatal("private path leaked")}
 if !strings.Contains(r.Prompt,"LAB") {t.Fatal("missing LAB instruction")}
}
func TestUnknownAndBounds(t *testing.T){
 r:=Analyze([]Evidence{{Code:"unknown",Detail:"secret"}})
 if len(r.Suggestions)!=0 || strings.Contains(r.Prompt,"secret") {t.Fatal("unknown evidence not safely handled")}
 a:=make([]Evidence,50);for i:=range a {a[i]=Evidence{Code:"access_denied"}}
 r=Analyze(a)
 if !r.Truncated || len(r.Evidence)!=32 {t.Fatalf("bad limits: %d",len(r.Evidence))}
 if _,err:=json.Marshal(r);err!=nil {t.Fatal(err)}
}
func TestOtherCases(t *testing.T){
 r:=Analyze([]Evidence{{Code:"clamav_exit_2"},{Code:"max_file_size"},{Code:"quick_scan_small"}})
 if len(r.Suggestions)!=3 {t.Fatalf("got %d",len(r.Suggestions))}
}
