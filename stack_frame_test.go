package erru

import (
	"fmt"
	"strings"
	"testing"
)

func TestStackFrame(t *testing.T) {
	t.Run("Getters", func(t *testing.T) {
		sf := &stackFrame{
			function:    "function",
			file:        "file",
			line:        10,
			isNamedGoFn: true,
		}

		if sf.Function() != sf.function {
			t.Fatalf("Was expecting %v but found %v", sf.function, sf.Function())
		}
		if sf.File() != sf.file {
			t.Fatalf("Was expecting %v but found %v", sf.file, sf.File())
		}
		if sf.Line() != sf.line {
			t.Fatalf("Was expecting %v but found %v", sf.line, sf.Line())
		}
		if sf.IsNamedGoFn() != sf.isNamedGoFn {
			t.Fatalf("Was expecting %v but found %v", sf.isNamedGoFn, sf.IsNamedGoFn())
		}
	})
	t.Run("FmtFunc", func(t *testing.T) {
		sf := &stackFrame{
			function:    "function",
			isNamedGoFn: true,
		}

		if sf.FmtFunc() != sf.function {
			t.Fatalf("Was expecting %v but found %v", sf.function, sf.FmtFunc())
		}

		sf.function = ""
		if sf.FmtFunc() != sf.function {
			t.Fatalf("Was expecting %v but found %v", sf.function, sf.FmtFunc())
		}

		sf.isNamedGoFn = false
		if strings.HasPrefix(sf.FmtFunc(), "<") == false {
			t.Fatalf("Was expecting a formatted func name but found %v", sf.FmtFunc())
		}
	})
	t.Run("FmtFileLine", func(t *testing.T) {
		sf := &stackFrame{
			file: "file",
			line: 10,
		}

		expected := fmt.Sprintf("%v:%v", sf.file, sf.line)
		if expected != sf.FmtFileLine() {
			t.Fatalf("was expecting %v but found %v", expected, sf.FmtFileLine())
		}

		sf.file = ""
		if strings.HasPrefix(sf.FmtFileLine(), "<") == false {
			t.Fatalf("Was expecting a formatted func name but found %v", sf.FmtFileLine())
		}
	})
}
