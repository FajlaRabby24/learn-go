package main

import "fmt"

func main() {
	sugar := 2

	makeCoffee := func() {
		coffee := "Cappuccino"
		sugar := 3 // modify
		fmt.Printf("Making %s with %d spoon of sugar \n", coffee, sugar)
		fmt.Println("Value of inner sugar", sugar)

	}

	makeCoffee()

	fmt.Println("Value of outer sugar", sugar)
}
