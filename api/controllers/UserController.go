package controllers

import (
	"api/models"
	"api/utils"
)

func RegisterUserAccount(username string, password string, email string) {
	userID := utils.GenerateID(15)
	models.CreateUser(userID, username, password, email, "", "", false)
}
