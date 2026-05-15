// * un-buffered channel

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	var ch = make(chan string) // un-buffered channel
	go uploadFile(ch)

	wg.Wait()
	// fileUrl := <-ch // * blocking call -> wait for file url
	// fmt.Println("Ch", ch)
	// fmt.Println("File url", fileUrl)
}

func uploadFile(c chan string) {
	defer wg.Done()

	fmt.Println("Uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File uploaded done")
	fileUrl := "https://cloudinary.com"
	c <- fileUrl

}
