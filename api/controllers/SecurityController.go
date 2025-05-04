package controllers

import (
	"api/models"
	"log"
)

func hasOwnership() {
	x := 0
	x = x + 1
}

func AuthorizeUser(username string, password string) bool {

	userData, err := models.GetUserByUsername(username)

	if err != nil {
		log.Fatal(err)
	}

	usr_username := userData.Username
	usr_password := userData.Password

	if (username == usr_username) && (password == usr_password) {
		return true
	}

	return false
}
