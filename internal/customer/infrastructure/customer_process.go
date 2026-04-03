package infrastructure

import (
	"context"
	"database/sql"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/application"
)

// RunCustomerProcess gestiona la inicialización y ejecución del proceso de importación de clientes
func RunCustomerProcess(ctx context.Context, db *sql.DB, cfg *config.Database) error {
	// Ruta al archivo CSV (Configurable si se prefiere mover a config)
	csvPath := "data/Anexo 11 - Clientes - Tarea 3.csv"

	repo := NewPostgresCustomerRepository(db, cfg)
	parser := NewCSVParser(csvPath)

	// Orquestar caso de uso
	etl := application.NewCrearCliente(repo, parser)

	// Ejecutar y retornar resultado
	return etl.Execute(ctx)
}
