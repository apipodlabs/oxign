package main

import (
	"net/http"

	_ "github.com/apipodlabs/oxign/api/oxign"
	"github.com/apipodlabs/oxign/internal/auth"
	handler "github.com/apipodlabs/oxign/internal/handler"
	"github.com/apipodlabs/oxign/internal/health"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"

	_ "github.com/golang-migrate/migrate/v4/source/github"
)

//	@title			Oxign API Documentation
//	@version		0.0.1
//	@description	Restful API Documentation on Oxign Platform.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Apipodlabs
//	@contact.url	https://github.com/apipodlabs/oxign/issues
//	@contact.email	muhammaddaffadinaya@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:1323
// @BasePath	/
// @schemes	http
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
	jwtAuthMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	})

	app.GET("/", handler.HelloWorld)
	app.GET("/health", health.HealthCheck)
	app.GET("/login", auth.Authenticate)
	app.GET("/verify", auth.JWTAuth, jwtAuthMiddleware)

	app.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	app.Logger.Fatal(app.Start(":1323"))
}
