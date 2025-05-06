package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Printf("Running server!")
	//ROUTES
	http.HandleFunc("/registerUser", RegisterUserWrapper)
	http.HandleFunc("/authorizeUser", AuthorizeUserWrapper)

	//misc
	http.ListenAndServe(":8080", nil)
}
