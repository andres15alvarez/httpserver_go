package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andres15alvarez/go_http_server/models"
)

var users []models.User = []models.User{
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
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
