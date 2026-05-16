//*  Comma OK Idiom & Assertion Handling

package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	message string
	code    int
}

func (cu *CustomError) Error() string {
	return cu.message
}

func login(password string) error {
	if password != "1234" {
		// return &CustomError{
		// 	message: "Password do not match",
		// 	code:    401,
		// }

		return errors.New("Password do not match")
	}
	return nil
}

func main() {
	err := login("2345")
	if err != nil {
		// customErr, ok := err.(*CustomError)
		// fmt.Println(customErr) // by default it's print message
		// fmt.Println(customErr.message)
		// fmt.Println(customErr.code)

		// if !ok {
		// 	fmt.Println(customErr)
		// } else {
		// 	fmt.Println(customErr.message)
		// }

		if customErr, ok := err.(*CustomError); ok {
			fmt.Println(customErr.code)
		}

		// fmt.Println("Error", err, "Code", err.code)
	}

	users := map[int]string{
		1: "John",
		2: "Jane",
		3: "Jack",
	}

	name, ok := users[3]
	if ok {
		fmt.Println(name)
	}

	fmt.Println("Main ends")
}
