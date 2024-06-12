package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	dir     string
	file    string
	run     []string
	ignore  []string
)

func init() {
	rootCmd.AddCommand(versionCmd)

	/*
	 *  Global Flags
	 */
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "./", "Specify the absolute path to the working directory.")

	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "./", "Specify the absolute path to the target file.")

	rootCmd.PersistentFlags().StringArrayVarP(&ignore, "ignore", "i", []string{}, "Exclude files/directories matching the provided glob pattern.")

	rootCmd.PersistentFlags().StringArrayVarP(&run, "run", "r", []string{}, "List services/commands to run and reload on changes.")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging for debugging purposes.")
}

var rootCmd = &cobra.Command{
	Use:   "mon",
	Short: "Monitor restart command/service on filesystem changes effortlessly",
	Long: `Monitor is a cli-tool built for development workflows where changes 
in configuration files or code requires restarting a service.`,
	Example: "mon ./ \"echo hello\"",
	Run:     handleRoot,
}

func handleRoot(cmd *cobra.Command, args []string) {
	// fmt.Println(cmd.Flags())
	fmt.Println(args)
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
