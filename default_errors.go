package cerror

import (
	"fmt"
	"reflect"
)

// InvalidCastError Create an instance of Invalid Cast Error with object types
func InvalidCastError(from, to interface{}) Error {
	return New(ApplicationError, fmt.Sprintf("Cannot convert the '%v' value to a %s.", from, reflect.TypeOf(to).Name()))
}
