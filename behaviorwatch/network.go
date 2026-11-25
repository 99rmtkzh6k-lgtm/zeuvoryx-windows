package behaviorwatch
import("net/netip";"sort";"strings")
type NetworkSignal struct{GUID,Image,DestinationIP string}
// CorrelateNetwork joins only process creation and network telemetry by process GUID.
// Lookup is a local reputation callback, never a connection to the destination.
func CorrelateNetwork(events []Event,lookup func(string)(bool,error))([]NetworkSignal,int){
 images:=map[string]string{};seen:=map[string]bool{};out:=[]NetworkSignal{};unknown:=0
 for _,e:=range events{g:=strings.ToLower(strings.TrimSpace(e.GUID));if e.ID==1&&g!=""&&e.Image!=""{images[g]=e.Image}}
 if lookup==nil{return out,0}
 for _,e:=range events{
  g:=strings.ToLower(strings.TrimSpace(e.GUID));if e.ID!=3||images[g]==""{continue}
  ip,err:=netip.ParseAddr(e.DestinationIP);if err!=nil||ip.Zone()!=""{unknown++;continue}
  key:=g+"|"+ip.Unmap().String();if seen[key]{continue};seen[key]=true
  hit,err:=lookup(ip.Unmap().String());if err!=nil{unknown++;continue}
  if hit{out=append(out,NetworkSignal{g,images[g],ip.Unmap().String()})}
 }
 sort.Slice(out,func(i,j int)bool{if out[i].GUID==out[j].GUID{return out[i].DestinationIP<out[j].DestinationIP};return out[i].GUID<out[j].GUID})
 return out,unknown
}
