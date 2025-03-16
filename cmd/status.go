package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/markusrut/mit/fs"
	"github.com/spf13/cobra"
)

var (
	path string
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of working directory",
	Long:  `Show the status of the working directory.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	if verbose {
	// 		fmt.Printf("Verbose mode enabled\n")
	// 	}

	// 	fmt.Printf("Status\n")

	// 	if len(args) > 0 {
	// 		fmt.Println("Additional arguments:", args)
	// 	}
	// },

	RunE: func(cmd *cobra.Command, args []string) error {
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
	// Local flags for the greet command
	statusCmd.Flags().StringVar(&path, "path", "testfiles", "Path to the directory to show the status of")
}
