//go:build !windows

package signatureav

import "os/exec"

func configure(c *exec.Cmd) {}

func openProcessSampler(int) (func() (Activity, bool), func()) { return nil, nil }
