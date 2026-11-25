package globalintel
import("context";"encoding/json";"errors";"os";"path/filepath";"sync";"time")
type cache314 struct{Attempt time.Time; V4,V6 []byte}
type Daily struct{mu sync.Mutex;path string;state cache314;data Drop;problem string;busy bool}
func OpenDaily(path string)*Daily{
 d:=&Daily{path:path,data:Bundled(time.Now())}
 st,e:=os.Stat(path);if e!=nil||st.Size()>24<<20{return d}
 b,e:=os.ReadFile(path);if e!=nil||json.Unmarshal(b,&d.state)!=nil{return d}
 now:=time.Now();a,e:=parseDrop(d.state.V4,false,now);if e!=nil{return d};v,e:=parseDrop(d.state.V6,true,now);if e!=nil{return d}
 d.data=Drop{V4:a,V6:v};return d
}
func(d *Daily) Status()Status{
 d.mu.Lock();defer d.mu.Unlock();s:=d.data.Status();s.Problem=d.problem
 if s.Updated!=""{t,e:=time.Parse(time.RFC3339,s.Updated);s.Stale=e!=nil||time.Since(t)>26*time.Hour}else{s.Stale=true}
 return s
}
func(d *Daily) Match(ip string,now time.Time)(bool,error){d.mu.Lock();defer d.mu.Unlock();return d.data.Match(ip,now)}
func(d *Daily) persist()error{
 if d.path==""{return errors.New("cache path unavailable")}
 if e:=os.MkdirAll(filepath.Dir(d.path),0700);e!=nil{return e}
 f,e:=os.CreateTemp(filepath.Dir(d.path),"drop-*.tmp");if e!=nil{return e};defer os.Remove(f.Name())
 e=json.NewEncoder(f).Encode(d.state);if e==nil{e=f.Sync()};ce:=f.Close();if e==nil{e=ce};if e!=nil{return e};return os.Rename(f.Name(),d.path)
}
func(d *Daily) Refresh(ctx context.Context)(status Status,err error){
 d.mu.Lock();now:=time.Now()
 if d.busy||(!d.state.Attempt.IsZero()&&now.Sub(d.state.Attempt)<24*time.Hour){d.mu.Unlock();return d.Status(),errors.New("Limite de uma tentativa por 24 horas; confira a data da base")}
 d.busy=true;d.state.Attempt=now
 if e:=d.persist();e!=nil{d.busy=false;d.problem=e.Error();d.mu.Unlock();return d.Status(),e}
 d.mu.Unlock()
 defer func(){d.mu.Lock();d.busy=false;if err!=nil{d.problem=err.Error()}else{d.problem=""};d.mu.Unlock();status=d.Status()}()
 c:=client();a,e:=fetch(ctx,c,drop4);if e!=nil{return Status{},e};v4,e:=parseDrop(a,false,now);if e!=nil{return Status{},e}
 b,e:=fetch(ctx,c,drop6);if e!=nil{return Status{},e};v6,e:=parseDrop(b,true,now);if e!=nil{return Status{},e}
 d.mu.Lock();defer d.mu.Unlock();previous:=d.state;d.state.V4=a;d.state.V6=b
 if e=d.persist();e!=nil{d.state=previous;return Status{},e}
 d.data=Drop{V4:v4,V6:v6};return Status{},nil
}
