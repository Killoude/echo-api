package api

import (
	"api/handlers"
	"github.com/labstack/echo"
)

//MainGroup grouping
func MainGroup(e *echo.Echo) {
	e.GET("/hello", handlers.Hello)
	e.GET("/login", handlers.Login)
	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)
	e.POST("/hamster", handlers.AddHamster)

	//Dogs CRUD
	e.POST("/dogs", handlers.AddDog) 
	e.GET("/dogs/:dog_id", handlers.GetDog)
	e.GET("/dogs", handlers.ListDog)
	e.PUT("/dogs/:dog_id", handlers.UpdateDog)
	e.DELETE("/dogs/:dog_id", handlers.DeleteDog)
}
