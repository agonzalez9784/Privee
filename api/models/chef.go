package models

import (
	"database/sql"
	"log"
)

type Chef struct {
	chefID, userID      string
	firstName, lastName string
	profilePhoto        string
	description         string
	ordersFulfilled     int
}

func CreateChef(chefID string, userID string, profilePhoto string, firstName string, lastName string, description string, ordersFulfilled int) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "Chef" (chefID, userID, profilePhoto, firstName, lastName, description, ordersFulfilled) VALUES ($1, $2, $3, $4, $5, $6, $7);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(chefID, userID, profilePhoto, firstName, lastName, description, ordersFulfilled)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetChefsByLocation(city string, state string) ([]Chef, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT "Chef".*, "User"."city", "User"."location" FROM "Chef" INNER JOIN "User" ON "Chef"."userid" = "User"."userid" WHERE "User"."city" = $1 AND "User"."location" = $2;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var chefs []Chef

	rows, err := stmt.Query(city, state)

	for rows.Next() {
		var chef Chef
		if err := rows.Scan(&chef.chefID, &chef.userID, &chef.profilePhoto, &chef.firstName, &chef.lastName, &chef.description, &chef.ordersFulfilled); err != nil {
			log.Fatal(err)
		}
		chefs = append(chefs, chef)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return chefs, nil
}

func GetChef(chefID string) (Chef, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "Chef" WHERE chefID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var chef Chef

	err = stmt.QueryRow(chefID).Scan(&chef.chefID, &chef.userID, &chef.profilePhoto, &chef.firstName, &chef.lastName, &chef.description, &chef.ordersFulfilled)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return chef, nil
}

func UpdateChef(chefID string, profilePhoto string, firstName string, lastName string, description string, ordersFulfilled int) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "Chef" SET profilePhoto = $1, firstName = $2, lastName = $3, description = $4, ordersFulfilled = $5 WHERE chefID = $6∂∂`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(profilePhoto, firstName, lastName, description, ordersFulfilled, chefID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteChef(chefID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "Chef" WHERE chefID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(chefID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
