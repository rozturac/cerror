package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rozturac/cerror"
	"github.com/rozturac/cerror/extensions"
	"net/http"
)

type EchoErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.Use(extensions.AddEchoErrorHandlingMiddlewareWithMapper(func(errorCode, message string, httpStatusCode int) interface{} {
		return EchoErrorResponse{
			Code:    errorCode,
			Message: message,
		}
	}))

	e.Use(middleware.CORS())

	e.GET("/example", example)

	e.Start(":8080")
}

func example(c echo.Context) error {
	if name := c.Param("Name"); len(name) == 0 {
		panic(cerror.NullReferenceError("name"))
	}

	return c.String(http.StatusOK, "Hello, World!")
}
