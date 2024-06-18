package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/maranix/monitor/pkg/fsutil"
	"github.com/maranix/monitor/pkg/observer"
	"github.com/maranix/monitor/pkg/runner"
	"github.com/spf13/cobra"
)

var cfg Config

var rootCmd = &cobra.Command{
	Use:     "mon",
	Short:   "Monitor restart command/service on filesystem changes effortlessly",
	Long:    `Monitor is a cli-tool built for development workflows where changes in configuration files or code requires restarting a service.`,
	Example: "mon ./ \"echo hello\"",
	Run:     handleRootRun,
}

func init() {
	cfg := createConfig()

	rootCmd.AddCommand(versionCmd)

	/*
	 *  Global Flags
	 */
	rootCmd.PersistentFlags().Float32VarP(&cfg.debounce, "debounce", "d", 0.3, "Exclude files/directories matching the provided glob pattern.")

	rootCmd.PersistentFlags().StringSliceVarP(&cfg.ignore, "ignore", "i", []string{}, "Exclude files/directories matching the provided glob pattern.")

	rootCmd.PersistentFlags().StringVarP(&cfg.target, "target", "t", "./", "Specify the absolute path of the target directory or file.")

	rootCmd.PersistentFlags().StringVarP(&cfg.run, "run", "r", "", "List services/commands to run and reload on changes.")

	rootCmd.PersistentFlags().BoolVarP(&cfg.verbose, "verbose", "v", false, "Enable verbose logging for debugging purposes.")
}

func handleRootRun(cmd *cobra.Command, args []string) {
	err := validateArgs(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	obs, err := observer.NewObserver(&cfg)
	if err != nil {
		fmt.Println(err.Error())
	}

	obs.Observe()
}

func validateArgs(args []string) error {
	// Positional parameters takes precedence over flags in case of target and run
	//
	// t = target, r = run
	tPos, rPos := args[0], args[1]

	err := resolveAndValidateTarget(cfg.target, tPos)
	if err != nil {
		return err
	}

	err = resolveAndValidateRunner(cfg.run, rPos)
	if err != nil {
		return err
	}

	return nil
}

func resolveAndValidateTarget(def string, pos string) error {
	if def == "" && pos == "" {
		return errors.New("**Missing Target:**\nPlease specify a target to monitor using the `--target` (or `-t`) flag. See the help documentation for details.")
	}

	if pos != "" {
		cfg.target = pos
	}

	err := fsutil.Validate(cfg.target)
	if err != nil {
		return err
	}

	cfg.target, err = fsutil.AbsPath(cfg.target)
	if err != nil {
		return err
	}

	return nil
}

func resolveAndValidateRunner(def string, pos string) error {
	if def == "" && pos == "" {
		return errors.New("**Missing Runner Target:**\nPlease specify a runner target to monitor using the `--run` (or `-r`) flag. See the help documentation for details.")
	}

	if pos != "" {
		cfg.run = pos
	}

	err := runner.Validate(cfg.run)
	if err != nil {
		return err
	}

	return nil
}
