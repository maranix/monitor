package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

const releaseVersion = "pre-alpha"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version number",
	Long:  `Prints the version number of Monitor CLI`,
	Args:  cobra.MaximumNArgs(0),
	Run:   handleVersionRun,
}

func handleVersionRun(cmd *cobra.Command, _ []string) {
	fmt.Println(releaseVersion)
}
