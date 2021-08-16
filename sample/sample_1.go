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

func lenAndUpper_2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func lenAndUpper_3(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeat(words ...string) {
	fmt.Println(words)
}

func addNumbers(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
	return 1
}

func canIfAge(age int) bool {
	if koreanAge := age + 1; koreanAge > 10 {
		return false
	}
	return true
}

func canSwitchAge(age int) bool {
	switch {
	case 18 < age:
		return false
	case 50 < age:
		return true
	}

	// switch age {
	// case 18:
	// 	return true
	// case 50:
	// 	return true
	// }

	return true
}

func main() {
	// name := "Jay"
	// fmt.Println("Hello World " + name)
	// fmt.Println(multiply(10, 8))

	// fmt.Println(lenAndUpper(name))
	// length, _ := lenAndUpper(name)
	// fmt.Println(length)

	// repeat("1", "2", "3", "4")

	// name = "Hahahaha"
	// length, uppercase := lenAndUpper_3(name)
	// fmt.Println(length, uppercase)

	// addNumbers(1, 2, 3, 4, 5)

	// fmt.Println(canIfAge(18))
	// fmt.Println(canSwitchAge(20))
}
