package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/markusrut/mit/fs"
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
		structure, err := fs.ReadFileStructure(path)
		if err != nil {
			return fmt.Errorf("failed to read file structure: %w", err)
		}

		// Output as JSON for demonstration (can be modified for other formats)
		data, err := json.MarshalIndent(structure, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal to JSON: %w", err)
		}

		fmt.Println(string(data))
		return nil
	},
}

func init() {
	statusCmd.Flags().StringVarP(&path, "path", "p", defaultPath, "Path to the directory to show the status of")

	rootCmd.AddCommand(statusCmd)
}
