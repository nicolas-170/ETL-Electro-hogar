package repository

import (
	"context"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/sale/domain"
)

// SaleRepository define las operaciones de persistencia para el módulo de Ventas
type SaleRepository interface {
	Save(ctx context.Context, sale domain.Sale) error
	Count(ctx context.Context) (int, error)
}

// CSVParser define la abstracción para extraer datos desde archivos planos
type CSVParser interface {
	Parse() ([]domain.Sale, error)
}
