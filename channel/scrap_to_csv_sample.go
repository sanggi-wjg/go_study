package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
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

func main() {
	totalPages := getTotalPages()
	// fmt.Println("Total Pages:", totalPages)
	jobs := []ScrappedJobResult{}
	c := make(chan []ScrappedJobResult)

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
		// jobs = append(jobs, scrappedJobs...)
	}
	for i := 0; i < totalPages; i++ {
		job := <-c
		jobs = append(jobs, job...)
	}

	// fmt.Println(jobs)
	resultToCSV(jobs)
}

func resultToCSV(jobs []ScrappedJobResult) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	output := csv.NewWriter(file)
	defer output.Flush()
	defer file.Close()

	headers := []string{"URL", "ID", "Title", "Name", "Location", "Salary"}
	err = output.Write(headers)
	checkErr(err)

	for _, job := range jobs {
		slice := []string{job.pageURL, job.id, job.title, job.name, job.salary}
		err = output.Write(slice)
		checkErr(err)
	}
}

func requestPage(requestURL string) *http.Response {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)

	return res
}

func getPage(page int, mainc chan []ScrappedJobResult) {
	jobs := []ScrappedJobResult{}
	c := make(chan ScrappedJobResult)

	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	res := requestPage(pageURL)

	defer res.Body.Close() // io 여서 close() 해줘야 함
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	jobCards := doc.Find("#mosaic-provider-jobcards > a")
	jobCards.Each(func(i int, card *goquery.Selection) {
		go scrapJob(card, c)
	})

	for i := 0; i < jobCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainc <- jobs
}

func scrapJob(card *goquery.Selection, c chan ScrappedJobResult) {
	id, _ := card.Attr("data-jk")
	title := card.Find(".jobTitle > span").Text()
	companyName := card.Find(".company_location > pre > span").Text()
	location := card.Find(".company_location > pre > div").Text()
	salary, _ := card.Find(".salary-snippet-container > span").Attr("aria-label")
	jobURL := "https://kr.indeed.com/%EC%B1%84%EC%9A%A9%EB%B3%B4%EA%B8%B0?jk=" + id
	// fmt.Println(id, title, companyName, location, salary)

	// return NewScrappedJobResult(jobURL, id, title, companyName, location, salary)
	c <- ScrappedJobResult{pageURL: jobURL, id: id, title: title, name: companyName, location: location, salary: salary}
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
