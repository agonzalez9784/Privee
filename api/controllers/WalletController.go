package controllers

import (
	"api/models"
	"log"
	"strconv"
)

func AddFunds(userID string, amountToAdd float64) {
	wallet, err := models.GetWalletByUserID(userID)

	if err != nil {
		log.Fatal(err)
	}

	walletID := wallet.WalletID
	amount := wallet.Amount + amountToAdd
	amountConverted := strconv.FormatFloat(amount, 'f', -1, 64)
	models.UpdateWallet(walletID, userID, amountConverted)
}
