package repositories

import (
	"errors"
	"log"

	"github.com/andres15alvarez/go_http_server/models"
	"github.com/andres15alvarez/go_http_server/utils"
)

func ListUsers() ([]models.User, error) {
	db := utils.ConnectDB()
	defer db.Close()
	var users []models.User
	err := db.Model(&users).Select()
	if err != nil {
		log.Printf(err.Error())
		return []models.User{}, errors.New(err.Error())
	}
	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	db := utils.ConnectDB()
	defer db.Close()
	_, err := db.Model(&user).Insert()
	if err != nil {
		log.Printf(err.Error())
		return models.User{}, errors.New(err.Error())
	}
	return user, nil
}
