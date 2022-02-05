package main

import "fmt"

func main() {
	// array length 지정
	names := [5]string{"a", "b", "c"}
	fmt.Println(names)

	names[3] = "d"
	names[4] = "e"
	fmt.Println(names)

	// array length 미지정
	names_2 := []string{"ㄱ", "ㄴ", "ㄷ"}

	names_2 = append(names_2, "ㄹ")
	fmt.Println(names_2)
}
