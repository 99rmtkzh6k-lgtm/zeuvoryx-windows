package main
import "testing"
func Test310DriveArguments(t *testing.T){
 for _,r:=range []string{"C:\\","D:\\"}{a,e:=drive310Args(r);if e!=nil||a!=r[:2]+" /O /U /V"{t.Fatal(r,a,e)}}
 for _,r:=range []string{"C:","C:\\Windows","\\\\server\\share","C: /X","1:\\",""}{if _,e:=drive310Args(r);e==nil{t.Fatal("unsafe volume accepted",r)}}
}
