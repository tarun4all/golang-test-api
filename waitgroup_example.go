package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)

			if err != nil {
				fmt.Fprint(w, "%+v\n", err)
			}

			fmt.Fprintf(w, "%+v\n", resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
