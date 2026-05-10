package main

import "fmt"

func calculate(a, b int, operation func(x, y int) int) int {
	result := operation(a, b)
	return result
}

func multiflyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	add := func(n1 int, n2 int) int {
		return n1 + n2
	}

	fmt.Println(calculate(10, 20, add))

	double := multiflyBy(2)
	fmt.Println(double(3))
	fmt.Println(double(5))
}
