package cerror

import (
	"fmt"
	"net/http"
	"reflect"
)

// InvalidCastError Create an instance of Invalid Cast Error with object types
func InvalidCastError(from, to interface{}) Error {
	return NewWithHttpStatusCode(ApplicationError, fmt.Sprintf("Cannot convert the '%v' value to %s.", from, reflect.TypeOf(to).Name()), http.StatusBadRequest)
}

// NullReferenceError Create an instance of Null Reference Error with object name
func NullReferenceError(fieldName string) Error {
	return NewWithHttpStatusCode(ApplicationError, fmt.Sprintf("%s is null!", fieldName), http.StatusBadRequest)
}
