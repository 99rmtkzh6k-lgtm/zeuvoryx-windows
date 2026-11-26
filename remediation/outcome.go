package remediation

import (
 "crypto/rand"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "errors"
 "fmt"
 "io"
 "os"
 "path/filepath"
 "strings"
 "time"
)

const EICARSHA256 = "275a021bbfb6489e54d471899f7db9d1663fc695ec2fe2a2c4538aabf651fd0f"

type Outcome struct {
 ID string `json:"event_id"`
 At time.Time `json:"at"`
 Path string `json:"original_path"`
 SHA256 string `json:"sha256,omitempty"`
 Detector string `json:"detector"`
 Executor string `json:"executor,omitempty"`
 Signature string `json:"signature,omitempty"`
 ScanStarted time.Time `json:"scan_started"`
 TestArtifact bool `json:"test_artifact"`
 Action string `json:"action"`
 Status string `json:"status"`
 QuarantineID string `json:"quarantine_id,omitempty"`
 OriginalAbsent bool `json:"original_absent"`
 CopyVerified bool `json:"copy_verified"`
 Error string `json:"error,omitempty"`
}

// An absent pathname alone never establishes who acted on it.
func CheckQuarantine(original, payload, expected string) (absent, verified bool, err error) {
 if _, e := os.Lstat(original); os.IsNotExist(e) { absent=true } else if e!=nil { return false,false,e }
 if len(expected)!=64 { return absent,false,errors.New("SHA-256 indisponível") }
 if _,e:=hex.DecodeString(expected);e!=nil{return absent,false,e}
 st,e:=os.Lstat(payload);if e!=nil{return absent,false,e}
 if !st.Mode().IsRegular(){return absent,false,errors.New("cópia não regular")}
 f,e:=os.Open(payload);if e!=nil{return absent,false,e};defer f.Close()
 opened,e:=f.Stat();if e!=nil{return absent,false,e}
 if !os.SameFile(st,opened){return absent,false,errors.New("cópia mudou")}
 h:=sha256.New();n,e:=io.Copy(h,f);if e!=nil{return absent,false,e}
 end,e:=f.Stat();if e!=nil{return absent,false,e};current,e:=os.Lstat(payload);if e!=nil{return absent,false,e}
 if n!=st.Size()||!os.SameFile(st,end)||!os.SameFile(st,current)||end.Size()!=st.Size()||current.Size()!=st.Size()||!end.ModTime().Equal(st.ModTime())||!current.ModTime().Equal(st.ModTime())||!strings.EqualFold(hex.EncodeToString(h.Sum(nil)),expected){return absent,false,errors.New("integridade da cópia não confirmada")}
 if _,e=os.Lstat(original);os.IsNotExist(e){absent=true}else if e==nil{absent=false}else{return false,false,e}
 if !absent{return false,true,errors.New("caminho original ainda presente")}
 return true,true,nil
}

func Record(dir string, o Outcome) (Outcome,error) {
 if e:=os.MkdirAll(dir,0700);e!=nil{return o,e}
 st,e:=os.Lstat(dir);if e!=nil{return o,e};if !st.IsDir()||st.Mode()&os.ModeSymlink!=0{return o,errors.New("diretório de histórico inválido")}
 var id [16]byte;if _,e=rand.Read(id[:]);e!=nil{return o,e}
 o.ID=hex.EncodeToString(id[:]);o.At=time.Now().UTC()
 raw,e:=json.Marshal(o);if e!=nil{return o,e}
 p:=filepath.Join(dir,o.ID+".json")
 f,e:=os.OpenFile(p,os.O_CREATE|os.O_EXCL|os.O_WRONLY,0600);if e!=nil{return o,e}
 n,w:=f.Write(raw);if w==nil&&n!=len(raw){w=io.ErrShortWrite};e=errors.Join(w,f.Sync(),f.Close());if e!=nil{return o,e}
 return o,nil
}

// Bounded history; invalid records are explicitly counted, not presented as successes.
type History struct { Events []Outcome `json:"events"`; Invalid int `json:"invalid"`; Truncated bool `json:"truncated"` }
func ReadHistory(dir string) (History,error) {
 out:=History{Events:[]Outcome{}}
 st,e:=os.Lstat(dir);if os.IsNotExist(e){return out,nil};if e!=nil{return out,e};if !st.IsDir()||st.Mode()&os.ModeSymlink!=0{return out,errors.New("histórico inválido")}
 f,e:=os.Open(dir);if e!=nil{return out,e};defer f.Close()
 entries,e:=f.ReadDir(1001);if e!=nil&&!errors.Is(e,io.EOF){return out,e};if len(entries)>1000{out.Truncated=true;entries=entries[:1000]}
 for _,v:=range entries{if !v.Type().IsRegular()||!strings.HasSuffix(v.Name(),".json"){continue};info,e:=v.Info();if e!=nil||info.Size()>65536{out.Invalid++;continue};p:=filepath.Join(dir,v.Name());h,e:=os.Open(p);if e!=nil{out.Invalid++;continue};raw,e:=io.ReadAll(io.LimitReader(h,65537));ce:=h.Close();var o Outcome;if e!=nil||ce!=nil||len(raw)>65536||json.Unmarshal(raw,&o)!=nil||o.ID+".json"!=v.Name()||len(o.ID)!=32||o.At.IsZero(){out.Invalid++;continue};out.Events=append(out.Events,o)}
 return out,nil
}

func OutcomeForQuarantine(o Outcome, payload string, actionErr error) Outcome {
 o.Action="quarantine"
 if actionErr!=nil {o.Status="failed_or_unconfirmed";o.Error=actionErr.Error();return o}
 a,v,e:=CheckQuarantine(o.Path,payload,o.SHA256);o.OriginalAbsent=a;o.CopyVerified=v
 if e!=nil{o.Status="failed_or_unconfirmed";o.Error=e.Error();return o}
 o.Status="quarantine_verified";o.Executor="ZEUVORYX";return o
}

func (o Outcome) ResultError() error { if o.Status!="quarantine_verified"{return fmt.Errorf("quarentena não confirmada: %s",o.Error)};return nil }
