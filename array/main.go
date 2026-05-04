package main

import "fmt"

func main() {
	var numbers [6]int
	fmt.Println(numbers)

	numbers[1] = 10
	fmt.Println(numbers)

	fmt.Println(len(numbers))
	fmt.Println(numbers[0])

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
