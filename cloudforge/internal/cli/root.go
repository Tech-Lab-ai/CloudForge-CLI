package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cloudforge",
	Short: "CloudForge is a CLI for managing cloud infrastructure",
	Long:  `A powerful and flexible CLI tool to provision, sync, and manage your cloud environments with ease.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(destroyCmd)
}
