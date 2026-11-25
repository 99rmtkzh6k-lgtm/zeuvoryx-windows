// Local audit: authenticated records, not a remote or immutable ledger.
package audit
import("crypto/hmac";"crypto/rand";"crypto/sha256";"encoding/hex";"encoding/json";"errors";"io";"os";"path/filepath";"sort";"strings";"time")
const Version="3.0.9"
const maxRecord=64<<20
type Event struct { ID,ScanID,Kind,Mode,Version,Error string; At time.Time; Stars,Files,Skipped,Alerts,ErrorCount int; Bytes int64; Cancelled,Partial bool }
type Envelope struct { Data json.RawMessage; MAC string }
type History struct { Events []Event; Invalid int; Truncated bool; Rating,Finished int; Files,Bytes int64 }
type Store struct { dir string; key []byte }
func NewID()(string,error){var b [16]byte;_,e:=rand.Read(b[:]);return hex.EncodeToString(b[:]),e}
func validID(id string)bool{if len(id)!=32{return false};_,e:=hex.DecodeString(id);return e==nil&&id==strings.ToLower(id)}
func Default()(*Store,error){d,e:=os.UserConfigDir();if e!=nil{return nil,e};return Open(filepath.Join(d,"ZEUVORYX","lab305-audit"))}
func Open(dir string)(*Store,error){
 if e:=os.MkdirAll(dir,0700);e!=nil{return nil,e}
 p:=filepath.Join(dir,"key.bin");b,e:=os.ReadFile(p)
 if os.IsNotExist(e){
 var k [32]byte;if _,e=rand.Read(k[:]);e!=nil{return nil,e}
 f,err:=os.OpenFile(p,os.O_WRONLY|os.O_CREATE|os.O_EXCL,0600)
 if err==nil{_,err=f.Write(k[:]);if err==nil{err=f.Sync()};ce:=f.Close();if err==nil{err=ce};if err!=nil{return nil,err};b=k[:];e=nil
 }else if os.IsExist(err){b,e=os.ReadFile(p)}else{return nil,err}
 }
 if e!=nil{return nil,e};if len(b)!=32{return nil,errors.New("invalid audit key")};return &Store{dir,b},nil
}
func (s *Store) seal(data []byte)([]byte,error){
 if len(data)>maxRecord||!json.Valid(data){return nil,errors.New("invalid or oversized audit record")}
 var r json.RawMessage=data;b,e:=json.Marshal(r);if e!=nil{return nil,e}
 h:=hmac.New(sha256.New,s.key);h.Write(b);return json.Marshal(Envelope{b,hex.EncodeToString(h.Sum(nil))})
}
func(s *Store) verify(raw []byte)(json.RawMessage,error){
 var v Envelope;if e:=json.Unmarshal(raw,&v);e!=nil{return nil,e};want,e:=hex.DecodeString(v.MAC);if e!=nil{return nil,e}
 h:=hmac.New(sha256.New,s.key);h.Write(v.Data);if !hmac.Equal(want,h.Sum(nil)){return nil,errors.New("audit authentication failed")};return v.Data,nil
}
func(s *Store) write(name string,data []byte)error{
 b,e:=s.seal(data);if e!=nil{return e};f,e:=os.CreateTemp(s.dir,".pending-");if e!=nil{return e};defer os.Remove(f.Name())
 _,e=f.Write(b);if e==nil{e=f.Sync()};ce:=f.Close();if e==nil{e=ce};if e!=nil{return e}
 dest:=filepath.Join(s.dir,name);if _,e=os.Lstat(dest);!os.IsNotExist(e){return errors.New("audit destination exists or inaccessible")};return os.Rename(f.Name(),dest)
}
func(s *Store) read(name string)(json.RawMessage,error){
 f,e:=os.Open(filepath.Join(s.dir,name));if e!=nil{return nil,e};defer f.Close()
 b,e:=io.ReadAll(io.LimitReader(f,maxRecord+4097));if e!=nil{return nil,e};if len(b)>maxRecord+4096{return nil,errors.New("record too large")};return s.verify(b)
}
func(s *Store) Add(v Event)(Event,error){
 if v.ID!=""{return v,errors.New("invalid preassigned event ID")}
 if v.Kind=="rating"{if v.Stars<1||v.Stars>5{return v,errors.New("rating must be 1–5")}}else if !validID(v.ScanID){return v,errors.New("invalid scan ID")}
 id,e:=NewID();if e!=nil{return v,e};v.ID=id;v.At=time.Now().UTC();v.Version=Version;b,e:=json.Marshal(v);if e!=nil{return v,e};return v,s.write("event-"+id+".json",b)
}
func(s *Store) SaveReport(id string,report any)error{if !validID(id){return errors.New("invalid scan ID")};b,e:=json.Marshal(report);if e!=nil{return e};return s.write("report-"+id+".json",b)}
func(s *Store) ReadReport(id string)(json.RawMessage,error){if !validID(id){return nil,errors.New("invalid scan ID")};return s.read("report-"+id+".json")}
func(s *Store) Finish(id,mode string,report any,v Event)error{if e:=s.SaveReport(id,report);e!=nil{return e};v.Kind="scan_finished";v.ScanID=id;v.Mode=mode;_,e:=s.Add(v);return e}
func(s *Store) Rating(stars int)(Event,error){return s.Add(Event{Kind:"rating",Stars:stars})}
func(s *Store) History()(History,error){
 out:=History{Events:[]Event{}};entries,e:=os.ReadDir(s.dir);if e!=nil{return out,e};var rate time.Time;complete:=map[string]bool{}
 for _,entry:=range entries{
 n:=entry.Name();if !strings.HasPrefix(n,"event-")||!strings.HasSuffix(n,".json"){continue};raw,err:=s.read(n);var v Event
 if err!=nil||json.Unmarshal(raw,&v)!=nil||!validID(v.ID)||n!="event-"+v.ID+".json"{out.Invalid++;continue}
 if v.Kind=="rating"&&v.Stars>=1&&v.Stars<=5&&v.At.After(rate){out.Rating=v.Stars;rate=v.At}
 if v.Kind=="scan_finished"&&!complete[v.ScanID]{complete[v.ScanID]=true;out.Finished++;out.Files+=int64(v.Files);out.Bytes+=v.Bytes}
 out.Events=append(out.Events,v)
 }
 sort.Slice(out.Events,func(i,j int)bool{return out.Events[i].At.After(out.Events[j].At)})
 if len(out.Events)>100{out.Events=out.Events[:100];out.Truncated=true};return out,nil
}
