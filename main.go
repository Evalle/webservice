package main

import (
	"fmt"

	"github.com/evalle/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "John",
		LastName:  "Doe",
	}
	fmt.Println(u)
}
