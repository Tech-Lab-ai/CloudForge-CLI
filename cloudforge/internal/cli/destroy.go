package cli

import (
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroys the infrastructure",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement destroy logic
	},
}
