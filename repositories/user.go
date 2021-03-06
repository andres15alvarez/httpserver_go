package repositories

import (
	"errors"
	"log"

	"github.com/andres15alvarez/go_http_server/models"
	"github.com/andres15alvarez/go_http_server/utils"
)

func GetUserByEmail(email string) (models.User, error) {
	db := utils.ConnectDB()
	defer db.Close()
	var user models.User
	err := db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		log.Printf(err.Error())
		return models.User{}, errors.New(err.Error())
	}
	return user, nil
}

func ListUsers() ([]models.UserCreated, error) {
	db := utils.ConnectDB()
	defer db.Close()
	var users []models.UserCreated
	err := db.Model(&users).Column("id", "name", "email").Select()
	if err != nil {
		log.Printf(err.Error())
		return []models.UserCreated{}, errors.New(err.Error())
	}
	return users, nil
}

func CreateUser(user models.User) (models.UserCreated, error) {
	db := utils.ConnectDB()
	defer db.Close()
	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf(err.Error())
		return models.UserCreated{}, errors.New(err.Error())
	}
	user.Password = passwordHashed
	_, insertErr := db.Model(&user).Insert()
	if insertErr != nil {
		log.Printf(insertErr.Error())
		return models.UserCreated{}, errors.New(insertErr.Error())
	}
	return models.UserCreated{UserBase: user.UserBase}, nil
}
