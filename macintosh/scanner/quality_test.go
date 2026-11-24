package scanner
import("context";"os";"path/filepath";"testing";"bytes")
func TestMacBenignAndFullHash(t *testing.T){
 d:=t.TempDir();b:=bytes.Repeat([]byte("ordinary document\n"),65536);p:=filepath.Join(d,"ordinary.txt")
 if e:=os.WriteFile(p,b,0600);e!=nil{t.Fatal(e)}
 r,e:=Scan(context.Background(),d);if e!=nil||r.Files!=1||r.Bytes!=int64(len(b))||r.Partial||len(r.Findings)!=0{t.Fatal(r,e)}
 got,e:=os.ReadFile(p);if e!=nil||!bytes.Equal(got,b){t.Fatal("file changed",e)}
}
func FuzzMacReadOnly(f *testing.F){
 f.Add([]byte("ordinary document"));f.Add([]byte{});f.Add([]byte{0xcf,0xfa,0xed,0xfe});f.Add(marker)
 f.Fuzz(func(t *testing.T,b []byte){
  if len(b)>65536{t.Skip()};p:=filepath.Join(t.TempDir(),"fixture.bin");if e:=os.WriteFile(p,b,0600);e!=nil{t.Fatal(e)}
  r,e:=Scan(context.Background(),p);if e!=nil||r.Files!=1||r.Partial{t.Fatal(r,e)}
  if len(r.Findings)>0&&(len(r.Findings)!=1||r.Findings[0].Kind!="test_eicar"||!bytes.Equal(b,marker)){t.Fatal("unexpected verdict",r)}
  got,e:=os.ReadFile(p);if e!=nil||!bytes.Equal(got,b){t.Fatal("file changed",e)}
 })
}
