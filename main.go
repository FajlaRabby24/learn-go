package main

import "fmt"

// * data types
// int, int8 int16 int32 int64
// uint, uint8 uint16 uint32 uint64
// float32 float64
// string
// bool

// isAdmin := true  // ! wrong way

func main() {

	// name := "fajla" // ? most used in real projects
	// name string := "Fajla" // ! wrong way

	// var name string = "Fajla"
	// var name = "Fajla"

	// * grouped valiables
	// var (
	// 	firstName = "Fajla"
	// 	lastName  = "Rabby"
	// )

	// * multiple variables declaration
	// var x, y int

	// x = 25
	// y = 30

	// var x, y int = 25, 30

	// isAdmin := true

	// const pi = 3.1416
	// pi = 12123 // !can't be changed

	// user_name := "elma"
	// user-name := "elma" // ! wrong
	// 1user_name := "elma" // ! wrong

	// fmt.Println(name)
	// fmt.Println(firstName, lastName)
	// fmt.Println(x, y)
	// fmt.Println(isAdmin)
	// fmt.Println(user_name)

	// var age int
	// fmt.Println(age) // 0

	// var name string
	// fmt.Println(name) // ""

	var is_admin bool
	fmt.Println(is_admin) // false

	var score float64
	fmt.Println(score) // 0

	var name string = "Fajla Rabby"
	var age int = 25
	var rating = 4.43
	// fmt.Println("My name is", name)
	// fmt.Printf("My name is %s and I am %d years old. and my rating is %.2f", name, age, rating)
	formattedString := fmt.Sprintf("My name is %s and I am %d years old. and my rating is %.2f", name, age, rating)
	fmt.Println(formattedString)
	myCoffe := makeCoffee("espresso", true)
	fmt.Println(myCoffe)
}

func makeCoffee(kind string, isSuger bool) string {
	// fmt.Printf("Making %s coffee...\n", kind)
	// fmt.Println("Suger added", isSuger)

	coffee := fmt.Sprintf("%s coffee! And Sugar added: %t", kind, isSuger)

	return coffee
}
