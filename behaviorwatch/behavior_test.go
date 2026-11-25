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

func TestBoundedCorrelation315(t *testing.T){
 if _,e:=Parse(strings.NewReader("<Events>"+strings.Repeat(" ",8<<20)+"</Events>"));e==nil{t.Fatal("oversized XML accepted")}
 es:=[]Event{{ID:1,GUID:"ABC",Image:"C:\\Users\\lab\\Downloads\\sample.exe"}}
 for _,p:=range []string{"a","b","c"}{es=append(es,Event{ID:11,GUID:"abc",Target:p})}
 for i:=0;i<4;i++{es=append(es,Event{ID:13,GUID:"abc",Target:"HKU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run\\example"})}
 got:=Analyze(es);if len(got)!=1||got[0].StartupWrites!=1{t.Fatalf("duplicate evidence %+v",got)}
}
