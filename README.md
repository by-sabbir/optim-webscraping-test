# Optimizely Content Intelligence and the Web

[![Go Report Card](https://goreportcard.com/badge/github.com/by-sabbir/optim-webscraping-test)](https://goreportcard.com/report/github.com/by-sabbir/optim-webscraping-test) [![codecov](https://codecov.io/gh/by-sabbir/optim-webscraping-test/graph/badge.svg?token=9LBFSUT3JZ)](https://codecov.io/gh/by-sabbir/optim-webscraping-test)

## Installation

Docker is required for next steps, please follow [docker install guide](https://docs.docker.com/engine/install/) for installing Docker

- Build the image

```bash
make build
```

This will build a docker image with tag- `scraper:v0.0.1` which contains the `scraper` binary.

## Usage

```bash
docker run -it --rm scraper:v0.0.1 scraper

output:
--------------
2023/11/05 11:20:54 INFO running scraper version=v0.0.1

Optimizely Content Intelligence and the Web

Usage:
  scrape [flags]

Flags:
  -h, --help            help for scrape
  -p, --parser string   the supported parser (default "guardian")
  -u, --url string      this should be the url you want to parse (default "https://www.theguardian.com/world/2023/nov/03/dozens-killed-and-injured-by-magnitude-64-earthquake-in-nepal")
```

Version `v0.0.1` supports `guardian` and `cnn` as parser configuration. They should be able to scrap any blog page for those domains.

**Example**

```bash
docker run -it --rm scraper:v0.0.1 scraper scrape -p guardian -u https://www.theguardian.com/world/2023/nov/03/dozens-killed-and-injured-by-magnitude-64-earthquake-in-nepal
```

This will scrap the spcified URL and return json with following template -

```go
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
```

## Design Decisions

#### Language and Libraries

- Go
  - Scraping system often requires high-scalability, concurrent execution for speed, and ability to process large volume of sreams. Considering these features Go seemed the best choice for me.
- Colly
  - The most reliable web-scraping framework for Go
- goquery
  - jQuery like DOM selector, also used in Colly

#### Design Pattern

Clearly the problem falls into "Creational Design Pattern" as We will have a scraper which should be able to scrap different types of web pages. So, our primary object/class will provide an interface to create subclasses that will be allowed to alter the html pages as the subclassses are deligated to.

Before considering the specific design patter, I first set a few rules/policies for the project -

- The implementation should be concrete in terms of data
- Ease of configuration (Rate Limiting, Domain whitelising, API Key etc)
- Robust error handling and logging
- Flexibility and extensiblity without loosing readability, often scraping for different domain loose readability with time. I wanted to avoid this.
- Should be easily testable.

with these policies in mind, I went for factory pattern to implement the scraper. In this assessment, I have implemented two example factories for scraping CNN and The Guardian.

#### Implementation

The `ScraperFactory` in [`./scraper/scraper.go`](https://github.com/by-sabbir/optim-webscraping-test/blob/efe00f45029d3c56409e4105f32236551cc882d9/scraper/scraper.go) acts as an factory allowing us to create objects that implement the `ScraperFactory` interface without exposing the details of concrete types. If we need to create a new scraper in the future, we can do so by adding a new implementation and updating the factory function without changing the existing client code. It also allows us to set different configuration for different factories - 

```go
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
```

Finally, I have wrapped the whole solution into a `cli` tool. This same solution can easily be integrated in a HTTP or GRPC server proving the flexibility.

#### Testing

**Requires Go installation** - [guide](https://go.dev/doc/install)

```bash
make test
```

Expected Output:
```bash
?       github.com/by-sabbir/optim-webscraping-test     [no test files]
?       github.com/by-sabbir/optim-webscraping-test/cmd [no test files]
ok      github.com/by-sabbir/optim-webscraping-test/scraper     10.534s coverage: 97.4% of statements
```