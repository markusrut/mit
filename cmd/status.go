package cmd

import (
	"fmt"

	"github.com/markusrut/mit/entity"
	"github.com/spf13/cobra"
)

var (
	path        string
	defaultPath = "testfiles"
)

var statusCmd = &cobra.Command{
	Use:   "status [path]",
	Short: "Status of working directory",
	Long:  `Show the status of the working directory.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 && path == defaultPath {
			path = args[0]
		}

		if verbose {
			fmt.Printf("Verbose mode enabled\n")
			fmt.Printf("Path set to %s\n", path)
		}

		// Read the file structure
		structure, err := entity.ReadFileStructure(path)
		if err != nil {
			return fmt.Errorf("failed to read file structure: %w", err)
		}

		structure.Print()

		return nil
	},
}

func init() {
	statusCmd.Flags().StringVarP(&path, "path", "p", defaultPath, "Path to the directory to show the status of")

	rootCmd.AddCommand(statusCmd)
}
