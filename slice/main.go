package main

import "fmt"

func main() {
	// var numbers = [6]int{1, 2, 3}               // partial array [1, 2, 3, 0, 0, 0]
	// var orders = [6]int{10, 23, 45, 43, 23, 21} // slice

	// // slice := orders[2:5] // [45, 43, 23]
	// slice := orders[:3] // [10, 23, 45]
	// // slice := orders[1:] //  [23, 45, 43, 23, 21]
	// slice := orders[:] // [10, 23, 45, 43, 23, 21] // *len, cap, pointer
	// // fmt.Println(slice)
	// // slice[0] = 100

	// fmt.Println("orders", orders)

	// slice = append(slice, 100) // [10, 23, 45, 100]
	// // slice = append(slice, 200, 300)
	// slice = append(slice, 200)
	// slice = append(slice, 400)
	// slice = append(slice, 500)
	// slice = append(slice, 600)

	// fmt.Println("slice", slice)
	// fmt.Println("orders", orders)

	// fmt.Println("ths length of the slice is ", len(slice))
	// fmt.Println("ths cap of the slick is ", cap(slice))

	var slice = []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)      // [1 2 3 4]
	fmt.Println(len(slice)) // 4
	fmt.Println(cap(slice)) //6
}
