package main

import "fmt"

// type user struct {
// 	name       string
// 	age        int
// 	isloggedIn bool
// 	greet      func()
// }

type user struct {
	name       string
	age        int
	isloggedIn bool
}

func main() {
	// user1 := user{
	// 	name:       "John",
	// 	age:        20,
	// 	isloggedIn: false,
	// }

	// user1.greet = func() {
	// 	println("Hello", user1.name)
	// }

	// user1.greet()

	user1 := user{
		name:       "John",
		age:        20,
		isloggedIn: false,
	}
	// user2 := user{
	// 	name:       "Doe",
	// 	age:        20,
	// 	isloggedIn: false,
	// }

	// user1.email = "john@gmail"  // ! wrong

	user1.greet()
	// user2.greet()
	// pointerUser1 := &user1
	// pointerUser1.login()
	user1.login()

	fmt.Printf("%+v", user1)
}

func (u *user) login( /*password string*/ ) {
	fmt.Println("login called")
	// u.isloggedIn = true
	(*u).isloggedIn = true
}

//* receiver function
func (u user) greet() {
	println("Hello!", u.name)
}
