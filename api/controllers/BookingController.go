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

func AcceptBooking(bookingID string) {

	booking, err := models.GetBooking(bookingID)

	if err != nil {
		log.Fatal(err)
	}

	models.UpdateBooking(bookingID, booking.CreatedAt, "accepted", booking.UserID, booking.ChefID, booking.StartTime, booking.EndTime, booking.Rate)
}

func ConfirmArrival(bookingID string) {
	booking, err := models.GetBooking(bookingID)

	if err != nil {
		log.Fatal(err)
	}

	models.UpdateBooking(bookingID, booking.CreatedAt, "in progress", booking.UserID, booking.ChefID, booking.StartTime, booking.EndTime, booking.Rate)
}

func DeclineBooking(bookingID string) {

	booking, err := models.GetBooking(bookingID)

	if err != nil {
		log.Fatal(err)
	}

	models.UpdateBooking(bookingID, booking.CreatedAt, "declined", booking.UserID, booking.ChefID, booking.StartTime, booking.EndTime, booking.Rate)
}

func noShow(bookingID string) {
	booking, err := models.GetBooking(bookingID)

	if err != nil {
		log.Fatal(err)
	}

	models.UpdateBooking(bookingID, booking.CreatedAt, "no show", booking.UserID, booking.ChefID, booking.StartTime, booking.EndTime, booking.Rate)
}
