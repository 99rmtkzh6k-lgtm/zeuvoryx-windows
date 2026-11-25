package gamecatalog
import("testing";"strings")
func TestCatalog314(t *testing.T){
 if len(approved)!=10{t.Fatal("catalog count")}
 for k,v:=range approved{s,e:=Resolve(k);if e!=nil||s!=v||!strings.HasPrefix(s,"https://"){t.Fatal(k,s,e)}}
 for _,x:=range []string{"","https://example.com","file:///C:/","javascript:alert(1)","stargiraf&calc","../stargiraf"}{if _,e:=Resolve(x);e==nil{t.Fatal("unapproved ID",x)}}
}
