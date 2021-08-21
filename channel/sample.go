package main

import (
	"fmt"
	"net/http"
)

type HitResult struct {
	url         string
	status_code int
	status      string
}

func NewHitResult(url string, status_code int, status string) *HitResult {
	result := HitResult{
		url:         url,
		status_code: status_code,
		status:      status,
	}
	return &result
}

// func (r *Result) AddResult(status string, status_code int) {
// 	r.status = status
// 	r.status_code = status_code
// }

func main() {
	urls := []string{
		"https://nomadcoders.co/",
		"https://naver.com",
		"https://www.google.com",
		"https://fmkorea.com",
		"https://www.reddit.com",
	}
	results := make([]HitResult, len(urls))
	channel := make(chan HitResult)

	for _, url := range urls {
		go hitURL(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		results[i] = <-channel
	}

	for _, r := range results {
		fmt.Println(r)
	}
}

func hitURL(url string, c chan HitResult) {
	response, err := http.Get(url)

	if err != nil || response.StatusCode >= 400 {
		c <- *NewHitResult(url, response.StatusCode, "Failed")
	}
	c <- *NewHitResult(url, response.StatusCode, "Success")
}
