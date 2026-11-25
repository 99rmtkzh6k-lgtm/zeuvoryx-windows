//go:build !windows

package signatureav
import "os/exec"
func configure(c *exec.Cmd){}
