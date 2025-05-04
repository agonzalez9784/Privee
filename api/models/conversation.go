package models

/*
CREATE TABLE IF NOT EXISTS "Conversation"(
    conversationID VARCHAR(15) NOT NULL UNIQUE,
    userID VARCHAR(15) NOT NULL,
    chefUserID VARCHAR(15) NOT NULL
);
*/

import (
	"database/sql"
	"log"
)

type Conversation struct {
	conversationID string
	userID         string
	chefUserID     string
}

func CreateConversation(conversationID string, userID string, chefUserID string) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO "Conversation" (conversationID, userID, chefUserID) VALUES ($1, $2, $3);`)

	if err != nil {

		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(conversationID, userID, chefUserID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func GetConversation(conversationID string) (Conversation, error) {
	db, err := dbConnector()

	defer db.Close()

	stmt, err := db.Prepare(`SELECT * FROM "Conversation" WHERE conversationID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var conversation Conversation

	err = stmt.QueryRow(conversationID).Scan(&conversation.conversationID, &conversation.userID, &conversation.chefUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return conversation, nil
}

func UpdateConversation(conversationID string, userID string, chefUserID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`UPDATE "Conversation" userID = $1, chefUserID = $2 WHERE conversationID = $3`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID, chefUserID, conversationID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}

func DeleteConversation(conversationID string) {
	db, err := dbConnector()

	stmt, err := db.Prepare(`DELETE FROM "Conversation" WHERE conversationID = $1;`)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(conversationID)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	return
}
