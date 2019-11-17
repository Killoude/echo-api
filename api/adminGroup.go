package api

import (
	"api/handlers"
	"github.com/labstack/echo"
)

//AdminGroup grouping
func AdminGroup(g *echo.Group) {
	
	g.GET("/main", handlers.MainAdmin)
}
