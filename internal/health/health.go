package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck godoc
//
//	@Summary		Show the status of server.
//	@Description	get the status of server.
//	@Tags			root
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/health [get]
func HealthCheck(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
