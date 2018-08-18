package erru_test

import (
	"errors"
	"fmt"

	"github.com/clavoie/erru"
)

func DbCall1() error {
	// do some work
	// err := db.Query() etc
	return nil
}

var expectedError = errors.New("expected error")

func DbCall2() error {
	// do some work
	return expectedError
}

func ExampleMultiplexer() {
	multiplexer := erru.NewMultiplexer()
	multiplexer.Add(DbCall1, DbCall2)

	// splits each call out into a seperate goroutine and
	// blocks until they all complete
	err := multiplexer.Go()
	fmt.Println(err)
	// Output: expected error
}
