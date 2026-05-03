package main

func main() {
	age := 20
	if age >= 18 {
		println("You are an adult")
	} else {
		println("You are not an adult")
	}

	score := 100
	println("outside if", score)

	if score := 95; score >= 90 {
		println("You got an A")
	} else if score >= 80 {
		println("You got a B")
	} else {
		println("You got a C")
	}

	// if err := saveToDb(); err != nil {
	// 	println(err)
	// }
}
