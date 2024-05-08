package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	collector := colly.NewCollector()
	collector.OnHTML(".card", func(e *colly.HTMLElement) {
		// Extract name and price from the child element
		name := strings.TrimSpace(e.ChildText(".title"))
		price := strings.TrimSpace(e.ChildText(".price"))
		desc := strings.TrimSpace(e.ChildText(".description"))

		// Print the name and price
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Price: %s\n", price)
		fmt.Printf("Description: %s\n", desc)
		fmt.Println("---------------------------------")
	})
	collector.Visit("https://webscraper.io/test-sites/e-commerce/static")
}
