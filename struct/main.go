package main

import "fmt"

// type additionalInfo struct {
// 	phone   int
// 	address string
// }

// type user struct {
// 	name     string
// 	email    string
// 	metaInfo additionalInfo // embedding struct
// }

type user struct {
	name string
	age  int
	role string
}

func main() {
	// john := user{name: "John", email: "john@gmail.com"}

	// fmt.Printf("%+v", john)

	// john.email = "john@go.dev"

	// fmt.Println(john.name)
	// fmt.Println(john.email)

	// var user1 user

	// user1.name = "John"
	// user1.email = "john@gmail.com"
	// fmt.Println(user1)

	// john := user{name: "John", email: "john@gmail.com", metaInfo: additionalInfo{
	// 	phone: 123456, address: "Dhaka",
	// }}

	// fmt.Println(john)
	// fmt.Printf("%+v", john) // %+v -> print all struct fields

	newUser := func(name string, age int, role string) user {
		if age <= 0 {
			age = 18
		}

		return user{
			name: name,
			age:  age,
			role: role,
		}

	}

	jamal := newUser("Fajla", 0, "Student")
	fmt.Print(jamal)
}

// crating struct type
// Different ways to create struct
// using printf to print struct (%+v)
// updating struct fields
// Accessing struct fields

// embedding struct
// constructor function
