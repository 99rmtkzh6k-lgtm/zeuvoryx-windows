package signatureav
import("context";"strings")
type observerKey314 struct{}
func WithObserver(ctx context.Context,fn func(Finding))context.Context{return context.WithValue(ctx,observerKey314{},fn)}
type lineObserver314 struct{fn func(Finding);pending []byte;discard bool}
func(o *lineObserver314) feed(p []byte){
 if o.fn==nil{return}
 for _,b:=range p{
  if b=='\n'{if !o.discard{o.emit(string(o.pending))};o.pending=o.pending[:0];o.discard=false;continue}
  if o.discard{continue};if len(o.pending)>=65536{o.discard=true;o.pending=o.pending[:0];continue};o.pending=append(o.pending,b)
 }
}
func(o *lineObserver314) emit(line string){
 line=strings.TrimSpace(line);if !strings.HasSuffix(line," FOUND"){return}
 v:=strings.TrimSuffix(line," FOUND");i:=strings.LastIndex(v,": ");if i<=0||i+2>=len(v){return}
 label:=v[i+2:];lower:=strings.ToLower(label);kind:="signature_detection"
 if strings.Contains(lower,"eicar"){kind="test_eicar"}else if strings.Contains(lower,"heuristics.")||strings.Contains(lower,"encrypted"){kind="heuristic_or_limit"}
 o.fn(Finding{Path:v[:i],Signature:label,Kind:kind})
}
