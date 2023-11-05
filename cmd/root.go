package cmd

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/by-sabbir/optim-webscraping-test/scraper"
	"github.com/spf13/cobra"
)

var url, parser string
var logger = slog.New(slog.Default().Handler())

var rootCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Optimizely Content Intelligence and the Web",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		if err := scrape(url, parser); err != nil {
			logger.Error("something went wrong", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "https://www.theguardian.com/world/2023/nov/03/dozens-killed-and-injured-by-magnitude-64-earthquake-in-nepal", "this should be the url you want to parse")
	rootCmd.PersistentFlags().StringVarP(&parser, "parser", "p", "guardian", "the supported parser")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("something went wrong", "error", err)
		os.Exit(1)
	}
}

func scrape(url string, scraperName string) error {
	logger := slog.New(slog.Default().Handler())
	scraperServcie, err := scraper.NewScraperService(scraperName)
	if err != nil {
		logger.Error("service initiation faild", "error", err)
		return err
	}

	items, err := scraperServcie.ScrapePage(url)
	if err != nil {
		logger.Error("error sracping", "error", err)
		return err
	}

	b, err := json.Marshal(items.Metadata)
	if err != nil {
		logger.Error("could not marshall json", "error", err)
		return err
	}
	logger.Info("result", "data", string(b))
	return nil
}
