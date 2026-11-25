package scanplan
import "log"
func CurrentWorkers(cpus int)int{
 known,total,available:=Memory()
 n:=Workers(cpus,known,total,available)
 log.Printf("scanplan workers=%d logical_cpu=%d ram_known=%t ram_total=%d ram_available=%d",n,cpus,known,total,available)
 return n
}
