package infrastructure

import (
	"context"
	"database/sql"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/sale/application"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/sale/infrastructure/csv"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/sale/infrastructure/db"
)

// SaleTask representa la tarea ETL para la tabla de hechos Fact Ventas
type SaleTask struct{}

func (t *SaleTask) Name() string {
	return "Ventas (Sales Fact Module)"
}

func (t *SaleTask) Run(ctx context.Context, database *sql.DB, cfg *config.Database) error {
	csvPath := "data/Anexo 12 - Ventas - Tarea 3.csv"

	repo := db.NewPostgresSaleRepository(database, cfg)
	parser := csv.NewCSVParser(csvPath)

	// Orquestar caso de uso
	etl := application.NewCreateSale(repo, parser)

	// Ejecutar proceso
	return etl.Execute(ctx)
}
