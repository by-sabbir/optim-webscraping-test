package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_GUARDIAN_SCRAPER = "guardian"
	TEST_GUARDIAN_URL     = "https://www.theguardian.com/world/2023/nov/03/dozens-killed-and-injured-by-magnitude-64-earthquake-in-nepal"
	TEST_GUARDIAN_BAD_URL = "https://www.theguardian.com/world/2022/nov/03/dozens-killed-and-injured-by-magnitude-64-earthquake-in-nepal"
)

func TestGuardian(t *testing.T) {
	svc, err := NewScraperService(TEST_GUARDIAN_SCRAPER)

	assert.Equal(t, err, nil)

	items, err := svc.ScrapePage(TEST_GUARDIAN_URL)

	assert.Equal(t, err, nil)

	assert.Equal(t, items.Title, "Search for survivors in western Nepal after earthquake kills at least 157 people")
}

func TestGuardianBadUrl(t *testing.T) {

	svc, err := NewScraperService(TEST_GUARDIAN_SCRAPER)
	assert.Equal(t, err, nil)

	_, err = svc.ScrapePage(TEST_GUARDIAN_BAD_URL)
	assert.Error(t, err, "Not Found")

}
