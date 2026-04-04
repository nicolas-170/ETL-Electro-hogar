package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/time/domain"
)

type PostgresTimeRepository struct {
	db       *sql.DB
	configDb *config.Database
}

func NewPostgresTimeRepository(db *sql.DB, configDb *config.Database) *PostgresTimeRepository {
	return &PostgresTimeRepository{db: db, configDb: configDb}
}

func (r *PostgresTimeRepository) Save(ctx context.Context, t domain.Time) error {
	// Inserción directa (no procedimiento almacenado)
	query := fmt.Sprintf(`
		INSERT INTO %s.TIEMPO (
			id_tiempo, 
			fecha, 
			dia, 
			mes, 
			trimestre, 
			anio, 
			dia_semana, 
			es_fin_semana
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id_tiempo) DO NOTHING`, r.configDb.Schema)

	_, err := r.db.ExecContext(ctx,
		query,
		t.ID,
		t.Fecha,
		t.Dia,
		t.Mes,
		t.Trimestre,
		t.Anio,
		t.DiaSemana,
		t.EsFinSemana,
	)

	if err != nil {
		return fmt.Errorf("error al insertar registro en TIEMPO: %w", err)
	}

	return nil
}
