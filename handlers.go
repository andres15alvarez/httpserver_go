package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users []User = []User{
	{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	},
	{
		Name:  "Bob",
		Email: "bob@example.com",
	},
	{
		Name:  "Alice",
		Email: "alice@example.com",
	},
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Home{"Hello, World"})
}

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
