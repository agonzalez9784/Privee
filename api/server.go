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
	

	//main app functions

	//message/chefID
	//message/chefID with content
	//channel/conversationID
	//bookAppointment
	//

	//autohrization 


	//misc
	http.ListenAndServe(":8080", nil)
}
