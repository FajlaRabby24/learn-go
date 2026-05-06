package main

import "fmt"

type user struct {
	name       string
	age        int
	isloggedIn bool
}

func main() {
	// var myMap map[string]int
	// myMap := make(map[string]int)

	// myMap["userScore"] = 4
	// myMap["user1Score"] = 84

	// delete(myMap, "user1Score")

	// fmt.Println(myMap)
	// fmt.Println(myMap["userScore"])
	// fmt.Println(myMap["user1Score"])

	// myMap := map[string]string{"name": "Fajla", "success": "ok"}

	// delete(myMap, "name")
	// fmt.Println(myMap)
	// fmt.Println(myMap["success"])

	myMap := map[string]user{
		"data": {
			name:       "Fajla",
			age:        23,
			isloggedIn: true,
		},
	}

	fmt.Println(myMap)

}

// crating map (using  make and literal)
// accessing map values
// adding new key-value pair to map
// deleting key-value pair from map
// map of struct
// iterating over map
