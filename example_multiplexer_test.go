package erru_test

import (
	"errors"
	"fmt"

	"github.com/clavoie/erru"
)

func DbCall1(errChan chan error) {
	// do some work
	errChan <- nil
}

var expectedError = errors.New("expected error")

func DbCall2(errChan chan error) {
	// do some work
	errChan <- expectedError
}

func ExampleMultiplexer() {
	multiplexer := erru.NewMultiplexer()
	multiplexer.Add(DbCall1)
	multiplexer.Add(DbCall2)

	// splits each call out into a seperate goroutine and
	// blocks until they all complete
	err := multiplexer.Go()
	fmt.Println(err)
	// Output: expected error
}
