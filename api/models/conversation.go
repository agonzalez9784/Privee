package models

/*
CREATE TABLE IF NOT EXISTS "Conversation"(
    conversationID VARCHAR(15) NOT NULL UNIQUE,
    userID VARCHAR(15) NOT NULL,
    chefUserID VARCHAR(15) NOT NULL
);
*/

type Conversation struct {
	conversationID string
	userID         string
	chefUserID     string
}
