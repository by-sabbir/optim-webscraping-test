// Scraper provides interface for multiple domain specific scraping service
package scraper

import (
	"errors"
	"log/slog"
	"time"

	"github.com/gocolly/colly"
)

var ErrNotImplemented = errors.New("specified scraper service not yet implemented")

type GuardianScraperService struct {
	Name      string
	Logger    *slog.Logger
	Collector *colly.Collector
}

type CNNScraperService struct {
	Name      string
	Logger    *slog.Logger
	Collector *colly.Collector
}

type ScraperFactory interface {
	ScrapePage(url string) (*ScrapedItem, error)
}

type Metadata struct {
	Description string `json:"description"`
	Tags        string `json:"tags"`
}
type ScrapedItem struct {
	Metadata `json:"metadata"`
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Images   []string `json:"images"`
}

// initiates a new scraper service
func NewScraperService(name string) (ScraperFactory, error) {

	logger := slog.New(slog.Default().Handler())
	switch {
	case name == "guardian":
		c := colly.NewCollector(
			colly.AllowedDomains("www.theguardian.com"),
		)
		c.Limit(&colly.LimitRule{
			RandomDelay: 2 * time.Second,
		})
		return &GuardianScraperService{
			Name:      name,
			Logger:    logger,
			Collector: c,
		}, nil
	case name == "cnn":
		c := colly.NewCollector()
		return &CNNScraperService{
			Name:      name,
			Logger:    logger,
			Collector: c,
		}, nil
	default:
		return nil, ErrNotImplemented
	}
}
