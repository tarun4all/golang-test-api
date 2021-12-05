package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"isDeveloper"`
}

func main() {
	book := Book{Title: "Test", Author: Author{Name: "Tarun", Age: 25, Developer: true}}
	byteArray, err := json.MarshalIndent(book, "", "   ")

	fmt.Println(string(byteArray), err)
}
