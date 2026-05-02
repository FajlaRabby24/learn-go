package main

import "fmt"

func makeCoffee(kind string, tk int) (coffee string, price int) { // named return value
	// coffee := fmt.Sprintf("%s ", kind)
	// return coffee, price

	coffee = fmt.Sprintf("%s", kind)
	price = tk

	return
}

func main() {
	myCoffee, mybill := makeCoffee("espresso", 25)
	myCoffee2, mybill2 := makeCoffee("black", 20)
	fmt.Printf("Your %s coffee is ready ,Your bill is %d", myCoffee, mybill)
	fmt.Println()
	fmt.Printf("Your %s coffee is ready ,Your bill is %d", myCoffee2, mybill2)
}
