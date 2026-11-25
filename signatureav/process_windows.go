package signatureav
import ("os/exec";"syscall")
func configure(c *exec.Cmd){c.SysProcAttr=&syscall.SysProcAttr{HideWindow:true}}
