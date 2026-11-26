// Package scanplan chooses conservative scheduling, not a malware verdict.
package scanplan

import "time"
const GiB uint64=1<<30
// Serial avoids concurrent heavyweight scanners on resource-constrained hosts.
// Unknown or inconsistent measurements use the conservative path.
func Serial(known bool,total,available uint64)bool{
 return !known||total==0||available>total||total<=6*GiB||available<3*GiB
}

// Workers caps the scanner's file-reading concurrency using the measured RAM
// and available logical CPUs. Unknown measurements never choose the fast path.
func Workers(cpus int, known bool,total,available uint64)int{
 n:=1
 if known&&total>0&&available<=total&&available>=GiB/2{
  switch{case total<=2*GiB:n=1
  case total<=4*GiB||available<2*GiB:n=2
  case total<=8*GiB||available<4*GiB:n=3
  default:n=4}
 }
 if cpus>0&&cpus<n{n=cpus}
 return n
}
// Standard ClamAV database has a documented RAM recommendation of 3 GiB.
// Below this physical capacity, keep the own scanner usable and report the
// optional signature stage as unavailable instead of falsely claiming coverage.
func ClamCapacity(known bool,total uint64)bool{return !known||total>=3*GiB}

// SignatureBudget limits the independent signature stage without cancelling the own scanner.
func SignatureBudget(mode string) time.Duration {
	switch mode {
	case "quick":
		return 5 * time.Minute
	case "full":
		return time.Hour
	default:
		return 20 * time.Minute
	}
}
