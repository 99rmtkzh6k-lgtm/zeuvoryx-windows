package releaseguard
import("bytes";"crypto/ed25519";"crypto/rand";"crypto/sha256";"encoding/hex";"encoding/json";"errors";"testing";"time")
type failedReader struct{}
func(failedReader)Read([]byte)(int,error){return 0,errors.New("simulated interrupted stream")}
func TestAuthorizedUpdate(t *testing.T){
 pub,priv,e:=ed25519.GenerateKey(rand.Reader);if e!=nil{t.Fatal(e)}
 other,_,e:=ed25519.GenerateKey(rand.Reader);if e!=nil{t.Fatal(e)}
 now:=time.Unix(1800000000,0)
 data:=[]byte("inert test package")
 hash:=sha256.Sum256(data)
 good:=Manifest{2,"3.0.3","beta","darwin-universal","https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/releases/download/mac-v3.0.3/test.zip",hex.EncodeToString(hash[:]),int64(len(data)),now.Add(time.Hour).Unix()}
 sign:=func(m Manifest)([]byte,[]byte){b,e:=json.Marshal(m);if e!=nil{t.Fatal(e)};return b,ed25519.Sign(priv,append([]byte(Domain),b...))}
 raw,sig:=sign(good)
 m,e:=Verify(raw,sig,pub,1,"beta",now);if e!=nil{t.Fatal(e)}
 if e=VerifyPackage(bytes.NewReader(data),m);e!=nil{t.Fatal(e)}
 if VerifyPackage(nil,m)==nil || VerifyPackage(failedReader{},m)==nil {t.Fatal("accepted invalid stream")}
 if VerifyPackage(bytes.NewReader(data),Manifest{Size:0})==nil {t.Fatal("accepted zero size")}
 if _,e=Verify(raw,sig,nil,1,"beta",now);e==nil{t.Fatal("accepted missing key")}
 if _,e=Verify(raw,nil,pub,1,"beta",now);e==nil{t.Fatal("accepted missing signature")}
 for _,bad:=range [][]byte{[]byte("inert test packagE"),data[:len(data)-1],append(append([]byte{},data...),0)}{if VerifyPackage(bytes.NewReader(bad),m)==nil{t.Fatal("accepted modified package")}}
 if _,e=Verify(raw,sig,other,1,"beta",now);e==nil{t.Fatal("accepted wrong owner")}
 if _,e=Verify(raw,sig,pub,2,"beta",now);e==nil{t.Fatal("accepted rollback/replay")}
 if _,e=Verify(raw,sig,pub,1,"stable",now);e==nil{t.Fatal("accepted wrong channel")}
 tampered:=append([]byte{},raw...);tampered[0]='['
 if _,e=Verify(tampered,sig,pub,1,"beta",now);e==nil{t.Fatal("accepted tampered bytes")}
 for _,change:=range []func(*Manifest){
  func(m *Manifest){m.Expires=now.Unix()},
  func(m *Manifest){m.Expires=now.Add(31*24*time.Hour).Unix()},
  func(m *Manifest){m.URL="https://github.com.evil.example/test.zip"},
  func(m *Manifest){m.URL="https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/releases/download/../test.zip"},
  func(m *Manifest){m.Platform="linux-amd64"},
  func(m *Manifest){m.Version="3.0.4"},
  func(m *Manifest){m.Version="03.0.3"},
  func(m *Manifest){m.URL="https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/releases/download/mac-v3.0.3/extra/test.zip"},
  func(m *Manifest){m.URL="https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/releases/download/mac-v3.0.3/test.zip?"},
  func(m *Manifest){m.URL="https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/releases/download/mac-v3.0.3/test.txt"},
  func(m *Manifest){m.URL="https://github.com/99rmtkzh6k-lgtm/zeuvoryx-windows/releases/download/mac-v3.0.3/test%2f.zip"},
  func(m *Manifest){m.Size=MaxPackage+1},
  func(m *Manifest){m.SHA256="not a hash"},
 }{bad:=good;change(&bad);b,s:=sign(bad);if _,e=Verify(b,s,pub,1,"beta",now);e==nil{t.Fatalf("accepted bad manifest: %+v",bad)}}
 duplicate:=append([]byte{},raw[:len(raw)-1]...);duplicate=append(duplicate,[]byte(",\"sequence\":2}")...)
 if _,e=Verify(duplicate,ed25519.Sign(priv,append([]byte(Domain),duplicate...)),pub,1,"beta",now);e==nil{t.Fatal("accepted duplicate field")}
}

func FuzzSignedManifest(f *testing.F){
 pub,priv,e:=ed25519.GenerateKey(rand.Reader);if e!=nil{f.Fatal(e)}
 f.Add([]byte("{}"));f.Add([]byte("null"));f.Add([]byte("invalid"))
 f.Fuzz(func(t *testing.T,b []byte){
  if len(b)>MaxManifest+1{return}
  sig:=ed25519.Sign(priv,append([]byte(Domain),b...))
  m,e:=Verify(b,sig,pub,0,"beta",time.Unix(1800000000,0))
  if e==nil{canonical,_:=json.Marshal(m);if !bytes.Equal(b,canonical)||m.Sequence==0{t.Fatal("invalid accepted manifest")}}
 })
}
