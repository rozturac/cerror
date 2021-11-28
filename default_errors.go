package cerror

import (
	"fmt"
	"reflect"
)

func InvalidCastError(from, to interface{}) Error {
	return New(ApplicationError, fmt.Sprintf("Cannot convert the '%v' value to a %s.", from, reflect.TypeOf(to).Name()))
}
