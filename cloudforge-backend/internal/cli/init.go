package cli

import (
	"context"
	"fmt"

	"cloudforge/proto"
	"github.com/spf13/cobra"
)

// initCmd representa o comando `cloudforge init`
var initCmd = &cobra.Command{
	Use:   "init [project-id] [provider]",
	Short: "Inicializa um novo projeto CloudForge.",
	Long: `Cria o arquivo de configuração inicial (cloudforge.yaml) e prepara
   o diretório para o gerenciamento com o CloudForge. Você deve especificar
   um ID para o projeto e um provedor de nuvem (ex: gcp, aws).
    `,
	Args: cobra.ExactArgs(2), // Exige exatamente 2 argumentos
	Run: func(cmd *cobra.Command, args []string) {
        projectID := args[0]
        provider := args[1]

		req := &proto.InitRequest{
			ProjectId:   projectID,
			Provider:    provider,
            Environment: environment, // Passa o ambiente da flag global
		}

		fmt.Printf("🚀 Inicializando projeto '%s' com o provedor '%s'...\n", projectID, provider)

		resp, err := grpcClient.Init(context.Background(), req)
		if err != nil {
			fmt.Printf("❌ Erro ao inicializar o projeto: %v\n", err)
			return
		}

		fmt.Printf("✅ %s\n", resp.Message)
	},
}
