package scraper

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func (s *GuardianScraperService) ScrapePage(url string) (*ScrapedItem, error) {

	var gItems ScrapedItem

	c := colly.NewCollector(colly.AllowedDomains("www.theguardian.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
		fmt.Println("visiting: ", r.URL, r.URL.RawPath, r.URL.RawQuery)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("status code: ", r.StatusCode)
		fmt.Println("error: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response code: ", r.StatusCode)
		fmt.Println("headers: ", r.Headers)

		// fmt.Println("body: ", string(r.Body))
		fmt.Println("=====================================================")
	})

	c.OnHTML("article", func(h *colly.HTMLElement) {

		imgs := h.DOM.Find("img")
		title := h.DOM.Find("h1").Text()
		body := h.DOM.Find("#maincontent").Text()
		imgs.Each(func(i int, s *goquery.Selection) {
			fmt.Println("img no: ", i)
			val, ok := s.Attr("src")
			if !ok {
				fmt.Println("src not found for img-", i)
			} else {
				gItems.Images = append(gItems.Images, val)
			}

		})

		gItems.Title = title
		gItems.Body = body
	})

	if err := c.Visit("https://www.theguardian.com/politics/2018/aug/19/brexit-tory-mps-warn-of-entryism-threat-from-leave-eu-supporters"); err != nil {
		fmt.Println("error visiting: ", err)
		return &gItems, err
	}

	defer fmt.Println("done scraping")

	return &gItems, nil
}
