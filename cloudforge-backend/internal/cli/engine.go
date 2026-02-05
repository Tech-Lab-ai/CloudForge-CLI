package cli

import (
	"fmt"
	"os"

	"cloudforge/internal/engine"
	"github.com/spf13/cobra"
)

// engineCmd representa o comando `cloudforge engine`
var engineCmd = &cobra.Command{
	Use:   "engine",
	Short: "Inicia a Engine do CloudForge em modo foreground (daemon).",
	Long: `Este comando inicia o servidor gRPC da Engine, que fica escutando por
requisições da CLI. Ele é o coração do CloudForge, responsável por executar
as operações de provisionamento e gerenciamento de estado.

Normalmente, você o executaria em um terminal separado ou como um serviço.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Inicializando o ciclo de vida da Engine...")

		// Cria e inicia o ciclo de vida da engine
		lifecycle, err := engine.NewLifecycle(cfgFile, environment)
		if err != nil {
			fmt.Printf("❌ Erro ao inicializar a Engine: %v\n", err)
			os.Exit(1)
		}

		// O método Start() é bloqueante e gerencia o graceful shutdown
		if err := lifecycle.Start(); err != nil {
			fmt.Printf("❌ A Engine encontrou um erro fatal: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Engine finalizada.")
	},
}
