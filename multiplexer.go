package erru

// Multiplexer takes a collection of functions that return
// errors and runs them in parallel
type Multiplexer interface {
	// Add adds one or more functions to the multiplexer
	Add(fns ...func() error)

	// Go runs each func added to the multiplexer in parallel,
	// waits for each to complete, then returns the first non-nil
	// error encountered. If no error is encountered then nil is
	// returned. If only one func has been added to the multiplexer
	// then nothing is run in parallel.
	//
	// This function blocks until all channels complete.
	Go() error
}

// multiplexer is an implementation of Multiplexer
type multiplexer struct {
	funcs []func() error
}

// NewMultiplexer returns a new instance of Multiplexer
func NewMultiplexer() Multiplexer {
	return &multiplexer{
		funcs: make([]func() error, 0, 2),
	}
}

func (m *multiplexer) Add(fns ...func() error) {
	if fns == nil {
		return
	}

	for _, fn := range fns {
		if fn == nil {
			continue
		}

		m.funcs = append(m.funcs, fn)
	}
}

func (m *multiplexer) Go() error {
	fnSize := len(m.funcs)

	if fnSize == 0 {
		return nil
	}

	if fnSize == 1 {
		return m.funcs[0]()
	}

	errChan := make(chan error, fnSize)
	for _, fn := range m.funcs {
		funcToRun := fn
		go func() {
			errChan <- funcToRun()
		}()
	}

	var err error
	for count := 0; count < fnSize; count += 1 {
		innerErr := <-errChan
		if innerErr != nil && err == nil {
			err = innerErr
		}
	}

	return err
}
