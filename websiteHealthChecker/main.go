package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkWebsiteUrl(url string, ch chan Result) {
	res, err := http.Get(url)

	if err != nil {
		ch <- Result{
			url:    url,
			status: "is down",
			err:    err,
		}
		return
	}

	ch <- Result{
		url:    url,
		status: "is up and running",
		err:    nil,
	}
	defer res.Body.Close()
}

type Result struct {
	url    string
	status string
	err    error
}

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.slashdot.org/",
		"https://www.udemy.com/",
		"https://www.wrong.com/",
		"https://cinetube-client-mu.vercel.app/",
		"https://medistore-client-omega.vercel.app/",
	}

	ch := make(chan Result)

	start := time.Now()

	for _, url := range urls {
		// checkWebsiteUrl(url)
		go checkWebsiteUrl(url, ch)

	}

	for range urls {
		result := <-ch
		if result.err != nil {
			fmt.Println(result.url, result.err, "Error")
			continue
		}

		fmt.Println(result.url, result.status)
	}

	fmt.Println("Took", time.Since(start))
	fmt.Println("All url")
}
