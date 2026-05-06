package main

import "fmt"

func main() {
	myMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	for _, value := range myMap {
		fmt.Println(value)
	}

	myArr := [3]string{"value1", "value2", "value3"}
	colors := [3]string{"red", "green", "blue"}

	for i, value := range myArr {
		fmt.Println(i, value)
	}

	for i, color := range colors {
		fmt.Println(i, color)
	}

	name := "nextlavel"

	var byteSlice = []byte(name)
	fmt.Println(byteSlice)

	for i, value := range name {
		fmt.Println(i, value)
	}
}

// array
// slice
// string
// map
// channel
