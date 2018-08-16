package erru_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/clavoie/erru"
)

func TestHttpErr(t *testing.T) {
	testFmt := "test: %v"
	testArgs := []interface{}{1}
	expectedMsg := fmt.Sprintf(testFmt, testArgs...)

	testStatusCodeHelper := func(t *testing.T, statusCode int, constructor func(string, ...interface{}) erru.HttpErr) {
		err := constructor(testFmt, testArgs...)

		if err == nil {
			t.Fatal("Was expecting non-nil err")
		}

		if err.StatusCode() != statusCode {
			t.Fatalf("Was expecting %v, but instead got %v", statusCode, err.StatusCode())
		}

		if err.Err().Error() != expectedMsg {
			t.Fatalf("Was expecting %v, but instead got %v", expectedMsg, err.Err().Error())
		}
	}
	t.Run("BadRequest", func(t *testing.T) {
		testStatusCodeHelper(t, http.StatusBadRequest, erru.NewHttpBadRequest)
	})
	t.Run("Forbidden", func(t *testing.T) {
		testStatusCodeHelper(t, http.StatusForbidden, erru.NewHttpForbidden)
	})
	t.Run("InternalServerError", func(t *testing.T) {
		testStatusCodeHelper(t, http.StatusInternalServerError, erru.NewHttpInternalServerErr)
	})
	t.Run("NotFound", func(t *testing.T) {
		testStatusCodeHelper(t, http.StatusNotFound, erru.NewHttpNotFound)
	})
	t.Run("Unauthorized", func(t *testing.T) {
		testStatusCodeHelper(t, http.StatusUnauthorized, erru.NewHttpUnauthorized)
	})
}
