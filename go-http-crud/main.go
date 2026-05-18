package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/users", usersHandler)

	fmt.Println("Server is running at port 5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of users")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server again")
}
