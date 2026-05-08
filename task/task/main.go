package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {
	//* Problem 1 — Array Sum
	arr := [5]int{10, 20, 30, 40, 50}
	var sum int
	for _, value := range arr {
		sum += value
	}
	fmt.Println(sum)

	//* Problem 2 — Slice Append
	names := []string{"Rahim", "Karim", "Jamal"}
	names = append(names, "Fajla")
	fmt.Println(names)

	//* Problem 3 — Even Numbers
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNums []int
	for _, value := range numbers {
		if value%2 == 0 {
			evenNums = append(evenNums, value)
		}
	}
	fmt.Println(evenNums)

	//* Problem 4 — Map Print
	myMap := map[string]int{"Rahim": 25, "Karim": 30, "Fajla": 22}
	for key, value := range myMap {
		fmt.Println(key, "→", value)
	}

	//* Problem 5 — Map Lookup
	mapLookUp := map[string]string{"bangladesh": "Dhaka", "japan": "Tokyo"}
	var userInput string
	fmt.Printf("Write your country name: ")
	fmt.Scan(&userInput)
	result := mapLookUp[userInput]
	if result == "" {
		fmt.Println("পাওয়া যায়নি")
	} else {
		fmt.Println(result)
	}

	//* Problem 6 — Divisible by 3
	var divisibleBy3 []int
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			divisibleBy3 = append(divisibleBy3, i)
		}
	}
	fmt.Println(divisibleBy3)

	//* Problem 7 — Range with Index
	fruitsName := []string{"apple", "banana", "mango"}
	for i, value := range fruitsName {
		fmt.Println(i, "→", value)
	}

	//* Problem 8 — Positive Negative Zero
	var inputNum int
	fmt.Printf("Input you number: ")
	fmt.Scan(&inputNum)
	if inputNum == 0 {
		fmt.Println("Zero")
	} else if inputNum > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}

	//* Problem 9 — Weekday Name
	var day int
	fmt.Printf("Input you day: ")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("Wednesday")

	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	//* Problem 10 — Grade Calculator
	var marks int
	fmt.Printf("Input you marks: ")
	fmt.Scan(&marks)
	if marks >= 90 && marks <= 100 {
		fmt.Println("A")
	} else if marks >= 80 {
		fmt.Println("B")
	} else if marks >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}

	//* Problem 11 — Student Info
	type Student struct {
		name  string
		age   int
		grade string
	}
	students := []Student{{"Rahim", 20, "A"}, {"Karim", 22, "B"}, {"Fajla", 21, "A+"}}
	for _, student := range students {
		fmt.Println("Name:", student.name, "|", "Age:", student.age, "|", "Grade:", student.grade)
	}

	//* Problem 12 — Rectangle Area
	Area := func(values Rectangle) int {
		return values.width * values.height
	}
	area := Rectangle{width: 5, height: 10}
	fmt.Println(Area(area))

	//* Problem 13 — Shape Interface
	circle := Circle{Radius: 7}
	area2 := Rectangle{width: 4, height: 5}

	getArea(circle)
	getArea(area2)

	//* Problem 14 — Max & Min
	randomNums := []int{3, 67, 12, 45, 8, 99, 23}
	var maxNum int
	minNum := randomNums[0]
	for _, value := range randomNums {
		if value > maxNum {
			maxNum = value
		}
		if value < minNum {
			minNum = value
		}
	}
	fmt.Println(maxNum)
	fmt.Println(minNum)

	//* Problem 15 — Adult Filter
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "John", Age: 12},
		{Name: "Jane", Age: 25},
		{Name: "Bob", Age: 40},
		{Name: "Alice", Age: 13}}

	adults := []Person{}
	for _, person := range people {
		if person.Age >= 18 {
			adults = append(adults, person)
		}
	}
	fmt.Println(adults)

}

//* Problem 13 — Shape Interface
type Shape interface {
	Area()
}

func getArea(s Shape) {
	s.Area()
}

type Circle struct {
	Radius int
}

func (c Circle) Area() {
	pi := 3.14
	// return pi * float64(c.Radius) * float64(c.Radius)
	fmt.Println(pi * float64(c.Radius) * float64(c.Radius))

}

func (a Rectangle) Area() {
	// return a.width * a.height
	fmt.Println(a.width * a.height)
}
