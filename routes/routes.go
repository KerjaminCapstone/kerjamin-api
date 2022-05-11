package routes

import (
	authRoute "github.com/KerjaminCapstone/kerjamin-backend-v1/routes/auth"
	clientRoute "github.com/KerjaminCapstone/kerjamin-backend-v1/routes/client"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const CSRFTokenHeader = "X-CSRF-Token"
const CSRFKey = "csrf"

func Init() *echo.Echo {
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Group Auth
	authGroup := e.Group("/auth")
	authRoute.AuthSubRoute(authGroup)

	// Group Client
	clientGroup := e.Group("/client")
	clientRoute.ClientSubRoute(clientGroup)

	return e
}
