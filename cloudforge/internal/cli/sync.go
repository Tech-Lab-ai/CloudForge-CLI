package cli

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes the local environment with the cloud",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement sync logic
	},
}
