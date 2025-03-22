// excel_parser.go
package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Aula struct {
	Edificio  string
	Numero    string
	Capacidad int
	Horarios  []Horario
}

type Horario struct {
	Dia        string
	HoraInicio string
	Asignatura string
}

// Función que utiliza un filePath por defecto
func ParseDefaultExcel() ([]Aula, error) {

	filePath := `C:\Users\MSI\Desktop\CMS\internal\src\Aulero 2024 - Sede Capital.xlsx`

	return parseExcel(filePath)
}

func parseExcel(filePath string) ([]Aula, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer f.Close()

	var aulas []Aula

	// Iterar sobre todas las hojas
	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil {
			log.Printf("Error al leer hoja %s: %v", sheet, err)
			continue
		}

		// Buscar filas que representan aulas (ej: "AULA 201")
		for rowIdx, row := range rows {
			if len(row) < 5 || row[0] == "" {
				continue // Saltar filas vacías o sin datos relevantes
			}

			// Ejemplo: Extraer datos de la fila "AULA 201 | 50.0 | ..."
			if esFilaDeAula(row) {
				aula, err := parseAula(row, rows[rowIdx+1:])
				if err != nil {
					log.Printf("Error en fila %d: %v", rowIdx, err)
					continue
				}
				aulas = append(aulas, aula)
			}
		}
	}

	return aulas, nil
}

// Helper para detectar filas de aula (personaliza según tu estructura)
func esFilaDeAula(row []string) bool {
	matched, _ := regexp.MatchString(`^AULA \d+`, row[0])
	return matched
}

// Convertir celdas a estructura Aula
func parseAula(headerRow []string, horarioRows [][]string) (Aula, error) {
	capacidad, _ := strconv.Atoi(headerRow[2])
	aula := Aula{
		Edificio:  headerRow[0],
		Numero:    headerRow[1],
		Capacidad: capacidad,
	}

	// Mapear horarios (ej: columnas 4 en adelante)
	for _, row := range horarioRows {
		for colIdx, cell := range row {
			if colIdx < 4 || cell == "" {
				continue // Saltar columnas no relevantes
			}
			// Ejemplo: Extraer día y hora de la cabecera
			dia, horaInicio := parseDiaYHora(headerRow[colIdx])
			aula.Horarios = append(aula.Horarios, Horario{
				Dia:        dia,
				HoraInicio: horaInicio,
				Asignatura: cell,
			})
		}
	}

	return aula, nil
}

// Convertir "8.0" -> "8:00" y detectar día
func parseDiaYHora(cabecera string) (string, string) {
	// Ejemplo: Si la cabecera es "LUNES | 8.0"
	re := regexp.MustCompile(`([A-Z]+)\s*\|\s*(\d+\.\d+)`)
	matches := re.FindStringSubmatch(cabecera)
	if len(matches) == 3 {
		hora := fmt.Sprintf("%s:%02d", matches[2][:1], 0) // "8.0" -> "8:00"
		return matches[1], hora
	}
	return "DESCONOCIDO", "00:00"
}
