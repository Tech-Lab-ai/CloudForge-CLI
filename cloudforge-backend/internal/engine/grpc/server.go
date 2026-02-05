package grpc_server

import (
	"cloudforge/internal/state"
	"cloudforge/proto"
)

// Server implementa a interface gRPC `CloudForgeServiceServer`.
// Ele atua como o ponto de entrada para todas as requisições da CLI,
// orquestrando as operações de estado, provisionamento e sincronização.
type Server struct {
	proto.UnimplementedCloudForgeServiceServer // Para compatibilidade futura

	stateManager *state.StateManager
	// Adicionar outras dependências aqui (ex: provisioner, syncer)
}

// NewServer cria uma nova instância do servidor gRPC.
// Recebe todas as dependências necessárias para processar as requisições.
func NewServer(sm *state.StateManager) *Server {
	return &Server{
		stateManager: sm,
	}
}
