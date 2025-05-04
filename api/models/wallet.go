package models

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
