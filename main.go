package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	pages := getPages()
	fmt.Println("Total Pages:", pages)

	for i := 0; i < pages; i++ {
		pageURL := baseURL + "&start=" + strconv.Itoa(i*50)
		fmt.Println(pageURL)
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close() // io 여서 close() 해줘야 함
	doc, err := goquery.NewDocumentFromReader(res.Body)
	doc.Find(".pagination .pagination-list").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html()) // fmt.Println(s.Find("a").Length())
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("request failed with ", res.StatusCode, res.Status)
	}
}
