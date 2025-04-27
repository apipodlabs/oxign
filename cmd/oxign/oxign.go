package main

import (
	"github.com/apipodlabs/oxign/internal/auth"
	handler "github.com/apipodlabs/oxign/internal/handler"
	"github.com/apipodlabs/oxign/internal/health"
	"github.com/labstack/echo/v4"

	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"

	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	// m, err := migrate.New(
	// 	"file://db/migrations",
	// 	"cockroachdb://cockroach:@localhost:26257/example?sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := m.Up(); err != nil {
	// 	log.Fatal(err)
	// }
	// TODO: create middleware for license auth
	// TODO: selective path for JWT middleware to apply for several path only
	// TODO: good-to-have create oauth2 and open id connect compliance API

	app := echo.New()

	app.GET("/", handler.HelloWorld)
	app.GET("/health", health.HealthCheck)
	app.POST("/login", handler.HelloWorld)
	app.GET("/me", handler.HelloWorld, auth.AuthMiddleware)

	app.Logger.Fatal(app.Start("localhost:8888"))
}
