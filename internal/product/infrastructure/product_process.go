package infrastructure

import (
	"context"
	"database/sql"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/product/application"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/product/infrastructure/csv"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/product/infrastructure/db"
)

// ProductTask representa la tarea ETL para la entidad Producto
type ProductTask struct{}

func (t *ProductTask) Name() string {
	return "Productos (Product Module)"
}

func (t *ProductTask) Run(ctx context.Context, database *sql.DB, cfg *config.Database) error {
	csvPath := "data/Anexo 9 - Productos - Tarea 3.csv"

	repo := db.NewPostgresProductRepository(database, cfg)
	parser := csv.NewCSVParser(csvPath)

	// Orquestar caso de uso
	etl := application.NewCreateProduct(repo, parser)

	// Ejecutar proceso
	return etl.Execute(ctx)
}
