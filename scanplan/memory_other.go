//go:build !windows

package scanplan
func Memory()(bool,uint64,uint64){return false,0,0}
