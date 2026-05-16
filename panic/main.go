package main

import (
	"fmt"
	"log"
)

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

func doAnotherSomething() {
	defer func() {
		fmt.Println("Defered func ran")
	}()

	log.Fatal("Something very big happend")
}

func main() {
	// dosomthing()
	doAnotherSomething()

	fmt.Println("Main completed normally")
}
