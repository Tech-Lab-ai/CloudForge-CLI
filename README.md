CloudForge CLI

CloudForge CLI é uma ferramenta CLI especializada em Golang para provisionamento automatizado de infraestrutura em nuvem, sincronização de ambientes e padronização DevOps, projetada para alta performance, extensibilidade e excelente Developer Experience (DX).

Desenvolvido por: Vini Amaral

Visão Geral

O CloudForge CLI permite que desenvolvedores e times DevOps criem, sincronizem e gerenciem ambientes dev / staging / produção usando comandos simples, evitando configurações manuais em painéis de cloud.

O projeto foi concebido para atuar como um orquestrador inteligente local, com comunicação via gRPC, engine desacoplada e suporte a múltiplos provedores de nuvem.

Principais Objetivos

Automatizar provisionamento de recursos em nuvem

Sincronizar ambientes locais e remotos

Reduzir erros humanos em deploys

Padronizar infraestrutura como código

Servir como base para CI/CD, GitOps e SaaS

Stack Tecnológica
Core

Golang — binário único, rápido e multiplataforma

Cobra — estrutura de comandos CLI

Viper — gerenciamento de configurações

gRPC + Protocol Buffers — comunicação eficiente e tipada

Infraestrutura

Docker — execução isolada da engine

Docker Compose — ambientes replicáveis

Cloud SDK Abstraction Layer — AWS / GCP / Azure (plugável)

Arquitetura do Sistema
┌─────────────────────┐
│   CloudForge CLI    │
│   (Go Binary)       │
└─────────┬───────────┘
          │ gRPC
┌─────────▼───────────┐
│ CloudForge Engine   │
│ (Local Daemon)      │
│                     │
│ • Providers         │
│ • Provisioners      │
│ • Sync Engine       │
│ • State Manager     │
└───────┬─────┬───────┘
        │     │
     Docker  Cloud SDK
        │     │
     Infra   Providers

Estrutura de Diretórios
cloudforge/
├── cmd/
│   └── cloudforge/
│       └── main.go
├── internal/
│   ├── cli/
│   │   ├── root.go
│   │   ├── init.go
│   │   ├── deploy.go
│   │   ├── sync.go
│   │   ├── status.go
│   │   └── destroy.go
│   ├── engine/
│   │   ├── daemon.go
│   │   ├── grpc_server.go
│   │   └── lifecycle.go
│   ├── providers/
│   │   ├── aws/
│   │   ├── gcp/
│   │   ├── azure/
│   │   └── provider.go
│   ├── provisioner/
│   │   ├── compute.go
│   │   ├── storage.go
│   │   ├── network.go
│   │   └── database.go
│   ├── sync/
│   │   ├── fsync.go
│   │   ├── envsync.go
│   │   └── secrets.go
│   ├── state/
│   │   └── state_manager.go
│   └── config/
│       └── loader.go
├── proto/
│   └── cloudforge.proto
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
├── cloudforge.yaml
├── go.mod
└── README.md

Comandos Disponíveis
Inicializar Projeto
cloudforge init


Cria arquivo de configuração

Detecta provider

Inicializa estado local

Provisionar Infraestrutura
cloudforge deploy


Cria recursos em nuvem

Aplica templates

Versiona o estado

Sincronizar Ambientes
cloudforge sync


Sincroniza arquivos

Atualiza containers/VMs

Injeta variáveis e secrets

Ver Status
cloudforge status


Infra ativa

Drift detection

Diferenças entre ambientes

Destruir Infra
cloudforge destroy


Remove recursos

Limpa estado

Executa rollback seguro

Arquivo de Configuração (cloudforge.yaml)
project: insightai
provider: gcp

environments:
  dev:
    region: us-central1
    compute:
      type: container
      replicas: 2

  prod:
    region: us-east1
    compute:
      type: vm
      size: n2-standard-4

sync:
  include:
    - src/**
  exclude:
    - .git
    - node_modules

secrets:
  source: env

gRPC — Contrato Base
syntax = "proto3";

service CloudForge {
  rpc Deploy(DeployRequest) returns (DeployResponse);
  rpc Sync(SyncRequest) returns (SyncResponse);
  rpc Status(StatusRequest) returns (StatusResponse);
  rpc Destroy(DestroyRequest) returns (DestroyResponse);
}

Diferenciais Técnicos

Cloud-agnostic real

Engine desacoplada da CLI

Comunicação gRPC (pronto para UI, agentes e SaaS)

Estado versionado (inspirado em Terraform)

Alta extensibilidade via providers e plugins

Casos de Uso

Startups SaaS

Produtos de IA

Ambientes DevOps padronizados

CI/CD e GitOps

Infraestrutura como produto

Roadmap

Sistema de plugins

UI Web opcional

Execução remota distribuída

Drift auto-correction

Modo SaaS multi-tenant

Autor

Vini Amaral
Engenharia de Software • Tech Labss • Arquitetura Cloud • Golang

Projeto desenvolvido com foco em automação real, infraestrutura moderna e liberdade criativa para desenvolvedores.# CloudForge-CLI
