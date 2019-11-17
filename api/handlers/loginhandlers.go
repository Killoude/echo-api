package handlers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//JwtClaims Token object
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
//Login func
func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	//check db
	if username == "jack" && password == "1234" {
		cookie := &http.Cookie{} //cookie := new(http.Cookie)
		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)
		// create jwt token
		token, err := createJwtToken()
		if err != nil {
			log.Println("Error Creating JWT token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "you were logged in",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "username or password is wrong")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	key, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)

	if err != nil {
		return "", err
	}
	token, err := rawToken.SignedString(key)

	if err != nil {
		return "", err
	}
	return token, nil
}
