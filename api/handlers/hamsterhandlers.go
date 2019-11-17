package handlers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

//Hamster Object
type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

//AddHamster 3
func AddHamster(c echo.Context) error {
	hamster := Hamster{}
	defer c.Request().Body.Close()
	err := c.Bind(&hamster)

	if err != nil {
		log.Printf("failing to process addHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster !")
}
