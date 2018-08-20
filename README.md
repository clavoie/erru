# erru [![GoDoc Reference](https://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/clavoie/erru) [![Build Status](https://travis-ci.org/clavoie/erru.svg?branch=master)](https://travis-ci.org/clavoie/erru) [![codecov](https://codecov.io/gh/clavoie/erru/branch/master/graph/badge.svg)](https://codecov.io/gh/clavoie/erru) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/erru)](https://goreportcard.com/report/github.com/clavoie/erru)

Error utilities for Go.

* [Stack Traces](#stack-traces)
* [Http](#http)
* [Multiplexer](#multiplexer)
* [Dependency Injection](#dependency-injection)

## Stack Traces
Stack traces can be added to an existing error by calling Wrap. A context message for the error along with the stack trace can be added by calling WrapF. New errors that contains a stack trace can be created by calling Errorf.

```go
  func DoWork(arg interface{}) error {
    err := somethingThatCanFail1(arg)
    
    if err != nil {
      return erru.Wrap(err)
    }
    
    err = somethingThatCanFail2(arg)
    
    if err != nil {
      return erru.WrapF(err, "context message with for error: %v", arg)    
    }
    
    success := somethingThatReturnsTrueOrFalse(arg)
    
    if !success {
      return erru.Errorf("error message with stack trace: %v", arg)
    }
    
    return nil
  }
```

[All error helper functions here](https://godoc.org/github.com/clavoie/erru#StackErr)

Below is an example of how the stack traces are formatted. If erru detects it is wrapping another erru error with additional context, it will indent the inner error by another level:

```
outer error message
  inner error message
  functionName1
    path/to/func.go:14
  functionName2
    path/to/func.go:34
functionName3
  path/to/func.go:54
functionName4
  path/to/func.go:64
```

## Http
A definition is provided for an HttpErr, which is an error with a stack trace and an http status code value. There are several convenience functions for common http error codes, or you can create your own:

```go
  badRequestErr := erru.NewHttpBadRequest("input %v is not valid", someValue)
  internalServerErr := erru.NewHttpInternalServerErr("cannot process request: %v", someValue)
  methodNotAllowedErr := erru.NewHttpError(http.StatusMethodNotAllowed, "http method PUT is not allowed")
  
  func onHttpErr(err error, w http.ResponseWriter, r *http.Request) {
	logger := NewLogger(r)
	logger.LogRequestErr(err)

	httpErr, isHttpErr := err.(erru.HttpErr)
	if isHttpErr {
		w.WriteHeader(httpErr.StatusCode())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
```
[All HTTP helper functions here](https://godoc.org/github.com/clavoie/erru#HttpErr)

## Multiplexer

Sometimes multiple db or external calls need to be done in parallel, and the higher level code only cares about if any of those calls returned an error or not. Multiplexer is a utility that runs several funcs in separate goroutines, waits for them to end, collects the first error and returns it to the caller.

```go
  func loadData() error {
    data1 := new(Type1)
    data2 := new(Type2)
    // etc
    
    multiplexer := erru.NewMultiplexer()
    multiplexer.Add(func () error {
      return db.Query(&data1)
    }, func () error {
      return db.Query(&data2)
    })
    // etc
    
    // splits each func out into a seperate goroutine and
    // blocks until they all complete. the first non-nil error
    // is returned, or nil if all funcs completed successfully
    return multiplexer.Go()
  }
```

## Dependency Injection

An interface is provided to wrap all top level package functions. This interface can be injected the into your code instead of calling the package functions directly. A function to hook these dependencies into the [di dependency injection system](https://github.com/clavoie/di) is provided, but the constructors for all wrappers are open in case you would like to use another:

```go
  resolver, err := di.NewResolver(errHandler, erru.NewDiDefs())
  // etc
  
  func InjectableFunc(impl erru.Impl, multiplexer erru.Multiplexer) error {
    multiplexer.Add(DbCall, LoadFileFromS3)
    return impl.Wrap(multiplexer.Go())
  }
  
  err = resolver.Invoke(InjectableFunc)
  // log err, etc...
```
