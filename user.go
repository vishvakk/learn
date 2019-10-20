package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string 
	Email string
}

func IntialMigration(){
	db, err = gorm.Open("sqlite3", "ts.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Fail to connect")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}
func AllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "All users")
}


func NewUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "All users")

}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "All users")

}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "All users")

}