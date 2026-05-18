package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
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

	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintln(w, "User created")

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalide request body")
		return
	}

	newUser.Id = len(users) + 1
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(newUser)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server again")
}
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	// users, _ := json.Marshal(users)
	// w.Write(users)
	encoder := json.NewEncoder(w)
	encoder.Encode(users)
}
