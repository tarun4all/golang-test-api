package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All users endpont hit")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user endpont hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete user endpont hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user endpont hit")
}
