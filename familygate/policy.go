package familygate
import("crypto/rand";"crypto/sha256";"crypto/pbkdf2";"crypto/subtle";"encoding/json";"errors";"fmt";"os";"path/filepath";"sync";"time")
type Policy struct{Enabled bool `json:"enabled"`;MaxAge int `json:"max_age"`;Sex bool `json:"sex"`;Violence bool `json:"violence"`;Drugs bool `json:"drugs"`;Gambling bool `json:"gambling"`;Language bool `json:"language"`}
type Rating struct{Age int `json:"age"`;Sex bool `json:"sex"`;Violence bool `json:"violence"`;Drugs bool `json:"drugs"`;Gambling bool `json:"gambling"`;Language bool `json:"language"`;Source string `json:"source"`}
type record struct{Policy Policy;Ratings map[string]Rating;Salt,Hash []byte}
type Status struct{Policy Policy `json:"policy"`;Configured bool `json:"configured"`;Problem string `json:"problem,omitempty"`;Ratings map[string]Rating `json:"ratings"`}
type Gate struct{mu sync.Mutex;path string;r record;problem string;next time.Time}
func validAge(n int)bool{switch n{case 0,10,12,14,16,18:return true};return false}
func Open(path string)*Gate{
 g:=&Gate{path:path,r:record{Policy:Policy{MaxAge:12,Sex:true,Violence:true,Drugs:true,Gambling:true},Ratings:map[string]Rating{}}}
 if path==""{g.problem="Configuração local indisponível";return g}
 st,e:=os.Stat(path);if os.IsNotExist(e){return g};if e!=nil||st.Size()>65536{g.problem="Configuração familiar não pôde ser lida";return g}
 b,e:=os.ReadFile(path);if e!=nil||json.Unmarshal(b,&g.r)!=nil||len(g.r.Salt)!=16||len(g.r.Hash)!=32||!validAge(g.r.Policy.MaxAge){g.problem="Configuração familiar inválida; acesso online bloqueado";return g}
 if g.r.Ratings==nil{g.r.Ratings=map[string]Rating{}}
 for _,v:=range g.r.Ratings{if !validAge(v.Age)||v.Source!="parent_review"{g.problem="Classificação local inválida";break}}
 return g
}
func(g *Gate) Status()Status{g.mu.Lock();defer g.mu.Unlock();s:=Status{Policy:g.r.Policy,Configured:len(g.r.Hash)==32,Problem:g.problem,Ratings:map[string]Rating{}};for k,v:=range g.r.Ratings{s.Ratings[k]=v};return s}
func key(password string,salt []byte)([]byte,error){return pbkdf2.Key(sha256.New,password,salt,600000,32)}
func(g *Gate) authenticate(password string)([]byte,[]byte,error){
 if g.problem!=""{return nil,nil,errors.New(g.problem)}
 if time.Now().Before(g.next){return nil,nil,errors.New("Aguarde antes de tentar novamente")}
 if len(password)<10||len(password)>128{return nil,nil,errors.New("Senha do responsável: 10 a 128 caracteres")}
 if len(g.r.Hash)==0{salt:=make([]byte,16);if _,e:=rand.Read(salt);e!=nil{return nil,nil,e};h,e:=key(password,salt);return salt,h,e}
 h,e:=key(password,g.r.Salt);if e!=nil{return nil,nil,e};if subtle.ConstantTimeCompare(h,g.r.Hash)!=1{g.next=time.Now().Add(3*time.Second);return nil,nil,errors.New("Senha incorreta")}
 return g.r.Salt,g.r.Hash,nil
}
func(g *Gate) persist(r record)error{
 if e:=os.MkdirAll(filepath.Dir(g.path),0700);e!=nil{return e};f,e:=os.CreateTemp(filepath.Dir(g.path),"family-*.tmp");if e!=nil{return e};defer os.Remove(f.Name())
 e=json.NewEncoder(f).Encode(r);if e==nil{e=f.Sync()};ce:=f.Close();if e==nil{e=ce};if e!=nil{return e};if e=os.Rename(f.Name(),g.path);e!=nil{return e};g.r=r;return nil
}
func(g *Gate) Save(p Policy,password string)error{g.mu.Lock();defer g.mu.Unlock();if !validAge(p.MaxAge){return errors.New("Idade inválida")};salt,h,e:=g.authenticate(password);if e!=nil{return e};r:=g.r;r.Policy=p;r.Salt=salt;r.Hash=h;return g.persist(r)}
func(g *Gate) Rate(id string,rating Rating,password string)error{g.mu.Lock();defer g.mu.Unlock();if len(g.r.Hash)==0{return errors.New("Configure a senha do responsável primeiro")};if !validAge(rating.Age){return errors.New("Idade inválida")};_,_,e:=g.authenticate(password);if e!=nil{return e};r:=g.r;r.Ratings=map[string]Rating{};for k,v:=range g.r.Ratings{r.Ratings[k]=v};rating.Source="parent_review";r.Ratings[id]=rating;return g.persist(r)}
func(g *Gate) Allowed(id string)error{g.mu.Lock();defer g.mu.Unlock();if g.problem!=""{return errors.New(g.problem)};if !g.r.Policy.Enabled{return nil};r,ok:=g.r.Ratings[id];if !ok{return errors.New("Classificação não confirmada: responsável deve revisar o jogo")};p:=g.r.Policy;if r.Age>p.MaxAge||p.Sex&&r.Sex||p.Violence&&r.Violence||p.Drugs&&r.Drugs||p.Gambling&&r.Gambling||p.Language&&r.Language{return fmt.Errorf("Restrição do perfil familiar")};return nil}
