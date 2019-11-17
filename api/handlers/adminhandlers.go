package handlers

import (
	"net/http"
	"github.com/labstack/echo"
)

//MainAdmin page
func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK,"hello your are in Admin page")
}