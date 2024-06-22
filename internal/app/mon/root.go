package mon

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
	Args:    cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		err := handleArgs(args)
		if err != nil {
			cmd.Help()
			fmt.Println(err.Error())
			os.Exit(1)
		}

		obs, err := observer.NewObserver(cfg)
		if err != nil {
			cmd.Help()
			fmt.Println(formatError(err))
			os.Exit(1)
		}

		defer obs.Close()

		obs.Observe()
	},
}

func init() {
	cfg = createConfig()

	rootCmd.AddCommand(versionCmd)

	/*
	 *  Global Flags
	 */
	rootCmd.PersistentFlags().Float32VarP(&cfg.debounce, "debounce", "d", 0.3, "Time to wait in-Seconds between events before restarting the runner")

	rootCmd.PersistentFlags().StringSliceVarP(&cfg.ignore, "ignore", "i", []string{}, "Exclude files/directories matching the provided glob pattern.")

	rootCmd.PersistentFlags().StringVarP(&cfg.target, "target", "t", "", "Path of the target directory or file to monitor.")

	rootCmd.PersistentFlags().StringVarP(&cfg.run, "run", "r", "", "service to run and re-run on filesytem changes.")

	rootCmd.PersistentFlags().BoolVarP(&cfg.verbose, "verbose", "v", false, "Enable verbose logging for debugging purposes.")
}

func handleArgs(args []string) error {
	if err := validateArgs(args); err != nil {
		return formatError(err)
	}

	if err := resolveAbsTargetPath(cfg.target); err != nil {
		return formatError(err)
	}

	if err := validateTarget(cfg.target); err != nil {
		return formatError(err)
	}

	if err := validateRunner(cfg.run); err != nil {
		return formatError(err)
	}

	return nil
}

func validateArgs(args []string) error {
	argsCount := len(args)
	if argsCount > 0 && argsCount != 2 {
		return errors.New("**Invalid Args:**\n\nExpected to receive exactly 2 positional args.")
	}

	// Positional parameters take precedence over flags in case of target and run
	// are provided via positonal parameters
	//
	// t = target, r = run
	if argsCount > 0 {
		cfg.target = args[0]
		cfg.run = args[1]
	}

	return nil
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

func resolveAbsTargetPath(path string) error {
	absPath, err := fsutil.AbsPath(path)
	if err != nil {
		return err
	}

	cfg.target = absPath
	return nil
}

func formatError(err error) error {
	msg := fmt.Sprintf("\n%s", err.Error())
	return errors.New(msg)
}
