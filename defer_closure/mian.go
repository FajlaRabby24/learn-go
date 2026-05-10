package main

import "fmt"

func example() int {
	result := 10

	defer func() {
		result += 100
		fmt.Println("deferred result:", result)
	}()

	fmt.Println("I am from example fn:", result)

	result += 100
	return result
}

func namedExample() (result int) {
	result = 10

	defer func() {
		result += 100
		fmt.Println("deferred result:", result)
	}()

	fmt.Println("I am from example fn:", result)

	result += 100
	return result
}

func connectDB() {
	defer fmt.Println("closing db connection")
	fmt.Println("connecting to db...")
}

func main() {
	fmt.Println("return result", example())
	fmt.Println("-----------------")
	fmt.Println("named return result", namedExample())

	// defer execute by LIFO -> last in first out
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
