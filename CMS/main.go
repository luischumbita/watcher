package main

import (
	"log"
)

func main() {
	excelPath := `C:\Users\MSI\Desktop\CMS\internal\src\Aulero 2024 - Sede Capital.xlsx`
	jsonPath := "aulas.json"

	// Procesamiento inicial
	aulas, err := parseExcel(excelPath)
	if err != nil {
		log.Fatal("Error al parsear Excel:", err)
	}

	if err := generateJSON(aulas, jsonPath); err != nil {
		log.Fatal("Error al generar JSON:", err)
	}

	// Monitorear cambios
	setupWatcher(excelPath, func() {
		aulas, err := parseExcel(excelPath)
		if err != nil {
			log.Println("Error al volver a parsear:", err)
			return
		}

		if err := generateJSON(aulas, jsonPath); err != nil {
			log.Println("Error al regenerar JSON:", err)
		}
	})
}
