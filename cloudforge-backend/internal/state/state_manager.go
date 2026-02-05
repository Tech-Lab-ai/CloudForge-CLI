package state

import (
	"fmt"
	"sync"
    "cloudforge/internal/config"
)

// State representa a configuração e o estado atual da infraestrutura gerenciada.
// Contém todos os recursos provisionados, suas versões e metadados.
type State struct {
	Version     string            `json:"version"`
	ProjectID   string            `json:"project_id"`
	Provider    string            `json:"provider"`
	Environment string            `json:"environment"`
	Resources   []ResourceState   `json:"resources"`
}

// ResourceState representa o estado de um único recurso na nuvem.
type ResourceState struct {
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	ProviderID string                 `json:"provider_id"` // ID do recurso no provedor cloud
	Attributes map[string]interface{} `json:"attributes"`
}

// StateManager gerencia o ciclo de vida do arquivo de estado (cloudforge.tfstate).
// É responsável por carregar, salvar, versionar e prevenir condições de corrida.
type StateManager struct {
	mu       sync.Mutex
	config   *config.Config
	currentState *State
    filePath string
}

// NewStateManager cria um novo gerenciador de estado para um ambiente específico.
// Ele garante que todas as operações no estado sejam atômicas.
func NewStateManager(cfg *config.Config) (*StateManager, error) {
    if cfg == nil {
        return nil, fmt.Errorf("a configuração não pode ser nula")
    }
    
    // O caminho do arquivo de estado pode ser derivado da configuração
    // Por exemplo: .cloudforge/states/dev.tfstate
    filePath := fmt.Sprintf(".cloudforge/states/%s.tfstate", cfg.CurrentEnvironment)

	return &StateManager{
		config:   cfg,
		filePath: filePath,
	}, nil
}

// Load carrega o estado do arquivo local. Se o arquivo não existir, inicializa um novo estado.
func (sm *StateManager) Load() (*State, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	state, err := readStateFile(sm.filePath)
	if err != nil {
        // Se o arquivo não existe, é uma operação válida; inicializamos um estado vazio
        if isNotExistError(err) {
            return sm.initState(), nil
        }
		return nil, fmt.Errorf("falha ao carregar o arquivo de estado: %w", err)
	}
    sm.currentState = state
	return sm.currentState, nil
}

// initState cria uma estrutura de estado vazia, pronta para ser usada.
func (sm *StateManager) initState() *State {
    sm.currentState = &State{
        Version:     "1.0.0", // Versão inicial do schema de estado
        ProjectID:   sm.config.Project,
        Provider:    sm.config.Provider,
        Environment: sm.config.CurrentEnvironment,
        Resources:   []ResourceState{},
    }
    return sm.currentState
}

// Save persiste o estado atual no arquivo em disco.
// A operação é atômica para evitar corrupção de dados.
func (sm *StateManager) Save(state *State) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if state == nil {
		return fmt.Errorf("tentativa de salvar um estado nulo")
	}

	if err := writeStateFile(sm.filePath, state); err != nil {
		return fmt.Errorf("falha ao salvar o arquivo de estado: %w", err)
	}
    sm.currentState = state
	return nil
}

// CurrentState retorna uma cópia segura do estado atual em memória.
func (sm *StateManager) CurrentState() (*State, error) {
    sm.mu.Lock()
    defer sm.mu.Unlock()

    if sm.currentState == nil {
        return sm.Load()
    }
    
    // Retorna uma cópia para evitar modificações externas acidentais
    stateCopy := *sm.currentState
    return &stateCopy, nil
}

// AddResource adiciona um novo recurso ao estado.
func (sm *StateManager) AddResource(resource ResourceState) {
    sm.mu.Lock()
    defer sm.mu.Unlock()

    sm.currentState.Resources = append(sm.currentState.Resources, resource)
}

// RemoveResource remove um recurso do estado com base no seu nome e tipo.
func (sm *StateManager) RemoveResource(name, resourceType string) error {
    sm.mu.Lock()
    defer sm.mu.Unlock()

    index := -1
    for i, r := range sm.currentState.Resources {
        if r.Name == name && r.Type == resourceType {
            index = i
            break
        }
    }

    if index == -1 {
        return fmt.Errorf("recurso '%s' do tipo '%s' não encontrado no estado", name, resourceType)
    }

    // Remove o elemento mantendo a ordem
    sm.currentState.Resources = append(sm.currentState.Resources[:index], sm.currentState.Resources[index+1:]...)
    return nil
}
