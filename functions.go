package cerror

import (
	"fmt"
	"runtime"
	"strings"
)

func getStackTraces() string {
	var traces strings.Builder
	rpc := make([]uintptr, 10)
	n := runtime.Callers(4, rpc[:])
	if n <= 0 {
		return traces.String()
	}

	frames := runtime.CallersFrames(rpc)
	for i := 0; i < n; i++ {
		frame, _ := frames.Next()
		if frame.Func != nil && !strings.Contains(frame.File, "go/src/runtime") {
			traces.WriteString(fmt.Sprintf("[%s] %s:%d\n", frame.Function, frame.File, frame.Line))
		}
	}
	return traces.String()
}
