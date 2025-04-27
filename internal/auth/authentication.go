package auth

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("TEST")
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader != "" {
			const prefixBearer = "Bearer "
			const prefixLicense = "License "

			if strings.HasPrefix(authHeader, prefixBearer) {
				token := strings.TrimPrefix(authHeader, prefixBearer)
				fmt.Println("Token:", token)
			} else if strings.HasPrefix(authHeader, prefixLicense) {
				token := strings.TrimPrefix(authHeader, prefixLicense)
				fmt.Println("Token:", token)
			}
		}
		return next(c)
	}
}
