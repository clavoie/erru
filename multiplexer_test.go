package erru_test

import (
	"errors"
	"testing"

	"github.com/clavoie/erru"
)

func TestMultiplexer(t *testing.T) {
	t.Run("AddNilNoErr", func(t *testing.T) {
		multiplexer := erru.NewMultiplexer()
		multiplexer.Add(nil)
		err := multiplexer.Go()

		if err != nil {
			t.Fatalf("Was expecting nil err but instead got: %v", err)
		}
	})
	t.Run("GoNoFns", func(t *testing.T) {
		multiplexer := erru.NewMultiplexer()
		err := multiplexer.Go()

		if err != nil {
			t.Fatalf("Was expecting nil err but instead got: %v", err)
		}
	})
	t.Run("GoOneFn", func(t *testing.T) {
		multiplexer := erru.NewMultiplexer()
		expectedErr := errors.New("expected err")
		multiplexer.Add(func(c chan error) { c <- expectedErr })
		err := multiplexer.Go()

		if err != expectedErr {
			t.Fatalf("Was expecting %v err but instead got: %v", expectedErr, err)
		}
	})
	t.Run("GoManyErrFn", func(t *testing.T) {
		multiplexer := erru.NewMultiplexer()
		expectedErr1 := errors.New("expected err 1")
		expectedErr2 := errors.New("expected err 2")

		multiplexer.Add(func(c chan error) { c <- nil })
		multiplexer.Add(func(c chan error) { c <- expectedErr2 })
		multiplexer.Add(func(c chan error) { c <- expectedErr1 })

		err := multiplexer.Go()

		if err != expectedErr2 && err != expectedErr1 {
			t.Fatalf("Was expecting a known err err but instead got: %v", err)
		}
	})
}
