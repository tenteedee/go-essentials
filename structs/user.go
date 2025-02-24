package structs

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) String() string {
	return fmt.Sprintf(`
User: %s %s
Birthdate: %s
Created at: %s
`, u.firstName, u.lastName, u.birthdate, u.createdAt.Format(time.RFC1123))
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName, and birthdate are required fields")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

type Admin struct {
	User
	email    string
	password string
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required fields")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "User",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}, nil
}
