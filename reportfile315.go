package main

import (
 "bytes"
 "encoding/json"
 "fmt"
 "html/template"
 "io"
 "os"
 "path/filepath"
 "sync"
)

// Save to a unique real file. No partial file survives a failed write.
func writeReport315(dir,prefix string,write func(io.Writer)error)(string,error){return writeReportExt315(dir,prefix,".json",write)}
func writeReportExt315(dir,prefix,ext string,write func(io.Writer)error)(string,error){
 if ext!=".json"&&ext!=".html"{return "",fmt.Errorf("Invalid report extension")}
 if e:=os.MkdirAll(dir,0700);e!=nil{return "",e}
 f,e:=os.CreateTemp(dir,prefix+"-*"+ext+".tmp");if e!=nil{return "",e}
 temp:=f.Name();name:=temp[:len(temp)-len(".tmp")];ok:=false
 defer func(){if !ok{_ = f.Close();_ = os.Remove(temp)}}()
 if e=write(f);e!=nil{return "",e}
 if e=f.Sync();e!=nil{return "",e}
 if e=f.Close();e!=nil{return "",e}
 st,e:=os.Stat(temp);if e!=nil{return "",e};if !st.Mode().IsRegular()||st.Size()==0{return "",fmt.Errorf("Relatório vazio ou inválido")}
 // Publish only after the full write is closed and synced; interrupted writes
 // leave a .tmp, never a misleading completed .json. Same-directory rename.
 if _,e=os.Lstat(name);!os.IsNotExist(e){return "",fmt.Errorf("Destino de relatório já existe ou não pôde ser verificado")}
 if e=os.Rename(temp,name);e!=nil{return "",e}
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

const maxReportText315=2<<20
// Do not truncate JSON or silently change evidence. Large reports remain files.
func formatRawReportText315(data []byte)(string,error){
 if len(data)>maxReportText315{return "",fmt.Errorf("Relatório grande demais para a caixa de texto; use Salvar arquivo JSON")}
 if !json.Valid(data){return "",fmt.Errorf("Relatório JSON inválido")}
 var out bytes.Buffer
 if e:=json.Indent(&out,data,"","  ");e!=nil{return "",e}
 if out.Len()>maxReportText315{return "",fmt.Errorf("Relatório grande demais para a caixa de texto; use Salvar arquivo JSON")}
 return out.String(),nil
}
func formatReportText315(v any)(string,error){
 b,e:=json.Marshal(v);if e!=nil{return "",e};return formatRawReportText315(b)
}
func readLastReportText315()(string,error){
 path,e:=lastReport315();if e!=nil{return "",e}
 f,e:=os.Open(path);if e!=nil{return "",e};defer f.Close()
 b,e:=io.ReadAll(io.LimitReader(f,maxReportText315+1));if e!=nil{return "",e}
 return formatRawReportText315(b)
}

var htmlReportTemplate315=template.Must(template.New("report").Parse(`<!doctype html><html lang="pt-BR"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>ZEUVORYX LAB — Relatório</title><style>body{font:16px system-ui;margin:2rem auto;max-width:1000px;padding:0 1rem;color:#182438}pre{white-space:pre-wrap;overflow-wrap:anywhere;background:#f0f3f7;padding:1rem;border-radius:8px}dl{display:grid;grid-template-columns:max-content 1fr;gap:8px}dt{font-weight:bold}</style><h1>ZEUVORYX LAB — Relatório local</h1><p>Scanner sob demanda. Sem bloqueio preventivo. Resultados parciais não comprovam ausência de ameaças.</p><h2>Resumo</h2><dl>{{range .Summary}}<dt>{{.Name}}</dt><dd>{{.Value}}</dd>{{end}}</dl><h2>Dados completos (findings, skipped_details e demais evidências)</h2><pre>{{.JSON}}</pre></html>`))
func writeHTMLReport315(dir,prefix string,v any)(string,error){
 var buf bytes.Buffer;enc:=json.NewEncoder(&buf);enc.SetEscapeHTML(false);enc.SetIndent("","  ");if e:=enc.Encode(v);e!=nil{return "",e};raw:=buf.Bytes()
 var fields map[string]any;if e:=json.Unmarshal(raw,&fields);e!=nil{return "",e}
 type pair struct{Name,Value string}
 data:=struct{Summary []pair;JSON string}{JSON:string(raw)}
 for _,k:=range []string{"scan_id","started","finished","partial","files","skipped","error_count"}{if value,ok:=fields[k];ok{data.Summary=append(data.Summary,pair{k,fmt.Sprint(value)})}}
 return writeReportExt315(dir,prefix,".html",func(w io.Writer)error{return htmlReportTemplate315.Execute(w,data)})
}
func htmlReport315(prefix string,v any)(string,error){
 home,e:=os.UserHomeDir();if e!=nil{return "",e};return writeHTMLReport315(filepath.Join(home,"ZEUVORYX-Relatorios"),prefix,v)
}
