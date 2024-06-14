package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mon",
	Short: "Monitor restart command/service on filesystem changes effortlessly",
	Long: `Monitor is a cli-tool built for development workflows where changes 
in configuration files or code requires restarting a service.`,
	Example: "mon ./ \"echo hello\"",
	Run:     handleRootRun,
}

var (
	// Debounce duration in seconds
	// Duration to wait before re-executing the command on detecting
	// subsequent changes in a short succession.
	//
	// Defaults to 300ms.
	debounce float32

	// A slice of glob patterns or path to dirs/files
	//
	// Cannot be a combination of two, keep it as simple as it can be.
	//
	// Defaults to an empty slice.
	ignore []string

	// Command/Service to run
	//
	// Default to an empty string.
	run string

	// Path to the target to watch
	//
	// Can be either a directory or a file
	//
	// Default to an empty string.
	target string

	// Verbose logging for debugging
	//
	// Default to false.
	verbose bool
)

func init() {
	rootCmd.AddCommand(versionCmd)

	/*
	 *  Global Flags
	 */
	rootCmd.PersistentFlags().Float32VarP(&debounce, "debounce", "d", 0.3, "Exclude files/directories matching the provided glob pattern.")

	rootCmd.PersistentFlags().StringSliceVarP(&ignore, "ignore", "i", []string{}, "Exclude files/directories matching the provided glob pattern.")

	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "./", "Specify the absolute path of the target directory or file.")

	rootCmd.PersistentFlags().StringVarP(&run, "run", "r", "", "List services/commands to run and reload on changes.")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging for debugging purposes.")
}

func handleRootRun(cmd *cobra.Command, args []string) {
	// TODO: Validate and replace vars if needed and then pass necessary
	//       vars to observer module.
	//
	// Failure: In-case of failure log the reason of failure and exit
	//

	// pathArg := args[0]
	// cmdArg := args[1]
	//
	// obs, err := observer.Create(pathArg, cmdArg)
	// if err != nil {
	// 	slog.Error("Failed to create an observable", err)
	// }
	//
	// obs.Observe()
}
