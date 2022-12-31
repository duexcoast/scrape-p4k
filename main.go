package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"github.com/zolamk/colly-postgres-storage/colly/postgres"
)
type Review struct {
	URL string
	Album string
	Artist string
	Author string
	Source string
	Rating string
	Date time.Time
	
}
func NewReviewPage(url string) *Review {
	return &Review{URL: url, Source: "Pitchfork"}
}
func (r *Review) Fetch() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PG_PASSWORD:= os.Getenv("PG_PASSWORD")

	c := colly.NewCollector()
		
	storage := &postgres.Storage{
		URI: "postgres://postgres:" + PG_PASSWORD + "@localhost:5432/pitchfork-reviews?sslmode=disable", 
		VisitedTable: "colly_visited",
		CookiesTable: "colly_cookies",
	}

	if err := c.SetStorage(storage); err != nil {
		panic(err)
	}
	c := colly.NewCollector(colly.AllowedDomains("pitchfork.com/reviews/albums/", "www.pitchfork.com/reviews/albums/"))

	c.OnHTML("h1", func(h *colly.HTMLElement) {
		r.Album = trim(h.Text)
	})
	c.OnHTML("div[class=SplitScreenContentHeaderArtist-lfLzFP]", func(h *colly.HTMLElement) {
		r.Artist = trim(h.text)
	})
	c.OnHTML("a[class=byline__name-link]", func(h *colly.HTMLElement) {
		r.Author = trim(h.text)
	})
	c.OnHTML("p[class=InfoSliceValue-gTHtZf", func(h *colly.HTMLElement) {
		date := trim(h.text)

		// Extracted date should be in this form
		const dateForm = "January 2, 2006"

		// Parse into time type
		t, err := time.Parse(dateForm, date)

		updatedDate = t.Format("2006-01-02")

		parseError := err != nil

		datesMatch := currentDate == updatedDate

		r.Date = updatedDate
	})
	c.OnHTML("p[class=Rating-cIWDua]", func(h *colly.HTMLElement) {
		r.Rating = trim(h.text)
	})
	
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PG_PASSWORD:= os.Getenv("PG_PASSWORD")

	c := colly.NewCollector()
		
	storage := &postgres.Storage{
		URI: "postgres://postgres:" + PG_PASSWORD + "@localhost:5432/pitchfork-reviews?sslmode=disable", 
		VisitedTable: "colly_visited",
		CookiesTable: "colly_cookies",
	}

	if err := c.SetStorage(storage); err != nil {
		panic(err)
	}
	
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		e.
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Starting", r.URL)
	})

	c.Visit("http://golang.org")
}
