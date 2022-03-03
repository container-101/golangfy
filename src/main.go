package main

import (
	"golwee/src/scrapper"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("public/index.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	if term == "" {
		return c.String(http.StatusBadRequest, "No term provided")
	}
	filePath := "public/" + term + "-jobs.csv"
	defer os.Remove(filePath)
	scrapper.Scrape(term)
	return c.Attachment(filePath, term+"-jobs.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
