package adaptive
import "testing"
func TestReviewAlert314(t *testing.T){
 if ReviewAlert(Vector{1,1,0,0,0},.99)||!ReviewAlert(Vector{0,1,.5,.5,0},.85)||ReviewAlert(Vector{0,1,.5,.5,0},.5){t.Fatal("review gate")}
}
