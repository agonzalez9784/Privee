package main

import (
	"api/controllers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"api/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	//Values
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	userID := utils.GenerateID(15)

	//Connection
	db, err := utils.DBConnector()

	defer db.Close()

	if err != nil {
		http.Error(w, "Oops! Something went wrong!", http.StatusInternalServerError)
	}

	//Prepared Statement
	stmt, err := db.Prepare(`INSERT INTO "User" (userID, username, password, email, city, state, isChef) VALUES ($1, $2, $3, $4, $5, $6, $7);`)

	if err != nil {
		http.Error(w, "Oops! Something went wrong!", http.StatusInternalServerError)
	}

	defer stmt.Close()

	//Execution
	_, err := stmt.Exec(userID, username, password, email, "", "", false)

	if err != nil {
		http.Error(w, "Oops! Something went wrong!", http.StatusInternalServerError)
	}

	//Response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func AuthorizeUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	urlString := r.URL.String()
	parsedURL, err := url.Parse(urlString)

	if err != nil {
		fmt.Println("Something went wrong! oops")
	}

	queryParams := parsedURL.Query()

	username := queryParams.Get("username")
	password := queryParams.Get("password")

	fmt.Println(username)
	fmt.Println(password)

	token, err := controllers.AuthorizeUser(username, password)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, token)

}

func searchForChefs(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	urlString := r.URL.String()
	parsedURL, err := url.Parse(urlString)

	if err != nil {
		fmt.Println("Something went wrong! oops")
	}

	queryParams := parsedURL.Query()

	city := queryParams.Get("city")
	state := queryParams.Get("state")

	chefs, err := controllers.GetChefs(city, state)

	jsonData, err := json.MarshalIndent(chefs, "", "  ")

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

}

func declineBooking(w http.ResponseWriter, r *http.Request) {
	
	bookingID := r.FormValue("bookingID")
	userID := r.FormValue("userID")

	if r.Method != http.MethodPost {
		
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	controllers.DeclineBooking(bookingID)
}

func addExperience(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	controllers.AddExperience()
}

func updatePortfolio(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/* REDACTED CAN JUST USE CONDITIONS IN THE SEND MESSAGE FUNC TO CHECK IF THE CONVERSATIONS EXISTS
func newConversation(w http.ResponseWriter, r *http.Request) {

}

func conversation(w http.ResponseWriter, r *http.Request) {

}*/

func sendMessage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func approveBooking(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	controllers.AcceptBooking()
}

func confirmArrival(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	controllers.ConfirmArrival()
}

func noShow(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	controllers.noShow()
}

func disputeBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := ""
	token := r.FormValue("token")

	

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if(!controllers.ChefPermissionsForBookings(bookingID, chefID)){
		http.Error(w, "You don't have permission to do that!", http.StatusMethodNotAllowed)
	}

	
}

func review(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func requestRefund(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func bookAChef(w http.ResponseWriter, r *http.Request) { //NOTE need to subtract tokens from userAccount TODO
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	urlString := r.URL.String()
	parsedURL, err := url.Parse(urlString)

	if err != nil {
		fmt.Println("Something went wrong! oops")
	}

	queryParams := parsedURL.Query()

	userID := queryParams.Get("userID")
	chefID := queryParams.Get("chefID")
	startTime := queryParams.Get("startTime")
	endTime := queryParams.Get("endTime")

	rateStr := queryParams.Get("rate")
	rate, err := strconv.ParseFloat(rateStr, 64)

	booking, err := controllers.BookAChef(userID, chefID, startTime, endTime, rate)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, booking)

}

func addFundsToWallet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	controllers.AddFunds()
}
