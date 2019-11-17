package main

import (
	"fmt"
	"router"
	//"github.com/gorilla/mux"
	"net/http"
)

//MuxPage init
func MuxPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This page is accessed by mux library")

}

func main() {

	fmt.Println("Welcome to the server")
	e := router.New()
	// //Mux router
	// r := mux.NewRouter()
	// r.HandleFunc("/mux",MuxPage).Methods("GET")
	// http.ListenAndServe(":8080", r)
	e.Start(":8080")
}
