package providers

import (
	"fmt"

	"cloudforge/internal/config"
	"cloudforge/internal/providers/gcp"
    // Placeholder para outros provedores
    // "cloudforge/internal/providers/aws"
    // "cloudforge/internal/providers/azure"
)

// NewProviderFactory retorna uma instância do provedor apropriado com base na configuração.
// Esta é a maneira centralizada de inicializar um provedor para a aplicação.
func NewProviderFactory(cfg *config.Config) (Provider, error) {
	switch cfg.Provider {
	case "gcp":
		return gcp.NewGCPProvider(cfg.Project)
	case "aws":
		// return aws.NewAWSProvider(cfg.Region)
		return nil, fmt.Errorf("o provedor AWS ainda não é suportado")
	case "azure":
		// return azure.NewAzureProvider(cfg.SubscriptionID)
		return nil, fmt.Errorf("o provedor Azure ainda não é suportado")
	default:
		return nil, fmt.Errorf("provedor desconhecido: %s", cfg.Provider)
	}
}
