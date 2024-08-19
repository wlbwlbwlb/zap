package zap

import (
	"fmt"
	"go.uber.org/zap/internal/stacktrace"
)

func Caller(skip int) string {
	stack := stacktrace.Capture(skip+2, stacktrace.Full)
	defer stack.Free()
	frame, _ := stack.Next()
	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
}
