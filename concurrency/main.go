package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var fileUrl string

func main() {
	var start = time.Now()

	// * handle manually
	// wg.Add(1)
	// go uploadFile()
	// wg.Add(1)
	// go saveToDB()
	// wg.Add(1)
	// go sendEmail()

	// * handle go routine
	wg.Go(uploadFile)
	wg.Go(saveToDB)
	wg.Go(sendEmail)

	// time.Sleep(3 * time.Second)

	wg.Wait() // waiting .... until couter is 0

	fmt.Println("File url", fileUrl)
	fmt.Println("All task complete in", time.Since(start))
}

func uploadFile() {
	// defer wg.Done()
	fmt.Println("Uploading file...")
	time.Sleep(3 * time.Second)
	fmt.Println("File uploaded done")
	fileUrl = "https://cloudinary.com"
	// wg.Add(-1)
	// wg.Done()
}
func saveToDB() {
	// defer wg.Done()
	fmt.Println("Saving file...")
	time.Sleep(1 * time.Second)
	fmt.Println("File saved done")
	// wg.Add(-1)
	// wg.Done()
}
func sendEmail() {
	// defer wg.Done()
	fmt.Println("Sending file...")
	time.Sleep(2 * time.Second)
	fmt.Println("File sent done")
	// wg.Add(-1)
	// wg.Done()
}
