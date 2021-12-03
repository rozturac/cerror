package cerror

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

func TestInvalidCastError(t *testing.T) {
	x := "x25"
	y, _ := strconv.Atoi(x)
	err := InvalidCastError(x, y)
	expectedErrorCode := "InvalidCastError"
	expectedHttpStatusCode := http.StatusBadRequest
	expectedMessage := "Cannot convert the 'x25' value to int."

	if err.HttpStatusCode() != expectedHttpStatusCode {
		t.Error(fmt.Sprintf("Expected code: %d, received: %d", expectedHttpStatusCode, err.HttpStatusCode()))
	}

	if err.ErrorCode() != expectedErrorCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedErrorCode, err.ErrorCode()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}
}

func TestNullReferenceError(t *testing.T) {
	err := NullReferenceError("x")
	expectedErrorCode := "NullReferenceError"
	expectedHttpStatusCode := http.StatusBadRequest
	expectedMessage := "x is null!"

	if err.HttpStatusCode() != expectedHttpStatusCode {
		t.Error(fmt.Sprintf("Expected code: %d, received: %d", expectedHttpStatusCode, err.HttpStatusCode()))
	}

	if err.ErrorCode() != expectedErrorCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedErrorCode, err.ErrorCode()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}
}
