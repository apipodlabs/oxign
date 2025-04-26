package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
