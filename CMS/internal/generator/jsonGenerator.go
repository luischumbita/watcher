// json_generator.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func generateJSON(aulas []Aula, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error al crear archivo JSON: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(aulas); err != nil {
		return fmt.Errorf("error al serializar JSON: %v", err)
	}

	return nil
}
