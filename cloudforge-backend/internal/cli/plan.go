package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Generate and display an execution plan",
	Long:  `Generate and display an execution plan for the current project configuration.`,
	Run:   runPlanCmd,
}

func runPlanCmd(cmd *cobra.Command, args []string) {
	config, err := readConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	provider, err := getProvider(config.Provider)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	plan, err := provider.Plan()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(plan)
}
