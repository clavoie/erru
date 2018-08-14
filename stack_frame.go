package erru

import "fmt"

// StackFrame represents the details of one frame of
// a stack trace
type StackFrame struct {
	// Function is the package path-qualified function name of
	// this call frame. If non-empty, this string uniquely
	// identifies a single function in the program.
	// This may be the empty string if not known.
	Function string

	// File and Line are the file name and line number of the
	// location in this frame. For non-leaf frames, this will be
	// the location of a call. These may be the empty string and
	// zero, respectively, if not known.
	File string
	Line int

	// IsNamedGoFn indicates if the function in this frame is likely
	// to be a named go func, or if it is potentially inlined / external
	IsNamedGoFn bool
}

// FmtFunc returns a human readable version of the Function property,
// taking into account external or inlined code
func (sf *StackFrame) FmtFunc() string {
	funcName := sf.Function

	if funcName == "" && sf.IsNamedGoFn == false {
		funcName = "<inlined or external code>"
	}

	return funcName
}

// FmtFileLine returns a human readable version of the File:Line
// properties, taking into account external or inlined code.
func (sf *StackFrame) FmtFileLine() string {
	if sf.File == "" {
		return "<unknown>"
	}

	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
}
