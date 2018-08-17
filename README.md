# erru [![GoDoc Reference](https://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/clavoie/erru) [![Build Status](https://travis-ci.org/clavoie/erru.svg?branch=master)](https://travis-ci.org/clavoie/erru) [![codecov](https://codecov.io/gh/clavoie/erru/branch/master/graph/badge.svg)](https://codecov.io/gh/clavoie/erru) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/erru)](https://goreportcard.com/report/github.com/clavoie/erru)

Error utilities for go.

## Stack Traces

erru provides several convenience functions for creating or wrapping go errors in order to provide a stack trace.

```go
  func DoWork1() error {
    return erru.Wrap(somethingThatCanFail())
  }
  
  func DoWork2() error {
    if somethingFailed {
      return erru.Errorf("error message with stack trace: %v", someValue)
    }
  }
  
  func DoWork3() error {
    return erru.WrapF(somethingThatCanFail(), "formatted context msg: %v", someValue)
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
erru provides a definition for an HttpErr, which is an error with a stack trace and an http status code value. There are several convenience functions for common http error codes, or you can create your own:

```go
  badRequestErr := erru.NewHttpBadRequest("input %v is not valid", someValue)
  internalServerErr := erru.NewHttpInternalServerErr("cannot process request: %v", someValue)
  methodNotAllowedErr := erru.NewHttpError(http.StatusMethodNotAllowed, "http method PUT is not allowed")
```
[All HTTP helper functions here](https://godoc.org/github.com/clavoie/erru#HttpErr)

## Multiplexer
