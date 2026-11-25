package scriptguard
import("encoding/binary";"strings";"testing";"unicode/utf16")
func encode(s string, order binary.ByteOrder) []byte {
 b:=[]byte{0xff,0xfe};if order==binary.BigEndian {b=[]byte{0xfe,0xff}}
 for _,v:=range utf16.Encode([]rune(s)){var a [2]byte;order.PutUint16(a[:],v);b=append(b,a[:]...)}
 return b
}
func TestUnicodeScripts(t *testing.T){
 for _,o:=range []binary.ByteOrder{binary.LittleEndian,binary.BigEndian}{
  for _,s:=range []string{"Write-Output 'hello'","texto português 日本語 🌞","Example indicator: Invoke-Expression"}{
   got,ok:=DecodeBOM(encode(s,o));if !ok||got!=s{t.Fatalf("decode: %q %v",got,ok)}
  }
 }
}
func TestRejectMalformed(t *testing.T){
 for _,b:=range [][]byte{nil,[]byte("MZ binary"),{0xff,0xfe,0x00},{0xff,0xfe,0,0},{0xff,0xfe,0,0xd8},{0xfe,0xff,0xdc,0},{0xff,0xfe,1,0},{0xff,0xfe}}{if _,ok:=DecodeBOM(b);ok{t.Fatalf("accepted %x",b)}}
}
func TestBound(t *testing.T){
 b:=encode(strings.Repeat("a",MaxDecodedBytes),binary.LittleEndian)
 got,ok:=DecodeBOM(b);if !ok||len(got)!=MaxDecodedBytes/2{t.Fatalf("size %d valid %v",len(got),ok)}
 b=encode(strings.Repeat("a",MaxDecodedBytes/2-1)+"🌞",binary.LittleEndian)
 got,ok=DecodeBOM(b);if !ok||len(got)!=MaxDecodedBytes/2-1{t.Fatalf("split pair %d valid %v",len(got),ok)}
}
