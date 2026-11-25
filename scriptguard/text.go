// Package scriptguard recognizes bounded Unicode text for the local scanner.
package scriptguard

import (
 "encoding/binary"
 "bytes"
 "unicode/utf8"
 "unicode/utf16"
 "strings"
)

const MaxDecodedBytes = 256 << 10

// DecodeBOM accepts well-formed BOM-marked UTF-8, UTF-16 and UTF-32 text. A bounded prefix
// is not full-file coverage. This function never executes content.
func DecodeBOM(b []byte) (string, bool) {
 // Test four-byte signatures before UTF-16, whose LE mark is a prefix.
 if bytes.HasPrefix(b,[]byte{0xff,0xfe,0,0}) {return decode32(b[4:],binary.LittleEndian)}
 if bytes.HasPrefix(b,[]byte{0,0,0xfe,0xff}) {return decode32(b[4:],binary.BigEndian)}
 if bytes.HasPrefix(b,[]byte{0xef,0xbb,0xbf}) {return decode8(b[3:])}
 if len(b)<2 || len(b)%2!=0 { return "",false }
 var order binary.ByteOrder
 switch {
 case b[0]==0xff && b[1]==0xfe: order=binary.LittleEndian
 case b[0]==0xfe && b[1]==0xff: order=binary.BigEndian
 default: return "",false
 }
 b=b[2:]
 if len(b)==0 {return "",false}
 clipped:=len(b)>MaxDecodedBytes
 if clipped {b=b[:MaxDecodedBytes]}
 // Decode directly, avoiding intermediate UTF-16 and rune allocations.
 var out strings.Builder
 out.Grow(len(b)/2)
 for i:=0;i<len(b);i+=2{
  v:=rune(order.Uint16(b[i:i+2]))
  if v>=0xd800&&v<=0xdbff{
   if i+2>=len(b){if clipped{break};return "",false}
   low:=rune(order.Uint16(b[i+2:i+4]))
   if low<0xdc00||low>0xdfff{return "",false}
   v=utf16.DecodeRune(v,low);i+=2
  }else if v>=0xdc00&&v<=0xdfff{return "",false}
  if !allowed(v){return "",false}
  out.WriteRune(v)
 }
 return out.String(),out.Len()>0
}

func allowed(r rune) bool {
 return r!=0 && (r>=32 || r==9 || r==10 || r==12 || r==13)
}
func decode32(b []byte, order binary.ByteOrder)(string,bool){
 if len(b)==0 || len(b)%4!=0{return "",false}
 if len(b)>MaxDecodedBytes{b=b[:MaxDecodedBytes]}
 out:=make([]rune,0,len(b)/4)
 for i:=0;i<len(b);i+=4{
  v:=order.Uint32(b[i:i+4])
  if v>utf8.MaxRune || !utf8.ValidRune(rune(v)) || !allowed(rune(v)){return "",false}
  out=append(out,rune(v))
 }
 return string(out),true
}
func decode8(b []byte)(string,bool){
 clipped:=len(b)>MaxDecodedBytes
 if clipped{b=b[:MaxDecodedBytes]}
 end:=0
 for end<len(b){
  if clipped&&!utf8.FullRune(b[end:]){break}
  r,n:=utf8.DecodeRune(b[end:])
  if (r==utf8.RuneError&&n==1)||!allowed(r){return "",false}
  end+=n
 }
 if end==0{return "",false}
 return string(b[:end]),true
}
