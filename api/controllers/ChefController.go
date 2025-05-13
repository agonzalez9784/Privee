package controllers

import (
	"api/models"
	"api/utils"
	"log"
)

func GetChefs(city string, state string) ([]models.Chef, error) {
	chefs, err := models.GetChefsByLocation(city, state)

	if err != nil {
		log.Fatal(err)
	}

	return chefs, err
}

func RegisterChefAccount(userID, description) error{
	
	user, err := models.GetUser(userID)

	if err != nil {
		log.Fatal(err)
	}

	chefID := utils.GenerateID(15)
	
	models.CreateChef(chefID, user.UserID, "", , , description, 0) 

	return err
}