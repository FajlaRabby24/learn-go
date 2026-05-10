package main

import "fmt"

func deferred(result int) {
	fmt.Println("deferred result:", result)
}

func example() int {
	result := 10

	defer deferred(result)

	fmt.Println("I am from example fn:", result)

	result = 100
	return result
}

func main() {
	// defer fmt.Println("I am from deferred print call")

	// fmt.Println("I am from main function")

	example()
}
