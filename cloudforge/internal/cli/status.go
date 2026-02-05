package cli

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows the status of the infrastructure",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement status logic
	},
}
