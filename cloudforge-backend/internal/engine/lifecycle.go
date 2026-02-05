package engine

import (
	"fmt"

	"cloudforge/internal/config"
	"cloudforge/internal/engine/grpc_server"
	"cloudforge/internal/state"
    "cloudforge/proto"
)

// EngineLifecycle gerencia a inicialização e o ciclo de vida dos componentes da engine.
// Ele amarra a configuração, o estado e o servidor gRPC.
type EngineLifecycle struct {
	config *config.Config
	daemon *Daemon
	stateManager *state.StateManager
    grpcServer *grpc_server.Server
}

// NewLifecycle cria e configura uma nova instância do ciclo de vida da engine.
// Carrega a configuração, inicializa o state manager e prepara o servidor gRPC.
func NewLifecycle(configPath, environment string) (*EngineLifecycle, error) {
    // Carrega a configuração principal da aplicação
	cfg, err := config.LoadConfig(configPath, environment)
	if err != nil {
		return nil, fmt.Errorf("falha ao carregar configuração para a engine: %w", err)
	}

    // Inicializa o gerenciador de estado
    sm, err := state.NewStateManager(cfg)
    if err != nil {
        return nil, fmt.Errorf("falha ao inicializar o state manager: %w", err)
    }

    // Cria o daemon que gerenciará os processos da engine
	daemon := NewDaemon()

    // Cria o servidor gRPC, injetando as dependências necessárias
    grpcServer := grpc_server.NewServer(sm) // Passa o state manager para o servidor

    // Registra o serviço CloudForge no servidor gRPC do daemon
    proto.RegisterCloudForgeServiceServer(daemon.GrpcServer(), grpcServer)

	return &EngineLifecycle{
		config: cfg,
		daemon: daemon,
        stateManager: sm,
        grpcServer: grpcServer,
	}, nil
}

// Start inicia a engine, que por sua vez inicia o daemon.
// Esta é uma operação bloqueante que espera por um sinal de interrupção.
func (e *EngineLifecycle) Start() error {
	fmt.Println("▶️  Iniciando a Engine do CloudForge...")
    fmt.Printf("    - Projeto: %s\n", e.config.Project)
    fmt.Printf("    - Provedor: %s\n", e.config.Provider)
    fmt.Printf("    - Ambiente: %s\n", e.config.CurrentEnvironment)

	return e.daemon.Start()
}
