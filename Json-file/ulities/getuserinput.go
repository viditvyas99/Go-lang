package ulities

import (
	"errors"
	"fmt"
)

func GetUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var userInput string

	fmt.Scanln(&userInput)

	if userInput == "" {
		return "", errors.New("Pls enter value")
	}

	return userInput, nil

}
