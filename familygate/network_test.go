package familygate
import("context";"errors";"net/netip";"testing")
func TestDNSCheck314(t *testing.T){
 good:=func(ctx context.Context,network,host string)([]netip.Addr,error){if network!="ip4"{t.Fatal(network)};if host=="example.com"{return []netip.Addr{netip.MustParseAddr("93.184.216.34")},nil};return []netip.Addr{netip.IPv4Unspecified()},nil}
 if r:=CheckDNS(context.Background(),good);r.Status!="test_passed"{t.Fatal(r)}
 failed:=func(context.Context,string,string)([]netip.Addr,error){return nil,errors.New("offline")}
 if r:=CheckDNS(context.Background(),failed);r.Status!="inconclusive"{t.Fatal(r)}
 unfiltered:=func(context.Context,string,string)([]netip.Addr,error){return []netip.Addr{netip.MustParseAddr("93.184.216.34")},nil}
 if r:=CheckDNS(context.Background(),unfiltered);r.Status!="test_not_blocked"{t.Fatal(r)}
 mixed:=func(ctx context.Context,n,h string)([]netip.Addr,error){if h=="example.com"{return good(ctx,n,h)};return []netip.Addr{netip.IPv4Unspecified(),netip.MustParseAddr("93.184.216.34")},nil}
 if r:=CheckDNS(context.Background(),mixed);r.Status!="inconclusive"{t.Fatal(r)}
}
