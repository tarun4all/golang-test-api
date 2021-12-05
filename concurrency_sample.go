package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Concurrency Example")

	// call syncronous
	// compute(5)
	// compute(5)

	// call async

	go compute(5)
	go compute(5)

	fmt.Scanln()
}
