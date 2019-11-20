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
	//e.POST("/hamster", handlers.AddHamster)

	//Dogs CRUD json
	e.POST("/dogs", handlers.AddDog)
	e.GET("/dogs/:dog_id", handlers.GetDog)
	e.GET("/dogs", handlers.ListDog)
	e.PUT("/dogs/:dog_id", handlers.UpdateDog)
	e.DELETE("/dogs/:dog_id", handlers.DeleteDog)

	//Hamster CRUD form-data
	e.POST("/hamster", handlers.AddHamster)
	e.GET("/hamster/:hamster_id", handlers.GetHamster)
	e.GET("/hamster", handlers.ListHamster)
	e.PUT("/hamster/:hamster_id", handlers.UpdateHamster)
	e.DELETE("/hamster/:hamster_id", handlers.DeleteHamster)
}
