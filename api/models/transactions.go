package models

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
