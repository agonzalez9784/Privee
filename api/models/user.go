package models

import (
	"database/sql"
	"log"
)

/*
CREATE TABLE IF NOT EXISTS "User" (
    userID VARCHAR(15) NOT NULL UNIQUE,
    username VARCHAR(25) NOT NULL UNIQUE,
    password VARCHAR(20) NOT NULL,
    email VARCHAR(25) UNIQUE,
    city VARCHAR(25),
    state VARCHAR(2),
    isChef BOOLEAN
);
*/

type User struct {
	userID   string
	username string
	password string
	email    string
	city     string
	state    string
	isChef   bool
}

func CreateUser(userID string, username string, password string, email string, city string, state string, isChef bool) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "User" (userID, username, password, email, city, state, isChef) VALUES ($1, $2, $3, $4, $5, $6, $7);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID, username, password, email, city, state, isChef)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetUser(userID string) (User, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "User" WHERE userID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var user User

	err = stmt.QueryRow(userID).Scan(&user.userID, &user.username, &user.password, &user.email, &user.city, &user.state, &user.isChef)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return user, nil
}

func UpdateUser(userID string, username string, password string, email string, city string, state string, isChef bool) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "User" SET username=$1, password=$2, email=$3, city=$4, state=$5, isChef=$6 WHERE walletID = $7;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID, username, password, email, city, state, isChef)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteUser(userID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "User" WHERE userID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
