package repository

import (
	"context"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/product/domain"
)

// ProductRepository define las operaciones de persistencia para el módulo de Producto
type ProductRepository interface {
	Save(ctx context.Context, product domain.Product) error
	Count(ctx context.Context) (int, error)
}

// CSVParser define la abstracción para extraer datos desde archivos planos
type CSVParser interface {
	Parse() ([]domain.Product, error)
}
