package controllers

import (
	"api/models"
	"api/utils"
	"log"
)

func BookAChef(userID string, chefID string, startTime string, endTime string, rate float64) (models.Booking, error) {

	bookingID := utils.GenerateID(15)
	createdAt := utils.GetCurrentTimeStamp()
	models.CreateBooking(bookingID, createdAt, "pending", userID, chefID, startTime, endTime, rate)

	booking, err := models.GetBooking(bookingID)

	if err != nil {
		log.Fatal(err)
	}

	return booking, nil
}

func AcceptBooking() {

	models.UpdateBooking()
}

func DeclineBooking() {

	models.UpdateBooking()
}
