package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		result := getPage(i)
		jobs = append(jobs, result...)
	}
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs), "jobs")
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

func getPage(pageIdx int) []extractedJob {
	jobs := make([]extractedJob, 0)
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
		job := extractJob(s)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(s *goquery.Selection) extractedJob {
	url, _ := s.Attr("data-jk")
	title := cleanString(s.Find("h2>span").Text())
	location := cleanString(s.Find(".companyLocation").Text())
	return extractedJob{
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
