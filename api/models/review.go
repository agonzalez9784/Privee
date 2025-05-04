package models

/*
CREATE TABLE IF NOT EXISTS "Review" (
    reviewID VARCHAR(15) NOT NULL UNIQUE,
    chefID VARCHAR(15) NOT NULL,
    userID VARCHAR(15) NOT NULL,
    stars INT NOT NULL,
    review TEXT
);
*/

type Review struct {
	reviewID string
	chefID   string
	userID   string
	stars    int
	review   string
}
