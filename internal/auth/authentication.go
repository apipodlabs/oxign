package auth

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Authenticate(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World!\n")
}

func JWTAuth(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return errors.New("failed to cast claims as jwt.MapClaims")
	}
	return c.JSON(http.StatusOK, claims)
}

func DefaultSkipper(context echo.Context) bool {
	path := context.Path()
	whitelist := []string{
		"/login",
	}
	fmt.Printf("%s\n", path)
	return slices.Contains(whitelist, path)
}
