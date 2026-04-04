package infrastructure

import (
	"context"
	"database/sql"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/application"
)

// CustomerTask representa la tarea ETL para clientes
type CustomerTask struct{}

func (t *CustomerTask) Name() string {
	return "Clientes (Customer Module)"
}

func (t *CustomerTask) Run(ctx context.Context, db *sql.DB, cfg *config.Database) error {
	// Ruta al archivo CSV
	csvPath := "data/Anexo 11 - Clientes - Tarea 3.csv"

	repo := NewPostgresCustomerRepository(db, cfg)
	parser := NewCSVParser(csvPath)

	// Orquestar caso de uso
	etl := application.NewCreateCustomer(repo, parser)

	// Ejecutar y retornar resultado
	return etl.Execute(ctx)
}
