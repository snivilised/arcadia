package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version number",
	Long:  fmt.Sprintf("Print the version number of %v", ApplicationName),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v %v version: '%v'\n", AppEmoji, ApplicationName, Version)
	},
}

func init() {
	// CobraCommandContainer.RegisterRootedCommand(versionCmd)
	rootCmd.AddCommand(versionCmd)
}
