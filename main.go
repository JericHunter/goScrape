package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"github.com/gocolly/colly"
)
type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
func main() {
	// Instantiate default collector
	c:= colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".factsList li", func(e *colly.HTMLElement) {
		// link := e.Attr("href")
		factID, err := strconv.Atoi(e.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}
		factDesc := e.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.factretriever.com/spider-facts")
}
func writeJSON(data []Fact) {

}