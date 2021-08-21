package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type ScrappedJobResult struct {
	pageURL  string
	id       string
	title    string
	name     string
	location string
	salary   string
}

func NewScrappedJobResult(pageURL string, id string, title string, name string, location string, salary string) *ScrappedJobResult {
	result := ScrappedJobResult{
		pageURL:  pageURL,
		id:       id,
		title:    title,
		name:     name,
		location: location,
		salary:   salary,
	}
	return &result
}

func (r *ScrappedJobResult) AddSalary(salary string) {
	r.salary = salary
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
var jobURL string = "https://kr.indeed.com/%EC%B1%84%EC%9A%A9%EB%B3%B4%EA%B8%B0?jk="

func main() {
	// totalPages := getTotalPages()
	// fmt.Println("Total Pages:", totalPages)

	for i := 0; i < 1; i++ {
		getPage(i)
	}
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

func requestPage(requestURL string) *http.Response {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)

	return res
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println(pageURL)

	res := requestPage(pageURL)
	defer res.Body.Close() // io 여서 close() 해줘야 함
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	doc.Find("#mosaic-provider-jobcards > a").Each(func(i int, card *goquery.Selection) {
		// result := NewScrappedJobResult(jobURL+id, id, title, companyName, location)
		// if salaryExist {
		// 	result.AddSalary(salary)
		// }
		job := scrapJob(card)
	})
}

func scrapJob(card *goquery.Selection) *ScrappedJobResult {
	id, _ := card.Attr("data-jk")
	title := card.Find(".jobTitle > span").Text()
	companyName := card.Find(".company_location > pre > span").Text()
	location := card.Find(".company_location > pre > div").Text()
	salary, _ := card.Find(".salary-snippet-container > span").Attr("aria-label")
	// fmt.Println(id, title, companyName, location, salary)
	return NewScrappedJobResult(jobURL+id, id, title, companyName, location, salary)
}

func getTotalPages() int {
	pages := 0
	res := requestPage(baseURL)

	defer res.Body.Close() // io 여서 close() 해줘야 함
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	doc.Find(".pagination .pagination-list").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html()) // fmt.Println(s.Find("a").Length())
		pages = s.Find("a").Length()
	})

	return pages
}
