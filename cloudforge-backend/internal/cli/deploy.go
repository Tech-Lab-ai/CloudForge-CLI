package cli

import (
	"context"
	"fmt"

	"cloudforge/proto"
	"github.com/spf13/cobra"
)

// deployCmd representa o comando `cloudforge deploy`
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Provisiona ou atualiza a infraestrutura de acordo com a configuração.",
	Long: `O comando deploy analisa o arquivo de configuração, compara com o estado
   atual e aplica as mudanças necessárias no provedor de nuvem para que a
   infraestrutura real corresponda à configuração desejada.
    `,
	Run: func(cmd *cobra.Command, args []string) {
		req := &proto.DeployRequest{
			Environment: environment,
		}

		fmt.Println("🚀 Disparando o processo de deploy...")

		resp, err := grpcClient.Deploy(context.Background(), req)
		if err != nil {
			fmt.Printf("❌ Erro durante o deploy: %v\n", err)
			return
		}

		fmt.Printf("✅ %s (Versão do Estado: %d)\n", resp.Message, resp.StateVersion)
	},
}
