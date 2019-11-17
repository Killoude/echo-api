package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SetAdminMiddlewares for admin page middlewares
func SetAdminMiddlewares(g *echo.Group) {
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${status} ${method} ${host} ${path} ${latency_human} + \n`,
	}))

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		//check in the database
		if username == "jack" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))
}
