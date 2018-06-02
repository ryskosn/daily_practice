package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const url = "https://qiita.com/Yaruki00/items/b50e346551690b158a79"

func useGoQuery() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Find("head > title").Text())
}

func collyExample1() {
	c := colly.NewCollector(
		colly.AllowedDomains("go-colly.org"),
		colly.MaxDepth(2),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("link found: %q -> %s\n", e.Text, link)
		// e.Request.Visit(e.Attr("href"))
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("http://go-colly.org/")
}

func storeXkcd() {
	fileName := "xkcd_store_items.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fileName, err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	// write csv header
	writer.Write([]string{"Name", "Price", "URL", "Image URL"})

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("store.xkcd.com"),
	)

	// Extract product details
	c.OnHTML(".product-grid-item", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildAttr("a", "title"),
			e.ChildText("span"),
			e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
			"https" + e.ChildAttr("img", "src"),
		})
	})

	// Find and visit next page links
	c.OnHTML(`.next a[href]`, func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	url := "https://store.xkcd.com/collections/everything"
	c.Visit(url)

	log.Printf("Scraping finished, check file %q for results\n", fileName)

	// Display collector's statistics'
	log.Println(c)
}

func main() {
	// useGoQuery()
	// collyExample1()
	storeXkcd()
}
