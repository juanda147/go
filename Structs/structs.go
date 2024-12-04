package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user
	//appUser = user {} -- null value for user struct type

	appUser = user{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}

	// ... do something awesome with that gathered data!
	outputUserDetails(appUser)
}

func outputUserDetails(u user) {
	//...
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
