package mon

import (
	"context"
	"log/slog"
	"os"

	"github.com/maranix/monitor/pkg/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cli.ExecuteContext(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
