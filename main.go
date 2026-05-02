// package main

// import "fmt"

// func makeCoffee(kind string, tk int) (coffee string, price int) { // named return value
// 	// coffee := fmt.Sprintf("%s ", kind)
// 	// return coffee, price

// 	coffee = fmt.Sprintf("%s", kind)
// 	price = tk

// 	return
// }

// func main() {
// 	myCoffee, mybill := makeCoffee("espresso", 25)
// 	myCoffee2, mybill2 := makeCoffee("black", 20)
// 	fmt.Printf("Your %s coffee is ready ,Your bill is %d", myCoffee, mybill)
// 	fmt.Println()
// 	fmt.Printf("Your %s coffee is ready ,Your bill is %d", myCoffee2, mybill2)
// }

package main

import (
	"fmt"
)

func coffeOrder(cus_name string, coffee_type string, price int) string {
	order := fmt.Sprintf("Order for %s: %s coffee costs %d taka", cus_name, coffee_type, price)
	return order
}

func sum(num1 int, num2 int) string {
	sum := num1 + num2
	result := fmt.Sprintf("Sum is %d", sum)
	return result
}

func adultOrMinor(age int) string {
	result := age > 17
	if result {
		return "Adult"
	}
	return "Minor"
}

func largestNum(a int, b int, c int) string {
	if a > b && a > c {
		return fmt.Sprintf("Largest is %d", a)
	} else if b > a && b > c {
		return fmt.Sprintf("Largest is %d", b)
	} else {
		return fmt.Sprintf("Largest is %d", c)
	}
}

func ractangleArea(width int, height int) string {
	multiply := width * height
	ractangle := fmt.Sprintf("Area is %d", multiply)
	return ractangle
}

func evenOrOdd(num int) string {
	value := num%2 == 0
	if value {
		return fmt.Sprintf("%d is Even", num)
	}
	return fmt.Sprintf("%d is Odd", num)
}

func calculator(num1 int, operator string, num2 int) string {
	switch operator {
	case "+":
		sum := num1 + num2
		return fmt.Sprintf("Result: %d", sum)
	case "-":
		minus := num1 - num2
		return fmt.Sprintf("Result: %d", minus)
	case "*":
		mult := num1 * num2
		return fmt.Sprintf("Result: %d", mult)
	case "/":
		if num2 == 0 {
			return "Error: cannot divide by zero"
		}
		div := num1 / num2
		return fmt.Sprintf("Result: %d", div)
	}

	return "Invalid operator"
}

func main() {
	order := coffeOrder("Fajla", "Latte", 150)
	sum := sum(10, 20)
	adultOrMinor := adultOrMinor(18)
	ractangleArea := ractangleArea(5, 5)
	evenOrOdd := evenOrOdd(7)
	largestNum := largestNum(10, 45, 88)
	calculator := calculator(0, "/", 5)

	fmt.Println(order)
	fmt.Println(sum)
	fmt.Println(adultOrMinor)
	fmt.Println(ractangleArea)
	fmt.Println(evenOrOdd)
	fmt.Println(largestNum)
	fmt.Println(calculator)
}
