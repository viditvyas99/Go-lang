package main

import (
	"fmt"

	"example.com/json-file/ulities"
)

func main() {

	name, nameErr := ulities.GetUserInput("Pls enter your name")

	if nameErr != nil {
		fmt.Println(nameErr.Error())
		return
	}
	fmt.Println(name)

}
