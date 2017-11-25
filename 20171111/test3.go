package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome !")
}

func Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "hello "+params.ByName("name"))
}

func GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	fmt.Fprintf(w, " get user id = "+id)
}

func ModifyUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	fmt.Fprintf(w, " modify user id = "+id)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	fmt.Fprintf(w, " delete user id = "+id)
}

func AddUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	fmt.Fprintf(w, " add user id = "+id)
}

func main() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/user/:id", GetUser)
	router.PUT("/user/:id", ModifyUser)
	router.DELETE("/user/:id", DeleteUser)
	router.POST("/user/:id", AddUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
