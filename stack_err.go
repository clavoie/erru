package erru

import (
	"fmt"
	"runtime"
	"strings"
)

// StackErrSize is the default number of stack frames
// to pull back when creating a new StackErr
const StackErrSize = 10

// StackErr represents an error with a stack trace. StackErr
// implements the error interface
type StackErr struct {
	// Err is the original error
	Err error

	// Msg is an additional user message adding context to which
	// the error was raised
	Msg string

	// Frames contains the stack trace leading up to the
	// point that the err was raised
	Frames []*StackFrame
}

// Wrap creates a new err with a stack trace
// from an existing err. If err is nil then nil is returned,
// otherwise an err is returned containing the original
// err and the stack trace leading up to the point that the
// err was raised.
func Wrap(err error) *StackErr { return wrapInternal(err, StackErrSize, "") }

// WrapF acts exactly like Wrap except you can specify an additional context
// message to give context to the error
func WrapF(err error, format string, fmtArgs ...interface{}) *StackErr {
	return wrapInternal(err, StackErrSize, format, fmtArgs...)
}

// WrapN acts exactly like Wrap except you can specify how many
// stack frames display in the stack trace
func WrapN(err error, stackSize int) *StackErr { return wrapInternal(err, stackSize, "") }

// WrapNf acts exactly like Wrap except you can specify how many
// stack frames display in the stack trace, and an additional context message
func WrapNf(err error, stackSize int, format string, fmtArgs ...interface{}) *StackErr {
	return wrapInternal(err, stackSize, format, fmtArgs...)
}

// wrapInternal is a common implementation of of wrapping an err called
// by higher level functions
func wrapInternal(err error, stackSize int, format string, fmtArgs ...interface{}) *StackErr {
	if err == nil {
		return nil
	}

	stackErr, isStackErr := err.(*StackErr)
	if isStackErr && format == "" {
		return stackErr
	}

	if stackSize < 1 {
		stackSize = 1
	}

	pcs := make([]uintptr, stackSize)
	pcsWritten := runtime.Callers(3, pcs)
	pcs = pcs[:pcsWritten]
	runtimeFrames := runtime.CallersFrames(pcs)
	stackFrames := make([]*StackFrame, 0, pcsWritten)

	for {
		runtimeFrame, more := runtimeFrames.Next()
		stackFrames = append(stackFrames, &StackFrame{
			Function:    runtimeFrame.Function,
			File:        runtimeFrame.File,
			Line:        runtimeFrame.Line,
			IsNamedGoFn: runtimeFrame.Func != nil,
		})

		if more == false {
			break
		}
	}

	return &StackErr{
		Err:    err,
		Msg:    fmt.Sprintf(format, formatArgs...),
		Frames: stackFrames,
	}
}

// Error returns the original error message with a stack trace
func (se *StackErr) Error() string {
	messageSpaces := 2

	if se.Msg != "" {
		messageSpaces += 2
	}

	msg := make([]string, (len(se.Frames)*2)+messageSpaces)
	baseIndex := 0

	if se.Msg != "" {
		msg[0] = se.Msg
		msg[1] = ""
		baseIndex = 2
	}

	msg[baseIndex] = se.Err.Error()
	msg[baseIndex+1] = ""
	baseIndex += 2

	for index, frame := range se.Frames {
		msg[index+baseIndex] = frame.FmtFunc()
		msg[index+baseIndex+1] = frame.FmtFileLine()
	}

	return strings.Join(msg, "\n")
}
