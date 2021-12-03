package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Blog struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Blogs []Blog

func allBlogs(w http.ResponseWriter, r *http.Request) {
	blogs := Blogs{Blog{Title: "First API", Desc: "First API", Content: "Hello World"}}
	fmt.Println("All blogs API hit")

	json.NewEncoder(w).Encode(blogs)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Working")
}

// TODO: add custom handler for future
func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/blogs", allBlogs).Methods("GET")
	// myRouter.HandleFunc("/blogs", allBlogs).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	handleRequest()
}
