package main

import (
	"encoding/json"
	"fmt"
)

type SendorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	jsonString := `{"name": "battery", "capacity": 40, "time": "2021-12-03T05:15:47Z"}`

	var read SendorReading
	// var read map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &read)

	fmt.Printf("%+v\n", read)
	fmt.Println(err)
}
