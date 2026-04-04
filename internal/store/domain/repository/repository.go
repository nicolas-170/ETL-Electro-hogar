package repository

import (
	"context"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/store/domain"
)

// StoreRepository define las operaciones de persistencia para el módulo de Tienda
type StoreRepository interface {
	Save(ctx context.Context, store domain.Store) error
	Count(ctx context.Context) (int, error)
}

// CSVParser define la abstracción para extraer datos desde archivos planos
type CSVParser interface {
	Parse() ([]domain.Store, error)
}
