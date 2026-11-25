package adaptive
func ReviewAlert(x Vector,score float64)bool{
 if !valid(x)||score<0.8||score>1{return false}
 n:=0;for _,i:=range []int{0,2,3,4}{if x[i]>=0.5{n++}}
 return n>=2 && (x[2]>=0.5||x[3]>=0.5||x[4]>=0.5)
}
