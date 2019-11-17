package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SetMainMiddlewares as middleware
func SetMainMiddlewares(e *echo.Echo) {
	//to serve static page
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../../static",
	}))

	e.Use(setHeader)
}

func setHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("notReallyHeader", "this has not meaning")
		return next(c)
	}
}
