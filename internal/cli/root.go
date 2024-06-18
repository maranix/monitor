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

type Config struct {
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

var rootCmd = &cobra.Command{
	Use:     "mon",
	Short:   "Monitor restart command/service on filesystem changes effortlessly",
	Long:    `Monitor is a cli-tool built for development workflows where changes in configuration files or code requires restarting a service.`,
	Example: "mon ./ \"echo hello\"",
	Run:     handleRootRun,
}

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

// Creates a new Config based on Global Flag vars.
func createConfig() *Config {
	cfg := Config{
		debounce: debounce,
		ignore:   ignore,
		run:      run,
		target:   target,
		verbose:  verbose,
	}

	return &cfg
}

func handleRootRun(cmd *cobra.Command, args []string) {
	err := validateArgs(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cfg := createConfig()
	obs, err := observer.NewObserver(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}

	obs.Observe()
}

func validateArgs(args []string) error {
	// Positional parameters takes precedence over flags in case of target and run
	//
	// t = target, r = run
	t, r := args[0], args[1]

	err := resolveAndValidateTarget(t)
	if err != nil {
		return err
	}

	err = resolveAndValidateRunner(r)
	if err != nil {
		return err
	}

	return nil
}

func resolveAndValidateTarget(t string) error {
	if t == "" && target == "" {
		return errors.New("**Missing Target:**\nPlease specify a target to monitor using the `--target` (or `-t`) flag. See the help documentation for details.")
	}

	if t != "" {
		target = t
	}

	err := fsutil.Validate(target)
	if err != nil {
		return err
	}

	target, err = fsutil.AbsPath(target)
	if err != nil {
		return err
	}

	return nil
}

func resolveAndValidateRunner(r string) error {
	if r == "" && run == "" {
		return errors.New("**Missing Runner Target:**\nPlease specify a runner target to monitor using the `--run` (or `-r`) flag. See the help documentation for details.")
	}

	if r != "" {
		run = r
	}

	err := runner.Validate(run)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) GetDebounce() float32 {
	return c.debounce
}

func (c *Config) GetIgnoreTarget() []string {
	return c.ignore
}

func (c *Config) GetRunner() string {
	return c.run
}

func (c *Config) GetTarget() string {
	return c.target
}

func (c *Config) GetVerbose() bool {
	return c.verbose
}
