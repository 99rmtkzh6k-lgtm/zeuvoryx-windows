// Package scriptguard recognizes bounded Unicode text for the local scanner.
package scriptguard

import (
 "encoding/binary"
 "unicode/utf16"
)

const MaxDecodedBytes = 256 << 10

// DecodeBOM accepts well-formed BOM-marked UTF-16 text. A bounded prefix
// is not full-file coverage. This function never executes content.
func DecodeBOM(b []byte) (string, bool) {
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
 u:=make([]uint16,0,len(b)/2)
 for i:=0;i<len(b);i+=2 {u=append(u,order.Uint16(b[i:i+2]))}
 if clipped && len(u)>0 && u[len(u)-1]>=0xd800 && u[len(u)-1]<=0xdbff {u=u[:len(u)-1]}
 for i:=0;i<len(u);i++ {
  c:=u[i]
  if c>=0xd800 && c<=0xdbff {if i+1>=len(u)||u[i+1]<0xdc00||u[i+1]>0xdfff{return "",false};i++;continue}
  if c>=0xdc00 && c<=0xdfff{return "",false}
  if c==0 || (c<32 && c!=9 && c!=10 && c!=12 && c!=13) {return "",false}
 }
 s:=string(utf16.Decode(u))
 return s,s!=""
}
