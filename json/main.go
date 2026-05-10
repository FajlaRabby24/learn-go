// json
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"peronName"` // struct tag
	Age  int    `json:"-"`         // hide field
	City string `json:"city"`
}

func main() {
	p := person{
		Name: "John",
		Age:  20,
		City: "New York",
	}

	rawJson, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(rawJson))
}
