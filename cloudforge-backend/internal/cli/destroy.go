package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the infrastructure created by CloudForge",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := readConfig()
		if err != nil {
			log.Fatalf("error reading config: %v", err)
		}

		provider, err := getProvider(config.Provider)
		if err != nil {
			log.Fatalf("error getting provider: %v", err)
		}

		fmt.Println("Destroying infrastructure...")
		provider.Destroy()
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
