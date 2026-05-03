package main

func makeCoffee(i int) {
	println("Making coffee", i)
}

func main() {
	// for i := 1; i <= 5; i++ { // for loop
	// 	makeCoffee(i)
	// }

	// i := 1
	// for i <= 5 { // while loop
	// 	makeCoffee(i)
	// 	i++
	// }

	// for { // infinite loop
	// 	makeCoffee(i)
	// }

	// for i := 1; i <= 10; i++ {  // break
	// 	makeCoffee(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	for i := 1; i <= 10; i++ { // continue
		if i%2 != 0 {
			continue
		}
		makeCoffee(i)
	}
}
