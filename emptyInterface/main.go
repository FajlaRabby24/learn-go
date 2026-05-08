package main

import "fmt"

// any == interface{ }
// type assertion
// ok idiom
func Process(data any) {
	// fmt.Println(data)

	// strData := data.(string) // type assertion
	// fmt.Println(strData)

	strData, ok := data.(string) // type assertion
	if ok {
		fmt.Println(len(strData))
		return
	}

	intData, ok := data.(int)
	if ok {
		fmt.Println(intData * 2)
		return
	}
}

func main() {
	// var data interface{}

	// data = "Fajla"
	// fmt.Println(data)

	// data = 26
	// fmt.Println(data)

	Process("Fajla")
	Process(26)
}
