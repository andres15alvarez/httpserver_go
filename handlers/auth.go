package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andres15alvarez/go_http_server/models"
	"github.com/andres15alvarez/go_http_server/repositories"
	"github.com/andres15alvarez/go_http_server/utils"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.SignIn
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	userFound, err := repositories.GetUserByEmail(user.Email)
	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !utils.CheckPasswordHash(user.Password, userFound.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := utils.GenerateToken(userFound)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.Token{Token: token})
}
