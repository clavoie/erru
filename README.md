# erru [![GoDoc Reference](https://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/clavoie/erru) [![Build Status](https://travis-ci.org/clavoie/erru.svg?branch=master)](https://travis-ci.org/clavoie/erru) [![codecov](https://codecov.io/gh/clavoie/erru/branch/master/graph/badge.svg)](https://codecov.io/gh/clavoie/erru) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/erru)](https://goreportcard.com/report/github.com/clavoie/erru)

Error utilities for Go.

## Stack Traces

You can add stack traces to an existing error by calling Wrap. If you'd like to add a context message along with the stack trace you can call WrapF. If you'd like to create a brand new error that contains a stack trace and message, call Errorf.

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

Below is an example of how the stack traces are formatted. If erru detects another erru error with additional context, it will indent the inner error by another level:

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
```
[All HTTP helper functions here](https://godoc.org/github.com/clavoie/erru#HttpErr)

## Multiplexer

Sometimes multiple db or external calls need to be done in parallel, and the higher level code only cares about if an error is returned or not. erru provides a multiplexing utility that runs several func in separate goroutines, waits for them to end, collects the first error and returns it to the caller.
