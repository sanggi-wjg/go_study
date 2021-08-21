package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	people := [2]string{"First", "Second"}

	for _, person := range people {
		go isSomething(person, channel)
	}

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func isSomething(person string, c chan bool) {
	time.Sleep(time.Second * 3)
	if person == "First" {
		c <- true
	} else {
		c <- false
	}
}
