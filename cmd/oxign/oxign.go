package main

import (
	"log"

	handler "github.com/apipodlabs/oxign/internal/handler"
	"github.com/golang-migrate/migrate/v4"
	"github.com/labstack/echo/v4"

	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"

	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	m, err := migrate.New(
		"file://db/migrations",
		"cockroachdb://cockroach:@localhost:26257/example?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

	app := echo.New()
	app.GET("/", handler.HelloWorld)

	app.Logger.Fatal(app.Start(":1323"))
}
