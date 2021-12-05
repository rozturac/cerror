package extensions

import (
	"encoding/json"
	"net/http"
)

func AddHttpErrorHandlingMiddlewareWithMapper(next http.HandlerFunc, mapper ErrorResponseMapper) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if recover := recover(); recover != nil {
				response, httpStatusCode := getErrorResponseAndStatusCode(recover, mapper)
				if buff, err := json.Marshal(response); err != nil {
					http.Error(writer, err.Error(), http.StatusInternalServerError)
				} else {
					http.Error(writer, string(buff), httpStatusCode)
				}
			} else {
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next(writer, request)
	}
}
