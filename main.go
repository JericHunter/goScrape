package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
    fmt.Println("Hi")
	// Instantiate default collector
	c := colly.NewCollector(
		// colly.AllowedDomains(""),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("table > tbody", func(e *colly.HTMLElement) {
		// link := e.Attr("href")
		
		// Print link
		fmt.Printf(e.Text)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://myanimelist.net/animelist/jeric616?status=7")
    fmt.Println("Bye")
}
