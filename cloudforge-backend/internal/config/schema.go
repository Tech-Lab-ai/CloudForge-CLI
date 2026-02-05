package config

// Config define a estrutura do arquivo de configuração `cloudforge.yaml`.
// Esta é a fonte da verdade para o provisionamento da infraestrutura.
type Config struct {
	Project            string              `mapstructure:"project"`
	Provider           string              `mapstructure:"provider"`
	DefaultEnvironment string              `mapstructure:"default_environment"`
	Environments       []*Environment      `mapstructure:"environments"`
    CurrentEnvironment string              `mapstructure:"-"` // Preenchido dinamicamente
}

// Environment representa a configuração para um ambiente específico (ex: dev, prod).
type Environment struct {
	Name    string         `mapstructure:"name"`
	Compute *ComputeConfig `mapstructure:"compute,omitempty"`
	Network *NetworkConfig `mapstructure:"network,omitempty"`
	Storage *StorageConfig `mapstructure:"storage,omitempty"`
}

// --- Configurações de Recursos ---

// ComputeConfig define os parâmetros para recursos de computação (ex: VMs, containers).
type ComputeConfig struct {
	InstanceType string `mapstructure:"instance_type"` // ex: t2.micro, n1-standard-1
	Image        string `mapstructure:"image"`         // ex: ami-1234, ubuntu-2004-lts
	Replicas     int    `mapstructure:"replicas"`
}

// NetworkConfig define os parâmetros para a configuração de rede.
type NetworkConfig struct {
	VPC_CIDR string   `mapstructure:"vpc_cidr"`
	Subnets  []string `mapstructure:"subnets"`
}

// StorageConfig define os parâmetros para recursos de armazenamento (ex: S3, GCS).
type StorageConfig struct {
	BucketName string `mapstructure:"bucket_name"`
	PublicRead bool   `mapstructure:"public_read"`
}

// GetEnvironment retorna a configuração para um ambiente específico pelo nome.
// Retorna nil se o ambiente não for encontrado.
func (c *Config) GetEnvironment(name string) *Environment {
	for _, env := range c.Environments {
		if env.Name == name {
			return env
		}
	}
	return nil
}
