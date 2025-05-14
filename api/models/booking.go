package models

import (
	"database/sql"
	"log"
)

type Booking struct {
	BookingID string
	CreatedAt string
	Status    string
	UserID    string
	ChefID    string
	StartTime string
	EndTime   string
	Rate      float64
}

func CreateBooking(bookingID string, createdAt string, status string, userID string, chefID string, startTime string, endTime string, rate float64) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "Booking" (bookingID, createdAt, status, userID, chefID, startID, endTime, rate) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(bookingID, createdAt, status, userID, chefID, startTime, endTime, rate)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetBooking(bookingID string) (Booking, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "Booking" WHERE bookingID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var booking Booking

	err = stmt.QueryRow(bookingID).Scan(&booking.BookingID, &booking.CreatedAt, &booking.Status, &booking.UserID, &booking.ChefID, &booking.StartTime, &booking.EndTime, &booking.Rate)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return booking, nil
}

func UpdateBooking(bookingID string, createdAt string, status string, userID string, chefID string, startTime string, endTime string, rate float64) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "Booking" SET createdAt = $1, status = $2, userID = $3, chefID = $4, startTime = $5, endTime = $6, rate = $7 WHERE bookingID = $8`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(createdAt, status, userID, chefID, startTime, endTime, rate, bookingID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteBooking(bookingID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "Booking" WHERE bookingID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(bookingID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
