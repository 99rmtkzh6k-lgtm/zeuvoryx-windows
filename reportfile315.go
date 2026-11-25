package main

import (
 "encoding/json"
 "fmt"
 "io"
 "os"
 "path/filepath"
 "sync"
)

// Save to a unique real file. No partial file survives a failed write.
func writeReport315(dir, prefix string, write func(io.Writer) error) (string,error) {
 if e:=os.MkdirAll(dir,0700);e!=nil{return "",e}
 f,e:=os.CreateTemp(dir,prefix+"-*.json");if e!=nil{return "",e}
 name:=f.Name();ok:=false
 defer func(){if !ok{_ = f.Close();_ = os.Remove(name)}}()
 if e=write(f);e!=nil{return "",e}
 if e=f.Sync();e!=nil{return "",e}
 if e=f.Close();e!=nil{return "",e}
 st,e:=os.Stat(name);if e!=nil{return "",e};if !st.Mode().IsRegular()||st.Size()==0{return "",fmt.Errorf("Relatório vazio ou inválido")}
 ok=true
 return name,nil
}
var transfer315 struct{sync.Mutex;path string;info os.FileInfo}
func saveReport315(prefix string,write func(io.Writer)error)(string,error){
 home,e:=os.UserHomeDir();if e!=nil{return "",e}
 path,e:=writeReport315(filepath.Join(home,"ZEUVORYX-Relatorios"),prefix,write);if e!=nil{return "",e}
 st,e:=os.Lstat(path);if e!=nil{return "",e}
 transfer315.Lock();transfer315.path=path;transfer315.info=st;transfer315.Unlock()
 return path,nil
}
func jsonReport315(prefix string,v any)(string,error){return saveReport315(prefix,func(w io.Writer)error{return json.NewEncoder(w).Encode(v)})}
func lastReport315()(string,error){
 transfer315.Lock();defer transfer315.Unlock()
 if transfer315.path==""||transfer315.info==nil{return "",fmt.Errorf("Salve um relatório primeiro")}
 st,e:=os.Lstat(transfer315.path);if e!=nil{return "",e}
 if !st.Mode().IsRegular()||!os.SameFile(st,transfer315.info)||st.Size()!=transfer315.info.Size()||!st.ModTime().Equal(transfer315.info.ModTime()){return "",fmt.Errorf("O arquivo foi alterado; salve novamente")}
 return transfer315.path,nil
}

// Automatic recovery uses separate files; it never changes the manual transfer selection.
func writeAutoReport315(dir,prefix string,v any)(string,error){
 return writeReport315(dir,prefix,func(w io.Writer)error{return json.NewEncoder(w).Encode(v)})
}
