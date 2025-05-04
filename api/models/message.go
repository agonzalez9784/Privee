package models

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
