package fieldlearn
import("encoding/json";"strings";"testing";"time";"fmt")
func TestObservation(t *testing.T){
 a:=time.Date(2026,11,25,12,0,0,0,time.UTC);b:=a.Add(time.Hour)
 event:=func(id,record int,at time.Time)string{return fmt.Sprintf("<Event><System><Provider Name=\"Microsoft-Windows-Windows Defender\"/><EventID>%d</EventID><EventRecordID>%d</EventRecordID><TimeCreated SystemTime=\"%s\"/></System><EventData><Data>PRIVATE_PATH_DO_NOT_EXPORT</Data></EventData></Event>",id,record,at.Format(time.RFC3339Nano))}
 d,e:=Parse(strings.NewReader("<Events>"+event(1000,1,a)+event(1116,2,a)+event(1116,2,a)+event(1001,3,b)+event(1116,4,a.Add(-time.Hour))+"</Events>"),a,b)
 if e!=nil||len(d.Events)!=3||d.Counts[1116]!=1{t.Fatal(d,e)}
 m:=[]Motor{{Name:"ZEUVORYX",Started:a,Finished:b,Files:Count(100)},{Name:"ClamAV",Started:a,Finished:b,Errors:Count(-1)},{Name:"old",Started:a.Add(-time.Hour)}}
 r:=Build(a,b,d,m,"");raw,e:=json.Marshal(r)
 if e!=nil||strings.Contains(string(raw),"PRIVATE_PATH")||r.TrainingEligible||len(r.Motors)!=2||r.Motors[1].Errors!=nil{t.Fatal(string(raw),e)}
 if Build(a,b,d,m,"unavailable").Comparison!="incomplete_evidence"{t.Fatal("missing evidence accepted")}
 d,e=Parse(strings.NewReader("<Events>"+strings.Repeat(event(1000,1,a),501)+"</Events>"),a,b);if e!=nil||!d.Truncated{t.Fatal(d,e)}
 d,e=Parse(strings.NewReader("<Events/>"),a,b);if e!=nil||d.State!="no_events_observed"{t.Fatal(d,e)}
}
