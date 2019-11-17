package middlewares

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)
//Set cookie through the session
func SetCookieMiddlewares(g *echo.Group) {
	g.Use(CheckCookie)
}
//CheckCookie valid or not
func CheckCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")

		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you don't have the right cookie")
			}
			log.Println(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "You don't have the right cookie")
	}
}
