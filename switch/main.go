package main

func main() {
	day := "sat"

	// if day == "sat"{
	// 	println("Today is saturday")
	// } else {
	// 	println("Today is sunday")
	// }

	// switch { // basic switch
	// case day == "sat":
	// 	println("Today is saturday")
	// case day == "sun":
	// 	println("Today is sunday")
	// default:
	// 	println("Today is not saturday or sunday")
	// }

	switch day { // tagged switch
	case "sat":
		println("Today is saturday")
	case "sun":
		println("Today is sunday")
	default:
		println("Today is not saturday or sunday")
	}
}
