package signatureav

import (
	"os/exec"
	"syscall"
	"time"
	"unsafe"
)

func configure(c *exec.Cmd) {
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var processIOCounters = kernel32.NewProc("GetProcessIoCounters")
var processMemoryInfo = kernel32.NewProc("K32GetProcessMemoryInfo")

type ioCounters struct {
	readOps, writeOps, otherOps       uint64
	readBytes, writeBytes, otherBytes uint64
}

type processMemoryCounters struct {
	size, pageFaults                                                 uint32
	peakWorkingSet, workingSet, peakPagedPool, pagedPool             uintptr
	peakNonPagedPool, nonPagedPool, pagefileUsage, peakPagefileUsage uintptr
}

func openProcessSampler(pid int) (func() (Activity, bool), func()) {
	const queryInformation = 0x0400
	handle, err := syscall.OpenProcess(queryInformation, false, uint32(pid))
	if err != nil {
		return nil, nil
	}
	return func() (Activity, bool) {
		var created, exited, kernel, user syscall.Filetime
		var io ioCounters
		if err := syscall.GetProcessTimes(handle, &created, &exited, &kernel, &user); err != nil {
			return Activity{}, false
		}
		ok, _, _ := processIOCounters.Call(uintptr(handle), uintptr(unsafe.Pointer(&io)))
		if ok == 0 {
			return Activity{}, false
		}
		ticks := func(t syscall.Filetime) uint64 { return uint64(t.HighDateTime)<<32 | uint64(t.LowDateTime) }
		result := Activity{
			ObservedAt: time.Now().UTC(), CPU100ns: ticks(kernel) + ticks(user),
			ReadBytes: io.readBytes, WriteBytes: io.writeBytes, OtherBytes: io.otherBytes,
		}
		var mem processMemoryCounters
		mem.size = uint32(unsafe.Sizeof(mem))
		ok, _, _ = processMemoryInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&mem)), uintptr(mem.size))
		if ok != 0 {
			working, peak := uint64(mem.workingSet), uint64(mem.peakWorkingSet)
			result.WorkingSetBytes, result.PeakWorkingBytes = &working, &peak
		}
		return result, true
	}, func() { _ = syscall.CloseHandle(handle) }
}
