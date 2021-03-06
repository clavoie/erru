package erru

import (
	"fmt"
	"runtime"
	"strings"
)

// StackErrSize is the default number of stack frames
// to pull back when creating a new StackErr
const StackErrSize = 10

// StackErr represents an error that contains a stack trace.
type StackErr interface {
	error

	// Err is the original error
	Err() error

	// Msg is an additional user message adding context in which
	// the error was raised
	Msg() string

	// Frames contains the stack trace leading up to the
	// point that the err was raised
	Frames() []StackFrame
}

// stackErr is an implementation of StackErr
type stackErr struct {
	err    error
	msg    string
	frames []StackFrame
}

func (se *stackErr) Err() error           { return se.err }
func (se *stackErr) Msg() string          { return se.msg }
func (se *stackErr) Frames() []StackFrame { return se.frames }

// Errorf is like fmt.Errorf except that the returned error has a stack trace
func Errorf(format string, fmtArgs ...interface{}) StackErr {
	return wrapInternal(fmt.Errorf(format, fmtArgs...), StackErrSize, nil, "")
}

func errorF(toFill *stackErr, format string, fmtArgs ...interface{}) {
	wrapInternal(fmt.Errorf(format, fmtArgs...), StackErrSize, toFill, "")
}

// Wrap creates a new err with a stack trace
// from an existing err. If err is nil then nil is returned,
// otherwise an err is returned containing the original
// err and the stack trace leading up to the point that the
// err was raised. If err is already a StackErr with no additional context
// then that err is returned
func Wrap(err error) StackErr { return wrapInternal(err, StackErrSize, nil, "") }

// WrapF acts exactly like Wrap except you can specify an additional context
// message to give context to the error
func WrapF(err error, format string, fmtArgs ...interface{}) StackErr {
	return wrapInternal(err, StackErrSize, nil, format, fmtArgs...)
}

// WrapN acts exactly like Wrap except you can specify how many
// stack frames display in the stack trace
func WrapN(err error, stackSize int) StackErr { return wrapInternal(err, stackSize, nil, "") }

// WrapNf acts exactly like Wrap except you can specify how many
// stack frames display in the stack trace, and an additional context message
func WrapNf(err error, stackSize int, format string, fmtArgs ...interface{}) StackErr {
	return wrapInternal(err, stackSize, nil, format, fmtArgs...)
}

// wrapInternal is a common implementation of of wrapping an err called
// by higher level functions
func wrapInternal(err error, stackSize int, toFill *stackErr, format string, fmtArgs ...interface{}) StackErr {
	if err == nil {
		return nil
	}

	stackErrInstance, isStackErr := err.(*stackErr)
	if isStackErr && format == "" {
		return stackErrInstance
	}

	if stackSize < 1 {
		stackSize = 1
	}

	pcs := make([]uintptr, stackSize)
	pcsWritten := runtime.Callers(3, pcs)
	pcs = pcs[:pcsWritten]
	runtimeFrames := runtime.CallersFrames(pcs)
	stackFrames := make([]StackFrame, 0, pcsWritten)

	for {
		runtimeFrame, more := runtimeFrames.Next()
		stackFrames = append(stackFrames, &stackFrame{
			function:    runtimeFrame.Function,
			file:        runtimeFrame.File,
			line:        runtimeFrame.Line,
			isNamedGoFn: runtimeFrame.Func != nil,
		})

		if more == false {
			break
		}
	}

	msg := fmt.Sprintf(format, fmtArgs...)
	if toFill == nil {
		return &stackErr{
			err:    err,
			msg:    msg,
			frames: stackFrames,
		}
	}

	toFill.err = err
	toFill.msg = msg
	toFill.frames = stackFrames
	return toFill
}

// Error returns the original error message with a stack trace
func (se *stackErr) Error() string {
	return se.errorInternal(0)
}

func (se *stackErr) errorInternal(indentLevel int) string {
	msg := make([]string, 0, (len(se.frames)*2)+4)
	indent := strings.Repeat(" ", indentLevel)
	add := func(str string) { msg = append(msg, indent+str) }

	if se.msg != "" {
		add(se.msg)
	}

	innerStackErr, isStackErr := se.err.(*stackErr)
	if isStackErr {
		add(innerStackErr.errorInternal(indentLevel + 2))
	}

	for _, frame := range se.frames {
		add(frame.FmtFunc())
		add("  " + frame.FmtFileLine())
	}

	return strings.Join(msg, "\n")
}
