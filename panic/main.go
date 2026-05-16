package main

import "fmt"

func dosomthing() {
	defer func() {
		fmt.Println("Defered func ran")
		r := recover()
		if r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	panic("Something really bad happend")
}

func main() {
	dosomthing()

	fmt.Println("Main completed normally")
}
