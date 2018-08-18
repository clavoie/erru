package erru

// Impl is a wrapper around all top level erru functions
// that can be used for dependency injection
type Impl interface {
	First(errs ...error) error

	// HttpErr
	NewHttpBadRequest(format string, fmtArgs ...interface{}) HttpErr
	NewHttpError(statusCode int, format string, fmtArgs ...interface{}) HttpErr
	NewHttpForbidden(format string, fmtArgs ...interface{}) HttpErr
	NewHttpInternalServerErr(format string, fmtArgs ...interface{}) HttpErr
	NewHttpNotFound(format string, fmtArgs ...interface{}) HttpErr
	NewHttpUnauthorized(format string, fmtArgs ...interface{}) HttpErr

	// StackErr
	Errorf(format string, fmtArgs ...interface{}) StackErr
	Wrap(err error) StackErr
	WrapF(err error, format string, fmtArgs ...interface{}) StackErr
	WrapN(err error, stackSize int) StackErr
	WrapNf(err error, stackSize int, format string, fmtArgs ...interface{}) StackErr

	// UserErrs
	NewUserErrs(errs ...string) UserErrs
}

// impl is an implementation of Impl
type impl struct{}

// NewImpl returns a new instance of Impl
func NewImpl() Impl { return new(impl) }

func (i *impl) First(errs ...error) error { return First(errs...) }
func (i *impl) NewHttpBadRequest(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpBadRequest(format, fmtArgs...)
}
func (i *impl) NewHttpError(statusCode int, format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpError(statusCode, format, fmtArgs...)
}
func (i *impl) NewHttpForbidden(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpForbidden(format, fmtArgs...)
}
func (i *impl) NewHttpInternalServerErr(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpInternalServerErr(format, fmtArgs...)
}
func (i *impl) NewHttpNotFound(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpNotFound(format, fmtArgs...)
}
func (i *impl) NewHttpUnauthorized(format string, fmtArgs ...interface{}) HttpErr {
	return NewHttpUnauthorized(format, fmtArgs...)
}
func (i *impl) Errorf(format string, fmtArgs ...interface{}) StackErr {
	return Errorf(format, fmtArgs...)
}
func (i *impl) Wrap(err error) StackErr { return Wrap(err) }
func (i *impl) WrapF(err error, format string, fmtArgs ...interface{}) StackErr {
	return WrapF(err, format, fmtArgs...)
}
func (i *impl) WrapN(err error, stackSize int) StackErr { return WrapN(err, stackSize) }
func (i *impl) WrapNf(err error, stackSize int, format string, fmtArgs ...interface{}) StackErr {
	return WrapNf(err, stackSize, format, fmtArgs...)
}
func (i *impl) NewUserErrs(errs ...string) UserErrs { return NewUserErrs(errs...) }
