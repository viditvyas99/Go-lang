package main

import (
	"fmt"
	"time"

	"example.com/File-Read-Write/user"
)

func main() {

	userName := getUserData("Enter User Name ")
	userLastName := getUserData("Enter User Last  Name ")
	dob := getUserData("Enter User D.O.B ")

	appUserData, err := user.NewUser(userName, userLastName, dob)

	if err != nil {
		fmt.Print(err)

		return
	}

	appUserData.ClearUserName()
	appUserData.OutPrintUserDetials()
	updatedFirstrName := getUserData("Enter the update first name")
	updatedLastName := getUserData("Enter the update Last name")
	appUserData.ReplaceUserName(updatedFirstrName, updatedLastName)
	appUserData.OutPrintUserDetials()
	outPrintUserInformationWithPointer(appUserData)
	outPrintUserInformation(appUserData)
	fmt.Print("Admin ----------------")
	email := getUserData("Enter User email ")
	password := getUserData("Enter User password ")

	admin, adminErr := user.NewAdmin(email, password, appUserData)

	if adminErr != nil {
		fmt.Print(err)

		return
	}

	admin.OutPrintAdminDetials()
}

func getUserData(PromptText string) string {
	fmt.Print(PromptText)
	var value string
	fmt.Scanln(&value)

	return value
}

func outPrintUserInformationWithPointer(userData *user.User) {
	fmt.Printf("User information\n name : %s\n Last name :  %s\n D.O.B:  %s\n time of cration : %s\n thisone using pointer", userData.FirstName, userData.LastName, userData.BirthDate, userData.CreatedAt.Format(time.RFC1123))
}

func outPrintUserInformation(userData *user.User) {
	fmt.Printf("User information\n name : %s\n Last name :  %s\n D.O.B:  %s\n time of cration : %s", userData.FirstName, userData.LastName, userData.BirthDate, userData.CreatedAt.Format(time.RFC1123))
}
