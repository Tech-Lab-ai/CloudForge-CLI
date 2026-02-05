package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new CloudForge project",
	Long:  `Initialize a new CloudForge project by creating a cloudforge.yaml file in the current directory.`,
	Run:   runInitCmd,
}

func runInitCmd(cmd *cobra.Command, args []string) {
	config := &Config{
		Project:  "my-project",
		Provider: "aws",
		Region:   "us-east-1",
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if err := os.WriteFile("cloudforge.yaml", data, 0644); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("CloudForge project initialized successfully.")
}
