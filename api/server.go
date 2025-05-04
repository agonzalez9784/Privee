package main

import (
	"api/models"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Server started on port 8080")
	//models.CreateChef("93jr039j", "38fh38h38", "default.png", "johnny", "sins", "your local chef", 0)
	res, err := models.GetChef("93jr039j")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(res)

}
