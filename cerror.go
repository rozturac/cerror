package cerror

import (
	"net/http"
	"reflect"
	"strings"
)

// Error defines custom error
type Error interface {
	ErrorType() ErrorType
	ErrorCode() string
	HttpStatusCode() int
	Error() string
	ErrorWithTrace() string
	With(err error) Error
}

type customError struct {
	Type           ErrorType
	message        string
	errorCode      string
	httpStatusCode int
	traces         string
	buildError     error
	appendedErrors []error
}

// HttpStatusCode Expose the StatusCode
func (c customError) HttpStatusCode() int {
	return c.httpStatusCode
}

// ErrorCode Expose the ErrorCode
func (c customError) ErrorCode() string {
	return c.errorCode
}

// ErrorType Expose the ErrorType
func (c customError) ErrorType() ErrorType {
	return c.Type
}

// ErrorWithTrace Expose the Error message
func (c customError) Error() string {
	return c.message
}

// ErrorWithTrace Expose the Error message with trace
func (c customError) ErrorWithTrace() string {
	var result strings.Builder
	result.WriteString("[")
	result.WriteString(c.errorCode)
	result.WriteString("] ")
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
	return NewWithErrorCodeAndHttpStatusCode(errType, message, getErrorCode(), http.StatusInternalServerError)
}

// NewWithErrorCode Create an instance of Error with errType, message, code
func NewWithErrorCode(errType ErrorType, message string, errCode string) Error {
	return NewWithErrorCodeAndHttpStatusCode(errType, message, errCode, http.StatusInternalServerError)
}

// NewWithHttpStatusCode Create an instance of Error with errType, message, code
func NewWithHttpStatusCode(errType ErrorType, message string, httpStatusCode int) Error {
	return NewWithErrorCodeAndHttpStatusCode(errType, message, getErrorCode(), httpStatusCode)
}

// NewWithErrorCodeAndHttpStatusCode Create an instance of Error with extra errorCode and httpStatusCode
func NewWithErrorCodeAndHttpStatusCode(errType ErrorType, message, errorCode string, httpStatusCode int) Error {
	return &customError{
		Type:           errType,
		message:        message,
		traces:         getStackTraces(),
		errorCode:      errorCode,
		httpStatusCode: httpStatusCode,
	}
}
