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
	Name string
}

type ScraperFactory interface {
	ScrapePage(url string) (*ScrapedItem, error)
}

type ScrapedItem struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Images []string `json:"images"`
}

func NewScraperService(name string) (ScraperFactory, error) {

	logger := slog.New(slog.Default().Handler())
	switch {
	case name == "guardian":
		return &GuardianScraperService{
			Name:   name,
			Logger: logger,
		}, nil
	case name == "cnn":
		return nil, ErrNotImplemented
	}

	return nil, ErrNotImplemented
}
