package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

// go get github.com/googollee/go-socket.io
func main() {
	fmt.Println("STarting server...")

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	server.On("Connection", func(so socketio.Socket) {
		log.Println("New Connection ")
	})

	http.Handle("/socket.io", server)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
