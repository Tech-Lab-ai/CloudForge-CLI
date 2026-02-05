package cli

import (
	"context"
	"fmt"
	"os"

	"cloudforge/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

var (
	// Flags globais
	cfgFile     string
	environment string

	// Cliente gRPC compartilhado entre os comandos
	grpcClient proto.CloudForgeServiceClient
    grpcConn *grpc.ClientConn
)

// rootCmd representa o comando base `cloudforge`
var rootCmd = &cobra.Command{
	Use:   "cloudforge",
	Short: "CloudForge é uma CLI para gerenciar infraestrutura de nuvem de forma declarativa.",
	Long: `CloudForge simplifica o provisionamento e o gerenciamento de recursos de nuvem
   através de um arquivo de configuração YAML, permitindo que você se concentre no
   desenvolvimento de sua aplicação.

   Ele opera em um modelo cliente-servidor, onde a CLI se comunica com uma Engine
   (daemon) em execução para executar as operações.
    `,
    // PersistentPreRunE é executado antes de qualquer comando filho.
    // É usado aqui para inicializar a conexão gRPC com a Engine.
    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
        // O comando 'engine' não precisa de uma conexão gRPC, ele inicia o servidor.
        if cmd.Name() == "engine" {
            return nil
        }

        conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
        if err != nil {
            fmt.Println("❌ Erro ao conectar-se à Engine do CloudForge. A Engine está em execução?")
            fmt.Println("   Você pode iniciá-la com o comando: `cloudforge engine`")
            return fmt.Errorf("falha ao conectar via gRPC: %w", err)
        }

        grpcClient = proto.NewCloudForgeServiceClient(conn)
        grpcConn = conn
        return nil
    },
    PersistentPostRun: func(cmd *cobra.Command, args []string) {
        // Fecha a conexão gRPC após a execução do comando
        if grpcConn != nil {
            grpcConn.Close()
        }
    },
}

// Execute é o ponto de entrada principal para a CLI.
// Ele adiciona todos os comandos filhos ao comando raiz e executa-o.
func Execute() {
	ctx := context.Background()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Adiciona as flags globais ao comando raiz
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "cloudforge.yaml", "Caminho para o arquivo de configuração (padrão é cloudforge.yaml)")
	rootCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "Ambiente a ser usado (ex: dev, prod). Sobrescreve o default_environment.")

	// Adiciona os subcomandos
	rootCmd.AddCommand(engineCmd)
    rootCmd.AddCommand(initCmd)
    rootCmd.AddCommand(deployCmd)
    rootCmd.AddCommand(statusCmd)
}
