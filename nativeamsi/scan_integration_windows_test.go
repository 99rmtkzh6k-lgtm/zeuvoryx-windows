//go:build windows && amsi_integration

package nativeamsi

import "testing"

func TestNativeBenign(t *testing.T) {
 v,e:=Scan([]byte("ZEUVORYX harmless plain text functional check."),"benign.txt")
 if e!=nil{t.Fatal(e)}
 if v.Status!="clean_by_provider"&&v.Status!="not_detected_by_provider"{t.Fatalf("unexpected benign verdict %+v",v)}
}
