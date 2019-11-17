package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)


//Dog struct
type Dog struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Owner *Owner `json:"owner"`
}

//Owner struct
type Owner struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Dogs slice
var Dogs []Dog

//AddDog method 2s
func AddDog(c echo.Context) error {
	dog := Dog{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	dog.ID = strconv.Itoa(rand.Intn(10000))

	Dogs = append(Dogs, dog)
	if err != nil {
		log.Printf("failing to process addDog request :%s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog : %v", dog)
	fmt.Println("====================")
	fmt.Println("your dog ID", dog.ID)
	fmt.Println("Owner", dog.Owner.Firstname, dog.Owner.Lastname)
	fmt.Println("=====================")
	// return c.String(http.StatusOK, "Your dog details:\nID:"+dog.ID+"\nName:"+dog.Name+"\nType:"+dog.Type+"\nOwner:")
	return json.NewEncoder(c.Response()).Encode(&dog)
}

//ListDog 全部の犬リスト
func ListDog(c echo.Context) error {
	//slice
	if len(Dogs) <= 0 {
		// return c.JSON(http.StatusBadRequest, map[string]string{
		// 	"error": "No data , please insert the data first",
		// })
		return c.String(http.StatusBadRequest, "Insert the data first")
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(Dogs)

	// u := &User{
	// 	Name:  "Jon",
	// 	Email: "jon@labstack.com",
	//   }
	//　return c.JSON(http.StatusOK, u)
}

//GetDog getby id
func GetDog(c echo.Context) error {
	//Get ID
	dogid := c.Param("dog_id")
	for _, item := range Dogs {
		//check Data in slice
		if item.ID == dogid {
			fmt.Println("here is your dog mate ", item.Name, item.Type)
			fmt.Println("Owner", item.Owner.Firstname, item.Owner.Lastname)
			return c.JSON(http.StatusOK, item)
		}
		log.Printf("No Dog found")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "No data found -- by get",
		})

	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "list is empty no data inserted yet",
	})
}

//UpdateDog update slice dogs
func UpdateDog(c echo.Context) error {

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	//Get ID
	dogid := c.Param("dog_id")
	for index, item := range Dogs {
		if item.ID == dogid {
			Dogs = append(Dogs[:index], Dogs[index+1:]...)
			dog := Dog{}

			defer c.Request().Body.Close()
			dog.ID = dogid

			err := json.NewDecoder(c.Request().Body).Decode(&dog)

			if err != nil {
				log.Printf("failing to process Update Dog request :%s", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			Dogs = append(Dogs, dog)
			log.Printf("Updated Dog value : %v", dog)
		}
	}
	return json.NewEncoder(c.Response()).Encode(Dogs)
}

//DeleteDog delete dog
func DeleteDog(c echo.Context) error {
	dogid := c.Param("dog_id")
	for index, item := range Dogs {
		if item.ID == dogid {
			Dogs = append(Dogs[:index], Dogs[index+1:]...)
			break
		}
		log.Printf("No Dog found")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "No data found -- by delete",
		})
	}
	return json.NewEncoder(c.Response()).Encode(Dogs)
}