package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id    int
	Name  string
	Age   int
	Email string
}

var users = []User{
	{
		Id:    1,
		Name:  "Kamal",
		Age:   23,
		Email: "kamal@gmail",
	},
	{
		Id:    2,
		Name:  "Fajla",
		Age:   23,
		Email: "fajla@gmail",
	},
	{
		Id:    3,
		Name:  "Jamal",
		Age:   23,
		Email: "jamal@gmail",
	},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	// * old version
	// mux.HandleFunc("/users", usersHandler)
	// * new version
	mux.HandleFunc("POST /users", usersHandler)
	mux.HandleFunc("GET /users", getUsersHandler)

	fmt.Println("Server is running at port 5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("Server error", err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// * old version
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	fmt.Fprintln(w, "Method not allowed")
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server again")
}
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	users, _ := json.Marshal(users)
	w.Write(users)
}
