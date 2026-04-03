package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/customer/domain"
)

type PostgresCustomerRepository struct {
	db *sql.DB
}

func NewPostgresCustomerRepository(db *sql.DB) *PostgresCustomerRepository {
	return &PostgresCustomerRepository{db: db}
}

func (r *PostgresCustomerRepository) Save(ctx context.Context, c domain.Customer) error {
	// Llamada al procedimiento almacenado
	query := `CALL insertar_cliente($1, $2, $3, $4, $5, $6, $7)`

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
