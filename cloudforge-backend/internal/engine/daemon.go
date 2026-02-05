package engine

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
    "google.golang.org/grpc/health"
    "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	DefaultGRPCPort = ":9000"
	DefaultHealthCheckPort = ":9001"
)

// Daemon gerencia o ciclo de vida do processo da Engine, incluindo
// o servidor gRPC e os endpoints de health check.	ype Daemon struct {
	grpcServer *grpc.Server
	healthServer *health.Server
	httpServer   *http.Server
	shutdown     chan os.Signal
}

// NewDaemon cria e configura uma nova instância do daemon da Engine.
func NewDaemon() *Daemon {
	grpcSrv := grpc.NewServer()
	healthSrv := health.NewServer()

	// O health server é usado pelo Kubernetes, Docker Swarm, etc.
    // para saber se o container está saudável.
    grpc_health_v1.RegisterHealthServer(grpcSrv, healthSrv)

	return &Daemon{
		grpcServer: grpcSrv,
		healthServer: healthSrv,
		shutdown:     make(chan os.Signal, 1),
	}
}

// Start inicia todos os serviços da Engine: o listener gRPC e o servidor HTTP para health checks.
// Ele bloqueia a execução até que um sinal de interrupção seja recebido.
func (d *Daemon) Start() error {
	// Canal para escutar sinais de shutdown do S.O.
	signal.Notify(d.shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Inicia o listener da porta gRPC
	lis, err := net.Listen("tcp", DefaultGRPCPort)
	if err != nil {
		return fmt.Errorf("falha ao escutar na porta gRPC %s: %w", DefaultGRPCPort, err)
	}

	// Inicia o servidor HTTP para health checks em uma goroutine
	go d.startHealthCheckServer()

	// Define o status inicial dos serviços como SERVING
    // TODO: Registrar os serviços gRPC reais aqui
    // d.healthServer.SetServingStatus("CloudForgeService", grpc_health_v1.HealthCheckResponse_SERVING)

	fmt.Printf("🚀 Engine gRPC escutando em %s\n", DefaultGRPCPort)

	// Inicia o servidor gRPC (bloqueante)
	go func() {
		if err := d.grpcServer.Serve(lis); err != nil {
			fmt.Printf("❌ Erro fatal no servidor gRPC: %v\n", err)
            os.Exit(1)
		}
	}()

    // Aguarda por um sinal de shutdown
    <-d.shutdown
    fmt.Println("\n🔌 Recebido sinal de shutdown. Desligando graciosamente...")

    // Inicia o processo de graceful shutdown
    return d.stop()
}

// startHealthCheckServer inicia um servidor HTTP simples para expor o status de saúde.
func (d *Daemon) startHealthCheckServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// A verificação pode ser expandida para checar a conexão com o provedor, etc.
        // Por enquanto, se o servidor está no ar, está saudável.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	d.httpServer = &http.Server{
		Addr:    DefaultHealthCheckPort,
		Handler: mux,
	}

	fmt.Printf("🩺 Health check escutando em %s\n", DefaultHealthCheckPort)
	if err := d.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("❌ Erro no servidor de health check: %v\n", err)
        os.Exit(1)
	}
}

// stop realiza o graceful shutdown do daemon, fechando todas as conexões ativas.
func (d *Daemon) stop() error {
    // Inicia o shutdown do servidor HTTP
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := d.httpServer.Shutdown(ctx); err != nil {
        fmt.Printf("Erro no shutdown do servidor HTTP: %v\n", err)
    }

    // Inicia o graceful stop do servidor gRPC
    // Ele aguardará as RPCs em andamento terminarem.
    d.grpcServer.GracefulStop()

    fmt.Println("✅ Engine desligada com sucesso.")
    return nil
}

// GrpcServer retorna a instância do servidor gRPC para registro de serviços externos.
func (d *Daemon) GrpcServer() *grpc.Server {
    return d.grpcServer
}
