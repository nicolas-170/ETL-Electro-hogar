package domain

import "time"

// Time define el modelo de dominio para la dimensión Tiempo en el proceso ETL
type Time struct {
	ID           string    `json:"id_tiempo"`
	Fecha        time.Time `json:"fecha"`
	Dia          int       `json:"dia"`
	Mes          int       `json:"mes"`
	Trimestre    int       `json:"trimestre"`
	Anio         int       `json:"anio"`
	DiaSemana    string    `json:"dia_semana"`
	EsFinSemana  bool      `json:"es_fin_semana"`
}
