package main

import "fmt"

// func makeCoffee() { // outside main function right
// }

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
