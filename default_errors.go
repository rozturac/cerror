package cerror

import (
	"fmt"
	"reflect"
)

// InvalidCastError Create an instance of Invalid Cast Error with object types
func InvalidCastError(from, to interface{}) Error {
	return New(ApplicationError, fmt.Sprintf("Cannot convert the '%v' value to %s.", from, reflect.TypeOf(to).Name()))
}

// NullReferenceError Create an instance of Null Reference Error with object name
func NullReferenceError(fieldName string) Error {
	return New(ApplicationError, fmt.Sprintf("%s is null!", fieldName))
}
