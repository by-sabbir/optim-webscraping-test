package main

import (
	"fmt"

	"github.com/by-sabbir/optim-webscraping-test/scraper"
)

func main() {
	gUrl := "https://www.theguardian.com/politics/2018/aug/19/brexit-tory-mps-warn-of-entryism-threat-from-leave-eu-supporters"

	guardianScraper, err := scraper.NewScraperService("guardian")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	items, err := guardianScraper.ScrapePage(gUrl)
	if err != nil {
		fmt.Println("error sracping: ", err)
		return
	}

	fmt.Println("items: ", items.Images, items.Title)

	cnnUrl := "https://edition.cnn.com/travel/airbus-overhead-airspace-l-bins/index.html"

	cnnScraper, err := scraper.NewScraperService("cnn")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	items, err = cnnScraper.ScrapePage(cnnUrl)
	if err != nil {
		fmt.Println("error sracping: ", err)
		return
	}

	fmt.Println("items: ", items.Images, items.Title)

	defer fmt.Println("done scraping")
}
