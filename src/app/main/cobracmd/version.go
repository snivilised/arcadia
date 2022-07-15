package cobracmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version number",
	Long:  fmt.Sprintf("Print the version number of %v", ApplicationName),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v %v version: '%v'\n", AppEmoji, ApplicationName, version)
	},
}

func init() {
	// CobraCommandContainer.RegisterRootedCommand(versionCmd)
	rootCmd.AddCommand(versionCmd)
}
