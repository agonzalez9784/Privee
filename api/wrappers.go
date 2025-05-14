package main

import (
	"api/controllers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func RegisterUserWrapper(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		controllers.RegisterUserAccount(username, password, email)
		fmt.Fprintln(w, "This is a POST request.")

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func AuthorizeUserWrapper(w http.ResponseWriter, r *http.Request) {
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
	controllers.DeclineBooking()
}

func addExperience(w http.ResponseWriter, r *http.Request) {
	controllers.AddExperience()
}

func updatePortfolio(w http.ResponseWriter, r *http.Request) {

}

func newConversation(w http.ResponseWriter, r *http.Request) {

}

func conversation(w http.ResponseWriter, r *http.Request) {

}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func approveBooking(w http.ResponseWriter, r *http.Request) {
	controllers.AcceptBooking()
}

func confirmArrival(w http.ResponseWriter, r *http.Request) {
	controllers.ConfirmArrival()
}

func noShow(w http.ResponseWriter, r *http.Request) {
	controllers.noShow()
}

func disputeBooking(w http.ResponseWriter, r *http.Request) {

}

func review(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func requestFefund(w http.ResponseWriter, r *http.Request) {

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

}

func requestRefund(w http.ResponseWriter, r *http.Request) {

}
