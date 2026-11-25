// Read-only correlation of recorded Sysmon events. Not a malware verdict.
package behaviorwatch
import("encoding/xml";"fmt";"io";"strings";"sort")
type Event struct{ID int; GUID,Image,Target string}
type Signal struct{GUID string;Image string;FileCreates int;StartupWrites int;Reason string}
type xmlEvent struct{ID int `xml:"System>EventID"`;Data []struct{Name string `xml:"Name,attr"`;Value string `xml:",chardata"`} `xml:"EventData>Data"`}
func Parse(r io.Reader)([]Event,error){
 const maxXML=8<<20
 const maxEvents=20000
 bounded:=&io.LimitedReader{R:r,N:maxXML+1}
 dec:=xml.NewDecoder(bounded);out:=[]Event{}
 for{tok,e:=dec.Token();if bounded.N<=0{return nil,fmt.Errorf("Sysmon XML exceeds %d bytes",maxXML)};if e==io.EOF{return out,nil};if e!=nil{return nil,e};st,ok:=tok.(xml.StartElement);if !ok||st.Name.Local!="Event"{continue};var x xmlEvent;if e=dec.DecodeElement(&x,&st);e!=nil{return nil,e};v:=Event{ID:x.ID};for _,d:=range x.Data{switch d.Name{case "ProcessGuid":v.GUID=d.Value;case "Image":v.Image=d.Value;case "TargetFilename","TargetObject":v.Target=d.Value}};if v.GUID!=""{if len(out)>=maxEvents{return nil,fmt.Errorf("Sysmon event count exceeds %d",maxEvents)};out=append(out,v)}}
}
func Analyze(events []Event)[]Signal{
 type state struct{image string;creates map[string]bool;writes map[string]bool;start bool}
 m:=map[string]*state{};out:=[]Signal{}
 for _,e:=range events{if e.GUID==""{continue};guid:=strings.ToLower(strings.TrimSpace(e.GUID));if guid==""{continue};v:=m[guid];if v==nil{v=&state{creates:map[string]bool{},writes:map[string]bool{}};m[guid]=v};if e.ID==1{v.start=true;v.image=e.Image};if e.ID==11&&e.Target!=""{v.creates[strings.ToLower(e.Target)]=true};if e.ID==13{p:=strings.ToLower(e.Target);if strings.Contains(p,"\\currentversion\\run\\")||strings.Contains(p,"\\currentversion\\runonce\\"){v.writes[p]=true}}}
 for guid,v:=range m{image:=strings.ToLower(v.image);if v.start&&len(v.writes)>0&&len(v.creates)>=3&&(strings.Contains(image,"\\appdata\\local\\temp\\")||strings.Contains(image,"\\downloads\\")){out=append(out,Signal{guid,v.image,len(v.creates),len(v.writes),"Correlacao de inicio de processo, criacao de arquivos e persistencia; investigar. Nao confirma malware."})}}
 sort.Slice(out,func(i,j int)bool{return out[i].GUID<out[j].GUID});return out
}
