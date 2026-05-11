package main

import (
	"encoding/json"
	"fmt"
)

func callbackFn(fn func()) {
	fn()
}

func calculate(num1, num2 int, operation func(x, y int) int) (result int) {
	result = operation(num1, num2)
	return result
}

func filter(nums []int, cb func(nums []int) []int) []int {
	result := cb(nums)
	return result
}

func makeCounter() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}

func makeMultiplier(x int) func(n int) int {
	return func(n int) int {
		return x * n
	}
}

func makeGreeter(greeting string) func(name string) string {
	return func(name string) string {
		return fmt.Sprintf("%s, %s!", greeting, name)
	}
}

func deferFn() {
	fmt.Println("First")
	defer fmt.Println("Second")
	fmt.Println("Third")
}

func deferCleanup() {
	fmt.Println("Opening file...")
	fmt.Println("Working with file...")
	defer fmt.Println("Closing file...")
}

func mapSlice(nums []int, cb func(num int) int) []int {
	var newSlice []int
	for _, element := range nums {
		newElement := cb(element)
		newSlice = append(newSlice, newElement)
	}
	return newSlice
}

func reduce(nums []int) int {
	var value int
	for _, item := range nums {
		value += item
	}

	return value
}

type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	//* Problem 1 — Simple Callback
	callback := func() {
		println("Hello from callback!")
	}
	callbackFn(callback)

	// //* Problem 2 — Math Callback
	add := func(x, y int) int {
		return x + y
	}
	multiply := func(x, y int) int {
		return x * y
	}
	fmt.Println("Result =", calculate(10, 5, add))
	fmt.Println("Result =", calculate(10, 5, multiply))

	// //* Problem 3 — Filter with Callback
	filterCB := func(nums []int) []int {
		var result []int
		for _, value := range nums {
			if value%2 == 0 {
				result = append(result, value)
			}
		}

		return result
	}
	fmt.Println(filter([]int{1, 2, 3, 4, 5, 6}, filterCB))

	// //* Problem 4 — Counter
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// //* Problem 5 — Multiplier
	double := makeMultiplier(2)
	fmt.Println(double(5))
	triple := makeMultiplier(3)
	fmt.Println(triple(5))

	// //* Problem 6 — Greeting Generator
	hello := makeGreeter("Hello")
	bonjour := makeGreeter("Bonjour")
	fmt.Println(hello("Rahim"))
	fmt.Println(bonjour("Kahim"))

	// //* Problem 7 — Simple Defer
	deferFn()

	// //* Problem 8 — Defer in Loop
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}

	// //* Problem 9 — Defer for Cleanup
	deferCleanup()

	//* Problem 10 — Days Enum
	// fmt.Println(Weekday. , Monday) // ! incomplete

	//* Problem 11 — Size Enum
	// ! incomplete

	//* Problem 12 — Map Function
	makeDouble := func(num int) int {
		return num * 2
	}
	fmt.Println(mapSlice([]int{1, 2, 3, 4, 5}, makeDouble))

	//* Problem 13 — Reduce Function
	fmt.Println("sum =", reduce([]int{1, 2, 3, 4, 5}))

	//* Problem 14 — Struct to JSON
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}
	person := Person{"Rahim", 25, "Dhaka"}
	rawJson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(rawJson))

	//* Problem 15 — JSON to Struct
	// ! incomplete

}
