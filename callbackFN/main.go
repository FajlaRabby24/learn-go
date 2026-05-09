package main

import "fmt"

// callback function
// first class citizen =>

func process(callback func()) {
	callback()
}

func calculate(a, b int, operation func(x, y int) int) int {
	result := operation(a, b)
	return result

}

func main() {
	greet := func() {
		fmt.Println("Hello world")
	}

	process(greet)

	add := func(n1 int, n2 int) int {
		return n1 + n2
	}
	multifly := func(n1 int, n2 int) int {
		return n1 * n2
	}

	fmt.Println(calculate(10, 20, add))
	fmt.Println(calculate(10, 20, multifly))
}
