package main

import (
	"api/controllers"
	"fmt"
	"net/http"
	"net/url"
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
