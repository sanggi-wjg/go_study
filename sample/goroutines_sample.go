package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go count("First")
// 	count("Second")
// }

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, i)
		time.Sleep(time.Second)
	}
}
