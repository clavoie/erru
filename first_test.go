package erru_test

import (
	"errors"
	"testing"

	"github.com/clavoie/erru"
)

func TestFirst(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		err := erru.First(nil, nil, nil)

		if err != nil {
			t.Fatalf("Was expecting nil, instead got: %v", err)
		}
	})
	t.Run("Not nil", func(t *testing.T) {
		err1 := errors.New("1")
		err2 := errors.New("2")

		err := erru.First(nil, err1, err2)
		if err != err1 {
			t.Fatalf("Was expecting %v, instead got: %v", err1, err)
		}

		err = erru.First(nil, err2, err1)
		if err != err2 {
			t.Fatalf("Was expecting %v, instead got: %v", err2, err)
		}
	})
}
