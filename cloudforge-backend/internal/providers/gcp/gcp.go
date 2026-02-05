package gcp

import (
	"context"
	"fmt"

	"cloudforge/internal/config"
	"cloudforge/internal/providers"
	"golang.org/x/oauth2/google"
)

// GCPProvider implementa a interface providers.Provider para o Google Cloud Platform.
type GCPProvider struct {
	ProjectID string
	// Adicionar clientes para os serviços do GCP aqui (ex: compute, storage)
}

// NewGCPProvider cria e configura um novo provedor para o GCP.
func NewGCPProvider(projectID string) (*GCPProvider, error) {
	if projectID == "" {
		return nil, fmt.Errorf("o ID do projeto GCP é obrigatório")
	}
	return &GCPProvider{ProjectID: projectID}, nil
}

// GetName retorna o nome do provedor.
func (p *GCPProvider) GetName() string {
	return "gcp"
}

// Authenticate verifica se as credenciais do GCP estão configuradas corretamente.
// Utiliza o Application Default Credentials (ADC).
func (p *GCPProvider) Authenticate() error {
	fmt.Println("🔐 Autenticando com o Google Cloud...")
	creds, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		return fmt.Errorf("falha ao encontrar as credenciais padrão do GCP (ADC): %w", err)
	}

	if creds.ProjectID == "" {
        fmt.Println("⚠️  Aviso: O Project ID não foi encontrado nas credenciais. Usando o ID do projeto da configuração.")
    } else {
        p.ProjectID = creds.ProjectID // Usa o project ID das credenciais se disponível
    }

	fmt.Println("✅ Autenticação com o GCP bem-sucedida.")
	return nil
}

// GetResourceStatus (ainda não implementado).
func (p *GCPProvider) GetResourceStatus(providerID string) (*providers.Resource, error) {
    return nil, fmt.Errorf("GetResourceStatus ainda não implementado para GCP")
}

// --- Métodos de Provisionamento (Stubs) ---

func (p *GCPProvider) ProvisionCompute(cfg *config.ComputeConfig) (*providers.Resource, error) {
	fmt.Printf("[GCP] Provisionando instância de computação: %v\n", cfg)
	// Lógica de provisionamento real seria aqui

	return &providers.Resource{
		ProviderID: "gcp-instance-12345",
		Type:       "gce_instance",
		Attributes: map[string]interface{}{"instance_type": cfg.InstanceType, "image": cfg.Image},
	}, nil
}

func (p *GCPProvider) ProvisionNetwork(cfg *config.NetworkConfig) (*providers.Resource, error) {
	fmt.Printf("[GCP] Provisionando rede: %v\n", cfg)
	return &providers.Resource{
		ProviderID: "gcp-vpc-12345",
		Type:       "gcp_vpc",
		Attributes: map[string]interface{}{"vpc_cidr": cfg.VPC_CIDR},
	}, nil
}

func (p *GCPProvider) ProvisionStorage(cfg *config.StorageConfig) (*providers.Resource, error) {
	fmt.Printf("[GCP] Provisionando bucket de armazenamento: %v\n", cfg)
	return &providers.Resource{
		ProviderID: "gcp-bucket-12345",
		Type:       "gcs_bucket",
		Attributes: map[string]interface{}{"bucket_name": cfg.BucketName},
	}, nil
}

// DeprovisionResource (ainda não implementado).
func (p *GCPProvider) DeprovisionResource(providerID string, resourceType string) error {
    return fmt.Errorf("DeprovisionResource ainda não implementado para GCP")
}
