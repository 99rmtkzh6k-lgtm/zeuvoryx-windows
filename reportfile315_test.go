package main
import("encoding/json";"errors";"io";"os";"path/filepath";"testing")
func TestReport315DurableFile(t *testing.T){
 dir:=filepath.Join(t.TempDir(),"ZEUVORYX-Relatórios")
 for _,partial:=range []bool{true,false}{
  name,e:=writeReport315(dir,"varredura",func(w io.Writer)error{return json.NewEncoder(w).Encode(map[string]any{"partial":partial,"files":42,"windows_local_314":map[string]any{"source":"Defender"}})})
  if e!=nil{t.Fatal(e)};b,e:=os.ReadFile(name);if e!=nil{t.Fatal(e)}
  var v struct{Partial bool;Files int};if e=json.Unmarshal(b,&v);e!=nil||v.Partial!=partial||v.Files!=42{t.Fatalf("JSON inválido %s %v",b,e)}
 }
 if _,e:=writeReport315(dir,"fail",func(w io.Writer)error{_,_=io.WriteString(w,"partial");return errors.New("write failure")});e==nil{t.Fatal("false success")}
 if _,e:=writeReport315(dir,"empty",func(w io.Writer)error{return nil});e==nil{t.Fatal("empty success")}
 files,e:=os.ReadDir(dir);if e!=nil||len(files)!=2{t.Fatalf("residue %v %v",files,e)}
}

func TestReport315DurableFileAutomatic(t *testing.T){
 dir:=t.TempDir()
 first,e:=writeAutoReport315(dir,"zeuvoryx-auto-parcial",map[string]any{"partial":true,"files":3});if e!=nil{t.Fatal(e)}
 final,e:=writeAutoReport315(dir,"zeuvoryx-auto-final",map[string]any{"partial":false,"files":4});if e!=nil{t.Fatal(e)}
 if first==final{t.Fatal("overwrote recovery file")}
 for _,p:=range []string{first,final}{b,e:=os.ReadFile(p);if e!=nil||!json.Valid(b){t.Fatalf("invalid JSON: %s %v",p,e)}}
 if _,e=writeAutoReport315(dir,"fail",make(chan int));e==nil{t.Fatal("encoding failure reported as success")}
 files,e:=os.ReadDir(dir);if e!=nil||len(files)!=2{t.Fatalf("unexpected residue: %v %v",files,e)}
}

func TestReport315DurableFileCopyText(t *testing.T){
 s,e:=formatReportText315(map[string]any{"partial":true,"files":42,"note":"ação"})
 if e!=nil||!json.Valid([]byte(s)){t.Fatalf("invalid display: %q %v",s,e)}
 var v struct{Partial bool;Files int;Note string};if e=json.Unmarshal([]byte(s),&v);e!=nil||!v.Partial||v.Files!=42||v.Note!="ação"{t.Fatal(v,e)}
 if _,e=formatRawReportText315([]byte("{"));e==nil{t.Fatal("accepted invalid JSON")}
 if _,e=formatRawReportText315(make([]byte,maxReportText315+1));e==nil{t.Fatal("accepted oversized report")}
}

func TestReport315DurableFilePublishAfterWrite(t *testing.T){
 dir:=t.TempDir()
 name,e:=writeReport315(dir,"atomic",func(w io.Writer)error{
  if _,e:=io.WriteString(w,"{");e!=nil{return e}
  paths,e:=filepath.Glob(filepath.Join(dir,"*.json"));if e!=nil||len(paths)!=0{t.Fatalf("unfinished JSON exposed: %v %v",paths,e)}
  _,e=io.WriteString(w,"\"partial\":true}");return e
 })
 if e!=nil{t.Fatal(e)}
 b,e:=os.ReadFile(name);if e!=nil||!json.Valid(b){t.Fatal("invalid published JSON",e)}
 tmp,e:=filepath.Glob(filepath.Join(dir,"*.tmp"));if e!=nil||len(tmp)!=0{t.Fatal("temporary residue",tmp,e)}
}
