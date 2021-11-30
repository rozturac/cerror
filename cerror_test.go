package cerror

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	expectedErrorType := DomainError
	expectedMessage := "Test Error Message"
	expectedCode := "CustomError"
	err := New(expectedErrorType, expectedMessage)

	if err.ErrorType() != expectedErrorType {
		t.Error(fmt.Sprintf("Expected ErrorType: %s, received: %s", expectedErrorType, err.ErrorType()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}

	if err.Code() != expectedCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedCode, err.Code()))
	}
}

func TestNewWithCode(t *testing.T) {
	expectedErrorType := ApplicationError
	expectedMessage := "Test Error Message"
	expectedCode := "TestErrorCode"
	err := NewWithCode(expectedErrorType, expectedMessage, expectedCode)

	if err.ErrorType() != expectedErrorType {
		t.Error(fmt.Sprintf("Expected ErrorType: %s, received: %s", expectedErrorType, err.ErrorType()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}

	if err.Code() != expectedCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedCode, err.Code()))
	}
}
