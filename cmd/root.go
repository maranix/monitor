package cmd

import (
	"log/slog"
	"os"

	"github.com/maranix/monitor/observe"
	"github.com/spf13/cobra"
)

const ReleaseVersion = "0.0.1-alpha"

var rootCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor restart command/service on filesystem changes effortlessly",
	Long: `Monitor is a cli-tool built for development workflows where changes 
in configuration files or code requires restarting a service.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pathArg := args[0]

		slog.Info(pathArg)

		obs, err := observe.Create(pathArg)
		if err != nil {
			slog.Error("Failed to create an observable", err)
		}

		obs.Observe()
	},
}

func addCommands() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	addCommands()

	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
