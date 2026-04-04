package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/product/domain"
)

type PostgresProductRepository struct {
	db       *sql.DB
	configDb *config.Database
}

func NewPostgresProductRepository(db *sql.DB, configDb *config.Database) *PostgresProductRepository {
	return &PostgresProductRepository{db: db, configDb: configDb}
}

func (r *PostgresProductRepository) Save(ctx context.Context, p domain.Product) error {
	query := fmt.Sprintf(`
		INSERT INTO %s.PRODUCTO (
			id_producto, 
			producto_nombre, 
			marca, 
			modelo, 
			categoria, 
			subcategoria, 
			precio
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id_producto) DO NOTHING`, r.configDb.Schema)

	_, err := r.db.ExecContext(ctx,
		query,
		p.ID,
		p.Nombre,
		p.Marca,
		p.Modelo,
		p.Categoria,
		p.Subcategoria,
		p.Precio,
	)

	if err != nil {
		return fmt.Errorf("error al insertar producto %s: %w", p.ID, err)
	}

	return nil
}

func (r *PostgresProductRepository) Count(ctx context.Context) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.PRODUCTO", r.configDb.Schema)
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error al contar registros en PRODUCTO: %w", err)
	}
	return count, nil
}
