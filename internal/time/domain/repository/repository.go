package repository

import (
	"context"
	"github.com/nicolas-170/ETL-Electro-hogar/internal/time/domain"
)

// TimeRepository define las operaciones de persistencia para el módulo de Tiempo
type TimeRepository interface {
	Save(ctx context.Context, time domain.Time) error
	Count(ctx context.Context) (int, error)
}

// CSVParser define la abstracción para extraer datos desde archivos planos
type CSVParser interface {
	Parse() ([]domain.Time, error)
}
