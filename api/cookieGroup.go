package api

import (
	"api/handlers"
	"github.com/labstack/echo"
)

//CookieGroup cookie
func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}

