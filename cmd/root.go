package cmd

import (
	"log/slog"

	"github.com/maranix/monitor/observer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor restart command/service on filesystem changes effortlessly",
	Long: `Monitor is a cli-tool built for development workflows where changes 
in configuration files or code requires restarting a service.`,
	Args: cobra.MinimumNArgs(2),
	Run:  handleRoot,
}

func handleRoot(cmd *cobra.Command, args []string) {
	pathArg := args[0]
	cmdArg := args[1]

	obs, err := observer.Create(pathArg, cmdArg)
	if err != nil {
		slog.Error("Failed to create an observable", err)
	}

	obs.Observe()
}
