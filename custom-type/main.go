package main

import "fmt"

type str string

func (test str) log() {
	fmt.Println(test)
}

func main() {

	var name str = "vidit"

	name.log()
}
