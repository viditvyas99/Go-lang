package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User
}

//attech function struct

func (userData *User) OutPrintUserDetials() {
	fmt.Printf("User information\n name : %s\n Last name :  %s\n D.O.B:  %s\n time of cration : %s\n this one is for function attech to struck ", userData.FirstName, userData.LastName, userData.BirthDate, userData.CreatedAt.Format(time.RFC1123))

}

func (userData *User) ClearUserName() {
	userData.FirstName = ""
	userData.LastName = ""
}

func (userData *User) ReplaceUserName(name, lastName string) {

	userData.FirstName = name
	userData.LastName = lastName

}

// consturtion / creation function

func NewUser(userName, userLastName, dob string) (*User, error) {

	if userLastName == "" || userName == "" || dob == "" {
		return nil, errors.New("pls enter name lastName and dob are zroori  ")
	}

	return &User{
		FirstName: userName,
		LastName:  userLastName,
		BirthDate: dob,
		CreatedAt: time.Now(),
	}, nil

}

func (adminData *Admin) OutPrintAdminDetials() {
	fmt.Printf("admin information\n Email : %s\n Password : %s\n  name : %s\n Last name :  %s\n D.O.B:  %s\n time of cration : %s\n this one is for function attech to struck ", adminData.Email, adminData.Email, adminData.FirstName, adminData.LastName, adminData.BirthDate, adminData.CreatedAt.Format(time.RFC1123))

}

func NewAdmin(email string, password string, user *User) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("password and email ")
	}

	return &Admin{
		Email:    email,
		Password: password,
		User:     *user,
	}, nil

}
