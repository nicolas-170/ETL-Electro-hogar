package etl

import (
	"context"
	"database/sql"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
)

// Task define la interfaz común para cualquier proceso ETL en el sistema
type Task interface {
	Name() string
	Run(ctx context.Context, db *sql.DB, cfg *config.Database) error
}
