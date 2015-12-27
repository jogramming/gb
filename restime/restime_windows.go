package restime

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	dll            = syscall.MustLoadDLL("kernel32.dll")
	queryCounter   = dll.MustFindProc("QueryPerformanceCounter")
	queryFrequency = dll.MustFindProc("QueryPerformanceFrequency")

	overhead Delta = 0
)

type Counter int64
type Delta int64
type Frequency int64

func (a Counter) Sub(b Counter) Delta {
	return Delta(a-b) - overhead
}

func Since(c Counter) time.Duration {
	return Now().Sub(c).Duration()
}

func (d Delta) Duration() time.Duration {
	return d.DurationFreq(Freq())
}
func (d Delta) DurationFreq(freq Frequency) time.Duration {
	return time.Duration((float64(d) / float64(freq)) * float64(time.Second))
}

func Now() (count Counter) {
	queryCounter.Call(uintptr(unsafe.Pointer(&count)))
	return
}

func Freq() (freq Frequency) {
	queryFrequency.Call(uintptr(unsafe.Pointer(&freq)))
	return
}

func Overhead() Delta {
	return overhead
}

func init() {
	start := Now()
	stop := Now()
	overhead = Delta(stop - start)
}
