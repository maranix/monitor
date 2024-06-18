package cli

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

// Creates a new empty Config.
func createConfig() *Config {
	cfg := Config{
		debounce: 0,
		ignore:   []string{},
		run:      "",
		target:   "",
		verbose:  false,
	}

	return &cfg
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

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
