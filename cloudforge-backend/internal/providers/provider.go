package providers

import (
	"cloudforge/internal/config"
)

// Provider define a interface universal para todos os provedores de nuvem suportados.
// Cada provedor (AWS, GCP, Azure) deve implementar esta interface para que o
// Provisioner possa interagir com eles de forma agnóstica.
type Provider interface {
	// GetName retorna o nome do provedor (ex: "aws", "gcp").
	GetName() string

	// Authenticate inicializa e valida as credenciais para o provedor.
	Authenticate() error

    // GetResourceStatus busca o estado atual de um recurso no provedor cloud.
    // Utilizado principalmente pela detecção de drift.
    GetResourceStatus(providerID string) (*Resource, error)

	// --- Métodos de Provisionamento ---

	// ProvisionCompute cria ou atualiza os recursos de computação (ex: VMs).
	ProvisionCompute(cfg *config.ComputeConfig) (*Resource, error)

	// ProvisionNetwork cria ou atualiza os recursos de rede (ex: VPCs, subnets).
	ProvisionNetwork(cfg *config.NetworkConfig) (*Resource, error)

	// ProvisionStorage cria ou atualiza os recursos de armazenamento (ex: buckets).
	ProvisionStorage(cfg *config.StorageConfig) (*Resource, error)

	// --- Métodos de Desprovisionamento ---

	// DeprovisionResource remove um recurso específico com base no seu ID do provedor.
	DeprovisionResource(providerID string, resourceType string) error
}

// Resource é uma representação genérica de um recurso provisionado na nuvem.
// Contém informações que são relevantes para o StateManager e o DriftDetector.
type Resource struct {
	ProviderID string // O ID único do recurso no provedor (ex: "i-12345" na AWS)
	Type       string // O tipo de recurso (ex: "instance", "bucket")
	Attributes map[string]interface{}
}
