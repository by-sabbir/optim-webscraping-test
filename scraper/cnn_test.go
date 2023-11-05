package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_CNN_SCRAPER = "cnn"
	TEST_CNN_URL     = "https://edition.cnn.com/travel/airbus-overhead-airspace-l-bins/index.html"
	TEST_CNN_BAD_URL = "https://edition.cnn.com/travel/overhead-airspace-l-bins/index.htm"
)

func TestCNN(t *testing.T) {
	svc, err := NewScraperService(TEST_CNN_SCRAPER)

	assert.Equal(t, err, nil)

	items, err := svc.ScrapePage(TEST_CNN_URL)

	assert.Equal(t, err, nil)

	assert.Equal(t, items.Title, "Search for survivors in western Nepal after earthquake kills at least 157 people")
}

func TestCNNBadUrl(t *testing.T) {

	svc, err := NewScraperService(TEST_CNN_SCRAPER)
	assert.Equal(t, err, nil)

	_, err = svc.ScrapePage(TEST_CNN_BAD_URL)
	assert.Error(t, err, "Not Found")

}
