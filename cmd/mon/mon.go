package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/maranix/monitor/internal/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cli.ExecuteWithContext(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
