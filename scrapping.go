package main

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"rsc.io/quote"
)

var prefix string = "https://en.wikipedia.org"
var prefix_wiki string = "https://en.wikipedia.org/wiki/"
var section string = "#mw-content-text"

func getQuote() string {
	return quote.Go()
}

func getWikiLinks(visit string) []string {

	short_mode := !strings.HasPrefix(visit, prefix_wiki)

	c := colly.NewCollector()
	c.SetRequestTimeout(30 * time.Second)

	links := []string{}

	// Find and visit all links
	c.OnHTML(section+" a[href]", func(e *colly.HTMLElement) {
		// Ignore subsections References and Final Tables
		if e.DOM.ParentsFiltered(".reflist").Is("div") || e.DOM.ParentsFiltered(".navbox").Is("div") {
			return
		}

		link := e.Attr("href")
		if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":") {
			// if !strings.HasPrefix(link, prefix) {
			// 	link = prefix + link
			// }
			// links = append([]string{link}, links...)
			if short_mode {
				link = strings.TrimPrefix(link, "/wiki/")
			} else {
				link = prefix + link
			}
			// fmt.Println(link)
			links = append(links, link)
		}
	})

	if short_mode {
		visit = prefix_wiki + visit
	}
	// fmt.Println("Visiting: ", visit)
	// fmt.Scanln()
	err := c.Visit(visit)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(links)
	return links
}

// mw-content-text
