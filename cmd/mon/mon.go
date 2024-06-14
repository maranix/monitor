package main

import (
	"log/slog"
	"os"

	"github.com/maranix/monitor/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
