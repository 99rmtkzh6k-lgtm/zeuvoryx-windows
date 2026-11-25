package behaviorwatch
import("strings";"testing")
func TestCorrelation313(t *testing.T){
 raw:="<Events><Event><System><EventID>1</EventID></System><EventData><Data Name=\"ProcessGuid\">a</Data><Data Name=\"Image\">C:\\Users\\lab\\Downloads\\app.exe</Data></EventData></Event></Events>"
 es,e:=Parse(strings.NewReader(raw));if e!=nil||len(es)!=1{t.Fatalf("%+v %v",es,e)}
 if len(Analyze(es))!=0{t.Fatal("start alone flagged")}
 for _,p:=range []string{"a","b","c"}{es=append(es,Event{ID:11,GUID:"a",Target:p})}
 es=append(es,Event{ID:13,GUID:"b",Target:"HKU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run\\name"})
 if len(Analyze(es))!=0{t.Fatal("different process correlated")}
 es=append(es,Event{ID:13,GUID:"a",Target:"HKU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run\\name"})
 if len(Analyze(es))!=1{t.Fatal("correlation not found")}
}
