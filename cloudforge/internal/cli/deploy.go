package cli

import (
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys the infrastructure",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement deploy logic
	},
}
