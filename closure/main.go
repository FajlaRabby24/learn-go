package main

import "fmt"

func multiflyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func makeCounter() func() int {
	count := 0

	inner := func() int {
		count++
		fmt.Println(&count)
		return count
	}

	return inner
}

func main() {
	double := multiflyBy(2)
	fmt.Println(double(3))
	fmt.Println(double(5))

	// closure
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
