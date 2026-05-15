//* buffered channel

package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string, 3) // by default channel is 1

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "File uploaded"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "File url saved"
	}()

	func() {
		time.Sleep(3 * time.Second)
		ch <- "Email sent"
	}()

	for range 3 {
		data := <-ch
		fmt.Println(data)
	}
}
