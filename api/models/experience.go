package models

/*
CREATE TABLE IF NOT EXISTS "Experience" (
    experienceID VARCHAR(15) NOT NULL UNIQUE,
    chefID VARCHAR(15) NOT NULL,
    jobTitle VARCHAR(25) NOT NULL,
    companyName VARCHAR(25) NOT NULL,
    startDate DATE,
    endDate DATE,
    description TEXT
);
*/

import (
	"database/sql"
	"log"
)

type Experience struct {
	experienceID string
	chefID       string
	jobTitle     string
	companyName  string
	startDate    string
	endDate      string
	description  string
}

func CreateExperience(experienceID string, chefID string, jobTitle string, companyName string, startDate string, endDate string, description string) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "Experience" (experienceID, chefID, jobTitle, companyName, startDate, endDate, description) VALUES ($1, $2, $3, $4, $5, $6, $7);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(experienceID, chefID, jobTitle, companyName, startDate, endDate, description)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetExperience(experienceID string) (Experience, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "Experience" WHERE experienceID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var experience Experience

	err = stmt.QueryRow(experienceID).Scan(&experience.experienceID, &experience.chefID, &experience.jobTitle, &experience.companyName, &experience.startDate, &experience.endDate, &experience.description)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return experience, nil
}

func UpdateExperience(experienceID string, chefID string, jobTitle string, companyName string, startDate string, endDate string, description string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "Experience" SET chefID=$1, jobTitle=$2, companyName=$3, startDate=$4, endDate=$5, description=$6 WHERE conversationID = $7;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(chefID, jobTitle, companyName, startDate, endDate, description, experienceID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteExperience(experienceID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "Experience" WHERE experienceID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(experienceID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
