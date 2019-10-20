package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/me", getUser).Methods("GET")

	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	router.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":9090", router))
}

func main() {
	fmt.Println("Listening on port 9090")
	IntialMigration()
	handleRequests()
}
