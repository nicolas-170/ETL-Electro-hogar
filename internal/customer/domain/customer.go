package domain

// Customer defines the domain model for a client in the ETL process
type Customer struct {
	ID       string `json:"id"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
	Segmento string `json:"segmento"`
	Ciudad   string `json:"ciudad"`
	Sexo     string `json:"sexo"`
	Edad     int    `json:"edad"`
}
