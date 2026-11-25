// Package scopeprogress measures logical file lengths, not disk allocation.
package scopeprogress
import("context";"io/fs";"path/filepath";"time")
type Inventory struct{
 State string `json:"state"`
 Files int64 `json:"files"`
 Bytes int64 `json:"bytes"`
 Errors int64 `json:"errors"`
 Excluded int64 `json:"excluded"`
 ReadyAt time.Time `json:"ready_at"`
 ScannerDone bool `json:"scanner_done"`
}
func Enumerate(ctx context.Context,roots []string,exclude func(string)bool,update func(Inventory))Inventory{
 r:=Inventory{State:"counting"};last:=time.Now()
 for _,root:=range roots{
  if ctx.Err()!=nil{break}
  err:=filepath.WalkDir(root,func(path string,d fs.DirEntry,e error)error{
   if ctx.Err()!=nil{return ctx.Err()}
   if time.Since(last)>=500*time.Millisecond&&update!=nil{update(r);last=time.Now()}
   if e!=nil{r.Errors++;return nil}
   if exclude!=nil&&exclude(path){r.Excluded++;if d.IsDir(){return filepath.SkipDir};return nil}
   if d.IsDir(){return nil}
   if !d.Type().IsRegular(){r.Excluded++;return nil}
   st,e:=d.Info();if e!=nil{r.Errors++;return nil}
   if !st.Mode().IsRegular()||st.Size()<0{r.Excluded++;return nil}
   if st.Size()>int64(^uint64(0)>>1)-r.Bytes{r.Errors++;return fs.SkipAll}
   r.Files++;r.Bytes+=st.Size();return nil
  })
  if err!=nil&&ctx.Err()==nil{r.Errors++}
 }
 r.ReadyAt=time.Now().UTC()
 switch{case ctx.Err()!=nil:r.State="cancelled";case len(roots)==0:r.State="unavailable";case r.Errors>0:r.State="partial";default:r.State="ready"}
 if update!=nil{update(r)}
 return r
}
