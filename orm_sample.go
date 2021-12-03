package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("OMG")
	}

	defer db.Close()
	db.AutoMigrate(&User{})
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("No Connection")
	}

	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("No Connection")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	db.Create(&User{Name: name})
	fmt.Fprintf(w, "User Created successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("No Connection")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Deleted successfully")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("No Connection")
	}

	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Name = "New Name"

	db.Save(&user)
	fmt.Fprintf(w, "Updated successfully")
}

func server() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}", NewUser).Methods("POST")
	router.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}", UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	InitialMigration()
	server()

	fmt.Println("Server started...")
}
