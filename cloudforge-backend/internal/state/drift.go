package state

import (
	"cloudforge/internal/providers"
    "fmt"
)

// Drift represents a discrepancy between the expected state and the actual state of a resource.	ype Drift struct {
	ResourceName string
	ResourceType string
	Details      string // Ex: "Attribute 'instance_type' changed from 't2.micro' to 't2.small'"
}

// DriftDetector is responsible for comparing the recorded state with the actual state
// of the infrastructure and reporting any discrepancies (drift).	ype DriftDetector struct {
	provider providers.Provider
}

// NewDriftDetector creates a new drift detector, configured for a specific cloud provider.
func NewDriftDetector(p providers.Provider) *DriftDetector {
	return &DriftDetector{provider: p}
}

// CheckDrift compares the given state with the actual infrastructure state.
// It iterates through each resource in the state file and queries the provider for its real status.
// Returns a list of drifts found.
func (d *DriftDetector) CheckDrift(state *State) ([]Drift, error) {
	var drifts []Drift

    if d.provider == nil {
        return nil, fmt.Errorf("o provedor não foi inicializado para o detector de drift")
    }

	// Itera sobre cada recurso registrado no estado
	for _, resourceState := range state.Resources {
		// Busca o estado real do recurso no provedor cloud
		actualResource, err := d.provider.GetResourceStatus(resourceState.ProviderID)
		if err != nil {
			// Se o recurso não existe mais, isso é um drift
			drifts = append(drifts, Drift{
				ResourceName: resourceState.Name,
				ResourceType: resourceState.Type,
				Details:      fmt.Sprintf("Recurso não encontrado no provedor (pode ter sido deletado manualmente)"),
			})
			continue
		}

		// Compara atributos-chave (exemplo simplificado)
        // Uma implementação real compararia todos os atributos gerenciados.
        if changed, details := compareAttributes(resourceState.Attributes, actualResource.Attributes); changed {
            drifts = append(drifts, Drift{
				ResourceName: resourceState.Name,
				ResourceType: resourceState.Type,
				Details:      details,
			})
        }
	}

	return drifts, nil
}

// compareAttributes é uma função auxiliar para comparar os atributos de um recurso.
// Retorna 'true' se houver diferenças e uma descrição delas.
func compareAttributes(expected, actual map[string]interface{}) (bool, string) {
    // Exemplo de lógica de comparação.
    // A implementação real seria mais robusta, iterando sobre chaves e valores.
    if fmt.Sprintf("%v", expected["size"]) != fmt.Sprintf("%v", actual["size"]) {
        return true, fmt.Sprintf("Atributo 'size' divergiu. Esperado: %v, Atual: %v", expected["size"], actual["size"])
    }

    return false, ""
}