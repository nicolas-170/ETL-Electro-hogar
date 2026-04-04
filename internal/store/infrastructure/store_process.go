package infrastructure

import (
	"context"
	"database/sql"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/store/application"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/store/infrastructure/csv"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/store/infrastructure/db"
)

// StoreTask representa la tarea ETL para la entidad Tienda/Ciudad
type StoreTask struct{}

func (t *StoreTask) Name() string {
	return "Tiendas (Store Module)"
}

func (t *StoreTask) Run(ctx context.Context, database *sql.DB, cfg *config.Database) error {
	csvPath := "data/Anexo 8 - Tiendas - Tarea 3.csv"

	repo := db.NewPostgresStoreRepository(database, cfg)
	parser := csv.NewCSVParser(csvPath)

	// Orquestar caso de uso
	etl := application.NewCreateStore(repo, parser)

	// Ejecutar proceso
	return etl.Execute(ctx)
}
