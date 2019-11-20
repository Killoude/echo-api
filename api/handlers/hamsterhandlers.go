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

//Hamster Object
type Hamster struct {
	ID    string       `json:"id" form:"id"`
	Name  string       `json:"name" form:"name"`
	Type  string       `json:"type" form:"type"`
	Owner HamsterOwner `json:"owner"`
}

//HamsterOwner for hamster
type HamsterOwner struct {
	Firstname string `json:"firstname" form:"firstname"`
	Lastname  string `json:"lastname" form:"lastname"`
}

//Hamsters slice
var Hamsters []Hamster

//AddHamster form-data
func AddHamster(c echo.Context) error {
	hamster := Hamster{}
	hamster.ID = strconv.Itoa(rand.Intn(10000))

	hamster.Name = c.FormValue("name")
	hamster.Type = c.FormValue("type")
	// incase *HamsterOwner harus nunjuk dlu ke struct yg mau baru bs diset
	//	temp := HamsterOwner{}
	// hamster.Owner = &temp
	hamster.Owner.Firstname = c.FormValue("firstname")
	hamster.Owner.Lastname = c.FormValue("lastname")
	Hamsters = append(Hamsters, hamster)

	//log.Printf("this is your hamster: %#v", hamster)
	fmt.Println("====================")
	fmt.Println("your Hamster ID", hamster.ID)
	fmt.Println("Hamster Owner", hamster.Owner.Firstname, hamster.Owner.Lastname)
	fmt.Println("=====================")
	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "Your Hamster details:\nID:"+hamster.ID+"\nName:"+hamster.Name+"\nType:"+hamster.Type+"\nOwner:"+hamster.Owner.Firstname+"\n"+hamster.Owner.Lastname)
}

//ListHamster listhamster
func ListHamster(c echo.Context) error {
	if len(Hamsters) <= 0 {
		return c.String(http.StatusOK, "input the hamster data first")
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(Hamsters)
}

//GetHamster for the specified
func GetHamster(c echo.Context) error {
	//getid
	id := c.Param("hamster_id")
	for _, item := range Hamsters {
		if item.ID == id {
			fmt.Println("here is your hamster mate ", item.Name, item.Type)
			fmt.Println("Owner", item.Owner.Firstname, item.Owner.Lastname)
			return c.JSON(http.StatusOK, item)
		}
	}
	return c.String(http.StatusBadRequest, "No id")
}

//UpdateHamster hamster 編集
func UpdateHamster(c echo.Context) error {
	id := c.Param("hamster_id")
	for index, item := range Hamsters {
		if item.ID == id {
			Hamsters = append(Hamsters[:index], Hamsters[index+1:]...)
			hamster := Hamster{}
			hamster.Name = c.FormValue("name")
			hamster.Type = c.FormValue("type")
			hamster.Owner.Firstname = c.FormValue("firstname")
			hamster.Owner.Lastname = c.FormValue("lastname")
			hamster.ID = id
			Hamsters = append(Hamsters, hamster)
			log.Printf("Updated Hamster value : %v", hamster)
			return c.String(http.StatusOK, "your hamster has been updated")
		}
	}
	return c.String(http.StatusBadRequest, "no hamster data")
}

//DeleteHamster Hamster data 削除
func DeleteHamster(c echo.Context) error {
	id := c.Param("hamster_id")
	for index, item := range Hamsters {
		if item.ID == id {
			Hamsters = append(Hamsters[:index], Hamsters[index+1:]...)
			break
		}
		log.Printf("no hamster found")
		return c.String(http.StatusBadRequest, "no hamster to delete")
	}
	return json.NewEncoder(c.Response()).Encode(Hamsters)
}
