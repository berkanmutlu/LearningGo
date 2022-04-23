package main

import (
	"fmt"
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
		FirstName:   "Berkan",
		LastName:    "Mutlu",
		PhoneNumber: "123456789",
		Age:         24,
		BirthDate:   time.Time{},
	}

	fmt.Println(user.FirstName, user.LastName, user.PhoneNumber, user.Age)
}
