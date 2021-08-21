package main

import (
	"fmt"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/jobs?q=django"

func main() {
	pages := getPages()
	fmt.Println(pages)
}

func getPages() int {
	response, err := http.Get(baseURL)

	return 0
}
