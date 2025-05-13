package controllers

import (
	"api/models"
	"api/utils"
	"fmt"
	"log"
)

func hasOwnership() {
	x := 0
	x = x + 1
}


func AuthorizeUser(username string, password string) (string, error) {

	userData, err := models.GetUserByUsername(username)

	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(userData)

	if err != nil {
		log.Fatal(err)
	}
	usr_ID := userData.UserID
	usr_username := userData.Username
	usr_password := userData.Password

	fmt.Println(usr_username)
	fmt.Println(usr_password)

	if (username == usr_username) && (password == usr_password) {
		fmt.Println("Successful but nothing ain't working with jwts")
		token, err := utils.CreateToken(usr_ID)

		fmt.Println(token)
		return token, err
	}

	return "", nil
}
