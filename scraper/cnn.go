package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func (s *CNNScraperService) ScrapePage(url string) (*ScrapedItem, error) {
	var cItems ScrapedItem

	s.Collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
		s.Logger.Info("visiting", "url", r.URL, "Query", r.URL.RawQuery)
	})

	s.Collector.OnError(func(r *colly.Response, err error) {
		s.Logger.Error("error", "status_code", r.StatusCode, "error", err, "url", url)
	})

	s.Collector.OnResponse(func(r *colly.Response) {
		s.Logger.Info("response", "status_code", r.StatusCode, "url", url)
	})

	s.Collector.OnHTML("meta", func(h *colly.HTMLElement) {
		doc := h.DOM
		doc.Each(func(i int, s *goquery.Selection) {
			if name, _ := s.Attr("name"); name == "description" {
				description, _ := s.Attr("content")
				cItems.Metadata.Description = description
			}
			if name, _ := s.Attr("property"); name == "article:tag" {
				tags, _ := s.Attr("content")
				cItems.Metadata.Tags = tags
			}
		})
	})

	s.Collector.OnHTML("h1", func(h *colly.HTMLElement) {
		title := h.Text
		cItems.Title = title
	})
	s.Collector.OnHTML(".article__content", func(h *colly.HTMLElement) {
		body := h.Text
		cItems.Body = body
	})
	s.Collector.OnHTML("img", func(h *colly.HTMLElement) {
		imgs := h.DOM
		imgs.Each(func(i int, sc *goquery.Selection) {
			val, ok := sc.Attr("src")
			if !ok {
				s.Logger.Warn("img_parsing", "msg", "src not found", "url", url)
			} else {
				cItems.Images = append(cItems.Images, val)
			}
		})
	})

	if err := s.Collector.Visit(url); err != nil {
		s.Logger.Error("could not visit page", "error", err, "url", url)
		return &cItems, err
	}

	defer s.Logger.Info("scrape done", "url", url)

	return &cItems, nil
}
