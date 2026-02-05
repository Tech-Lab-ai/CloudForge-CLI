package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the infrastructure defined in cloudforge.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := readConfig()
		if err != nil {
			log.Fatalf("error reading config: %v", err)
		}

		provider, err := getProvider(config.Provider)
		if err != nil {
			log.Fatalf("error getting provider: %v", err)
		}

		path, _ := cmd.Flags().GetString("path")

		fmt.Println("Deploying infrastructure...")
		provider.Deploy(path)
	},
}

func init() {
	deployCmd.Flags().String("path", ".", "The path to the application to be deployed")
	rootCmd.AddCommand(deployCmd)
}
