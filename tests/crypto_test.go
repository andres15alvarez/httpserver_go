package tests

import (
	"testing"

	"github.com/andres15alvarez/go_http_server/utils"
)

func TestHashPassword(t *testing.T) {
	password := "holahola"
	_, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("want hashed password got %q", err.Error())
	}
}

func TestCheckPassworHash(t *testing.T) {
	password := "holahola"
	hashedPassword := "$2a$14$YrUBnEALoP8ekZwT5iwUJe3euijN5JVYRanuxxIVrOy4/w0FRwriy"
	if !utils.CheckPasswordHash(password, hashedPassword) {
		t.Error("want true got false")
	}
}
