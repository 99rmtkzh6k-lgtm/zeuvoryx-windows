package portable
import("os";"strconv";"testing";"strings")
var fraudCases=[]struct{text string;want bool}{
 {"ENVIE SUA FRASE DE RECUPERAÇÃO",true},{"Retorno garantido sem risco",true},{"Pague taxa de desbloqueio",true},{"Urgente confirme seus dados",true},{"Instale AnyDesk",true},{"Conecte sua carteira",true},{"Envie seu CPF",true},{"Your account suspended",true},{"Reunião confirmada amanhã",false},{"Relatório de despesas disponível",false},
}
func TestFraudCases(t *testing.T){for _,v:=range fraudCases{r,e:=assessFraud(v.text);if e!=nil||(len(r.Signals)>0)!=v.want{t.Fatalf("case %q: %+v %v",v.text,r,e)}}}
func TestFraudInputLimits(t *testing.T){for _,v:=range []string{"",strings.Repeat("a",16385)}{if _,e:=assessFraud(v);e==nil{t.Fatal("invalid input accepted")}};r,e:=assessFraud("CÓDIGO\tSMS");if e!=nil||len(r.Signals)!=1{t.Fatal(r,e)}}
func TestIncidentPlans(t *testing.T){for _,v:=range []string{"crypto","identity","malware","ransomware"}{r,e:=incidentPlan(v);if e!=nil||len(r)<4{t.Fatal(v,r,e)}};if _,e:=incidentPlan("arbitrary");e==nil{t.Fatal("unknown plan")}}
func TestSyntheticFraudMillion(t *testing.T){if os.Getenv("ZEUVORYX_STRESS")!="1"{t.Skip("opt-in synthetic workload")};for i:=0;i<1000000;i++{v:=fraudCases[i%len(fraudCases)];r,e:=assessFraud(v.text+" ref "+strconv.Itoa(i));if e!=nil||(len(r.Signals)>0)!=v.want{t.Fatalf("case %d",i)}};t.Log("1000000 synthetic message evaluations; 800000 marked, 200000 unmarked; not live attacks")}
