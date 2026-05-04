package main

import (
	"fmt"
	"time"
)

func  main()  {

	done := make(chan bool) 
	 go greet()

	go slowGreet(done)
	go greet()
	 for  range done {
	 }
	
}

func greet () {
	fmt.Println("nice to meet you")
}


func slowGreet(doneChan chan bool) {
	time.Sleep(3*time.Second)
	fmt.Println("byee")

	doneChan <- true
	close(doneChan)
}

