//go:build windows

package main
import("testing";"time")
func Test308PerformanceMath(t *testing.T){
 a:=pTicks{100,200,100};b:=pTicks{140,260,140}
 if v:=pCPU(a,b);v==nil||*v!=60{t.Fatal("CPU math")}
 for _,c:=range []pTicks{a,{99,260,140},{200,201,101}}{if pCPU(a,c)!=nil{t.Fatal("invalid CPU accepted")}}
 if v:=pPercent(16,4);v==nil||*v!=75{t.Fatal("RAM math")}
 if pPercent(0,0)!=nil||pPercent(1,2)!=nil{t.Fatal("invalid memory")}
}
func Test308PerformanceNative(t *testing.T){
 var p pSampler;a:=p.sample()
 if a["ram_total"].(uint64)==0||a["ram_percent"]==nil||a["ram_available"].(uint64)>a["ram_total"].(uint64){t.Fatal("native RAM unavailable")}
 if a["cpu"]!=nil{t.Fatal("first CPU sample must be unknown")}
 if b:=p.sample();b["at"]!=a["at"]{t.Fatal("cache ignored")}
 time.Sleep(20*time.Millisecond);p.mu.Lock();p.last=time.Time{};p.mu.Unlock();b:=p.sample()
 if v,ok:=b["cpu"].(float64);ok&&(v<0||v>100){t.Fatal("invalid CPU")}
 if b["disk_percent"]!=nil&&(b["disk_total"].(uint64)==0||b["disk_free"].(uint64)>b["disk_total"].(uint64)){t.Fatal("invalid disk")}
}
