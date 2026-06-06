package scraper

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type ScrapedItem struct {
	Title   string `json:"title,omitempty"`
	Link    string `json:"link,omitempty"`
	Source  string `json:"source"`
	Scraped string `json:"scraped_at"`
}

type Result struct {
	Items []ScrapedItem `json:"items"`
}

func ScrapeCryptoNews() (*Result, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("cointelegraph.com", "www.cointelegraph.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
	)

	c.SetRequestTimeout(30 * time.Second)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1, Delay: 1 * time.Second})

	var items []ScrapedItem

	c.OnHTML("article", func(e *colly.HTMLElement) {
		title := ""
		link := ""

		title = e.ChildText("h2")
		if title == "" {
			title = e.ChildText("h3")
		}
		if title == "" {
			title = e.ChildText(".post-card-title")
		}
		if title == "" {
			title = e.ChildText("a")
		}

		link = e.ChildAttr("a", "href")
		if link != "" && !strings.HasPrefix(link, "http") {
			link = "https://cointelegraph.com" + link
		}

		title = strings.TrimSpace(title)
		if title != "" && len(title) > 10 {
			items = append(items, ScrapedItem{
				Title:   title,
				Link:    link,
				Source:  "CoinTelegraph",
				Scraped: time.Now().UTC().Format(time.RFC3339),
			})
		}
	})

	c.OnHTML("a[class*='post']", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		if title != "" && len(title) > 15 {
			link := e.Attr("href")
			if link != "" && !strings.HasPrefix(link, "http") {
				link = "https://cointelegraph.com" + link
			}
			items = append(items, ScrapedItem{
				Title:   title,
				Link:    link,
				Source:  "CoinTelegraph",
				Scraped: time.Now().UTC().Format(time.RFC3339),
			})
		}
	})

	err := c.Visit("https://cointelegraph.com")
	if err != nil {
		return nil, fmt.Errorf("scrape cointelegraph: %w", err)
	}

	c.Wait()

	seen := make(map[string]bool)
	unique := make([]ScrapedItem, 0, len(items))
	for _, item := range items {
		if !seen[item.Title] {
			seen[item.Title] = true
			unique = append(unique, item)
		}
	}

	if len(unique) > 10 {
		unique = unique[:10]
	}

	log.Printf("scraped %d news items from CoinTelegraph", len(unique))

	return &Result{Items: unique}, nil
}
