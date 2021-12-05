package extensions

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddEchoErrorHandlingMiddlewareWithMapper(mapper ErrorResponseMapper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			defer func() {
				if recover := recover(); recover != nil {
					response, httpStatusCode := getErrorResponseAndStatusCode(recover, mapper)
					if err := context.JSON(httpStatusCode, response); err != nil {
						context.NoContent(http.StatusInternalServerError)
					}
				} else {
					context.NoContent(http.StatusInternalServerError)
				}
			}()

			if err := next(context); err != nil {
				response, httpStatusCode := getErrorResponseAndStatusCode(err, mapper)
				if err := context.JSON(httpStatusCode, response); err != nil {
					context.NoContent(http.StatusInternalServerError)
				}
			}

			return nil
		}
	}
}
