package models

import (
	"database/sql"
	"log"
)

/*
CREATE TABLE IF NOT EXISTS "Wallet" (
    walletID VARCHAR(15) NOT NULL UNIQUE,
    userID VARCHAR(15) NOT NULL,
    amount FLOAT NOT NULL
);
*/

type Wallet struct {
	walletID string
	userID   string
	amount   float64
}

func CreateWallet(walletID string, userID string, amount string) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "Wallet" (walletID, userID, amount) VALUES ($1, $2, $3);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(walletID, userID, amount)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetWallet(walletID string) (Wallet, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "Wallet" WHERE walletID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var wallet Wallet

	err = stmt.QueryRow(walletID).Scan(&wallet.walletID, &wallet.userID, &wallet.amount)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return wallet, nil
}

func UpdateWallet(walletID string, userID string, amount string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "Wallet" SET userID = $1, amount=$2 WHERE walletID = $3;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID, amount, walletID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteWallet(walletID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "Wallet" WHERE walletID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(walletID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
