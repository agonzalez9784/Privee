package models

import (
	"database/sql"
	"log"
)

/*
CREATE TABLE IF NOT EXISTS "Review" (
    reviewID VARCHAR(15) NOT NULL UNIQUE,
    chefID VARCHAR(15) NOT NULL,
    userID VARCHAR(15) NOT NULL,
    stars INT NOT NULL,
    review TEXT
);
*/

type Review struct {
	reviewID string
	chefID   string
	userID   string
	stars    int
	review   string
}

func CreateReview(reviewID string, chefID string, userID string, stars int, review string) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "Review" (reviewID, chefID, userID, stars, review) VALUES ($1, $2, $3, $4, $5);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(reviewID, chefID, userID, stars, review)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetReview(reviewID string) (Review, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "Review" WHERE reviewID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var review Review

	err = stmt.QueryRow(reviewID).Scan(&review.reviewID, &review.chefID, &review.userID, &review.stars, &review.review)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return review, nil
}

func UpdateReview(reviewID string, chefID string, userID string, stars int, review string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "Review" SET chefID = $1, userID = $2, stars = $3, review = $4 WHERE reviewID = $5;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(chefID, userID, stars, review, reviewID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteReview(reviewID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "Review" WHERE reviewID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(reviewID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
