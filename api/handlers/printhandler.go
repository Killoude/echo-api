package handlers

import (
	"net/http"
	"github.com/labstack/echo"
)

//GetPrint  for printing word to screen
func GetPrint(c echo.Context) error {
	return c.String(http.StatusOK,"printing this page..")
}