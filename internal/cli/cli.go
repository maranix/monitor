package cli

import (
	"context"
)

func ExecuteWithContext(ctx context.Context) error {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}
