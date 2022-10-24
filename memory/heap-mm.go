package memory

import (
	"github.com/junyupL/sys/windows"
)

var HEAP windows.Handle

func HInit() {
	HEAP, _ = windows.GetProcessHeap()
}
