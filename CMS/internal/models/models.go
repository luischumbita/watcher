// models.go
package main

type Horario struct {
	Dia        string `json:"dia"`         // Ej: "LUNES"
	HoraInicio string `json:"hora_inicio"` // Ej: "8:00"
	HoraFin    string `json:"hora_fin"`    // Ej: "9:30"
	Asignatura string `json:"asignatura"`  // Ej: "Métodos de la Investigación Literaria"
}

type Aula struct {
	Edificio     string    `json:"edificio"`  // Ej: "MODULO 5"
	Numero       string    `json:"numero"`    // Ej: "101.0"
	Capacidad    int       `json:"capacidad"` // Ej: 20
	Equipamiento []string  `json:"equipamiento,omitempty"`
	Horarios     []Horario `json:"horarios"`
}
