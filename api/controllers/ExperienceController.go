package controllers

import (
	"api/models"
	"api/utils"
)

func AddExperience(chefID string, jobTitle string, companyName string, startDate string, endDate string, description string) {
	experienceID := utils.GenerateID(15)
	models.CreateExperience(experienceID, chefID, jobTitle, companyName, startDate, endDate, description)
}
