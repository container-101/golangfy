package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	url      string
	title    string
	location string
}

var LinkBaseURL string = "https://kr.indeed.com/viewjob?jk="
var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	jobs := make([]extractedJob, 0)
	c := make(chan []extractedJob)
	start := time.Now()

	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		jobs = append(jobs, <-c...)
	}

	writeJobs(jobs)
	end := time.Now()

	fmt.Println("Done, extracted", len(jobs), "jobs")
	fmt.Println("Execution time:", end.Sub(start))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("../jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	// buffer to actual file?
	defer w.Flush()

	header := []string{"Link", "Title", "Location"}
	w.Write(header)

	for _, job := range jobs {
		line := []string{LinkBaseURL + job.url, job.title, job.location}
		w.Write(line)
	}
}

func getPage(pageIdx int, mainChannel chan<- []extractedJob) {
	jobs := make([]extractedJob, 0)
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(pageIdx*50)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)

	// prevent memory leak
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob)
	}
	mainChannel <- jobs
}

func extractJob(s *goquery.Selection, c chan<- extractedJob) {
	url, _ := s.Attr("data-jk")
	title := cleanString(s.Find("h2>span").Text())
	location := cleanString(s.Find(".companyLocation").Text())
	c <- extractedJob{
		url:      url,
		title:    title,
		location: location,
	}
}

func getPages() int {
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)

	// prevent memory leak
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	fmt.Println(doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	}))

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
