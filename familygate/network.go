package familygate
import("context";"net";"net/netip";"time")
type DNSCheck struct {
 CheckedAt time.Time `json:"checked_at"`
 Connection string `json:"connection"`
 Adult string `json:"adult"`
 Malware string `json:"malware"`
 Status string `json:"status"`
}
type Lookup func(context.Context,string,string)([]netip.Addr,error)
func sample(ctx context.Context,lookup Lookup,host string)string{
 q,cancel:=context.WithTimeout(ctx,5*time.Second);defer cancel()
 ips,e:=lookup(q,"ip4",host)
 if e!=nil||len(ips)==0{return "inconclusive"}
 zero,other:=false,false
 for _,ip:=range ips{if ip.Unmap()==netip.IPv4Unspecified(){zero=true}else{other=true}}
 if zero&&!other{return "blocked"}
 if other&&!zero{return "resolved"}
 return "inconclusive"
}
// Test the system DNS, never a forced external resolver. Only harmless published test domains.
// A positive control avoids treating a network outage as proof of filtering.
func CheckDNS(ctx context.Context,lookup Lookup)DNSCheck{
 r:=DNSCheck{Connection:sample(ctx,lookup,"example.com"),Adult:sample(ctx,lookup,"nudity.testcategory.com"),Malware:sample(ctx,lookup,"malware.testcategory.com"),Status:"inconclusive"}
 r.CheckedAt=time.Now().UTC()
 if r.Connection=="resolved"{
 if r.Adult=="blocked"&&r.Malware=="blocked"{r.Status="test_passed"}else if r.Adult=="resolved"||r.Malware=="resolved"{r.Status="test_not_blocked"}
 }
 return r
}
func CheckSystemDNS(ctx context.Context)DNSCheck{return CheckDNS(ctx,net.DefaultResolver.LookupNetIP)}
