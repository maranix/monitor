package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version number",
	Long:  `Prints the version number of Monitor CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ReleaseVersion)
	},
}
