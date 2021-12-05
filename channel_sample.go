package main

import (
	"fmt"
	"time"
)

func SendValue(c chan string, num int) {
	fmt.Println("Executing go routine")
	time.Sleep(1 * time.Second)
	if num == 1 {
		c <- "one"
	} else {
		c <- "other than one"
	}
	fmt.Println("Finished go routine")

}

func main() {
	fmt.Println("Channel example")
	values := make(chan string, 2)

	defer close(values)

	// if we pass single values to 2 go routines then only one executes one blcoked
	go SendValue(values, 1)
	go SendValue(values, 2)

	value := <-values

	fmt.Println(value)
	time.Sleep(1 * time.Second)
}
