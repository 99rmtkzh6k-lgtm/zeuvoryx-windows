//go:build windows

package nativeamsi

import "testing"

func TestVerdicts(t *testing.T) {
 cases:=map[uint32]string{0:"clean_by_provider",1:"not_detected_by_provider",2:"provider_score",16383:"provider_score",16384:"policy_block",20479:"policy_block",20480:"provider_score",32767:"provider_score",32768:"malware",65535:"malware"}
 for raw,want:=range cases {if got:=Interpret(raw);got.Raw!=raw||got.Status!=want{t.Errorf("%d: %+v",raw,got)}}
}
func TestBounds(t *testing.T) {
 for _,b:=range [][]byte{nil,make([]byte,MaxContent+1)}{if _,e:=Scan(b,"test");e==nil{t.Fatal("invalid size accepted")}}
}
func TestNativeBenign(t *testing.T) {
 v,e:=Scan([]byte("ZEUVORYX harmless plain text functional check."),"benign.txt")
 if e!=nil{t.Fatal(e)}
 if v.Status!="clean_by_provider"&&v.Status!="not_detected_by_provider"{t.Fatalf("unexpected benign verdict %+v",v)}
}
