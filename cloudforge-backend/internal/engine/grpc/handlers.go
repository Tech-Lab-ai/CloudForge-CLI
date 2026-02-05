package grpc_server

import (
	"context"
	"fmt"

	"cloudforge/proto"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// Init é o handler gRPC para o comando `cloudforge init`.
// Valida a requisição e prepara o ambiente, mas não cria recursos.
func (s *Server) Init(ctx context.Context, req *proto.InitRequest) (*proto.InitResponse, error) {
    fmt.Printf("[gRPC] Recebida requisição Init para o projeto: %s\n", req.ProjectId)

    if req.ProjectId == "" || req.Provider == "" || req.Environment == "" {
        return nil, status.Error(codes.InvalidArgument, "ProjectID, Provider e Environment são obrigatórios")
    }

    // Lógica de inicialização (ex: criar a estrutura de diretórios .cloudforge)
    // Por enquanto, apenas confirmamos o recebimento.

    return &proto.InitResponse{
        Status:  "Success",
        Message: fmt.Sprintf("Projeto '%s' inicializado com sucesso para o ambiente '%s' com o provedor '%s'.", req.ProjectId, req.Environment, req.Provider),
    }, nil
}

// Deploy é o handler para o comando `cloudforge deploy`.
// Carrega o estado, planeja as mudanças e aplica-as usando o provisionador.
func (s *Server) Deploy(ctx context.Context, req *proto.DeployRequest) (*proto.DeployResponse, error) {
    fmt.Printf("[gRPC] Recebida requisição Deploy para o ambiente: %s\n", req.Environment)

    // 1. Carregar o estado atual
    state, err := s.stateManager.Load()
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Falha ao carregar o estado: %v", err)
    }

    // 2. Lógica de provisionamento (a ser implementada com os providers)
    // Por enquanto, vamos simular a criação de um recurso e adicioná-lo ao estado.
    fmt.Println("[Deploy] Lógica de provisionamento ainda não implementada.")

    // 3. Salvar o novo estado
    if err := s.stateManager.Save(state); err != nil {
        return nil, status.Errorf(codes.Internal, "Falha ao salvar o novo estado: %v", err)
    }

    return &proto.DeployResponse{
        Status:      "Success",
        Message:     "Deploy executado com sucesso (simulado).",
        StateVersion: state.Version,
    }, nil
}

// Status é o handler para o comando `cloudforge status`.
// Carrega o estado e o retorna para a CLI.
func (s *Server) Status(ctx context.Context, req *proto.StatusRequest) (*proto.StatusResponse, error) {
    fmt.Printf("[gRPC] Recebida requisição Status para o ambiente: %s\n", req.Environment)

    state, err := s.stateManager.Load()
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Falha ao carregar o estado: %v", err)
    }

    // A detecção de drift será implementada aqui posteriormente

    return &proto.StatusResponse{
        Status:       "Success",
        Provider:     state.Provider,
        StateVersion: state.Version,
        DriftDetected: false,
        // Resources: resources, // A ser preenchido
    }, nil
}

// --- Handlers a serem implementados ---

func (s *Server) Sync(ctx context.Context, req *proto.SyncRequest) (*proto.SyncResponse, error) {
    return nil, status.Error(codes.Unimplemented, "O método Sync ainda não foi implementado")
}

func (s *Server) Destroy(ctx context.Context, req *proto.DestroyRequest) (*proto.DestroyResponse, error) {
    return nil, status.Error(codes.Unimplemented, "O método Destroy ainda não foi implementado")
}
