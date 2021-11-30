package cerror

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInvalidCastError(t *testing.T) {
	x := "x25"
	y, _ := strconv.Atoi(x)
	err := InvalidCastError(x, y)
	expectedCode := "InvalidCastError"
	expectedMessage := "Cannot convert the 'x25' value to int."

	if err.Code() != expectedCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedCode, err.Code()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}
}

func TestNullReferenceError(t *testing.T) {
	err := NullReferenceError("x")
	expectedCode := "NullReferenceError"
	expectedMessage := "x is null!"

	if err.Code() != expectedCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedCode, err.Code()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}
}
