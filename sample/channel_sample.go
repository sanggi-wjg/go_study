package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	channel := make(chan string)
	people := [5]string{"First", "Second", "Third", "Fourth", "Fifth"}

	for _, person := range people {
		go isSomething(person, channel)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-channel)
	}

	// fmt.Println("Done")
}

func isSomething(person string, c chan string) {
	random := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(random))

	if person == "First" {
		c <- person + " is true sleep: " + strconv.Itoa(random)
	} else {
		c <- person + " is false sleep: " + strconv.Itoa(random)
	}
}
