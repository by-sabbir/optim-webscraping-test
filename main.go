package main

import (
	"fmt"
	"log/slog"

	"github.com/by-sabbir/optim-webscraping-test/scraper"
)

func main() {
	logger := slog.New(slog.Default().Handler())

	if err := run(); err != nil {
		logger.Error("something went wrong", "error", err)
	}
}

func run() error {
	guardianUrl := "https://www.theguardian.com/politics/2018/aug/19/brexit-tory-mps-warn-of-entryism-threat-from-leave-eu-supporters"
	cnnUrl := "https://edition.cnn.com/travel/airbus-overhead-airspace-l-bins/index.html"

	if err := scrape(guardianUrl, "guardian"); err != nil {
		return err
	}
	if err := scrape(cnnUrl, "cnn"); err != nil {
		return err
	}

	return nil
}

func scrape(url string, scraperName string) error {
	scraperServcie, err := scraper.NewScraperService(scraperName)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	items, err := scraperServcie.ScrapePage(url)
	if err != nil {
		fmt.Println("error sracping: ", err)
		return err
	}

	defer fmt.Println("scraped items: ", items.Images, items.Title)
	return nil
}
