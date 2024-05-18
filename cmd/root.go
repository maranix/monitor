package cmd

import (
	"log/slog"
	"os"

	"github.com/maranix/monitor/observer"
	"github.com/spf13/cobra"
)

const ReleaseVersion = "0.0.1-alpha"

var verbose bool
var rootCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor restart command/service on filesystem changes effortlessly",
	Long: `Monitor is a cli-tool built for development workflows where changes 
in configuration files or code requires restarting a service.`,
	Args: cobra.MinimumNArgs(2),
	Run:  handleRun,
}

func handleRun(cmd *cobra.Command, args []string) {
	if !verbose {
		slog.SetLogLoggerLevel(slog.LevelError)
	}

	pathArg := args[0]
	cmdArg := args[1]

	obs, err := observer.Create(pathArg, cmdArg)
	if err != nil {
		slog.Error("Failed to create an observable", err)
	}

	obs.Observe()
}

func addCommands() {
	rootCmd.AddCommand(versionCmd)
}

func addFlags() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "output additional information")
}

func Execute() {
	addCommands()
	addFlags()

	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
