package signatureav
import "testing"
func TestStreamingObserver314(t *testing.T){
 var hits []Finding
 o:=lineObserver314{fn:func(f Finding){hits=append(hits,f)}}
 o.feed([]byte("C:\\test.txt: Eicar-Test"));if len(hits)!=0{t.Fatal("premature")}
 o.feed([]byte("-Signature FOUND\r\nbenign: OK\n"));if len(hits)!=1||hits[0].Kind!="test_eicar"{t.Fatal(hits)}
}
