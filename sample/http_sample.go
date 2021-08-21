package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://nomadcoders.co/go-for-beginners/lectures/1519",
		"https://naver.com",
		"https://www.google.com",
		"https://fmkorea.com",
	}

	results := map[string]string{}

	for _, url := range urls {
		err := hitURL(url)
		if err == nil {
			results[url] = "Success"
		} else {
			results[url] = "Failed"
		}
	}

	fmt.Println(results)
}

func hitURL(url string) error {
	fmt.Println("Check URL:", url)
	response, err := http.Get(url)

	if err != nil || response.StatusCode >= 400 {
		return errors.New("request failed")
	}

	return nil
}
