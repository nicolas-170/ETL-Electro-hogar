package infrastructure

import (
	"context"
	"database/sql"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/time/application"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/time/infrastructure/csv"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/time/infrastructure/db"
)

// RunTimeProcess gestiona la inicialización y ejecución del proceso de importación de la dimensión tiempo
func RunTimeProcess(ctx context.Context, database *sql.DB, cfg *config.Database) error {
	// Ruta al archivo CSV
	csvPath := "data/Anexo 10- Tiempo - Tarea 3.csv"

	repo := db.NewPostgresTimeRepository(database, cfg)
	parser := csv.NewCSVParser(csvPath)

	// Orquestar caso de uso
	etl := application.NewCreateTime(repo, parser)

	// Ejecutar y retornar resultado
	return etl.Execute(ctx)
}
