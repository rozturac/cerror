package main

import (
	"github.com/rozturac/cerror"
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		panic(cerror.NullReferenceError("testProp"))
	})

	mapper := func(errorCode, message string, httpStatusCode int) interface{} {
		return ErrorResponse{
			Code:    errorCode,
			Message: message,
		}
	}

	http.Handle("/example", cerror.AddErrorHandlingMiddlewareWithMapper(handler, mapper))
	http.ListenAndServe(":8080", nil)
}
