package config

import (
	"fmt"
	"github.com/spf13/viper"
    "strings"
)

// LoadConfig carrega a configuração do arquivo `cloudforge.yaml` e valida sua estrutura.
// Utiliza o Viper para permitir flexibilidade (ex: override por variáveis de ambiente).
func LoadConfig(path string, environment string) (*Config, error) {
	v := viper.New()

	v.SetConfigFile(path) // Define o caminho explícito para o arquivo de config
    v.SetConfigType("yaml")

    // Habilita a leitura de variáveis de ambiente com o prefixo CLOUDFORGE
    v.SetEnvPrefix("CLOUDFORGE")
    v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo de configuração: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a configuração: %w", err)
	}

    // Valida a configuração carregada
    if err := validateConfig(&cfg); err != nil {
        return nil, err
    }

    // Define o ambiente atual que está sendo operado
    if env := cfg.GetEnvironment(environment); env == nil {
        return nil, fmt.Errorf("ambiente '%s' não encontrado no arquivo de configuração", environment)
    }
    cfg.CurrentEnvironment = environment

	return &cfg, nil
}

// validateConfig executa uma série de checagens para garantir que a config é válida.
func validateConfig(cfg *Config) error {
    if cfg.Project == "" {
        return fmt.Errorf("o campo 'project' é obrigatório no arquivo de configuração")
    }

    if cfg.Provider == "" {
        return fmt.Errorf("o campo 'provider' é obrigatório")
    }

    if len(cfg.Environments) == 0 {
        return fmt.Errorf("nenhum ambiente foi definido no arquivo de configuração")
    }

    // Valida se o provedor é um dos suportados
    supportedProviders := map[string]bool{"aws": true, "gcp": true, "azure": true}
    if !supportedProviders[cfg.Provider] {
        return fmt.Errorf("provedor '%s' não é suportado. Válidos: aws, gcp, azure", cfg.Provider)
    }

    return nil
}
