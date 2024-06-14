package cli

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
