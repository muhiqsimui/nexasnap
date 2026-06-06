package writer

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteJSON(path string, data any) error {

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, 0644)
}