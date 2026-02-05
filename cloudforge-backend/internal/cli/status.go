package cli

import (
	"context"
	"fmt"

	"cloudforge/proto"
	"github.com/spf13/cobra"
)

// statusCmd representa o comando `cloudforge status`
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Exibe o status atual da infraestrutura gerenciada.",
	Long: `O comando status carrega o estado atual salvo pelo CloudForge e o exibe
   de forma legível, mostrando os recursos provisionados, a versão do estado
   e o provedor em uso. Futuramente, também detectará drift.
    `,
	Run: func(cmd *cobra.Command, args []string) {
		req := &proto.StatusRequest{
			Environment: environment,
		}

		fmt.Println("🔎 Verificando o status da infraestrutura...")

		resp, err := grpcClient.Status(context.Background(), req)
		if err != nil {
			fmt.Printf("❌ Erro ao obter o status: %v\n", err)
			return
		}

        fmt.Println("--- Status do CloudForge ---")
		fmt.Printf("Status:        %s\n", resp.Status)
		fmt.Printf("Provedor:      %s\n", resp.Provider)
		fmt.Printf("Versão Estado: %d\n", resp.StateVersion)
        fmt.Printf("Drift:         %t\n", resp.DriftDetected)
        fmt.Println("---------------------------")

	},
}
