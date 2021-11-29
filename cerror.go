package cerror

import (
	"reflect"
	"strings"
)

// Error defines custom error
type Error interface {
	Error() string
	ErrorWithTrace() string
	With(err error) Error
}

type customError struct {
	Type           ErrorType
	message        string
	traces         string
	buildError     error
	appendedErrors []error
}

// ErrorWithTrace Expose the Error message
func (c customError) Error() string {
	return c.message
}

// ErrorWithTrace Expose the Error message with trace
func (c customError) ErrorWithTrace() string {
	var result strings.Builder
	result.WriteString(c.message)
	result.WriteString("\n")

	for _, x := range c.appendedErrors {
		result.WriteString(x.Error())
		result.WriteString("\n")
	}

	result.WriteString(c.buildError.Error())
	result.WriteString("\n")
	result.WriteString(c.traces)
	return result.String()
}

// With Append the low or high level Error
func (c customError) With(err error) Error {
	if reflect.TypeOf(err).Kind() == reflect.TypeOf(customError{}).Kind() {
		c.appendedErrors = append(c.appendedErrors, err.(error))
	} else {
		c.buildError = err
	}
	return c
}

// New Create an instance of Error with Type and Message
func New(errType ErrorType, message string) Error {
	return &customError{
		Type:    errType,
		message: message,
		traces:  getStackTraces(),
	}
}
