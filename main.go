package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	// const namFe string = "Jay"
	name := "Jay"
	fmt.Println("Hello World " + name)
	fmt.Println(multiply(10, 8))
	fmt.Println(lenAndUpper(name))
}
