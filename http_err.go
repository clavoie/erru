package erru

import (
	"net/http"
)

// HttpErr is an error implementation that returns an http status code
type HttpErr interface {
	StackErr

	// StatusCode returns the HTTP status code that should be written as the header of the
	// response writer
	StatusCode() int
}

// httpErr is an implementation of HttpErr
type httpErr struct {
	stackErr
	statusCode int
}

// NewHttpBadRequest is a convenience function for creating a new HttpErr that responds
// with status code 400
func NewHttpBadRequest(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpError(http.StatusBadRequest, format, fmtArgs)
}

// NewHttpUnauthorized is a convenience function for creating a new HttpErr that responds
// with status code 401
func NewHttpUnauthorized(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpError(http.StatusUnauthorized, format, fmtArgs)
}

// NewHttpForbidden is a convenience function for creating a new HttpErr that responds
// with status code 403
func NewHttpForbidden(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpError(http.StatusForbidden, format, fmtArgs)
}

// NewHttpNotFound is a convenience function for creating a new HttpErr that responds
// with status code 404
func NewHttpNotFound(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpError(http.StatusNotFound, format, fmtArgs)
}

// NewHttpInternalServerErr is a convenience function for creating a new HttpErr that responds
// with status code 500
func NewHttpInternalServerErr(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpError(http.StatusInternalServerError, format, fmtArgs)
}

// NewHttpError returns a new IHttpError from an HTTP status code and a format error message string
func NewHttpError(statusCode int, format string, fmtArgs ...interface{}) HttpErr {
	err := &httpErr{
		statusCode: statusCode,
	}
	errorF(&err.stackErr, format, fmtArgs...)
	return err
}

func (he *httpErr) StatusCode() int {
	return he.statusCode
}
