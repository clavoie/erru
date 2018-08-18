package erru_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/clavoie/erru"
)

func TestImpl(t *testing.T) {
	i := erru.NewImpl()
	err := errors.New("error")
	format := "format %v"
	fmtArgs := []interface{}{1}
	expectedFormat := fmt.Sprintf(format, fmtArgs...)
	expect := func(t *testing.T, actual, expected interface{}) {
		if actual != expected {
			t.Fatal(actual, expected)
		}
	}
	expectHttp := func(t *testing.T, statusCode int, httpErr erru.HttpErr) {
		if statusCode != httpErr.StatusCode() {
			t.Fatal(statusCode, httpErr.StatusCode())
		}

		if httpErr.Err().Error() != expectedFormat {
			t.Fatal(httpErr.Err().Error(), expectedFormat)
		}
	}

	t.Run("First", func(t *testing.T) {
		expect(t, i.First(nil, err), err)
	})

	t.Run("HttpBadRequest", func(t *testing.T) {
		httpErr := i.NewHttpBadRequest(format, fmtArgs...)
		expectHttp(t, http.StatusBadRequest, httpErr)
	})
	t.Run("HttpError", func(t *testing.T) {
		httpErr := i.NewHttpError(http.StatusTeapot, format, fmtArgs...)
		expectHttp(t, http.StatusTeapot, httpErr)
	})
	t.Run("HttpForbidden", func(t *testing.T) {
		httpErr := i.NewHttpForbidden(format, fmtArgs...)
		expectHttp(t, http.StatusForbidden, httpErr)
	})
	t.Run("HttpInternalServerErr", func(t *testing.T) {
		httpErr := i.NewHttpInternalServerErr(format, fmtArgs...)
		expectHttp(t, http.StatusInternalServerError, httpErr)
	})
	t.Run("HttpNotFound", func(t *testing.T) {
		httpErr := i.NewHttpNotFound(format, fmtArgs...)
		expectHttp(t, http.StatusNotFound, httpErr)
	})
	t.Run("HttpUnauthorized", func(t *testing.T) {
		httpErr := i.NewHttpUnauthorized(format, fmtArgs...)
		expectHttp(t, http.StatusUnauthorized, httpErr)
	})

	t.Run("Errorf", func(t *testing.T) {
		stackErr := i.Errorf(format, fmtArgs...)
		expect(t, stackErr.Err().Error(), expectedFormat)
	})
	t.Run("Wrap", func(t *testing.T) {
		stackErr := i.Wrap(err)
		expect(t, stackErr.Err(), err)
	})
	t.Run("WrapF", func(t *testing.T) {
		stackErr := i.WrapF(err, format, fmtArgs...)
		expect(t, stackErr.Err(), err)
		expect(t, stackErr.Msg(), expectedFormat)
	})
	t.Run("WrapN", func(t *testing.T) {
		stackErr := i.WrapN(err, 2)
		expect(t, stackErr.Err(), err)
		expect(t, len(stackErr.Frames()), 2)
	})
	t.Run("WrapNf", func(t *testing.T) {
		stackErr := i.WrapNf(err, 2, format, fmtArgs...)
		expect(t, stackErr.Err(), err)
		expect(t, stackErr.Msg(), expectedFormat)
		expect(t, len(stackErr.Frames()), 2)
	})

	t.Run("NewUserErrs", func(t *testing.T) {
		userErrs := i.NewUserErrs("a", "b")
		expect(t, userErrs.User()[0], "a")
		expect(t, userErrs.User()[1], "b")
	})
}
