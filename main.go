package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{
		FirstName:   "Henry",
		LastName:    "Lee",
		PhoneNumber: "1 555- 555- 1212",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate, "PhoneNumber:", user.PhoneNumber)

}
