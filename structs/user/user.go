package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastname  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetails() {
	fmt.Printf("%s\n%s\n%s\n%v\n", (*u).firstName, u.lastname, u.birthDate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastname = ""
}

func NewAdmin(userEmail, userPassword string) Admin {

	return Admin{
		email:    userEmail,
		password: userPassword,
		User: User{
			firstName: "ADMIN",
			lastname:  "ADMIN",
			birthDate: "1.1.1981",
			createdAt: time.Now(),
		},
	}
}

func New(firstName string, lastName string, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastname:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
