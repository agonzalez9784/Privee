package models

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
	passwod  string
	email    string
	city     string
	state    string
	isChef   bool
}
