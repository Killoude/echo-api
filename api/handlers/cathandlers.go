package handlers
import (
    "net/http"
    "fmt"
    "log"
    "io/ioutil"
    "encoding/json"

    "github.com/labstack/echo"
)

//Cat object
type Cat struct {
    Name        string    `json:"name"`
    Type        string    `json:"type"`
}
//GetCats public function
func GetCats(c echo.Context) error {
	//passing name and type
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is %s\nyour cat type is %s", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "invalid request",
	})

}

//AddCat method1  the fastest
func AddCat(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("failed recording the request body :%s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("failed Unmarshal in addCats :%s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat : %#v", cat)
	return c.String(http.StatusOK, "we got your cat!")
}