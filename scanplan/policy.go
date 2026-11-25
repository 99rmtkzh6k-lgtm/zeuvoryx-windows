// Package scanplan chooses conservative scheduling, not a malware verdict.
package scanplan
const GiB uint64=1<<30
// Serial avoids concurrent heavyweight scanners on resource-constrained hosts.
// Unknown or inconsistent measurements use the conservative path.
func Serial(known bool,total,available uint64)bool{
 return !known||total==0||available>total||total<=6*GiB||available<3*GiB
}
