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

var cfg *Config

var rootCmd = &cobra.Command{
	Use:     "mon",
	Short:   "Monitor restart command/service on filesystem changes effortlessly",
	Long:    `Monitor is a cli-tool built for development workflows where changes in configuration files or code requires restarting a service.`,
	Example: "mon ./ \"echo hello\"",
	Run:     handleRootRun,
}

func init() {
	cfg = createConfig()

	rootCmd.AddCommand(versionCmd)

	/*
	 *  Global Flags
	 */
	rootCmd.PersistentFlags().Float32VarP(&cfg.debounce, "debounce", "d", 0.3, "Time to wait in-Seconds between events before restarting the runner")

	rootCmd.PersistentFlags().StringSliceVarP(&cfg.ignore, "ignore", "i", []string{}, "Exclude files/directories matching the provided glob pattern.")

	rootCmd.PersistentFlags().StringVarP(&cfg.target, "target", "t", "./", "Specify the absolute path of the target directory or file.")

	rootCmd.PersistentFlags().StringVarP(&cfg.run, "run", "r", "", "List services/commands to run and reload on changes.")

	rootCmd.PersistentFlags().BoolVarP(&cfg.verbose, "verbose", "v", false, "Enable verbose logging for debugging purposes.")
}

func handleRootRun(cmd *cobra.Command, args []string) {
	err := validateArgs(args)
	if err != nil {
		cmd.Help()
		fmt.Println(fmt.Sprintf("\n%s", err.Error()))
		os.Exit(1)
	}

	obs, err := observer.NewObserver(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}

	obs.Observe()
}

func validateArgs(args []string) error {
	argsCount := len(args)

	if argsCount > 0 && argsCount != 2 {
		return errors.New("**Invalid Args:**\n\nExpected to receive exactly 2 positional args.")
	}

	// Positional parameters takes precedence over flags in case of target and run
	// are provided via positonal parameters
	//
	// t = target, r = run
	if argsCount > 0 {
		cfg.target = args[0]
		cfg.run = args[1]
	}

	// Seems much more cleaner, intuitive and easier to read than
	//  ```go
	//  err := validateTarget(someTarget)
	//  if err != nil {
	//      return err
	//  }
	//
	//  err = validateRunner(someRunner)
	//  if err != nil {
	//      return err
	//  }
	//
	//  return nil
	//  ```
	var err error
	err = validateTarget(cfg.target)
	err = validateRunner(cfg.run)

	return err
}

func validateTarget(target string) error {
	if target == "" {
		return errors.New("**Missing Target:**\n\nPlease specify a target to monitor using the `--target` (or `-t`) flag.")
	}

	err := fsutil.Validate(target)
	if err != nil {
		return err
	}

	return nil
}

func validateRunner(run string) error {
	if run == "" {
		return errors.New("**Missing Runner:**\n\nPlease specify a runner to execute using the `--run` (or `-r`) flag.")
	}

	err := runner.Validate(run)
	if err != nil {
		return err
	}

	return nil
}
