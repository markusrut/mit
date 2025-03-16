package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Flags that are available to all commands
	verbose bool
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "mit",
	Short: "My git, a simple git-like command line application",
	Long: `My git
	A simple git-like command line application`,
}

// Execute adds all child commands to the root command and sets flags
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbose output")

	// Add subcommands
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(statusCmd)
}
