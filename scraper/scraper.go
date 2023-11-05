// Scraper provides interface for multiple domain specific scraping service
package scraper

import (
	"errors"
	"log/slog"
)

var ErrNotImplemented = errors.New("specified scraper service not yet implemented")

type GuardianScraperService struct {
	Name   string
	Logger *slog.Logger
}

type CNNScraperService struct {
	Name   string
	Logger *slog.Logger
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
		return &GuardianScraperService{
			Name:   name,
			Logger: logger,
		}, nil
	case name == "cnn":
		return &CNNScraperService{
			Name:   name,
			Logger: logger,
		}, nil
	default:
		return nil, ErrNotImplemented
	}

	return nil, ErrNotImplemented
}
