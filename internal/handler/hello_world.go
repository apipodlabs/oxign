package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World!")
}
