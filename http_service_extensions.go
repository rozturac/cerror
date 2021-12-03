package cerror

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseMapper func(errorCode, message string, httpStatusCode int) interface{}

func AddErrorHandlingMiddlewareWithMapper(next http.HandlerFunc, mapper ErrorResponseMapper) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			response := mapper("UnExpectedError", http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			httpStatusCode := http.StatusInternalServerError
			if recover := recover(); recover != nil {
				if err, ok := recover.(error); ok {
					if customError, ok := err.(Error); ok {
						response = mapper(customError.ErrorCode(), customError.Error(), customError.HttpStatusCode())
						httpStatusCode = customError.HttpStatusCode()
					} else {
						response = mapper("UnExpectedError", err.Error(), http.StatusInternalServerError)
					}
				}
			}

			if buff, err := json.Marshal(response); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			} else {
				http.Error(writer, string(buff), httpStatusCode)
			}
		}()

		next(writer, request)
	}
}
