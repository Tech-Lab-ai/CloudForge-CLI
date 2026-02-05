package state

import (
	"encoding/json"
	"os"
	"path/filepath"
    "errors"
)

// readStateFile lê e decodifica o arquivo de estado JSON.
// Retorna um erro específico se o arquivo não existir.
func readStateFile(filePath string) (*State, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
        if os.IsNotExist(err) {
            return nil, err // Retorna o erro original para que o StateManager possa tratá-lo
        }
		return nil, err
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}

	return &state, nil
}

// writeStateFile codifica e escreve o estado no arquivo JSON.
// Garante que o diretório de destino exista.
func writeStateFile(filePath string, state *State) error {
	// Garante que o diretório onde o arquivo será salvo exista
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	// Escreve o arquivo com permissões restritivas
	return os.WriteFile(filePath, data, 0644)
}

// isNotExistError é uma função auxiliar para verificar se um erro é do tipo 'não encontrado'.
func isNotExistError(err error) bool {
    return errors.Is(err, os.ErrNotExist)
}
