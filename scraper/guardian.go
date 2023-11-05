// Scraper guardian scraper service
package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func (s *GuardianScraperService) ScrapePage(url string) (*ScrapedItem, error) {

	var gItems ScrapedItem

	c := colly.NewCollector(colly.AllowedDomains("www.theguardian.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
		s.Logger.Info("visiting", "url", r.URL, "Query", r.URL.RawQuery)
	})

	c.OnError(func(r *colly.Response, err error) {
		s.Logger.Error("error", "status_code", r.StatusCode, "error", err, "url", url)
	})

	c.OnResponse(func(r *colly.Response) {
		s.Logger.Info("response", "status_code", r.StatusCode, "url", url)
	})

	c.OnHTML("meta", func(h *colly.HTMLElement) {
		doc := h.DOM
		doc.Each(func(i int, s *goquery.Selection) {
			if name, _ := s.Attr("name"); name == "description" {
				description, _ := s.Attr("content")
				gItems.Metadata.Description = description
			}
			if name, _ := s.Attr("property"); name == "article:tag" {
				tags, _ := s.Attr("content")
				gItems.Metadata.Tags = tags
			}
		})
	})
	c.OnHTML("article", func(h *colly.HTMLElement) {

		imgs := h.DOM.Find("img")
		title := h.DOM.Find("h1").Text()
		body := h.DOM.Find("#maincontent").Text()
		imgs.Each(func(i int, sc *goquery.Selection) {
			val, ok := sc.Attr("src")
			if !ok {
				s.Logger.Warn("img_parsing", "msg", "src not found", "url", url)
			} else {
				gItems.Images = append(gItems.Images, val)
			}
		})

		gItems.Title = title
		gItems.Body = body
	})

	if err := c.Visit(url); err != nil {
		s.Logger.Error("could not visit page", "error", err, "url", url)
		return &gItems, err
	}

	defer s.Logger.Info("scrape done", "url", url)

	return &gItems, nil
}
