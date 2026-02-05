package provisioner

import (
	"fmt"

	"cloudforge/internal/config"
	"cloudforge/internal/providers"
	"cloudforge/internal/state"
)

// Provisioner é o motor central que traduz a configuração em recursos de nuvem.
// Ele orquestra as chamadas ao provedor e atualiza o estado da infraestrutura.
type Provisioner struct {
	provider providers.Provider
}

// NewProvisioner cria um novo provisionador, configurado com um provedor específico.
func NewProvisioner(p providers.Provider) *Provisioner {
	return &Provisioner{provider: p}
}

// Plan compara a configuração desejada com o estado atual e gera um plano de execução.
// (Implementação futura, por enquanto focamos no Apply direto).
func (p *Provisioner) Plan(cfg *config.Config, st *state.State) {
    fmt.Println("Funcionalidade de Planejamento (Plan) ainda não implementada.")
}

// Apply executa as ações necessárias para alinhar a infraestrutura com a configuração.
// Ele provisiona ou atualiza os recursos conforme necessário.
func (p *Provisioner) Apply(cfg *config.Config, st *state.State) error {
	fmt.Printf("🚀 Iniciando o provisionamento com o provedor: %s\n", p.provider.GetName())

    // Autentica o provedor antes de qualquer operação
    if err := p.provider.Authenticate(); err != nil {
        return fmt.Errorf("falha na autenticação do provedor: %w", err)
    }

    // Busca a configuração do ambiente atual
    envCfg := cfg.GetEnvironment(cfg.CurrentEnvironment)
    if envCfg == nil {
        return fmt.Errorf("configuração para o ambiente '%s' não encontrada", cfg.CurrentEnvironment)
    }

    // --- Lógica de Provisionamento ---
    // Aqui iteramos sobre a configuração e provisionamos cada tipo de recurso.

    if envCfg.Compute != nil {
        resource, err := p.provider.ProvisionCompute(envCfg.Compute)
        if err != nil {
            return fmt.Errorf("erro ao provisionar computação: %w", err)
        }
        st.AddResource(*resource, "compute", "virtual_machine")
    }

    if envCfg.Network != nil {
        resource, err := p.provider.ProvisionNetwork(envCfg.Network)
        if err != nil {
            return fmt.Errorf("erro ao provisionar rede: %w", err)
        }
        st.AddResource(*resource, "network", "vpc")
    }

    if envCfg.Storage != nil {
        resource, err := p.provider.ProvisionStorage(envCfg.Storage)
        if err != nil {
            return fmt.Errorf("erro ao provisionar armazenamento: %w", err)
        }
        st.AddResource(*resource, "storage", "bucket")
    }

    fmt.Println("✅ Provisionamento concluído com sucesso.")
	return nil
}
