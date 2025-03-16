package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name string
	age  int
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet a person",
	Long:  `Greet a person with customizable name and age parameters.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Printf("Verbose mode enabled\n")
			fmt.Printf("Name set to %s\n", name)
			fmt.Printf("Age set to %d\n", age)
		}

		if age > 0 {
			fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
		} else {
			fmt.Printf("Hello, %s!\n", name)
		}

		if len(args) > 0 {
			fmt.Println("Additional arguments:", args)
		}
	},
}

func init() {
	// Local flags for the greet command
	greetCmd.Flags().StringVar(&name, "name", "World", "a name to greet")
	greetCmd.Flags().IntVar(&age, "age", 0, "age of the person")
}
