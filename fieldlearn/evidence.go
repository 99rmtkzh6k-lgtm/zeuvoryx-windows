// Package fieldlearn records comparable observations, not automatic malware verdicts.
package fieldlearn
import("encoding/xml";"errors";"io";"sort";"time")
const MaxEvents=500
const MaxBytes=4<<20
type Event struct{ID int `json:"event_id"`;Record uint64 `json:"record_id"`;At time.Time `json:"at"`}
type EventLog struct{State string `json:"state"`;Events []Event `json:"events"`;Counts map[int]int `json:"counts"`;Truncated bool `json:"truncated"`}
type Motor struct{
 Name string `json:"name"`
 State string `json:"state"`
 Version string `json:"version,omitempty"`
 Started time.Time `json:"started"`
 Finished time.Time `json:"finished"`
 Files *int64 `json:"files"`
 Bytes *int64 `json:"bytes"`
 Alerts *int64 `json:"alerts"`
 Signatures *int64 `json:"signature_records"`
 Tests *int64 `json:"test_records"`
 Limits *int64 `json:"heuristic_or_limit_records"`
 Errors *int64 `json:"errors"`
 Partial bool `json:"partial"`
}
type Review struct{Schema int `json:"schema"`;Started time.Time `json:"started"`;Ended time.Time `json:"ended"`;Defender EventLog `json:"defender_events"`;Motors []Motor `json:"motors"`;Comparison string `json:"comparison"`;Problem string `json:"problem,omitempty"`;TrainingEligible bool `json:"training_eligible"`}
func Count(n int64)*int64{if n<0{return nil};return &n}
func known(id int)bool{switch id{case 1000,1001,1002,1116,1117,1118,1119:return true};return false}
func Parse(r io.Reader,start,end time.Time)(EventLog,error){
 out:=EventLog{State:"recorded_events",Counts:map[int]int{},Events:[]Event{}}
 if start.IsZero()||!end.After(start){return EventLog{},errors.New("invalid interval")}
 limited:=&io.LimitedReader{R:r,N:MaxBytes+1};d:=xml.NewDecoder(limited);seen:=map[uint64]bool{};n:=0
 for{
  tok,e:=d.Token();if limited.N<=0{return EventLog{},errors.New("event byte limit")}
  if e==io.EOF{break};if e!=nil{return EventLog{},e}
  st,ok:=tok.(xml.StartElement);if !ok||st.Name.Local!="Event"{continue}
  n++;if n>MaxEvents{out.Truncated=true;out.State="partial";break}
  var v struct{System struct{ID int `xml:"EventID"`;Record uint64 `xml:"EventRecordID"`;Time struct{At string `xml:"SystemTime,attr"`} `xml:"TimeCreated"`;Provider struct{Name string `xml:"Name,attr"`} `xml:"Provider"`} `xml:"System"`}
  if e=d.DecodeElement(&v,&st);e!=nil{return EventLog{},e}
  if limited.N<=0{return EventLog{},errors.New("event byte limit")}
  x:=v.System;if x.Provider.Name!="Microsoft-Windows-Windows Defender"||!known(x.ID){continue}
  at,e:=time.Parse(time.RFC3339Nano,x.Time.At);if e!=nil||x.Record==0{return EventLog{},errors.New("invalid event identity")}
  if at.Before(start)||at.After(end)||seen[x.Record]{continue}
  seen[x.Record]=true;out.Events=append(out.Events,Event{x.ID,x.Record,at.UTC()});out.Counts[x.ID]++
 }
 sort.Slice(out.Events,func(i,j int)bool{if out.Events[i].At.Equal(out.Events[j].At){return out.Events[i].Record<out.Events[j].Record};return out.Events[i].At.Before(out.Events[j].At)})
 if len(out.Events)==0&&!out.Truncated{out.State="no_events_observed"}
 return out,nil
}
func Build(start,end time.Time,d EventLog,motors []Motor,problem string)Review{
 r:=Review{Schema:1,Started:start.UTC(),Ended:end.UTC(),Defender:d,Motors:[]Motor{},Problem:problem,Comparison:"unpaired_observation",TrainingEligible:false}
 for _,m:=range motors{
  if m.Started.IsZero()||m.Started.Before(start)||m.Started.After(end){continue}
  if m.Finished.IsZero()||m.Finished.After(end){m.Partial=true}
  r.Motors=append(r.Motors,m)
 }
 if problem!=""||d.Truncated{r.Comparison="incomplete_evidence"}
 for _,m:=range r.Motors{if m.Partial{r.Comparison="incomplete_evidence"}}
 return r
}
