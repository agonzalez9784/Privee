package models

import (
	"database/sql"
	"log"
)

/*
CREATE TABLE IF NOT EXISTS "Transactions" (
    transactionID VARCHAR(15) NOT NULL UNIQUE,
    payersWalletID VARCHAR(15) NOT NULL,
    chefWalletID VARCHAR(15) NOT NULL,
    amount FLOAT,
    status VARCHAR(20) NOT NULL DEFAULT "pending"
)
*/

type Transactions struct {
	transactionID  string
	payersWalletID string
	chefWalletID   string
	amount         float64
	status         string
}

func CreateTransaction(transactionID string, payersWalletID string, chefWalletID string, amount float64, status string) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "Transactions" (transactionID, payersWalletID, chefWalletID, amount, status) VALUES ($1, $2, $3, $4, $5);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(transactionID, payersWalletID, chefWalletID, amount, status)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetTransaction(transactionID string) (Transactions, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "Transactions" WHERE transactionID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var transaction Transactions

	err = stmt.QueryRow(transactionID).Scan(&transaction.transactionID, &transaction.payersWalletID, &transaction.chefWalletID, &transaction.amount, &transaction.status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return transaction, nil
}

func UpdateTransaction(transactionID string, payersWalletID string, chefWalletID string, amount float64, status string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "Transactions" SET payersWalletID=$1, chefWalletID=$2, amount=$3, status=$4 WHERE transactionID = $5;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(payersWalletID, chefWalletID, amount, status, transactionID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteTransaction(transactionID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "Transactions" WHERE transactionID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(transactionID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
