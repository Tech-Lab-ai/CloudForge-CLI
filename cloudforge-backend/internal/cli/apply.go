package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the configuration to the specified cloud provider",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := readConfig()
		if err != nil {
			log.Fatalf("error reading config: %v", err)
		}

		provider, err := getProvider(config.Provider)
		if err != nil {
			log.Fatalf("error getting provider: %v", err)
		}

		fmt.Println("Applying configuration...")
		provider.Apply()
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
