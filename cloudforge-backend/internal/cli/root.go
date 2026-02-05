package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cloudforge",
	Short: "CloudForge is a tool for managing cloud infrastructure.",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(planCmd)
}
