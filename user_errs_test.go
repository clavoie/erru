package erru_test

import (
	"testing"

	"github.com/clavoie/erru"
)

func TestUserErrs(t *testing.T) {
	t.Run("ErrorEmpty", func(t *testing.T) {
		userErr := erru.NewUserErrs()
		errStr := userErr.Error()

		if errStr != "" {
			t.Fatalf("Was expecting the empty string but instead found: %v", errStr)
		}
	})
	t.Run("ErrorFirst", func(t *testing.T) {
		err1 := "error1"
		err2 := "error2"
		userErr := erru.NewUserErrs(err1, err2)
		errStr := userErr.Error()

		if errStr != err1 {
			t.Fatalf("Was expecting %v but instead found: %v", err1, errStr)
		}
	})
	t.Run("User", func(t *testing.T) {
		err1 := "error1"
		err2 := "error2"
		userErr := erru.NewUserErrs(err1, err2)
		userErrs := userErr.User()

		if userErrs[0] != err1 {
			t.Fatalf("Was expecting %v but instead found: %v", err1, userErrs[0])
		}
		if userErrs[1] != err2 {
			t.Fatalf("Was expecting %v but instead found: %v", err2, userErrs[1])
		}
	})
}
