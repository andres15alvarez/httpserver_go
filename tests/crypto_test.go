package tests

import (
	"testing"

	"github.com/andres15alvarez/go_http_server/utils"
)

func Test_hashPassword(t *testing.T) {
	password := "holahola"
	_, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("want hashed password got %q", err.Error())
	}
}
