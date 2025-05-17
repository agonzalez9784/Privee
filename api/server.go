package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Printf("Running server!")
	//ROUTES

	//USERS AND REGISTRATION

	//normal users
	http.HandleFunc("/registerUser", RegisterUser)
	http.HandleFunc("/authorizeUser", AuthorizeUser)

	//chefs

	http.HandleFunc("/approveBooking", approveBooking)
	http.HandleFunc("/declineBooking", declineBooking)
	http.HandleFunc("/addExperience", addExperience)
	http.HandleFunc("/updatePortfolio", updatePortfolio)
	//APPLICATION FUNCTIONS MAIN

	http.HandleFunc("/searchForChefs", searchForChefs)
	http.HandleFunc("/newConversation", newConversation)
	http.HandleFunc("/conversation", conversation)
	http.HandleFunc("/sendMessage", sendMessage)
	http.HandleFunc("/book", bookAChef)
	http.HandleFunc("/confirmArrival", confirmArrival)
	http.HandleFunc("/noShow", noShow)
	http.HandleFunc("/disputeBooking", disputeBooking) //for chefs too
	http.HandleFunc("/review", review)

	//WALLET AND TRANSACTIONS
	http.HandleFunc("/addFundsToWallet", addFundsToWallet)
	http.HandleFunc("/requestRefund", requestRefund)

	//SECURITY AND AUTHORIZATION

	//MISC
	http.ListenAndServe(":8080", nil)
}
