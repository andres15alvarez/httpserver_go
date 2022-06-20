package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/andres15alvarez/go_http_server/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Home{Data: "Hello, World"})
}
