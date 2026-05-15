package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(6, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}
