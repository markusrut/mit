package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version information
var (
	Version   = "0.1.0"
	BuildTime = "unknown"
	Commit    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version, build time and commit of this application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", Version)
		if verbose {
			fmt.Printf("Build Time: %s\n", BuildTime)
			fmt.Printf("Commit: %s\n", Commit)
		}
	},
}
