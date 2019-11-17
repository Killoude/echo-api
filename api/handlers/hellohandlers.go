package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//Hello public
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello from 200")
}
