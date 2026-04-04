package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/domain"
)

type PostgresCustomerRepository struct {
	db       *sql.DB
	configDb *config.Database
}

func NewPostgresCustomerRepository(db *sql.DB, configDb *config.Database) *PostgresCustomerRepository {
	return &PostgresCustomerRepository{db: db, configDb: configDb}
}

func (r *PostgresCustomerRepository) Save(ctx context.Context, c domain.Customer) error {
	// Llamada al procedimiento almacenado
	query := fmt.Sprintf(`CALL %s.insertar_cliente($1, $2, $3, $4, $5, $6, $7)`, r.configDb.Schema)
	_, err := r.db.ExecContext(ctx,
		query,
		c.Nombre,
		c.Email,
		c.Telefono,
		c.Segmento,
		c.Ciudad,
		c.Sexo,
		c.Edad,
	)

	if err != nil {
		return fmt.Errorf("error al ejecutar procedimiento insertar_cliente: %w", err)
	}

	return nil
}

func (r *PostgresCustomerRepository) Count(ctx context.Context) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.CLIENTE", r.configDb.Schema)
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error al contar registros en CLIENTE: %w", err)
	}
	return count, nil
}
