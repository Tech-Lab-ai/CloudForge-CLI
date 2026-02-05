package cli

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new CloudForge project",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement init logic
	},
}
