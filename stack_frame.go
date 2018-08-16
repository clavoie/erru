package erru

import "fmt"

// StackFrame represents the details of one frame of
// a stack trace
type StackFrame interface {
	// Function is the package path-qualified function name of
	// this call frame. If non-empty, this string uniquely
	// identifies a single function in the program.
	// This may be the empty string if not known.
	Function() string

	// File and Line are the file name and line number of the
	// location in this frame. For non-leaf frames, this will be
	// the location of a call. These may be the empty string and
	// zero, respectively, if not known.
	File() string
	Line() int
	// IsNamedGoFn indicates if the function in this frame is likely
	// to be a named go func, or if it is potentially inlined / external
	IsNamedGoFn() bool

	// FmtFunc returns a human readable version of the Function property,
	// taking into account external or inlined code
	FmtFunc() string
	// FmtFileLine returns a human readable version of the File:Line
	// properties, taking into account external or inlined code.
	FmtFileLine() string
}

// stackFrame is an implementation of StackFrame
type stackFrame struct {
	function    string
	file        string
	line        int
	isNamedGoFn bool
}

func (sf *stackFrame) Function() string  { return sf.function }
func (sf *stackFrame) File() string      { return sf.file }
func (sf *stackFrame) Line() int         { return sf.line }
func (sf *stackFrame) IsNamedGoFn() bool { return sf.isNamedGoFn }

func (sf *stackFrame) FmtFunc() string {
	funcName := sf.function

	if funcName == "" && sf.isNamedGoFn == false {
		funcName = "<inlined or external code>"
	}

	return funcName
}

func (sf *stackFrame) FmtFileLine() string {
	if sf.file == "" {
		return "<unknown>"
	}

	return fmt.Sprintf("%s:%d", sf.file, sf.line)
}
