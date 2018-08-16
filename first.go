package erru

// First returns the first error passed to it that is not nil,
// or nil if all errors are nil
func First(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
