package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"github.com/zolamk/colly-postgres-storage/colly/postgres"
)
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
	
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Starting", r.URL)
	})

	c.Visit("http://golang.org")
}
