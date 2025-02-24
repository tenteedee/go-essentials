package controller

import (
	"fmt"

	"github.com/tenteedee/go-essentials/structs"
)

func UserController() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// appUser := User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthdate: birthdate,
	// 	createdAt: time.Now(),
	// }

	appUser, err := structs.New(firstName, lastName, birthdate)
	adminUser, err := structs.NewAdmin("testemail@gmail.com", "password")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(appUser.String())
	appUser.ClearUsername()
	fmt.Println(appUser.String())

	fmt.Println(adminUser.String())
	adminUser.ClearUsername()
	fmt.Println(adminUser.String())

	// outputUserData(&appUser)

	// ... do something awesome with that gathered data!

}

// func outputUserData(user *structs.User) {
// 	fmt.Printf(`
// User: %s %s
// Birthdate: %s
// Created at: %s
// `, user.firstName, user.lastName, user.birthdate, user.createdAt.Format(time.RFC1123))
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
