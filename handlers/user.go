package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andres15alvarez/go_http_server/models"
	"github.com/andres15alvarez/go_http_server/repositories"
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

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usersDB, err := repositories.ListUsers()
	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(usersDB)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	userCreated, err := repositories.CreateUser(user)
	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userCreated)
}
