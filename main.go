package main

import (
	"log/slog"

	"github.com/by-sabbir/optim-webscraping-test/cmd"
)

var build = "dev"

func main() {
	logger := slog.New(slog.Default().Handler())

	logger.Info("running scraper", "version", build)
	cmd.Execute()
}
