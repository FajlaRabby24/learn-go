package main

import "fmt"

// func makeCoffee() { // outside main function right
// }

// * init -> First run before main
// func init(x int) string { // ! func init must have no arguments and no return values
func init() {
	fmt.Println("This is init function for calling db...")
}

func test() {
	fmt.Println("this is test function")
}

func main() {
	//  func makeCoffee () {  // ! inside main function wrong
	//  }

	// * anonymous function
	// coffeeBanao := func() {
	// 	fmt.Printf("Making coffee")
	// }

	// coffeeBanao()

	// * IIFE
	// func(coffeeType string) {
	// 	fmt.Printf("Making %s coffee. IIFE", coffeeType)
	// }("Latte")

	makeCoffee := func() {
		coffee := "Cappuccino"
		fmt.Printf("Making %s", coffee)
	}

	makeCoffee()
}
