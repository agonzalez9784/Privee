package controllers

import "api/models"

func addFunds(userID string, amountToAdd float64) {
	wallet, err := models.GetWalletByUserID(userID)
	walletID := wallet.WalletID
	amount := wallet.Amount
	models.UpdateWallet(walletID, userID, amount+amountToAdd)
}
