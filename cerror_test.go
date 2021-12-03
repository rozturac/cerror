package cerror

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	expectedErrorType := DomainError
	expectedMessage := "Test Error Message"
	expectedErrorCode := "CustomError"
	expectedHttpStatusCode := http.StatusInternalServerError
	err := New(expectedErrorType, expectedMessage)

	if err.HttpStatusCode() != expectedHttpStatusCode {
		t.Error(fmt.Sprintf("Expected code: %d, received: %d", expectedHttpStatusCode, err.HttpStatusCode()))
	}

	if err.ErrorType() != expectedErrorType {
		t.Error(fmt.Sprintf("Expected ErrorType: %s, received: %s", expectedErrorType, err.ErrorType()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}

	if err.ErrorCode() != expectedErrorCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedErrorCode, err.ErrorCode()))
	}
}

func TestNewWithErrorCode(t *testing.T) {
	expectedErrorType := ApplicationError
	expectedMessage := "Test Error Message"
	expectedErrorCode := "TestErrorCode"
	expectedHttpStatusCode := http.StatusInternalServerError
	err := NewWithErrorCode(expectedErrorType, expectedMessage, expectedErrorCode)

	if err.HttpStatusCode() != expectedHttpStatusCode {
		t.Error(fmt.Sprintf("Expected code: %d, received: %d", expectedHttpStatusCode, err.HttpStatusCode()))
	}

	if err.ErrorType() != expectedErrorType {
		t.Error(fmt.Sprintf("Expected ErrorType: %s, received: %s", expectedErrorType, err.ErrorType()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}

	if err.ErrorCode() != expectedErrorCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedErrorCode, err.ErrorCode()))
	}
}

func TestNewWithHttpStatusCode(t *testing.T) {
	expectedErrorType := ApplicationError
	expectedMessage := "Test Bad Gateway Message"
	expectedErrorCode := "CustomError"
	expectedHttpStatusCode := http.StatusBadGateway
	err := NewWithHttpStatusCode(expectedErrorType, expectedMessage, expectedHttpStatusCode)

	if err.HttpStatusCode() != expectedHttpStatusCode {
		t.Error(fmt.Sprintf("Expected code: %d, received: %d", expectedHttpStatusCode, err.HttpStatusCode()))
	}

	if err.ErrorType() != expectedErrorType {
		t.Error(fmt.Sprintf("Expected ErrorType: %s, received: %s", expectedErrorType, err.ErrorType()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}

	if err.ErrorCode() != expectedErrorCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedErrorCode, err.ErrorCode()))
	}
}

func TestNewWithErrorCodeAndHttpStatusCode(t *testing.T) {
	expectedErrorType := DomainError
	expectedMessage := "Test Conflict Message"
	expectedErrorCode := "Conflict Error"
	expectedHttpStatusCode := http.StatusConflict
	err := NewWithErrorCodeAndHttpStatusCode(expectedErrorType, expectedMessage, expectedErrorCode, expectedHttpStatusCode)

	if err.HttpStatusCode() != expectedHttpStatusCode {
		t.Error(fmt.Sprintf("Expected code: %d, received: %d", expectedHttpStatusCode, err.HttpStatusCode()))
	}

	if err.ErrorType() != expectedErrorType {
		t.Error(fmt.Sprintf("Expected ErrorType: %s, received: %s", expectedErrorType, err.ErrorType()))
	}

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}

	if err.ErrorCode() != expectedErrorCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedErrorCode, err.ErrorCode()))
	}
}

func TestCustomError_Error(t *testing.T) {
	expectedErrorType := DomainError
	expectedMessage := "Test Conflict Message"
	err := New(expectedErrorType, expectedMessage)

	if err.Error() != expectedMessage {
		t.Error(fmt.Sprintf("Expected message: %s, received: %s", expectedMessage, err.Error()))
	}
}

func TestCustomError_ErrorCode(t *testing.T) {
	expectedErrorType := DomainError
	expectedMessage := "Test Conflict Message"
	expectedErrorCode := "Conflict Error"
	err := NewWithErrorCode(expectedErrorType, expectedMessage, expectedErrorCode)

	if err.ErrorCode() != expectedErrorCode {
		t.Error(fmt.Sprintf("Expected code: %s, received: %s", expectedErrorCode, err.ErrorCode()))
	}
}

func TestCustomError_ErrorType(t *testing.T) {
	expectedErrorType := DomainError
	expectedMessage := "Test Conflict Message"
	err := New(expectedErrorType, expectedMessage)

	if err.ErrorType() != expectedErrorType {
		t.Error(fmt.Sprintf("Expected ErrorType: %s, received: %s", expectedErrorType, err.ErrorType()))
	}
}

func TestCustomError_HttpStatusCode(t *testing.T) {
	expectedErrorType := DomainError
	expectedMessage := "Test Conflict Message"
	expectedHttpStatusCode := http.StatusConflict
	err := NewWithHttpStatusCode(expectedErrorType, expectedMessage, expectedHttpStatusCode)

	if err.HttpStatusCode() != expectedHttpStatusCode {
		t.Error(fmt.Sprintf("Expected code: %d, received: %d", expectedHttpStatusCode, err.HttpStatusCode()))
	}
}
