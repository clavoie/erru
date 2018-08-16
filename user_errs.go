package erru

// UserErrs represents input errors from the user.
//
// UserErrs represents an error not with the system, but
// rather with the input the system received. It still
// indicates that the system should stop processing and
// return, but not that it needs to exit with a fatal or
// 500 status.
type UserErrs interface {
	// The implementation of error prints the first user
	// error found, or the empty string if there were no errors
	error

	// User returns the collection of input errors
	// coming from the user
	User() []string
}

// userErrs is an implementation of UserErrs
type userErrs struct {
	errs []string
}

// NewUserErrors creates and returns a new UserErrs
func NewUserErrors(errs ...string) UserErrs {
	return &userErrs{
		errs: errs,
	}
}

func (ue *userErrs) Error() string {
	if len(ue.errs) == 0 {
		return ""
	}

	return ue.errs[0]
}

func (ue *userErrs) User() []string {
	return ue.errs
}
