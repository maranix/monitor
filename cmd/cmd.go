package cmd

func Execute() error {
	err := rootCmd.Execute()
	return err
}
