package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const releaseVersion = "pre_release-dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version number",
	Long:  `Prints the version number of Monitor CLI`,
	Run:   handleVersionRun,
}

func handleVersionRun(cmd *cobra.Command, args []string) {
	fmt.Println(releaseVersion)
}
