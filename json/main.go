// json
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string                `json:"personName"` // struct tag
	Age  int    /*`json:"-"`*/ // hide field
	City string                `json:"city"`
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

	var p2 person
	jsonText := `{"personName":"John","city":"New York"}`
	error := json.Unmarshal([]byte(jsonText), &p2)
	if error != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(p2.Name)
	fmt.Println(p2.Age)
	fmt.Println(p2.City)
}
