package behaviorwatch
import("errors";"testing")
func TestNetworkCorrelationLab(t *testing.T){
 es:=[]Event{{ID:1,GUID:"A",Image:"example.exe"},{ID:3,GUID:"a",DestinationIP:"192.0.2.1"},{ID:3,GUID:"a",DestinationIP:"192.0.2.1"},{ID:3,GUID:"a",DestinationIP:"198.51.100.1"},{ID:3,GUID:"b",DestinationIP:"192.0.2.1"}}
 hits,unknown:=CorrelateNetwork(es,func(s string)(bool,error){if s=="198.51.100.1"{return false,errors.New("unavailable")};return s=="192.0.2.1",nil})
 if len(hits)!=1||unknown!=1||hits[0].Image!="example.exe"{t.Fatal(hits,unknown)}
}
