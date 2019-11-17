package router

import (
	"api"
	"api/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New function
func New() *echo.Echo {
	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	//set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
	}))

	//set main routes
	api.MainGroup(e)

	//set group routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	return e
}
