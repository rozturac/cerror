package extensions

import (
	"github.com/rozturac/cerror"
	"net/http"
)

type ErrorResponseMapper func(errorCode, message string, httpStatusCode int) interface{}

func getErrorResponseAndStatusCode(candidateError interface{}, mapper ErrorResponseMapper) (interface{}, int) {
	response := mapper("UnExpectedError", http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	httpStatusCode := http.StatusInternalServerError

	if err, ok := candidateError.(error); ok {
		if customError, ok := err.(cerror.Error); ok {
			response = mapper(customError.ErrorCode(), customError.Error(), customError.HttpStatusCode())
			httpStatusCode = customError.HttpStatusCode()
		} else {
			response = mapper("UnExpectedError", err.Error(), http.StatusInternalServerError)
		}
	}

	return response, httpStatusCode
}
