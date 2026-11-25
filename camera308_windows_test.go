//go:build windows

package main
import "testing"
func TestCameraEvidence(t *testing.T){
 cases:=[]struct{rows []cameraEvidence;partial bool;want string}{
 {nil,false,"unknown"},
 {[]cameraEvidence{{0,0,true}},false,"unknown"},
 {[]cameraEvidence{{5,4,true}},false,"unknown"},
 {[]cameraEvidence{{5,0,false}},false,"unknown"},
 {[]cameraEvidence{{5,6,true}},false,"no_activity_recorded"},
 {[]cameraEvidence{{5,6,true}},true,"unknown"},
 {[]cameraEvidence{{5,0,true}},true,"activity"},
 }
 for _,c:=range cases{if got:=cameraClassify(c.rows,c.partial);got!=c.want{t.Errorf("%+v: %s",c,got)}}
}
func TestCameraReadOnly(t *testing.T){
 var c cameraSampler;a:=c.sample();b:=c.sample()
 if a!=b||(a!="unknown"&&a!="activity"&&a!="no_activity_recorded"){t.Fatal("invalid camera status")}
}
