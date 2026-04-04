package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nicolas-170/ETL-Electro-hogar/internal/config"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/sale/domain"
)

type PostgresSaleRepository struct {
	db       *sql.DB
	configDb *config.Database
}

func NewPostgresSaleRepository(db *sql.DB, configDb *config.Database) *PostgresSaleRepository {
	return &PostgresSaleRepository{db: db, configDb: configDb}
}

func (r *PostgresSaleRepository) Save(ctx context.Context, s domain.Sale) error {
	query := fmt.Sprintf(`
		INSERT INTO %s.VENTA (
			id_venta, 
			id_tiempo, 
			id_cliente, 
			id_producto, 
			id_tienda, 
			cantidad, 
			precio_unitario, 
			total_venta, 
			costo_unitario, 
			margen_unitario, 
			margen_total
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (id_venta) DO NOTHING`, r.configDb.Schema)

	_, err := r.db.ExecContext(ctx,
		query,
		s.ID,
		s.IDTiempo,
		s.IDCliente,
		s.IDProducto,
		s.IDTienda,
		s.Cantidad,
		s.PrecioUnitario,
		s.TotalVenta,
		s.CostoUnitario,
		s.MargenUnitario,
		s.MargenTotal,
	)

	if err != nil {
		return fmt.Errorf("error al insertar venta %s: %w", s.ID, err)
	}

	return nil
}

func (r *PostgresSaleRepository) Count(ctx context.Context) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.VENTA", r.configDb.Schema)
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error al contar registros en VENTA: %w", err)
	}
	return count, nil
}
