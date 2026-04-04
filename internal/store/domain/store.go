package domain

// Store representa la entidad Tienda/Ciudad en el proceso ETL
type Store struct {
	ID           string `json:"id_ciudad"`
	Nombre       string `json:"ciudad"`
	Departamento string `json:"departamento"`
	Region       string `json:"region"`
}
