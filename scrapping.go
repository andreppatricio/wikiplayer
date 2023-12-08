package main

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"rsc.io/quote"
)

var prefix string = "https://en.wikipedia.org"

func getQuote() string {
	return quote.Go()
}

func getWikiLinks(visit string) []string {

	section := "#mw-content-text"
	c := colly.NewCollector()
	c.SetRequestTimeout(30 * time.Second)

	links := []string{}
	// linkSet := make(Set)

	// Find and visit all links
	c.OnHTML(section+" a[href]", func(e *colly.HTMLElement) {
		// Ignore subsections References and Final Tables

		if e.DOM.ParentsFiltered(".reflist").Is("div") || e.DOM.ParentsFiltered(".navbox").Is("div") {
			return
		}

		link := e.Attr("href")
		if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":") {
			if !strings.HasPrefix(link, prefix) {
				link = prefix + link
			}
			// links = append([]string{link}, links...)
			links = append(links, link)
		}
	})

	err := c.Visit(visit)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(links)
	return links
}

// mw-content-text
