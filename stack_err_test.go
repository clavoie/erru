package erru_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/clavoie/erru"
)

func testNamedFunc() error {
	innerErr := func() error {
		return func() error {
			return erru.WrapF(errors.New("inner error"), "inner error format")
		}()
	}()
	return erru.WrapF(innerErr, "outer error format")
}

func TestStackErr(t *testing.T) {
	t.Run("Errorf", func(t *testing.T) {
		format := "test err %v"
		args := []interface{}{1}
		expected := fmt.Sprintf(format, args...)
		stackErr := erru.Errorf(format, args...)

		if stackErr.Err().Error() != expected {
			t.Fatalf("Excpected %v but found %v", expected, stackErr.Err().Error())
		}

		if stackErr.Msg() != "" {
			t.Fatal("Expected found non-empty str")
		}

		if len(stackErr.Frames()) < 1 {
			t.Fatal("Expecting stack frames")
		}
	})
	t.Run("Wrap", func(t *testing.T) {
		err := errors.New("error")
		stackErr := erru.Wrap(err)

		if stackErr.Err() != err {
			t.Fatalf("Excpected %v but found %v", err, stackErr.Err())
		}

		if stackErr.Msg() != "" {
			t.Fatal("Expected found non-empty str")
		}

		if len(stackErr.Frames()) < 1 {
			t.Fatal("Expecting stack frames")
		}
	})
	t.Run("WrapF", func(t *testing.T) {
		err := errors.New("error")
		format := "test err %v"
		args := []interface{}{1}
		expected := fmt.Sprintf(format, args...)
		stackErr := erru.WrapF(err, format, args...)

		if stackErr.Err() != err {
			t.Fatalf("Excpected %v but found %v", err, stackErr.Err())
		}

		if stackErr.Msg() != expected {
			t.Fatalf("Expected %v found %v", expected, stackErr.Msg())
		}

		if len(stackErr.Frames()) < 1 {
			t.Fatal("Expecting stack frames")
		}
	})
	t.Run("WrapN", func(t *testing.T) {
		err := errors.New("error")
		expected := 2
		stackErr := erru.WrapN(err, 2)

		if stackErr.Err() != err {
			t.Fatalf("Excpected %v but found %v", err, stackErr.Err())
		}

		if stackErr.Msg() != "" {
			t.Fatal("Expected found non-empty str")
		}

		if len(stackErr.Frames()) != expected {
			t.Fatalf("Expecting %v stack frames but found %v", expected, len(stackErr.Frames()))
		}

		// min stack size
		stackErr = erru.WrapN(err, 0)
		expected = 1
		if len(stackErr.Frames()) != expected {
			t.Fatalf("Expecting %v stack frames but found %v", expected, len(stackErr.Frames()))
		}
	})
	t.Run("WrapNf", func(t *testing.T) {
		err := errors.New("error")
		format := "test err %v"
		args := []interface{}{1}
		expected := fmt.Sprintf(format, args...)
		stack := 2
		stackErr := erru.WrapNf(err, stack, format, args...)

		if stackErr.Err() != err {
			t.Fatalf("Excpected %v but found %v", err, stackErr.Err())
		}

		if stackErr.Msg() != expected {
			t.Fatalf("Expected %v found %v", expected, stackErr.Msg())
		}

		if len(stackErr.Frames()) != stack {
			t.Fatalf("Expecting %v stack frames but found %v", stack, len(stackErr.Frames()))
		}
	})
	t.Run("WrapNil", func(t *testing.T) {
		stackErr := erru.Wrap(nil)

		if stackErr != nil {
			t.Fatalf("Was expecting nil but found: %v", stackErr)
		}
	})
	t.Run("WrapExisting", func(t *testing.T) {
		err := errors.New("error")
		stackErr1 := erru.Wrap(err)
		stackErr2 := erru.Wrap(stackErr1)

		if stackErr2 != stackErr1 {
			t.Fatalf("Was expecting %v but found: %v", stackErr1, stackErr2)
		}

		stackErr2 = erru.WrapF(stackErr1, "additional context")

		if stackErr2 == stackErr1 {
			t.Fatalf("Was expecting not %v but found: %v", stackErr1, stackErr2)
		}
	})
	t.Run("Error", func(t *testing.T) {
		t.Log("\n" + testNamedFunc().Error())
	})
}
