package models

import (
	"database/sql"
	"log"
)

/*
CREATE TABLE IF NOT EXISTS "Message" (
    messageID VARCHAR(20) NOT NULL UNIQUE,
    conversationID VARCHAR(15) NOT NULL,
    userID VARCHAR(15) NOT NULL,
    messageContent VARCHAR(500)
);
*/

type Message struct {
	messageID      string
	conversationID string
	userID         string
	messageContent string
}

func CreateMessage(messageID string, conversationID string, userID string, messageContent string) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "Message" (messageID, conversationID, userID, messageContent) VALUES ($1, $2, $3, $4);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(messageID, conversationID, userID, messageContent)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetMessage(messageID string) (Message, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "Message" WHERE messageID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var message Message

	err = stmt.QueryRow(messageID).Scan(&message.messageID, &message.conversationID, &message.userID, &message.messageContent)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return message, nil
}

func UpdateMessage(messageID string, conversationID string, userID string, messageContent string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "Message" SET conversationID = $1, userID=$2, messageContent=$3  WHERE messageID = $4;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(conversationID, userID, messageContent, messageID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteMessage(messageID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "Message" WHERE messageID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(messageID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
