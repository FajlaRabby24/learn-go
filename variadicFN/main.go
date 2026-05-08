package main

import "fmt"

// variadic functions
// variadic parameters
// variadic arguments

// func sum(a, b, c int) int {
// 	return a + b + c
// }

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func greet(prefix string, mps ...string) {
	for _, mp := range mps {
		fmt.Println(prefix, mp)
	}
}

func main() {
	sum := sum(4, 3, 5, 456)
	fmt.Println(sum)

	// greet("Welcome", "Fajla", "kamal", "Jamal", "Samal")

	mps := []string{"Fajla", "kamal", "Jamal", "Samal"}
	// variadic arguments
	greet("Welcome", mps...)
}

// flexible amount of argument
// must be the last parameter
// internally slice
