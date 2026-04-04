package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/store/domain"
)

type PostgresStoreRepository struct {
	db       *sql.DB
	configDb *config.Database
}

func NewPostgresStoreRepository(db *sql.DB, configDb *config.Database) *PostgresStoreRepository {
	return &PostgresStoreRepository{db: db, configDb: configDb}
}

func (r *PostgresStoreRepository) Save(ctx context.Context, s domain.Store) error {
	query := fmt.Sprintf(`
		INSERT INTO %s.TIENDA (
			id_ciudad, 
			ciudad, 
			departamento, 
			region
		) VALUES ($1, $2, $3, $4)
		ON CONFLICT (id_ciudad) DO NOTHING`, r.configDb.Schema)

	_, err := r.db.ExecContext(ctx,
		query,
		s.ID,
		s.Nombre,
		s.Departamento,
		s.Region,
	)

	if err != nil {
		return fmt.Errorf("error al insertar tienda %s: %w", s.ID, err)
	}

	return nil
}

func (r *PostgresStoreRepository) Count(ctx context.Context) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.TIENDA", r.configDb.Schema)
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error al contar registros en TIENDA: %w", err)
	}
	return count, nil
}
